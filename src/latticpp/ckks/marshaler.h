// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/marshaler.h"

namespace latticpp {

    void marshalBinaryCiphertext(Ciphertext ct, std::ostream &stream);

    void marshalBinaryParameters(Parameters params, std::ostream &stream);

    void marshalBinarySecretKey(SecretKey sk, std::ostream &stream);

    void marshalBinaryPublicKey(PublicKey pk, std::ostream &stream);

    void marshalBinaryEvaluationKey(EvaluationKey evaKey, std::ostream &stream);

    void marshalBinaryRotationKeys(RotationKeys rotKeys, std::ostream &stream);

    Ciphertext unmarshalBinaryCiphertext(std::istream &stream);

    Parameters unmarshalBinaryParameters(std::istream &stream);

    SecretKey unmarshalBinarySecretKey(std::istream &stream);

    PublicKey unmarshalBinaryPublicKey(std::istream &stream);

    EvaluationKey unmarshalBinaryEvaluationKey(std::istream &stream);

    RotationKeys unmarshalBinaryRotationKeys(std::istream &stream);
}  // namespace latticpp
