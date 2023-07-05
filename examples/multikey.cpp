// SPDX-License-Identifier: Apache-2.0

#include "latticpp/latticpp.h"

#include <algorithm>
#include <cmath>
#include <iomanip>
#include <random>
#include <vector>

using namespace std;
using namespace latticpp;

struct TestContext {
  int numParties;

  Parameters params;
  Ring ringQ;
  RingQP ringQP;

  PRNG prng;

  Encoder encoder;
  Evaluator evaluator;

  vector<SecretKey> sk0Shards, sk1Shards;

  SecretKey sk0, sk1;

  PublicKey pk0, pk1;

  Encryptor encryptorPk0;
  Decryptor decryptorSk0, decryptorSk1;
};

struct PartyCKG {
  CKGProtocol ckgProtocol;
  SecretKey s;
  CKGShare s1;
};

struct PartyRKG {
  RKGProtocol rkgProtocol;
  SecretKey ephSk, sk;
  RKGShare share1, share2;
};

struct PartyCKS {
  CKSProtocol cksProtocol;
  SecretKey s0, s1;
  CKSShare share;
};

struct PartyRTG {
  RTGProtocol rtgProtocol;
  SecretKey s;
  RTGShare share;
};

Parameters generateParamsForTest() {
  BootstrappingParameters btpParams =
      getBootstrappingParams(N15QP880H16384H32);
  Parameters params = genParams(N15QP880H16384H32);

  cout << "CKKS parameters: logN = " << logN(params)
       << ", logSlots = " << logSlots(params)
       << ", h = " << ephemeralSecretWeight(btpParams)
       << ", logQP = " << logQP(params) << ", levels = " << qiCount(params)
       << ", scale = 2^" << log2(scale(params)) << ", sigma = " << sigma(params)
       << endl;

  return params;
}

TestContext generateTestContextForTest(const Parameters &params,
                                       int numParties) {
  TestContext res;

  res.numParties = numParties;
  res.params = params;
  res.ringQ = ringQ(params);
  res.ringQP = ringQP(params);

  res.prng = newPRNG();
  res.encoder = newEncoder(params);
  res.evaluator = newEvaluator(params, makeEmptyEvaluationKey());

  KeyGenerator kgen = newKeyGenerator(params);

  res.sk0Shards.resize(numParties);
  res.sk1Shards.resize(numParties);
  PolyQP tmp0 = newPolyQP(res.ringQP);
  PolyQP tmp1 = newPolyQP(res.ringQP);

  uint64_t levelQ = qiCount(params)-1;
  uint64_t levelP = piCount(params)-1;

  for (int j = 0; j < numParties; j++) {
    res.sk0Shards.at(j) = genSecretKey(kgen);
    res.sk1Shards.at(j) = genSecretKey(kgen);
    addLvl(res.ringQP, levelQ, levelP, tmp0, polyQP(res.sk0Shards.at(j)), tmp0);
    addLvl(res.ringQP, levelQ, levelP, tmp1, polyQP(res.sk1Shards.at(j)), tmp1);
  }

  res.sk0 = newSecretKey(params);
  res.sk1 = newSecretKey(params);
  PolyQP sk0Value = polyQP(res.sk0);
  PolyQP sk1Value = polyQP(res.sk1);
  copy(sk0Value, tmp0);
  copy(sk1Value, tmp1);

  res.pk0 = genPublicKey(kgen, res.sk0);
  res.pk1 = genPublicKey(kgen, res.sk1);

  res.encryptorPk0 = newEncryptor(params, res.pk0);
  res.decryptorSk0 = newDecryptor(params, res.sk0);
  res.decryptorSk1 = newDecryptor(params, res.sk1);

  return res;
}

void newTestVectors(const TestContext &testContext, const Encryptor &encryptor,
                    vector<double> &values, Plaintext &plaintext,
                    Ciphertext &ciphertext) {
  const Parameters &params = testContext.params;
  uint64_t ls = logSlots(params);
  values.resize(1 << ls);

  uniform_real_distribution<double> unif(-10, 10);
  default_random_engine re;

  for (size_t i = 0; i < values.size(); i++) {
    values.at(i) = unif(re);
  }

  plaintext = encodeNew(testContext.encoder, values, maxLevel(params), scale(params));
  ciphertext = encryptNew(encryptor, plaintext);
}

void verifyTestVectors(const TestContext &testContext,
                       const Decryptor &decryptor,
                       const vector<double> &expectedPT,
                       const Ciphertext &ciphertext) {
  vector<double> actualPT =
      decode(testContext.encoder, decryptNew(decryptor, ciphertext),
             logSlots(testContext.params));

  double err = 0;
  for (size_t i = 0; i < actualPT.size() && i < expectedPT.size(); i++) {
    err += abs(actualPT.at(i) - expectedPT.at(i));
  }

  cout << "Average error= " << err / min(actualPT.size(), expectedPT.size())
       << endl;
}

void testPublicKeyGen(const TestContext &testContext) {
  const Decryptor &decryptorSk0 = testContext.decryptorSk0;
  const vector<SecretKey> &sk0Shards = testContext.sk0Shards;
  const Parameters &params = testContext.params;

  vector<PartyCKG> ckgParties(testContext.numParties);
  for (size_t i = 0; i < ckgParties.size(); i++) {
    PartyCKG &p = ckgParties.at(i);
    p.ckgProtocol = newCKGProtocol(params);
    p.s = sk0Shards.at(i);
    p.s1 = ckgAllocateShare(p.ckgProtocol);
  }

  PartyCKG &p0 = ckgParties.at(0);
  CKGCRP crp = ckgSampleCRP(p0.ckgProtocol, testContext.prng);

  for (size_t i = 0; i < ckgParties.size(); i++) {
    PartyCKG &p = ckgParties.at(i);
    ckgGenShare(p.ckgProtocol, p.s, crp, p.s1);
    if (i > 0) {
      ckgAggregateShares(p0.ckgProtocol, p.s1, p0.s1, p0.s1);
    }
  }
  PublicKey pk = newPublicKey(params);
  ckgGenPublicKey(p0.ckgProtocol, p0.s1, crp, pk);

  Encryptor encryptorTester = newEncryptor(params, pk);
  vector<double> values;
  Plaintext plaintext;
  Ciphertext ciphertext;
  newTestVectors(testContext, encryptorTester, values, plaintext, ciphertext);

  verifyTestVectors(testContext, decryptorSk0, values, ciphertext);
}

void testRelinKeyGen(const TestContext &testContext) {
  const Encryptor &encryptorPk0 = testContext.encryptorPk0;
  const Decryptor &decryptorSk0 = testContext.decryptorSk0;
  const vector<SecretKey> &sk0Shards = testContext.sk0Shards;
  const Parameters &params = testContext.params;

  vector<PartyRKG> rkgParties(testContext.numParties);
  for (size_t i = 0; i < rkgParties.size(); i++) {
    PartyRKG &p = rkgParties.at(i);
    p.rkgProtocol = newRKGProtocol(params);
    p.sk = sk0Shards.at(i);
    p.ephSk = newSecretKey(params);
    p.share1 = newRKGShare();
    p.share2 = newRKGShare();
    rkgAllocateShare(p.rkgProtocol, p.ephSk, p.share1, p.share2);
  }

  PartyRKG &p0 = rkgParties.at(0);
  RKGCRP crp = rkgSampleCRP(p0.rkgProtocol, testContext.prng);

  for (size_t i = 0; i < rkgParties.size(); i++) {
    PartyRKG &p = rkgParties.at(i);
    rkgGenShareRoundOne(p.rkgProtocol, p.sk, crp, p.ephSk, p.share1);
    if (i > 0) {
      rkgAggregateShares(p0.rkgProtocol, p.share1, p0.share1, p0.share1);
    }
  }

  for (size_t i = 0; i < rkgParties.size(); i++) {
    PartyRKG &p = rkgParties.at(i);
    rkgGenShareRoundTwo(p.rkgProtocol, p.ephSk, p.sk, p0.share1, p.share2);
    if (i > 0) {
      rkgAggregateShares(p0.rkgProtocol, p.share2, p0.share2, p0.share2);
    }
  }

  RelinearizationKey rlk = newRelinearizationKey(params);
  rkgGenRelinearizationKey(p0.rkgProtocol, p0.share1, p0.share2, rlk);

  vector<double> values;
  Plaintext plaintext;
  Ciphertext ciphertext;
  newTestVectors(testContext, encryptorPk0, values, plaintext, ciphertext);

  for (double &val : values) {
    val *= val;
  }

  EvaluationKey evalKey = makeEmptyEvaluationKey();
  setRelinKeyForEvaluationKey(evalKey, rlk);
  Evaluator evaluator = evaluatorWithKey(testContext.evaluator, evalKey);

  mulRelin(evaluator, ciphertext, ciphertext, ciphertext);
  rescale(evaluator, ciphertext, scale(params), ciphertext);

  verifyTestVectors(testContext, decryptorSk0, values, ciphertext);
}

void testKeySwitching(const TestContext &testContext) {
  const Encryptor &encryptorPk0 = testContext.encryptorPk0;
  const Decryptor &decryptorSk1 = testContext.decryptorSk1;
  const vector<SecretKey> &sk0Shards = testContext.sk0Shards;
  const vector<SecretKey> &sk1Shards = testContext.sk1Shards;
  const Parameters &params = testContext.params;

  vector<double> values;
  Plaintext plaintext;
  Ciphertext origCiphertext;
  newTestVectors(testContext, encryptorPk0, values, plaintext, origCiphertext);

  double sigmaSmudging = 3.2;

  vector<uint64_t> dropBys = {0, level(origCiphertext)};
  for (uint64_t dropBy : dropBys) {
    Ciphertext ciphertext = copyNew(origCiphertext);
    dropLevel(testContext.evaluator, ciphertext, dropBy);

    vector<PartyCKS> cksParties(testContext.numParties);
    for (size_t i = 0; i < cksParties.size(); i++) {
      PartyCKS &p = cksParties.at(i);
      p.cksProtocol = newCKSProtocol(params, sigmaSmudging);
      p.s0 = sk0Shards.at(i);
      p.s1 = sk1Shards.at(i);
      p.share = cksAllocateShare(p.cksProtocol, level(ciphertext));
    }

    PartyCKS &p0 = cksParties.at(0);
    for (size_t i = 0; i < cksParties.size(); i++) {
      PartyCKS &p = cksParties.at(i);
      cksGenShare(p.cksProtocol, p.s0, p.s1, ciphertext, p.share);
      if (i > 0) {
        cksAggregateShares(p0.cksProtocol, p.share, p0.share, p0.share);
      }
    }

    Ciphertext ksCiphertext = newCiphertext(params, 1, level(ciphertext));
    cksKeySwitch(p0.cksProtocol, ciphertext, p0.share, ksCiphertext);

    verifyTestVectors(testContext, decryptorSk1, values, ksCiphertext);
  }
}

void testRotKeyGenCols(const TestContext &testContext) {
  const RingQP &ringQP = testContext.ringQP;
  const Encryptor &encryptorPk0 = testContext.encryptorPk0;
  const Decryptor &decryptorSk0 = testContext.decryptorSk0;
  const vector<SecretKey> &sk0Shards = testContext.sk0Shards;
  const Parameters &params = testContext.params;

  vector<PartyRTG> rtgParties(testContext.numParties);
  for (size_t i = 0; i < rtgParties.size(); i++) {
    PartyRTG &p = rtgParties.at(i);
    p.rtgProtocol = newRTGProtocol(params);
    p.s = sk0Shards.at(i);
    p.share = rtgAllocateShare(p.rtgProtocol);
  }

  vector<double> values;
  Plaintext plaintext;
  Ciphertext ciphertext;
  newTestVectors(testContext, encryptorPk0, values, plaintext, ciphertext);

  Ciphertext receiver = newCiphertext(params, degree(ciphertext), level(ciphertext));

  vector<uint64_t> galEls = galoisElementsForRowInnerSum(params);
  RotationKeys rotKeySet = newRotationKeys(params, galEls);

  PartyRTG &p0 = rtgParties.at(0);
  RTGCRP crp = rtgSampleCRP(p0.rtgProtocol, testContext.prng);

  for (uint64_t galEl : galEls) {
    for (size_t i = 0; i < rtgParties.size(); i++) {
      PartyRTG &p = rtgParties.at(i);
      rtgGenShare(p.rtgProtocol, p.s, galEl, crp, p.share);
      if (i > 0) {
        rtgAggregateShares(p.rtgProtocol, p.share, p0.share, p0.share);
      }
    }

    SwitchingKey rotKey = getRotationKey(rotKeySet, galEl);
    rtgGenRotationKey(p0.rtgProtocol, p0.share, crp, rotKey);
  }

  EvaluationKey evalKey = makeEmptyEvaluationKey();
  setRotKeysForEvaluationKey(evalKey, rotKeySet);
  Evaluator evaluator = evaluatorWithKey(testContext.evaluator, evalKey);

  for (int k = 1; k < 1 << logSlots(params); k <<= 1) {
    rotate(evaluator, ciphertext, k, receiver);

    vector<double> expected(values);
    rotate(expected.begin(), expected.begin() + k, expected.end());

    verifyTestVectors(testContext, decryptorSk0, expected, receiver);
  }
}

int main() {
  int numParties = 10;

  Parameters params = generateParamsForTest();

  TestContext testContext = generateTestContextForTest(params, numParties);

  testPublicKeyGen(testContext);
  testRelinKeyGen(testContext);
  testKeySwitching(testContext);
  testRotKeyGenCols(testContext);

  return 0;
}
