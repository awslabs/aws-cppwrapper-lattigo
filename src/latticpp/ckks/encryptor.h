// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshall/gohandle.h"
#include "cgo/encryptor.h"

namespace latticpp {

    Encryptor newEncryptorFromPk(const Parameters &params, const PublicKey &pk);

    Ciphertext encryptNew(const Encryptor &encryptor, const Plaintext &pt);
}  // namespace latticpp
