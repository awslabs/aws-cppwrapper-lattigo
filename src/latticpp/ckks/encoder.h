// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshall/gohandle.h"
#include "cgo/encoder.h"
#include <vector>

namespace latticpp {

    GoHandle<Encoder> newEncoder(const GoHandle<Parameters> &params);

    GoHandle<Plaintext> encodeNew(const GoHandle<Encoder> &encoder, std::vector<double> &values);

    std::vector<double> decode(const GoHandle<Encoder> &encoder, const GoHandle<Plaintext> &pt, uint64_t logSlots);
}  // namespace latticpp
