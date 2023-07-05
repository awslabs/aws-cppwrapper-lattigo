// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/params.h"
#include <vector>

namespace latticpp {

    enum NamedClassicalParams {
        // PN12QP109 is the index in DefaultParams for logQP = 109
        PN12QP109,
        // PN13QP218 is the index in DefaultParams for logQP = 218
        PN13QP218,
        // PN14QP438 is the index in DefaultParams for logQP = 438
        PN14QP438,
        // PN15QP880 is the index in DefaultParams for logQP = 880
        PN15QP880,
        // PN16QP1761 is the index in DefaultParams for logQP = 1761
        PN16QP1761
    };

    enum NamedPQParams {
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

    Parameters getDefaultClassicalParams(NamedClassicalParams paramId);

    Parameters getDefaultPQParams(NamedPQParams paramId);

    // logN is the log of the polynomial ring degree. Alternatively, it is log(num_slots) + 1
    Parameters newParameters(uint64_t logN, const std::vector<uint64_t> &qi, const std::vector<uint64_t> &pi, uint8_t logScale);

    // logN is the log of the polynomial ring degree. Alternatively, it is log(num_slots) + 1
    Parameters newParametersFromLogModuli(uint64_t logN, const std::vector<uint8_t> &logQi, const std::vector<uint8_t> &logPi, uint8_t logScale);

    uint64_t numSlots(const Parameters &params);

    uint64_t logN(const Parameters &params);

    Ring ringQ(const Parameters &params);

    Ring ringP(const Parameters &params);

    RingQP ringQP(const Parameters &params);

    uint64_t logQP(const Parameters &params);

    uint64_t maxLevel(const Parameters &params);

    double scale(const Parameters &params);

    double sigma(const Parameters &params);

    uint64_t qi(const Parameters &params, uint64_t i);

    uint64_t pi(const Parameters &params, uint64_t i);

    uint64_t qiCount(const Parameters &params);

    uint64_t piCount(const Parameters &params);

    uint64_t logQLvl(const Parameters &params, uint64_t lvl);

    uint64_t logSlots(const Parameters &params);

    uint64_t galoisElementForRowRotation(const Parameters &params);

    std::vector<uint64_t> galoisElementsForRowInnerSum(const Parameters &params);
    
    uint64_t inverseGaloisElement(const Parameters &params, uint64_t galEl);

    uint64_t rotationFromGaloisElement(const Parameters &params, uint64_t galEl);
    
    uint64_t noiseBound(const Parameters &params);
}  // namespace latticpp