// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "plaintext.h"

namespace latticpp {

    Plaintext newPlaintext(const Parameters &params, uint64_t level){
        return Plaintext(lattigo_newPlaintext(params.getRawHandle(), level));
    }

    Plaintext newPlaintext(const Parameters &params, const Poly &poly, uint64_t level){
        return Plaintext(lattigo_newPlaintextFromPoly(params.getRawHandle(), poly.getRawHandle(), level));
    }

    double scale(const Plaintext &ct) {
        return lattigo_plaintextGetScale(ct.getRawHandle());
    }

    void setScale(Plaintext &ct, double scale) {
        lattigo_plaintextSetScale(ct.getRawHandle(), scale);
    }    

    Poly poly(const Plaintext &plaintext) {
        return Poly(lattigo_getPlaintextPoly(plaintext.getRawHandle()));
    }

}  // namespace latticpp