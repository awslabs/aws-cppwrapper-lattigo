// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "keygen.h"

namespace latticpp {

    KeyGenerator newKeyGenerator(const Parameters &params) {
        return KeyGenerator(lattigo_newKeyGenerator(params.getRawHandle()));
    }

    KeyPairHandle genKeyPair(const KeyGenerator &keygen) {
        Lattigo_KeyPairHandle kp = lattigo_genKeyPair(keygen.getRawHandle());
        return KeyPairHandle { SecretKey(kp.sk), PublicKey(kp.pk) };
    }

    EvaluationKey genRelinKey(const KeyGenerator &keygen, const SecretKey &sk) {
        return EvaluationKey(lattigo_genRelinKey(keygen.getRawHandle(), sk.getRawHandle()));
    }

    RotationKeys genRotationKeysPow2(const KeyGenerator &keygen, const SecretKey &sk) {
        return RotationKeys(lattigo_genRotationKeysPow2(keygen.getRawHandle(), sk.getRawHandle()));
    }
}  // namespace latticpp
