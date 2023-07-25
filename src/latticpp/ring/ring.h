// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "cgo/ring.h"
#include "latticpp/marshal/gohandle.h"
#include <vector>

namespace latticpp {

    Ring newRing(uint64_t n, std::vector<uint64_t> moduli);

    Poly newPoly(const Ring &ring);

    PolyQP newPolyQP(const RingQP &ring);

    void copy(Poly &dst, const Poly &src);

    void copy(PolyQP &pTarget, const PolyQP &pSrc);

    PolyQP copyNew(const PolyQP &src);

    UniformSampler newUniformSampler(const PRNG &prng, const Ring &ring);

    Poly polyQ(const PolyQP &polyQp);

    Poly polyP(const PolyQP &polyQp);

    void copyLvl(uint64_t level, const Poly &sourcePoly, Poly &targetPoly);

    void copyLvlToOtherLvl(uint64_t srcLevel, uint64_t dstLevel, const Poly &srcPoly, Poly &dstPoly);

    BasisExtender newBasisExtender(const Ring &ringQ, const Ring &ringP);

    void modUpQtoP(const BasisExtender &ext, uint64_t levelQ, uint64_t levelP, const Poly &polQ, Poly &polP);

    void invNTTLvl(const RingQP &ringqp, uint64_t levelQ, uint64_t levelP, const PolyQP &pIn, PolyQP &pOut);

    void nttLvl(const RingQP &ringqp, uint64_t levelQ, uint64_t levelP, const PolyQP &pIn, PolyQP &pOut);

    void invNTTLvl(const Ring &ring, uint64_t level, const Poly &pIn, Poly &pOut);

    void nttLvl(const Ring &ring, uint64_t level, const Poly &pIn, Poly &pOut);

    void invMFormLvl(const RingQP &ringqp, uint64_t levelQ, uint64_t levelP, const PolyQP &pIn, PolyQP &pOut);

    void mFormLvl(const RingQP &ringqp, uint64_t levelQ, uint64_t levelP, const PolyQP &pIn, PolyQP &pOut);

    void invMFormLvl(const Ring &ring, uint64_t level, const Poly &pIn, Poly &pOut);

    void mFormLvl(const Ring &ring, uint64_t level, const Poly &pIn, Poly &pOut);

    uint64_t degree(const Poly &p);

    uint64_t ringN(const Ring &ring);

    std::vector<uint64_t> permuteNTTIndex(const Ring &ring, uint64_t galEl);

    void permuteNTTWithIndexLvl(const Ring &ring, uint64_t level, const Poly &polyIn, const std::vector<uint64_t> &index, Poly &polyOut);

    uint64_t log2OfInnerSum(uint64_t level, const Ring &ring, const Poly &poly);

    void addLvl(const RingQP &ring, uint64_t levelQ, uint64_t levelP, const PolyQP &p1, const PolyQP &p2, PolyQP &polyOut);

    void mulCoeffsMontgomeryAndAddLvl(const RingQP &ringQP, uint64_t levelQ, uint64_t levelP, const PolyQP &p1, const PolyQP &p2, PolyQP &polyOut);

    void mulCoeffsMontgomeryAndAddLvl(const Ring &ring, uint64_t level, const Poly &p1, const Poly &p2, Poly &polyOut);

    uint64_t equals(const Poly &p1, const Poly &p2);
    
} // namespace latticpp