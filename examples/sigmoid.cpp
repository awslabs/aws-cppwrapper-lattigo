// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "latticpp/ckks/ciphertext.h"
#include "latticpp/ckks/decryptor.h"
#include "latticpp/ckks/encoder.h"
#include "latticpp/ckks/encryptor.h"
#include "latticpp/ckks/evaluator.h"
#include "latticpp/ckks/keygen.h"
#include "latticpp/ckks/params.h"
#include "latticpp/marshal/gohandle.h"

#include <cmath>
#include <iostream>
#include <map>
#include <vector>
#include <iomanip>

using namespace std;
using namespace latticpp;

typedef map<uint64_t, Ciphertext> ChebyMap;
typedef pair<uint64_t, Ciphertext> ChebyKeyValuePair;

// generate a random vector of the given dimension, where each value is in the range [-maxNorm, maxNorm].
vector<double> randomVector(int dim, double maxNorm) {
    vector<double> x(dim);

    for (int i = 0; i < dim; i++) {
        // generate a random double between -maxNorm and maxNorm
        double a = -maxNorm + ((static_cast<double>(random())) / (static_cast<double>(RAND_MAX))) * (2 * maxNorm);
        x[i] = a;
    }
    return x;
}

// checkEnoughLevels checks that enough levels are available to evaluate the polynomial.
// Also checks if c is a gaussian integer or not. If not, then one more level is needed
// to evaluate the polynomial.
void checkEnoughLevels(int levels, int numCoeffs, double c) {

    int logDegree = int(log2(double(numCoeffs)) + 0.5);

    if (c != (double)((uint64_t)c)) {
        logDegree++;
    }

    if (levels < logDegree) {
        cout << levels << " levels < " << logDegree << " log(d) -> cannot evaluate" << endl;
        throw invalid_argument("checkEnoughLevels");
    }
}

struct ChebyshevInterpolation {
    int a;
    int b;
    int maxDeg;
    bool lead;
    vector<double> coeffs;

    int degree() {
        return coeffs.size() - 1;
    }
};

int bitCount(int x) {
    return ceil(log2((double)(x + 1)));
}

bool contains(const ChebyMap &c, uint64_t k) {
    // if iterator == c.end(), then `k` is not in the map
    // http://www.cplusplus.com/reference/map/map/find/
    return c.find(k) != c.end();
}

void computePowerBasisCheby(uint64_t n, ChebyMap &cMap, Evaluator &eval, EvaluationKey &evakey, Parameters &params) {

    // Given a hash table with the first three evaluations of the Chebyshev ring at x in the interval a, b:
    // C0 = 1 (actually not stored in the hash table)
    // C1 = (2*x - a - b)/(b-a)
    // C2 = 2*C1*C1 - C0
    // Evaluates the nth degree Chebyshev ring in a recursive manner, storing intermediate results in the hashtable.
    // Consumes at most ceil(sqrt(n)) levels for an evaluation at Cn.
    // Uses the following property: for a given Chebyshev ring Cn = 2*Ca*Cb - Cc, n = a+b and c = abs(a-b)

    if (!contains(cMap, n)) {

        // Computes the index required to compute the asked ring evaluation
        uint64_t a = (uint64_t)ceil(((double)n) / 2);
        uint64_t b = n >> 1;
        uint64_t c = (uint64_t)abs(((double)a) - ((double)b));

        // Recurses on the given indexes
        computePowerBasisCheby(a, cMap, eval, evakey, params);
        computePowerBasisCheby(b, cMap, eval, evakey, params);

        // Since cMap[0] is not stored (but rather seen as the constant 1), only recurses on c if c!= 0
        if (c != 0) {
            computePowerBasisCheby(c, cMap, eval, evakey, params);
        }

        // Computes cMap[n] = cMap[a]*cMap[b]
        cMap.insert(ChebyKeyValuePair(n, mulRelinNew(eval, cMap.at(a), cMap.at(b), evakey)));
        rescale(eval, cMap.at(n), scale(params), cMap.at(n));

        // Computes cMap[n] = 2*cMap[a]*cMap[b]
        add(eval, cMap.at(n), cMap.at(n), cMap.at(n));

        // Computes cMap[n] = 2*cMap[a]*cMap[b] - cMap[c]
        if (c == 0) {
            addConst(eval, cMap.at(n), -1, cMap.at(n));
        } else {
            sub(eval, cMap.at(n), cMap.at(c), cMap.at(n));
        }
    }
}

Ciphertext evaluatePolyFromPowerBasis(double targetScale, ChebyshevInterpolation &cheby, ChebyMap &cMap, Evaluator &eval, Parameters &params) {

    if (cheby.degree() == 0) {

        Ciphertext res = newCiphertext(params, 1, level(cMap.at(1)), targetScale);

        if (abs(cheby.coeffs[0]) > 1e-14) {
            addConst(eval, res, cheby.coeffs[0], res);
        }

        return res;
    }

    double currentQi = (double)qi(params, level(cMap.at(cheby.degree())));

    double ctScale = targetScale * currentQi;

    Ciphertext res = newCiphertext(params, 1, level(cMap.at(cheby.degree())), ctScale);

    if (abs(cheby.coeffs[0]) > 1e-14) {
        addConst(eval, res, cheby.coeffs[0], res);
    }

    for (int key = cheby.degree(); key > 0; key--) {

        if (key != 0 && abs(cheby.coeffs[key]) > 1e-14) {

            // Target scale * rescale-scale / power basis scale
            double constScale = targetScale * currentQi / scale(cMap.at(key));

            uint64_t cReal = (int64_t)(cheby.coeffs[key] * constScale);
            uint64_t cImag = 0; //int64(imag(coeffs.coeffs[key]) * constScale)

            multByGaussianIntegerAndAdd(eval, cMap.at(key), cReal, cImag, res);
        }
    }

    rescale(eval, res, scale(params), res);

    return res;
}


struct ChebyPair {
    ChebyshevInterpolation coeffsq;
    ChebyshevInterpolation coeffsr;
};

ChebyPair splitCoeffsCheby(ChebyshevInterpolation &cheby, uint64_t split) {
    // Splits a Chebyshev polynomial p such that p = q*C^degree + r, where q and r are a linear combination of a Chebyshev basis.
    ChebyshevInterpolation coeffsr;
    coeffsr.coeffs = vector<double>(split);
    if (cheby.maxDeg == cheby.degree()) {
        coeffsr.maxDeg = split - 1;
    } else {
        coeffsr.maxDeg = cheby.maxDeg - (cheby.degree() - split + 1);
    }

    for (uint64_t i = 0; i < split; i++) {
        coeffsr.coeffs[i] = cheby.coeffs[i];
    }

    ChebyshevInterpolation coeffsq;
    coeffsq.coeffs = vector<double>(cheby.degree() - split + 1);
    coeffsq.maxDeg = cheby.maxDeg;

    coeffsq.coeffs[0] = cheby.coeffs[split];

    uint64_t j = 1;
    for (uint64_t i = split + 1; i < cheby.degree() + 1; i++, j++) {
        coeffsq.coeffs[i-split] = 2 * cheby.coeffs[i];
        coeffsr.coeffs[split-j] -= cheby.coeffs[i];
    }

    if (cheby.lead) {
        coeffsq.lead = true;
    }

    return ChebyPair{coeffsq, coeffsr};
}

Ciphertext recurseCheby(double targetScale, int logSplit, int logDegree, ChebyshevInterpolation &cheby, ChebyMap &cMap, Evaluator &eval, EvaluationKey &evakey, Parameters &params) {
    // Recursively computes the evalution of the Chebyshev polynomial using a baby-set giant-step algorithm.
    if (cheby.degree() < (((uint64_t)1) << logSplit)) {
        if (cheby.lead && cheby.maxDeg > ((((uint64_t)1) << logDegree) - (((uint64_t)1) << (logSplit-1))) && logSplit > 1) {

            logDegree = bitCount(cheby.degree());
            logSplit = logDegree >> 1;

            return recurseCheby(targetScale, logSplit, logDegree, cheby, cMap, eval, evakey, params);
        }

        return evaluatePolyFromPowerBasis(targetScale, cheby, cMap, eval, params);
    }

    uint64_t nextPower = ((uint64_t)1) << logSplit;
    while (nextPower < ((cheby.degree() >> 1) + 1)) {
        nextPower <<= 1;
    }

    ChebyPair coeffs = splitCoeffsCheby(cheby, nextPower);
    ChebyshevInterpolation coeffsq = coeffs.coeffsq;
    ChebyshevInterpolation coeffsr = coeffs.coeffsr;

    int ctLevel = level(cMap.at(nextPower)) - 1;

    if (coeffsq.maxDeg >= (((uint64_t)1) << (logDegree-1)) && coeffsq.lead) {
        ctLevel++;
    }

    double currentQi = (double)qi(params, ctLevel);

    Ciphertext res = recurseCheby(targetScale*currentQi/scale(cMap.at(nextPower)), logSplit, logDegree, coeffsq, cMap, eval, evakey, params);

    Ciphertext tmp = recurseCheby(targetScale, logSplit, logDegree, coeffsr, cMap, eval, evakey, params);

    if (level(res) > level(tmp)) {
        while (level(res) != level(tmp)+1) {
            dropLevel(eval, res, 1);
        }
    }

    mulRelin(eval, res, cMap.at(nextPower), evakey, res);

    if (level(res) > level(tmp)) {
        rescale(eval, res, scale(params), res);
        add(eval, res, tmp, res);
    } else {
        add(eval, res, tmp, res);
        rescale(eval, res, scale(params), res);
    }

    return res;
}

// EvaluateCheby evaluates a polynomial in Chebyshev basis on the input Ciphertext in ceil(log2(deg+1))+1 levels.
// Returns an error if the input ciphertext does not have enough level to carry out the full polynomial evaluation.
// Returns an error if something is wrong with the scale.
// A change of basis ct' = (2/(b-a)) * (ct + (-a-b)/(b-a)) is necessary before the polynomial evaluation to ensure correctness.
Ciphertext evaluateCheby(Evaluator &eval, Ciphertext &op, ChebyshevInterpolation &cheby, EvaluationKey &evakey, Parameters &params) {

    checkEnoughLevels(level(op), cheby.coeffs.size(), 1);

    ChebyMap cMap;
    cMap.insert(ChebyKeyValuePair(1, copyNew(op)));

    int logDegree = bitCount(cheby.degree());
    int logSplit = (logDegree >> 1);

    for (uint64_t i = 2; i < (((uint64_t)1) << logSplit); i++) {
        computePowerBasisCheby(i, cMap, eval, evakey, params);
    }

    for (uint64_t i = logSplit; i < logDegree; i++) {
        computePowerBasisCheby(((uint64_t)1) << i, cMap, eval, evakey, params);
    }

    Ciphertext opOut = recurseCheby(scale(params), logSplit, logDegree, cheby, cMap, eval, evakey, params);

    return opOut;
}

double sigmoid(double x) {
    return 1 / (exp(-x) + 1);
}

int main() {
    // initialize random generator
    srand(time(nullptr));

    // Scheme params
    Parameters params = getParams(PN14QP438);

    Encoder encoder = newEncoder(params);

    // Keys
    KeyGenerator kgen = newKeyGenerator(params);
    struct KeyPairHandle kp = genKeyPair(kgen);

    // Relinearization key
    EvaluationKey rlk = genRelinKey(kgen, kp.sk);

    // Encryptor
    Encryptor encryptor = newEncryptorFromPk(params, kp.pk);

    // Decryptor
    Decryptor decryptor = newDecryptor(params, kp.sk);

    // Evaluator
    Evaluator evaluator = newEvaluator(params);

    // Values to encrypt
    uint64_t num_slots = numSlots(params);
    vector<double> values = randomVector(num_slots, 8);

    cout << "CKKS parameters: logN = " << logN(params) << ", logQ = " << logQP(params) << ", levels = "
         << (maxLevel(params) + 1) << ", scale = " << scale(params) << ", sigma = " << sigma(params) << endl;
    cout << "Values: " << values[0] << " " << values[1] << " " << values[2] << " " << values[3] << endl;

    // Plaintext creation and encoding process
    Plaintext plaintext = encodeNTTAtLvlNew(params, encoder, values, maxLevel(params), scale(params));

    // Encryption process
    Ciphertext ciphertext = encryptNew(encryptor, plaintext);

    cout << "Evaluation of the function 1/(exp(-x)+1) in the range [-8, 8] (degree of approximation: 32)" << endl;

    // the Lattigo sigmoid example computes these constants; I'm just hard-coding them here
    ChebyshevInterpolation cheby;
    cheby.a = -8;
    cheby.b = 8;
    cheby.maxDeg = 33;
    cheby.lead = true;
    cheby.coeffs = vector({0.5, 0.6190561250008967, -8.750149297941835e-17, -0.16826527295488336,
                       -1.1578322553169466e-16, 0.07086825147354787, -2.2487309706325084e-16,
                       -0.03212482622239708, -4.821032513153558e-16, 0.014825676596539702,
                       -3.0006251772512057e-16, -0.006875369709170067, -3.039273795870395e-17,
                       0.003192799517808038, 5.145304627624297e-16, -0.001483261054825451,
                       8.829774381021138e-16, 0.0006891491944659845, 1.9437289553627407e-15,
                       -0.00032020333178368437, 3.1327447784583217e-15, 0.0001487836210064365,
                       2.6020931860400084e-15, -6.91422920543282e-05, 2.1701805232359335e-15,
                       3.2151561941196054e-05, 2.7115705004067323e-15, -1.4993552804346294e-05,
                       3.673678752950693e-15, 7.084323706714906e-06, 3.290240290953957e-15,
                       -3.545062005157911e-06, 2.967304694823051e-15, 2.1925728430065914e-06});

    // Change of variable
    multByConst(evaluator, ciphertext, 2.0 / ((double)(cheby.b - cheby.a)), ciphertext);
    addConst(evaluator, ciphertext, (-cheby.a - cheby.b) / (cheby.b - cheby.a), ciphertext);
    rescale(evaluator, ciphertext, scale(params), ciphertext);

    // We evaluate the interpolated Chebyshev interpolant on the ciphertext
    ciphertext = evaluateCheby(evaluator, ciphertext, cheby, rlk, params);

    cout << "Done...Consumed levels: " << (maxLevel(params) - level(ciphertext)) << endl;

    // Computation of the reference values
    vector<double> expectedResult(num_slots);
    for (int i = 0; i < num_slots; i++) {
        expectedResult[i] = sigmoid(values[i]);
    }

    // Print results and comparison
    vector<double> homomResult = decode(encoder, decryptNew(decryptor, ciphertext), logSlots(params));
    cout << endl;
    cout << "Level: " << level(ciphertext) << " (logQ = " << logQLvl(params, level(ciphertext)) << ")" << endl;
    cout << "Scale: 2^" << log2(scale(ciphertext)) << endl;
    cout << "Expected Result: " << setprecision(3) << expectedResult[0] << " " << expectedResult[1] << " " << expectedResult[2] << " " << expectedResult[3] << endl;
    cout << "Actual Result:   " << setprecision(3) << homomResult[0] << " " << homomResult[1] << " " << homomResult[2] << " " << homomResult[3] << endl;
    cout << endl;

    return 0;
}
