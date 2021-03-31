// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "encryptor.h"

namespace latticpp {

    GoHandle<Encryptor> newEncryptorFromPk(const GoHandle<Parameters> &params, const GoHandle<PublicKey> &pk) {
        return GoHandle<Encryptor>(lattigo_newEncryptorFromPk(params.getRawHandle(), pk.getRawHandle()));
    }

    GoHandle<Ciphertext> encryptNew(const GoHandle<Encryptor> &encryptor, const GoHandle<Plaintext> &pt) {
        return GoHandle<Ciphertext>(lattigo_encryptNew(encryptor.getRawHandle(), pt.getRawHandle()));
    }
}  // namespace latticpp
