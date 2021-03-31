// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "keygen.h"

namespace latticpp {

    GoHandle<KeyGenerator> newKeyGenerator(const GoHandle<Parameters> &params) {
        return GoHandle<KeyGenerator>(lattigo_newKeyGenerator(params.getRawHandle()));
    }

    KeyPairHandle genKeyPair(const GoHandle<KeyGenerator> &keygen) {
        Lattigo_KeyPairHandle kp = lattigo_genKeyPair(keygen.getRawHandle());
        return KeyPairHandle { GoHandle<SecretKey>(kp.sk), GoHandle<PublicKey>(kp.pk) };
    }


    GoHandle<RelinKey> genRelinKey(const GoHandle<KeyGenerator> &keygen, const GoHandle<SecretKey> &sk) {
        return GoHandle<RelinKey>(lattigo_genRelinKey(keygen.getRawHandle(), sk.getRawHandle()));
    }
}  // namespace latticpp



