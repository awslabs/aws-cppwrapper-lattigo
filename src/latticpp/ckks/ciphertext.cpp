// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "ciphertext.h"

namespace latticpp {

    uint64_t level(const Ciphertext &ct) {
        return lattigo_level(ct.getRawHandle());
    }

    double scale(const Ciphertext &ct) {
        return lattigo_ciphertextScale(ct.getRawHandle());
    }

    Ciphertext copyNew(const Ciphertext &ct) {
        return lattigo_copyNew(ct.getRawHandle());
    }

    Ciphertext newCiphertext(const Parameters &params, uint64_t degree, uint64_t level, double scale) {
        return Ciphertext(lattigo_newCiphertext(params.getRawHandle(), degree, level, scale));
    }
}  // namespace latticpp
