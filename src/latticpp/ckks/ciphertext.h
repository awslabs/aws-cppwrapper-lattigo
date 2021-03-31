// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshall/gohandle.h"
#include "cgo/ciphertext.h"

namespace latticpp {

    uint64_t level(const GoHandle<Ciphertext> &ct);

    double scale(const GoHandle<Ciphertext> &ct);

    GoHandle<Ciphertext> copyNew(const GoHandle<Ciphertext> &ct);

    GoHandle<Ciphertext> newCiphertext(const GoHandle<Parameters> &params, uint64_t degree, uint64_t level, double scale);
}  // namespace latticpp
