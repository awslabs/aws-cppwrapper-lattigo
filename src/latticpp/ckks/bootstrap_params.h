// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/bootstrap_params.h"

namespace latticpp {

    // These correspond to the default bootstrapping parameters provided in Lattigo
    enum NamedBootstrappingParams {
        N16QP1767H32768H32,
        N16QP1788H32768H32,
        N16QP1793H32768H32,
        N15QP880H16384H32
    };

    BootstrappingParameters getBootstrappingParams(const NamedBootstrappingParams paramId);

    Parameters genParams(const NamedBootstrappingParams paramId);

    uint64_t ephemeralSecretWeight(const BootstrappingParameters &bootParams);
}  // namespace latticpp
