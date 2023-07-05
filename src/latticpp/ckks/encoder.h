// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/encoder.h"
#include <vector>

namespace latticpp {

    Encoder newEncoder(const Parameters &params);

    void encode(const Encoder &encoder, const std::vector<double> &values, Plaintext &outPt);

    Plaintext encodeNew(const Encoder &encoder, const std::vector<double> &values, uint64_t level, double scale);

    std::vector<double> decode(const Encoder &encoder, const Plaintext &pt, uint64_t logSlots);
}  // namespace latticpp
