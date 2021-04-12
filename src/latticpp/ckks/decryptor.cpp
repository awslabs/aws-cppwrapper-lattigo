// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "decryptor.h"

namespace latticpp {

    Decryptor newDecryptor(const Parameters &params, const SecretKey &sk) {
        return Decryptor(lattigo_newDecryptor(params.getRawHandle(), sk.getRawHandle()));
    }

    Plaintext decryptNew(const Decryptor &decryptor, const Ciphertext &ct) {
        return Plaintext(lattigo_decryptNew(decryptor.getRawHandle(), ct.getRawHandle()));
    }
}  // namespace latticpp
