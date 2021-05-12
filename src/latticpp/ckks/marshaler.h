// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/marshaler.h"

namespace latticpp {

    std::string marshalBinaryCiphertext(Ciphertext ct);

    std::string marshalBinaryParameters(Parameters params);

    std::string marshalBinarySecretKey(SecretKey sk);

    std::string marshalBinaryPublicKey(PublicKey pk);

    std::string marshalBinaryEvaluationKey(EvaluationKey evaKey);

    std::string marshalBinaryRotationKeys(RotationKeys rotKeys);

    Ciphertext unmarshalBinaryCiphertext(std::istream &stream);

    Parameters unmarshalBinaryParameters(std::istream &stream);

    SecretKey unmarshalBinarySecretKey(std::istream &stream);

    PublicKey unmarshalBinaryPublicKey(std::istream &stream);

    EvaluationKey unmarshalBinaryEvaluationKey(std::istream &stream);

    RotationKeys unmarshalBinaryRotationKeys(std::istream &stream);
}  // namespace latticpp
