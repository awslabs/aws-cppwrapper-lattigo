// SPDX-License-Identifier: Apache-2.0

#include "ring.h"

using namespace std;

namespace latticpp {

    Ring newRing(uint64_t n, vector<uint64_t> moduli) {
        return Ring(lattigo_newRing(n, moduli.data(), moduli.size()));
    }

    Poly newPoly(const Ring &ring) {
        return Poly(lattigo_newPoly(ring.getRawHandle()));
    }

    void add(const Ring &ring, const Poly &p1, const Poly &p2, Poly &p3) {
        lattigo_ringAdd(ring.getRawHandle(), p1.getRawHandle(), p2.getRawHandle(),
                        p3.getRawHandle());
    }

    void copy(Poly &pTarget, const Poly &pSrc) {
        lattigo_polyCopy(pTarget.getRawHandle(), pSrc.getRawHandle());
    }

    UniformSampler newUniformSampler(const PRNG &prng, const Ring &ring) {
        return UniformSampler(
            lattigo_newUniformSampler(prng.getRawHandle(), ring.getRawHandle()));
    }

    Poly readNewFromSampler(const UniformSampler &sampler) {
        return Poly(lattigo_readNewFromSampler(sampler.getRawHandle()));
    }
} // namespace latticpp
