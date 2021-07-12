
#include "latticpp/latticpp.h"

#include <cmath>
#include <iomanip>
#include <random>
#include <vector>

using namespace std;
using namespace latticpp;

// generate a random vector of the given dimension, where each value is in the range [-maxNorm, maxNorm].
vector<double> randomVector(int dim, double maxNorm) {
    vector<double> x(dim);
    random_device r;
    uniform_real_distribution<double> unif(-maxNorm, maxNorm);
    default_random_engine re(r());

    for (int i = 0; i < dim; i++) {
        // generate a random double between -maxNorm and maxNorm
        x[i] = unif(re);
    }
    return x;
}

vector<double> printDebug(const Parameters &params, const Ciphertext &ciphertext, const vector<double> &expectedPT, const Decryptor &decryptor, const Encoder encoder) {

    vector<double> actualPT = decode(encoder, decryptNew(decryptor, ciphertext), logSlots(params));

    cout << "Level: " << level(ciphertext) << " (logQ = " << logQLvl(params, level(ciphertext)) << ")" << endl;
    cout << "Scale: 2^" << log2(scale(ciphertext)) << endl;

    cout << "Actual Result:   " << setprecision(3) << actualPT[0] << " " << actualPT[1] << " " << actualPT[2] << " " << actualPT[3] << endl;
    cout << "Expected Result: " << setprecision(3) << expectedPT[0] << " " << expectedPT[1] << " " << expectedPT[2] << " " << expectedPT[3] << endl;

    string precStats = precisionStats(params, expectedPT, actualPT);

    cout << precStats << endl;

    return actualPT;
}

int main() {
    Parameters params = getParams(BootstrapParams0);
    BootstrappingParameters btpParams = getBootstrappingParams(BootstrapParams_Set2);

    cout << "CKKS parameters: logN = " << logN(params) << ", logSlots = " << logSlots(params)
         << ", h = " << bootstrap_h(btpParams) << ", logQP = " << logQP(params)
         << ", levels = " << qiCount(params) << ", scale = 2^" << log2(scale(params))
         << ", sigma = " << sigma(params) << endl;

    KeyGenerator kgen = newKeyGenerator(params);
    struct KeyPairHandle kp = genKeyPairSparse(kgen, bootstrap_h(btpParams));

    Encoder encoder = newEncoder(params);
    Decryptor decryptor = newDecryptor(params, kp.sk);
    Encryptor encryptor = newEncryptorFromPk(params, kp.pk);
    Evaluator evaluator = newEvaluator(params);

    EvaluationKey rlk = genRelinKey(kgen, kp.sk);

    cout << "Generating bootstrapping keys..." << endl;
    BootstrappingKey btpKey = genBootstrappingKey(kgen, logSlots(params), btpParams, kp.sk);
    Bootstrapper btp = newBootstrapper(params, btpParams, btpKey);
    cout << "Done" << endl;

    uint64_t num_slots = numSlots(params);
    vector<double> values = randomVector(num_slots, 1);

    Plaintext plaintext = encodeNTTAtLvlNew(params, encoder, values, maxLevel(params), scale(params));

    for (int i = 0; i < values.size(); i++) {
        values[i] *= values[i];
    }

    Ciphertext ciphertext1 = encryptNew(encryptor, plaintext);

    cout << "Scale after encryption: " << log2(scale(ciphertext1)) << endl;
    mulRelin(evaluator, ciphertext1, ciphertext1, rlk, ciphertext1);
    cout << "Scale after mul: " << log2(scale(ciphertext1)) << endl;
    rescaleMany(evaluator, ciphertext1, 1, ciphertext1);
    cout << "Scale after rescale: " << log2(scale(ciphertext1)) << endl;

    cout << "Precision of values vs. ciphertext" << endl;
    vector<double> valuesTest1 = printDebug(params, ciphertext1, values, decryptor, encoder);

    cout << "Bootstrapping..." << endl;
    Ciphertext ciphertext2 = bootstrap(btp, ciphertext1);
    cout << "Scale after bootstrap: " << log2(scale(ciphertext2)) << endl;
    cout << "Done" << endl;

    cout << "Precision of ciphertext vs. Bootstrapp(ciphertext)" << endl;
    printDebug(params, ciphertext2, valuesTest1, decryptor, encoder);

    return 0;
}
