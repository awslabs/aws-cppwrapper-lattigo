// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/decryptor.h"

namespace latticpp {

    Decryptor newDecryptor(const Parameters &params, const SecretKey &sk);

    Plaintext decryptNew(const Decryptor &decryptor, const Ciphertext &ct);
}  // namespace latticpp
