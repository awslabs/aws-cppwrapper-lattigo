// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/marshaler.h"

namespace latticpp {

    const std::string marshalBinaryCiphertext(Ciphertext ct);

    const std::string marshalBinaryParameters(Parameters params);

    const std::string marshalBinarySecretKey(SecretKey sk);

    const std::string marshalBinaryPublicKey(PublicKey pk);

    const std::string marshalBinaryEvaluationKey(EvaluationKey evaKey);

    const std::string marshalBinarySwitchingKey(SwitchingKey switchKey);

    const std::string marshalBinaryRotationKeys(RotationKeys rotKeys);

    Ciphertext unmarshalBinaryCiphertext(std::string buf);

    Parameters unmarshalBinaryParameters(std::string buf);

    SecretKey unmarshalBinarySecretKey(std::string buf);

    PublicKey unmarshalBinaryPublicKey(std::string buf);

    EvaluationKey unmarshalBinaryEvaluationKey(std::string buf);

    SwitchingKey unmarshalBinarySwitchingKey(std::string buf);

    RotationKeys unmarshalBinaryRotationKeys(std::string buf);
}  // namespace latticpp
