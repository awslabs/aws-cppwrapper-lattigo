// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/encryptor.h"

namespace latticpp {

    Encryptor newEncryptor(const Parameters &params, const SecretKey &sk);

    Encryptor newEncryptor(const Parameters &params, const PublicKey &pk);

    Ciphertext encryptNew(const Encryptor &encryptor, const Plaintext &pt);

    void encryptZeroQP(const Parameters &params, const SecretKey &sk, CiphertextQP &ctxQP);
}  // namespace latticpp
