// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "marshaler.h"
#include <vector>
#include <sstream>

using namespace std;

namespace latticpp {
    const string toString(vector<uint8_t> v) {
        ostringstream out;
        for (const uint8_t &c : v) {
            out << c;
        }
        return out.str();
    }

    const string marshalBinaryCiphertext(Ciphertext ct) {
        uint64_t serializedLength = lattigo_getDataLenCiphertext(ct.getRawHandle(), true);
        vector<uint8_t> buf(serializedLength);
        lattigo_marshalBinaryCiphertext(ct.getRawHandle(), buf.data());
        return toString(buf);
    }

    const string marshalBinaryParameters(Parameters params) {
        uint64_t serializedLength = lattigo_getDataLenParameters(params.getRawHandle(), true);
        vector<uint8_t> buf(serializedLength);
        lattigo_marshalBinaryParameters(params.getRawHandle(), buf.data());
        return toString(buf);
    }

    const string marshalBinarySecretKey(SecretKey sk) {
        uint64_t serializedLength = lattigo_getDataLenSecretKey(sk.getRawHandle(), true);
        vector<uint8_t> buf(serializedLength);
        lattigo_marshalBinarySecretKey(sk.getRawHandle(), buf.data());
        return toString(buf);
    }

    const string marshalBinaryPublicKey(PublicKey pk) {
        uint64_t serializedLength = lattigo_getDataLenPublicKey(pk.getRawHandle(), true);
        vector<uint8_t> buf(serializedLength);
        lattigo_marshalBinaryPublicKey(pk.getRawHandle(), buf.data());
        return toString(buf);
    }

    const string marshalBinaryEvaluationKey(EvaluationKey evaKey) {
        uint64_t serializedLength = lattigo_getDataLenEvaluationKey(evaKey.getRawHandle(), true);
        vector<uint8_t> buf(serializedLength);
        lattigo_marshalBinaryEvaluationKey(evaKey.getRawHandle(), buf.data());
        return toString(buf);
    }

    const string marshalBinarySwitchingKey(SwitchingKey switchKey) {
        uint64_t serializedLength = lattigo_getDataLenSwitchingKey(switchKey.getRawHandle(), true);
        vector<uint8_t> buf(serializedLength);
        lattigo_marshalBinarySwitchingKey(switchKey.getRawHandle(), buf.data());
        return toString(buf);
    }

    const string marshalBinaryRotationKeys(RotationKeys rotKeys) {
        uint64_t serializedLength = lattigo_getDataLenRotationKeys(rotKeys.getRawHandle(), true);
        vector<uint8_t> buf(serializedLength);
        lattigo_marshalBinaryRotationKeys(rotKeys.getRawHandle(), buf.data());
        return toString(buf);
    }

    Ciphertext unmarshalBinaryCiphertext(string buf);

    Parameters unmarshalBinaryParameters(string buf);

    SecretKey unmarshalBinarySecretKey(string buf);

    PublicKey unmarshalBinaryPublicKey(string buf);

    EvaluationKey unmarshalBinaryEvaluationKey(string buf);

    SwitchingKey unmarshalBinarySwitchingKey(string buf);

    RotationKeys unmarshalBinaryRotationKeys(string buf);
}  // namespace latticpp
