// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshall/gohandle.h"
#include "cgo/decryptor.h"

namespace latticpp {

    GoHandle<Decryptor> newDecryptor(const GoHandle<Parameters> &params, const GoHandle<SecretKey> &sk);

    GoHandle<Plaintext> decryptNew(const GoHandle<Decryptor> &decryptor, const GoHandle<Ciphertext> &ct);
}  // namespace latticpp
