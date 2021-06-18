// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "bootstrap_params.h"

using namespace std;

namespace latticpp {

    BootstrappingParameters getBootstrappingParams(const NamedBootstrappingParams paramId) {
        return BootstrappingParameters(lattigo_getBootstrappingParams(paramId));
    }

    uint64_t bootstrap_h(const BootstrappingParameters &bootParamHandle) {
        return lattigo_bootstrap_h(bootParamHandle.getRawHandle());
    }
}  // namespace latticpp
