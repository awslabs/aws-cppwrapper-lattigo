// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "cgo/ring.h"
#include "latticpp/marshal/gohandle.h"
#include <vector>

namespace latticpp {

    Ring newRing(uint64_t n, std::vector<uint64_t> moduli);

    Poly newPoly(const Ring &ring);

    void add(const Ring &ring, const Poly &p1, const Poly &p2, Poly &p3);

    void copy(Poly &pTarget, const Poly &pSrc);

    UniformSampler newUniformSampler(const PRNG &prng, const Ring &ring);

    Poly readNewFromSampler(const UniformSampler &sampler);
} // namespace latticpp
