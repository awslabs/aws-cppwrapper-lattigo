// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "encoder.h"
#include <cmath>
#include <stdexcept>
#include <iostream>

using namespace std;

namespace latticpp {

    Encoder newEncoder(const Parameters &params) {
        return Encoder(lattigo_newEncoder(params.getRawHandle()));
    }

    Plaintext encodeNew(const Encoder &encoder, vector<double> &values) {

        int len = values.size();
        int logLen = log2(len);

        if (len != pow(2, logLen)) {
            throw invalid_argument("Invalid input length for encodeNew");
        }

        return Plaintext(lattigo_encodeNew(encoder.getRawHandle(), values.data(), logLen));
    }

    vector<double> decode(const Encoder &encoder, const Plaintext &pt, uint64_t logSlots) {
        vector<double> coeffs(((uint64_t)1) << logSlots);
        lattigo_decode(encoder.getRawHandle(), pt.getRawHandle(), logSlots, coeffs.data());
        return coeffs;
    }
}  // namespace latticpp
