// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/bootstrap_params.h"

namespace latticpp {

    enum NamedBootstrappingParams {
        // Bootstrapping parameters index 0
        BootstrapParams_Set2,
        // Bootstrapping parameters index 1
        BootstrapParams_Set5,
        // Bootstrapping parameters index 2
        BootstrapParams_Set7,
        // Bootstrapping parameters index 3
        BootstrapParams_Set4
    };

    BootstrappingParameters getBootstrappingParams(const NamedBootstrappingParams paramId);

    uint64_t bootstrap_h(const BootstrappingParameters &bootParamHandle);
}  // namespace latticpp
