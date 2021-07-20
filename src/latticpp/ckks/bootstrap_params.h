// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/bootstrap_params.h"

namespace latticpp {

    // These correspond to the default bootstrapping parameters provided in Lattigo
    enum NamedBootstrappingParams {
        BootstrapParams_Set1,
        BootstrapParams_Set2,
        BootstrapParams_Set3,
        BootstrapParams_Set4,
        BootstrapParams_Set5
    };

    BootstrappingParameters getBootstrappingParams(const NamedBootstrappingParams paramId);

    Parameters genParams(const BootstrappingParameters &bootParams);

    // Get the secret key Hamming weight for which these bootstrapping parameters were created
    uint64_t secretHammingWeight(const BootstrappingParameters &bootParams);

    // The multiplicative depth of the bootstrapping circuit
    int bootstrapDepth(const BootstrappingParameters &bootParams);
}  // namespace latticpp
