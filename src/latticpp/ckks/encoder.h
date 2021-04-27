// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/encoder.h"
#include <vector>

namespace latticpp {

    Encoder newEncoder(const Parameters &params);

    Plaintext encodeNew(const Encoder &encoder, std::vector<double> &values);

    std::vector<double> decode(const Encoder &encoder, const Plaintext &pt, uint64_t logSlots);
}  // namespace latticpp
