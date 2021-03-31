// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshall/gohandle.h"
#include "cgo/keygen.h"

namespace latticpp {

    struct KeyPairHandle {
        GoHandle<SecretKey> sk;
        GoHandle<PublicKey> pk;
    };

    GoHandle<KeyGenerator> newKeyGenerator(const GoHandle<Parameters> &params);

    KeyPairHandle genKeyPair(const GoHandle<KeyGenerator> &keygen);

    GoHandle<RelinKey> genRelinKey(const GoHandle<KeyGenerator> &keygen, const GoHandle<SecretKey> &sk);

}  // namespace latticpp
