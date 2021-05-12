// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "marshaler.h"
#include <vector>
#include <sstream>

using namespace std;

namespace latticpp {
    string toString(vector<uint8_t> v) {
        ostringstream out;
        for (const uint8_t &c : v) {
            out << c;
        }
        return out.str();
    }

    string marshalBinaryCiphertext(Ciphertext ct) {
        uint64_t serializedLength = lattigo_getDataLenCiphertext(ct.getRawHandle(), true);
        vector<uint8_t> buf(serializedLength);
        lattigo_marshalBinaryCiphertext(ct.getRawHandle(), buf.data());
        return toString(buf);
    }

    string marshalBinaryParameters(Parameters params) {
        uint64_t serializedLength = lattigo_getDataLenParameters(params.getRawHandle(), true);
        vector<uint8_t> buf(serializedLength);
        lattigo_marshalBinaryParameters(params.getRawHandle(), buf.data());
        return toString(buf);
    }

    string marshalBinarySecretKey(SecretKey sk) {
        uint64_t serializedLength = lattigo_getDataLenSecretKey(sk.getRawHandle(), true);
        vector<uint8_t> buf(serializedLength);
        lattigo_marshalBinarySecretKey(sk.getRawHandle(), buf.data());
        return toString(buf);
    }

    string marshalBinaryPublicKey(PublicKey pk) {
        uint64_t serializedLength = lattigo_getDataLenPublicKey(pk.getRawHandle(), true);
        vector<uint8_t> buf(serializedLength);
        lattigo_marshalBinaryPublicKey(pk.getRawHandle(), buf.data());
        return toString(buf);
    }

    string marshalBinaryEvaluationKey(EvaluationKey evaKey) {
        uint64_t serializedLength = lattigo_getDataLenEvaluationKey(evaKey.getRawHandle(), true);
        vector<uint8_t> buf(serializedLength);
        lattigo_marshalBinaryEvaluationKey(evaKey.getRawHandle(), buf.data());
        return toString(buf);
    }

    string marshalBinaryRotationKeys(RotationKeys rotKeys) {
        uint64_t serializedLength = lattigo_getDataLenRotationKeys(rotKeys.getRawHandle(), true);
        vector<uint8_t> buf(serializedLength);
        lattigo_marshalBinaryRotationKeys(rotKeys.getRawHandle(), buf.data());
        return toString(buf);
    }

    Ciphertext unmarshalBinaryCiphertext(istream &stream) {
        vector<char> buffer( istreambuf_iterator<char>{stream},
                             istreambuf_iterator<char>() );
        return Ciphertext(lattigo_unmarshalBinaryCiphertext(buffer.data(), buffer.size()));
    }

    Parameters unmarshalBinaryParameters(istream &stream) {
        vector<char> buffer( istreambuf_iterator<char>{stream},
                             istreambuf_iterator<char>() );
        return Parameters(lattigo_unmarshalBinaryParameters(buffer.data(), buffer.size()));
    }

    SecretKey unmarshalBinarySecretKey(istream &stream) {
        vector<char> buffer( istreambuf_iterator<char>{stream},
                             istreambuf_iterator<char>() );
        return SecretKey(lattigo_unmarshalBinarySecretKey(buffer.data(), buffer.size()));
    }

    PublicKey unmarshalBinaryPublicKey(istream &stream) {
        vector<char> buffer( istreambuf_iterator<char>{stream},
                             istreambuf_iterator<char>() );
        return PublicKey(lattigo_unmarshalBinaryPublicKey(buffer.data(), buffer.size()));
    }

    EvaluationKey unmarshalBinaryEvaluationKey(istream &stream) {
        vector<char> buffer( istreambuf_iterator<char>{stream},
                             istreambuf_iterator<char>() );
        return EvaluationKey(lattigo_unmarshalBinaryEvaluationKey(buffer.data(), buffer.size()));
    }

    RotationKeys unmarshalBinaryRotationKeys(istream &stream) {
        vector<char> buffer( istreambuf_iterator<char>{stream},
                             istreambuf_iterator<char>() );
        return RotationKeys(lattigo_unmarshalBinaryRotationKeys(buffer.data(), buffer.size()));
    }
}  // namespace latticpp
