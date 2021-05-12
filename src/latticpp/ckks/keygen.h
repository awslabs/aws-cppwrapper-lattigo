// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/keygen.h"

namespace latticpp {

    struct KeyPairHandle {
        SecretKey sk;
        PublicKey pk;
    };

    KeyGenerator newKeyGenerator(const Parameters &params);

    KeyPairHandle genKeyPair(const KeyGenerator &keygen);

    EvaluationKey genRelinKey(const KeyGenerator &keygen, const SecretKey &sk);

    RotationKeys genRotationKeysPow2(const KeyGenerator &keygen, const SecretKey &sk);
}  // namespace latticpp
