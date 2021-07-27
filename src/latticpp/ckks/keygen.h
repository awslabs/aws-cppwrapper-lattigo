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

    KeyPairHandle genKeyPair(const KeyGenerator &keygen);

    KeyPairHandle genKeyPairSparse(const KeyGenerator &keygen, uint64_t hw);

    RelinearizationKey genRelinKey(const KeyGenerator &keygen, const SecretKey &sk);

    RotationKeys genRotationKeysForRotations(const KeyGenerator &keygen, const SecretKey &sk, std::vector<int> shifts);

    EvaluationKey makeEvaluationKey(const RelinearizationKey &relinKey, const RotationKeys &rotKeys);

    BootstrappingKey genBootstrappingKey(const KeyGenerator &keygen, const Parameters &params, const BootstrappingParameters &bootParams, const SecretKey &sk, const RelinearizationKey &relinKey, const RotationKeys &rotKeys);

    BootstrappingKey makeBootstrappingKey(const RelinearizationKey &relinKey, const RotationKeys &rotKeys);

}  // namespace latticpp
