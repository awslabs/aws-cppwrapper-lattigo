// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "bootstrap_params.h"

using namespace std;

namespace latticpp {

    BootstrappingParameters getBootstrappingParams(const NamedBootstrappingParams paramId) {
        return BootstrappingParameters(lattigo_getBootstrappingParams(paramId));
    }

    Parameters genParams(const BootstrappingParameters &bootParams) {
        return Parameters(lattigo_params(bootParams.getRawHandle()));
    }

    uint64_t secretHammingWeight(const BootstrappingParameters &bootParams) {
        return lattigo_bootstrap_h(bootParams.getRawHandle());
    }

    int bootstrapDepth(const BootstrappingParameters &bootParams) {
        return lattigo_bootstrap_depth(bootParams.getRawHandle());
    }
}  // namespace latticpp
