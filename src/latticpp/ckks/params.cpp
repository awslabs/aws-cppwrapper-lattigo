// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "params.h"

namespace latticpp {

    GoHandle<Parameters> getParams(HEParams paramId) {
        return GoHandle<Parameters>(lattigo_getParams(paramId));
    }

    uint64_t numSlots(const GoHandle<Parameters> &params) {
        return lattigo_numSlots(params.getRawHandle());
    }

    uint64_t logN(const GoHandle<Parameters> &params) {
        return lattigo_logN(params.getRawHandle());
    }

    uint64_t logQP(const GoHandle<Parameters> &params) {
        return lattigo_logQP(params.getRawHandle());
    }

    uint64_t maxLevel(const GoHandle<Parameters> &params) {
        return lattigo_maxLevel(params.getRawHandle());
    }

    double scale(const GoHandle<Parameters> &params) {
        return lattigo_paramsScale(params.getRawHandle());
    }

    double sigma(const GoHandle<Parameters> &params) {
        return lattigo_sigma(params.getRawHandle());
    }

    uint64_t getQi(const GoHandle<Parameters> &params, uint64_t i) {
        return lattigo_getQi(params.getRawHandle(), i);
    }

    uint64_t logQLvl(const GoHandle<Parameters> &params, uint64_t lvl) {
        return lattigo_logQLvl(params.getRawHandle(), lvl);
    }

    uint64_t logSlots(const GoHandle<Parameters> &params) {
        return lattigo_logSlots(params.getRawHandle());
    }
}  // namespace latticpp
