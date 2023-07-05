// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "latticpp/latticpp.h"

#include <cmath>
#include <chrono>
#include <iomanip>
#include <iostream>
#include <map>
#include <vector>

using namespace std;
using namespace std::chrono;
using namespace latticpp;

vector<double> printDebug(const Parameters &params, const Ciphertext &ciphertext, const vector<double> &valuesWant, const Decryptor &decryptor, const Encoder &encoder) {
  vector<double> valuesTest = decode(encoder, decryptNew(decryptor, ciphertext), logSlots(params));

  cout << endl;
  cout << "Level: " << level(ciphertext) << " (logQ = " << logQLvl(params, level(ciphertext)) << ")" << endl;
  cout << "Scale: 2^" << log2(scale(ciphertext)) << endl;
  cout << "ValuesTest: " << valuesTest[0] << ", " << valuesTest[1] << ", " << valuesTest[2] << ", " << valuesTest[3] << endl;
  cout << "ValuesWant: " << valuesWant[0] << ", " << valuesWant[1] << ", " << valuesWant[2] << ", " << valuesWant[3] << endl;
  cout << endl;

  string precStats = precisionStats(params, encoder, valuesWant, valuesTest);
  cout << precStats << endl;

  return valuesTest;
}

int main() {
  high_resolution_clock::time_point start;

  // Schemes parameters are created from scratch
  uint64_t logNVal = 14;
  vector<uint8_t> logQi = {55, 40, 40, 40, 40, 40, 40, 40};
  vector<uint8_t> logPi = {45, 45};
  uint8_t logScale = 40;
  Parameters params = newParametersFromLogModuli(logNVal, logQi, logPi, logScale);

  cout << endl;
  cout << "=========================================" << endl;
  cout << "         INSTANTIATING SCHEME            " << endl;
  cout << "=========================================" << endl;
  cout << endl;

  start = high_resolution_clock::now();

  KeyGenerator kgen = newKeyGenerator(params);

  SecretKey sk = genSecretKey(kgen);

  RelinearizationKey rlk = genRelinKey(kgen, sk);

  Encryptor encryptor = newEncryptor(params, sk);

  Decryptor decryptor = newDecryptor(params, sk);

  Encoder encoder = newEncoder(params);

  Evaluator evaluator = newEvaluator(params, makeEvaluationKey(rlk));

  cout << "Done in " << duration_cast<seconds>(high_resolution_clock::now() - start).count() << " seconds." << endl;

  cout << endl;
  cout << "CKKS parameters: logN = " << logN(params) << ", logSlots = " << logSlots(params) << ", logQP = " << logQP(params) << ", levels = " << maxLevel(params) + 1 << ", scale= " << scale(params) << ", sigma = " << sigma(params) << endl;

  cout << endl;
  cout << "=========================================" << endl;
  cout << "           PLAINTEXT CREATION            " << endl;
  cout << "=========================================" << endl;
  cout << endl;

  start = high_resolution_clock::now();

  double r = 16;

  double pi = 3.141592653589793;

  uint64_t slots = numSlots(params);

  vector<double> values(slots, 2*pi);

  Plaintext plaintext = newPlaintext(params, maxLevel(params));
  setScale(plaintext, scale(plaintext) / r);
  encode(encoder, values, plaintext);

  cout << "Done in " << duration_cast<seconds>(high_resolution_clock::now() - start).count() << " seconds." << endl;

  cout << endl;
  cout << "=========================================" << endl;
  cout << "              ENCRYPTION                 " << endl;
  cout << "=========================================" << endl;
  cout << endl;

  start = high_resolution_clock::now();

  Ciphertext ciphertext = encryptNew(encryptor, plaintext);

  cout << "Done in " << duration_cast<seconds>(high_resolution_clock::now() - start).count() << " seconds." << endl;

  printDebug(params, ciphertext, values, decryptor, encoder);

  cout << endl;
  cout << "===============================================" << endl;
  cout << "        EVALUATION OF i*x on " << slots << " values" << endl;
  cout << "===============================================" << endl;
  cout << endl;

  start = high_resolution_clock::now();

  multByConst(evaluator, ciphertext, 2, ciphertext);

  cout << "Done in " << duration_cast<seconds>(high_resolution_clock::now() - start).count() << " seconds." << endl;

  for (int i=0; i < values.size(); i++) {
    values[i] *= 2;
  }

  printDebug(params, ciphertext, values, decryptor, encoder);

  cout << endl;
  cout << "===============================================" << endl;
  cout << "       EVALUATION of x/r on " << slots << " values" << endl;
  cout << "===============================================" << endl;
  cout << endl;

  start = high_resolution_clock::now();

  setScale(ciphertext, scale(ciphertext) * r);

  cout << "Done in " << duration_cast<seconds>(high_resolution_clock::now() - start).count() << " seconds." << endl;

  for (int i=0; i < values.size(); i++) {
    values[i] /= r;
  }

  printDebug(params, ciphertext, values, decryptor, encoder);

  cout << endl;
  cout << "===============================================" << endl;
  cout << "         DECRYPTION & DECODING           " << endl;
  cout << "===============================================" << endl;
  cout << endl;

  start = high_resolution_clock::now();

  decode(encoder, decryptNew(decryptor, ciphertext), logSlots(params));

  cout << "Done in " << duration_cast<seconds>(high_resolution_clock::now() - start).count() << " seconds." << endl;

  printDebug(params, ciphertext, values, decryptor, encoder);

  return 0;
}
