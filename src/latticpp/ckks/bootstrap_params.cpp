// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "bootstrap_params.h"

using namespace std;

namespace latticpp {

    BootstrappingParameters getBootstrappingParams(const NamedBootstrappingParams paramId) {
        return BootstrappingParameters(lattigo_getBootstrappingParams(paramId));
    }

    Parameters genParams(const NamedBootstrappingParams paramId) {
        return Parameters(lattigo_params(paramId));
    }

    uint64_t ephemeralSecretWeight(const BootstrappingParameters &bootParams) {
        return lattigo_ephemeralSecretWeight(bootParams.getRawHandle());
    }
}  // namespace latticpp
