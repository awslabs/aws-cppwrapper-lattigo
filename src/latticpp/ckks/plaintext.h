// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/plaintext.h"

namespace latticpp {

    Plaintext newPlaintext(const Parameters &params, uint64_t level);

    Plaintext newPlaintext(const Parameters &params, const Poly &poly, uint64_t level);

    double scale(const Plaintext &ct);

    void setScale(Plaintext &ct, double scale);

    Poly poly(const Plaintext &plaintext);
}  // namespace latticpp