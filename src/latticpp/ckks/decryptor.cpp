// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "decryptor.h"

namespace latticpp {

    GoHandle<Decryptor> newDecryptor(const GoHandle<Parameters> &params, const GoHandle<SecretKey> &sk) {
        return GoHandle<Decryptor>(lattigo_newDecryptor(params.getRawHandle(), sk.getRawHandle()));
    }

    GoHandle<Plaintext> decryptNew(const GoHandle<Decryptor> &decryptor, const GoHandle<Ciphertext> &ct) {
        return GoHandle<Plaintext>(lattigo_decryptNew(decryptor.getRawHandle(), ct.getRawHandle()));
    }
}  // namespace latticpp
