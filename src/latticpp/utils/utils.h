
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "cgo/utils.h"
#include "latticpp/marshal/gohandle.h"
#include <vector>

namespace latticpp {

    PRNG newPRNG();

    PRNG newKeyedPRNG(std::vector<char> &key);
} // namespace latticpp
