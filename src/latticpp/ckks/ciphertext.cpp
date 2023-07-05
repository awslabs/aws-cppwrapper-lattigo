// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "ciphertext.h"

namespace latticpp {

    uint64_t level(const Ciphertext &ct) {
        return lattigo_level(ct.getRawHandle());
    }

    double scale(const Ciphertext &ct) {
        return lattigo_ciphertextGetScale(ct.getRawHandle());
    }

    void setScale(Ciphertext &ct, double scale) {
        lattigo_ciphertextSetScale(ct.getRawHandle(), scale);
    }

    uint64_t degree(const Ciphertext &ct) {
      return lattigo_ciphertextDegree(ct.getRawHandle());
    }

    Ciphertext copyNew(const Ciphertext &ct) {
        if (ct.getRawHandle() == 0) {
            return ct;
        }
        return lattigo_copyNew(ct.getRawHandle());
    }

    CiphertextQP copyNew(const CiphertextQP &ctQP) {
        if (ctQP.getRawHandle() == 0) {
            return ctQP;
        }
        return lattigo_copyNewCiphertextQP(ctQP.getRawHandle());
    }

    Ciphertext newCiphertext(const Parameters &params, uint64_t degree, uint64_t level) {
        return Ciphertext(lattigo_newCiphertext(params.getRawHandle(), degree, level));
    }

    CiphertextQP newCiphertextQP(const Parameters &params) {
        return CiphertextQP(lattigo_newCiphertextQP(params.getRawHandle()));
    }

    void setMetaData(Ciphertext &ctx, const MetaData &metaData) {
      lattigo_setCiphertextMetaData(ctx.getRawHandle(), metaData.getRawHandle());
    }

    void setMetaData(CiphertextQP &ctx, const MetaData &metaData) {
      lattigo_setCiphertextQPMetaData(ctx.getRawHandle(), metaData.getRawHandle());
    }

    MetaData getMetaData(const CiphertextQP &ctxQP) {
      return lattigo_getCiphertextQPMetaData(ctxQP.getRawHandle());
    }

    MetaData getMetaData(const Ciphertext &ctx) {
      return lattigo_getCiphertextMetaData(ctx.getRawHandle());
    }

    Poly poly(const Ciphertext &ctx, uint64_t i) {
      return lattigo_getCiphertextPoly(ctx.getRawHandle(), i);
    }

    PolyQP polyQP(const CiphertextQP &ctQP, uint64_t i) {
      return lattigo_getCiphertextPolyQP(ctQP.getRawHandle(), i);
    }

}  // namespace latticpp