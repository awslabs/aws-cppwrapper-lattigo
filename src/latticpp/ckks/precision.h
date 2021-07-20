// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/precision.h"
#include <vector>
#include <string>

namespace latticpp {

    std::string precisionStats(const Parameters &params, const Encoder &encoder, const std::vector<double> &expectedValues, const std::vector<double> &actualValues);
}  // namespace latticpp
