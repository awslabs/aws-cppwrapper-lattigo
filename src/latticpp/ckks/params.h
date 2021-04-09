// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshall/gohandle.h"
#include "cgo/params.h"

namespace latticpp {

    enum HEParams {
        // PN12QP109 is the index in DefaultParams for logQP = 109
        PN12QP109,
        // PN13QP218 is the index in DefaultParams for logQP = 218
        PN13QP218,
        // PN14QP438 is the index in DefaultParams for logQP = 438
        PN14QP438,
        // PN15QP880 is the index in DefaultParams for logQP = 880
        PN15QP880,
        // PN16QP1761 is the index in DefaultParams for logQP = 1761
        PN16QP1761,
        // PN12QP101pq is the index in DefaultParams for logQP = 101 (post quantum)
        PN12QP101pq,
        // PN13QP202pq is the index in DefaultParams for logQP = 202 (post quantum)
        PN13QP202pq,
        // PN14QP411pq is the index in DefaultParams for logQP = 411 (post quantum)
        PN14QP411pq,
        // PN15QP827pq is the index in DefaultParams for logQP = 827 (post quantum)
        PN15QP827pq,
        // PN16QP1654pq is the index in DefaultParams for logQP = 1654 (post quantum)
        PN16QP1654pq
    };

    Parameters getParams(HEParams paramId);

    uint64_t numSlots(const Parameters &params);

    uint64_t logN(const Parameters &params);

    uint64_t logQP(const Parameters &params);

    uint64_t maxLevel(const Parameters &params);

    double scale(const Parameters &params);

    double sigma(const Parameters &params);

    uint64_t getQi(const Parameters &params, uint64_t i);

    uint64_t logQLvl(const Parameters &params, uint64_t lvl);

    uint64_t logSlots(const Parameters &params);
}  // namespace latticpp
