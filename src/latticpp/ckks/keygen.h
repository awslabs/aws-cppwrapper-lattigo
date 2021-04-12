// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshall/gohandle.h"
#include "cgo/keygen.h"

namespace latticpp {

    struct KeyPairHandle {
        SecretKey sk;
        PublicKey pk;
    };

    KeyGenerator newKeyGenerator(const Parameters &params);

    KeyPairHandle genKeyPair(const KeyGenerator &keygen);

    RelinKey genRelinKey(const KeyGenerator &keygen, const SecretKey &sk);

}  // namespace latticpp
