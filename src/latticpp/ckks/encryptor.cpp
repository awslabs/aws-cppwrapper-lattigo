// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "encryptor.h"

namespace latticpp {

    Encryptor newEncryptor(const Parameters &params, const SecretKey &sk) {
        return Encryptor(lattigo_newEncryptorFromSk(params.getRawHandle(), sk.getRawHandle()));
    }

    Encryptor newEncryptor(const Parameters &params, const PublicKey &pk) {
        return Encryptor(lattigo_newEncryptorFromPk(params.getRawHandle(), pk.getRawHandle()));
    }

    Ciphertext encryptNew(const Encryptor &encryptor, const Plaintext &pt) {
        return Ciphertext(lattigo_encryptNew(encryptor.getRawHandle(), pt.getRawHandle()));
    }

    void encryptZeroQP(const Parameters &params, const SecretKey &sk, CiphertextQP &ctxQP){
        lattigo_encryptZeroQP(params.getRawHandle(), sk.getRawHandle(), ctxQP.getRawHandle());
    }

}  // namespace latticpp
