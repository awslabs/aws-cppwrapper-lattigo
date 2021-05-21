// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "marshaler.h"
#include <sstream>

using namespace std;

namespace latticpp {

    void writeToStream(void* ostreamPtr, void* data, uint64_t len) {
        (*((ostream*)ostreamPtr)).write((const char*)data, len);
    }

    void marshalBinaryCiphertext(Ciphertext ct, std::ostream &stream) {
        lattigo_marshalBinaryCiphertext(ct.getRawHandle(), &writeToStream, (void*)(&stream));
    }

    void marshalBinaryParameters(Parameters params, std::ostream &stream) {
        lattigo_marshalBinaryParameters(params.getRawHandle(), &writeToStream, (void*)(&stream));
    }

    void marshalBinarySecretKey(SecretKey sk, std::ostream &stream) {
        lattigo_marshalBinarySecretKey(sk.getRawHandle(), &writeToStream, (void*)(&stream));
    }

    void marshalBinaryPublicKey(PublicKey pk, std::ostream &stream) {
        lattigo_marshalBinaryPublicKey(pk.getRawHandle(), &writeToStream, (void*)(&stream));
    }

    void marshalBinaryEvaluationKey(EvaluationKey evaKey, std::ostream &stream) {
        lattigo_marshalBinaryEvaluationKey(evaKey.getRawHandle(), &writeToStream, (void*)(&stream));
    }

    void marshalBinaryRotationKeys(RotationKeys rotKeys, std::ostream &stream) {
        lattigo_marshalBinaryRotationKeys(rotKeys.getRawHandle(), &writeToStream, (void*)(&stream));
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
