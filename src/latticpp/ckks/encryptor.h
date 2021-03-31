// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshall/gohandle.h"
#include "cgo/encryptor.h"

namespace latticpp {

    GoHandle<Encryptor> newEncryptorFromPk(const GoHandle<Parameters> &params, const GoHandle<PublicKey> &pk);

    GoHandle<Ciphertext> encryptNew(const GoHandle<Encryptor> &encryptor, const GoHandle<Plaintext> &pt);
}  // namespace latticpp
