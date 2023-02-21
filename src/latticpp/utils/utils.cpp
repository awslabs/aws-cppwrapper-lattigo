// SPDX-License-Identifier: Apache-2.0

#include "utils.h"

using namespace std;

namespace latticpp {

    PRNG newPRNG() { return PRNG(lattigo_newPRNG()); }

    PRNG newKeyedPRNG(vector<char> &key) {
        return PRNG(lattigo_newKeyedPRNG(key.data(), key.size()));
    }
} // namespace latticpp
