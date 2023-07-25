// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/marshaler.h"

namespace latticpp {

    void marshalBinaryCiphertext(const Ciphertext &ct, std::ostream &stream);

    void marshalBinaryParameters(const Parameters &params, std::ostream &stream);

    void marshalBinaryBootstrapParameters(const BootstrappingParameters &btp_params, std::ostream &stream);

    void marshalBinarySecretKey(const SecretKey &sk, std::ostream &stream);

    void marshalBinaryPublicKey(const PublicKey &pk, std::ostream &stream);

    void marshalBinaryRelinearizationKey(const RelinearizationKey &relinKey, std::ostream &stream);

    void marshalBinaryRotationKeys(const RotationKeys &rotKeys, std::ostream &stream);

    Ciphertext unmarshalBinaryCiphertext(std::istream &stream);

    Parameters unmarshalBinaryParameters(std::istream &stream);

    BootstrappingParameters unmarshalBinaryBootstrapParameters(std::istream &stream);

    SecretKey unmarshalBinarySecretKey(std::istream &stream);

    PublicKey unmarshalBinaryPublicKey(std::istream &stream);

    RelinearizationKey unmarshalBinaryRelinearizationKey(std::istream &stream);

    RotationKeys unmarshalBinaryRotationKeys(std::istream &stream);
}  // namespace latticpp