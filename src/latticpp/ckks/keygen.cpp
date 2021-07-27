// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "keygen.h"

using namespace std;

namespace latticpp {

    KeyGenerator newKeyGenerator(const Parameters &params) {
        return KeyGenerator(lattigo_newKeyGenerator(params.getRawHandle()));
    }

    KeyPairHandle genKeyPair(const KeyGenerator &keygen) {
        Lattigo_KeyPairHandle kp = lattigo_genKeyPair(keygen.getRawHandle());
        return KeyPairHandle { SecretKey(kp.sk), PublicKey(kp.pk) };
    }

    KeyPairHandle genKeyPairSparse(const KeyGenerator &keygen, uint64_t hw) {
        Lattigo_KeyPairHandle kp = lattigo_genKeyPairSparse(keygen.getRawHandle(), hw);
        return KeyPairHandle { SecretKey(kp.sk), PublicKey(kp.pk) };
    }

    RelinearizationKey genRelinKey(const KeyGenerator &keygen, const SecretKey &sk) {
        return RelinearizationKey(lattigo_genRelinearizationKey(keygen.getRawHandle(), sk.getRawHandle()));
    }

    RotationKeys genRotationKeysForRotations(const KeyGenerator &keygen, const SecretKey &sk, vector<int> shifts) {
        // convert from variable-sized int to fixed-size SIGNED int64_t
        vector<int64_t> fixed_width_shifts(shifts.size());
        for (int i = 0; i < shifts.size(); i++) {
            fixed_width_shifts[i] = static_cast<int64_t>(shifts[i]);
        }
        return RotationKeys(lattigo_genRotationKeysForRotations(keygen.getRawHandle(), sk.getRawHandle(), fixed_width_shifts.data(), shifts.size()));
    }

    EvaluationKey makeEvaluationKey(const RelinearizationKey &relinKey, const RotationKeys &rotKeys) {
        return EvaluationKey(lattigo_makeEvaluationKey(relinKey.getRawHandle(), rotKeys.getRawHandle()));
    }

    BootstrappingKey genBootstrappingKey(const KeyGenerator &keygen, const Parameters &params, const BootstrappingParameters &bootParams, const SecretKey &sk, const RelinearizationKey &relinKey, const RotationKeys &rotKeys) {
        return BootstrappingKey(lattigo_genBootstrappingKey(keygen.getRawHandle(), params.getRawHandle(), bootParams.getRawHandle(), sk.getRawHandle(), relinKey.getRawHandle(), rotKeys.getRawHandle()));
    }

    BootstrappingKey makeBootstrappingKey(const RelinearizationKey &relinKey, const RotationKeys &rotKeys) {
        return BootstrappingKey(lattigo_makeBootstrappingKey(relinKey.getRawHandle(), rotKeys.getRawHandle()));
    }
}  // namespace latticpp
