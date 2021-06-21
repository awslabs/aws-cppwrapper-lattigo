// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/bootstrap.h"

namespace latticpp {

    Bootstrapper newBootstrapper(const Parameters &params, const BootstrappingParameters &bootParams, const BootstrappingKey &bootKey);

    Ciphertext bootstrap(const Bootstrapper &btp, const Ciphertext &ct);
}  // namespace latticpp
