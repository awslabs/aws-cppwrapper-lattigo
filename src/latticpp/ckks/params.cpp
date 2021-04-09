// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "params.h"

namespace latticpp {

    Parameters getParams(HEParams paramId) {
        return Parameters(lattigo_getParams(paramId));
    }

    uint64_t numSlots(const Parameters &params) {
        return lattigo_numSlots(params.getRawHandle());
    }

    uint64_t logN(const Parameters &params) {
        return lattigo_logN(params.getRawHandle());
    }

    uint64_t logQP(const Parameters &params) {
        return lattigo_logQP(params.getRawHandle());
    }

    uint64_t maxLevel(const Parameters &params) {
        return lattigo_maxLevel(params.getRawHandle());
    }

    double scale(const Parameters &params) {
        return lattigo_paramsScale(params.getRawHandle());
    }

    double sigma(const Parameters &params) {
        return lattigo_sigma(params.getRawHandle());
    }

    uint64_t getQi(const Parameters &params, uint64_t i) {
        return lattigo_getQi(params.getRawHandle(), i);
    }

    uint64_t logQLvl(const Parameters &params, uint64_t lvl) {
        return lattigo_logQLvl(params.getRawHandle(), lvl);
    }

    uint64_t logSlots(const Parameters &params) {
        return lattigo_logSlots(params.getRawHandle());
    }
}  // namespace latticpp
