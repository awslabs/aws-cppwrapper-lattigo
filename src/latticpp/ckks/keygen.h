// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/keygen.h"
#include <vector>

namespace latticpp {

    struct KeyPairHandle {
        SecretKey sk;
        PublicKey pk;
    };

    KeyGenerator newKeyGenerator(const Parameters &params);

    SwitchingKey getRotationKey(const RotationKeys &rtks, uint64_t galEl);

    void setRotationKey(const RotationKeys &rotKeys, const SwitchingKey &rotKey, uint64_t galEl);

    uint64_t rotationKeyExist(const RotationKeys &rotationKeys, uint64_t galEl);

    std::vector<uint64_t> getGaloisElementsOfRotationKeys(const RotationKeys &rotationKeys);

    SwitchingKey copyNew(const SwitchingKey &rotKey);

    uint64_t numOfDecomp(const SwitchingKey &rtk);

    uint64_t galoisElementForColumnRotationBy(const Parameters &params,
                                              uint64_t rotationStep);

    uint64_t rotationKeyIsCorrect(const SwitchingKey &rtk, uint64_t galEl,
                                  const SecretKey &sk, const Parameters &params,
                                  uint64_t log2Bound);

    CiphertextQP getCiphertextQP(const SwitchingKey& rtk, uint64_t i, uint64_t j);

    void setCiphertextQP(const SwitchingKey& rtk, const CiphertextQP& ctQP, uint64_t i, uint64_t j);

    SecretKey newSecretKey(const Parameters &params);

    SecretKey copyNewSecretKey(const SecretKey &sk);

    PolyQP polyQP(const SecretKey &sk);

    PublicKey newPublicKey(const Parameters &params);

    RelinearizationKey newRelinearizationKey(const Parameters &params);

    RotationKeys newRotationKeys(const Parameters &params, std::vector<uint64_t> galoisElements);

    SecretKey genSecretKey(const KeyGenerator &keygen);

    PublicKey genPublicKey(const KeyGenerator &keygen, const SecretKey &sk);

    KeyPairHandle genKeyPair(const KeyGenerator &keygen);

    KeyPairHandle genKeyPairSparse(const KeyGenerator &keygen, uint64_t hw);

    RelinearizationKey genRelinKey(const KeyGenerator &keygen, const SecretKey &sk);

    RotationKeys genRotationKeysForRotations(const KeyGenerator &keygen, const SecretKey &sk, std::vector<int> shifts);

    EvaluationKey makeEvaluationKey(const RelinearizationKey &relinKey);

    EvaluationKey makeEvaluationKey(const RelinearizationKey &relinKey, const RotationKeys &rotKeys);

    EvaluationKey makeEmptyEvaluationKey();

    void setRelinKeyForEvaluationKey(const EvaluationKey &evalKey,
                                     const RelinearizationKey &relinKey);

    void setRotKeysForEvaluationKey(const EvaluationKey &evalKey,
                                    const RotationKeys &rotKeys);

    BootstrappingKey genBootstrappingKey(const KeyGenerator &keygen, const Parameters &params, const BootstrappingParameters &bootParams, const SecretKey &sk, const RelinearizationKey &relinKey, const RotationKeys &rotKeys);

    PolyQP polyQP(const SecretKey &sk);

    SwitchingKey newSwitchingKey(const Parameters &params, uint64_t levelQ, uint64_t levelP);

    void switchKeys(const Evaluator &eval, const Ciphertext &ctxIn, const SwitchingKey &swk, const Ciphertext &ctxOut);
}  // namespace latticpp