// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "ciphertext.h"

namespace latticpp {

    uint64_t level(const GoHandle<Ciphertext> &ct) {
        return lattigo_level(ct.getRawHandle());
    }

    double scale(const GoHandle<Ciphertext> &ct) {
        return lattigo_ciphertextScale(ct.getRawHandle());
    }

    GoHandle<Ciphertext> copyNew(const GoHandle<Ciphertext> &ct) {
        return lattigo_copyNew(ct.getRawHandle());
    }

    GoHandle<Ciphertext> newCiphertext(const GoHandle<Parameters> &params, uint64_t degree, uint64_t level, double scale) {
        return GoHandle<Ciphertext>(lattigo_newCiphertext(params.getRawHandle(), degree, level, scale));
    }
}  // namespace latticpp
