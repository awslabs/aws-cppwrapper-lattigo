// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "cgo/dckks.h"
#include "latticpp/marshal/gohandle.h"
#include <vector>

namespace latticpp {

    CKGProtocol newCKGProtocol(const Parameters &params);

    CKGShare ckgAllocateShares(const CKGProtocol &protocol);

    void ckgGenShare(const CKGProtocol &protocol, const SecretKey &sk,
                    const Poly &crp, CKGShare &shareOut);

    void ckgAggregateShares(const CKGProtocol &protocol, const CKGShare &share1,
                            const CKGShare &share2, CKGShare &shareOut);

    void ckgGenPublicKey(const CKGProtocol &protocol, const CKGShare &roundShare,
                        const Poly &crp, PublicKey &pk);

    RKGProtocol newRKGProtocol(const Parameters &params);

    RKGShare newRKGShare();

    void rkgAllocateShares(const RKGProtocol &protocol, SecretKey &ephSk,
                        RKGShare &share1, RKGShare &share2);

    void rkgGenShareRoundOne(const RKGProtocol &protocol, const SecretKey &sk,
                            const std::vector<Poly> &crps, SecretKey &ephSkOut,
                            RKGShare &shareOut);

    void rkgGenShareRoundTwo(const RKGProtocol &protocol, const SecretKey &ephSk,
                            const SecretKey &sk, const RKGShare &round1,
                            const std::vector<Poly> &crps, RKGShare &shareOut);

    void rkgAggregateShares(const RKGProtocol &protocol, const RKGShare &share1,
                            const RKGShare &share2, RKGShare &shareOut);

    void rkgGenRelinearizationKey(const RKGProtocol &protocol,
                                const RKGShare &round1, const RKGShare &round2,
                                RelinearizationKey &rlnKeyOut);

    CKSProtocol newCKSProtocol(const Parameters &params, double sigmaSmudging);

    CKSShare cksAllocateShare(const CKSProtocol &protocol, int level);

    void cksGenShare(const CKSProtocol &protocol, const SecretKey &skInput,
                    const SecretKey &skOutput, const Ciphertext &ct,
                    CKSShare &shareOut);

    void cksAggregateShares(const CKSProtocol &protocol, const CKSShare &share1,
                            const CKSShare &share2, CKSShare &shareOut);

    void cksKeySwitch(const CKSProtocol &protocol, const CKSShare &combined,
                    const Ciphertext &ct, Ciphertext &ctOut);

    RTGProtocol newRTGProtocol(const Parameters &params);

    RTGShare rtgAllocateShares(const RTGProtocol &protocol);

    void rtgGenShare(const RTGProtocol &protocol, const SecretKey &sk,
                    uint64_t galEl, const std::vector<Poly> &crps,
                    RTGShare &shareOut);

    void rtgAggregate(const RTGProtocol &protocol, const RTGShare &share1,
                    const RTGShare &share2, RTGShare &shareOut);

    void rtgGenRotationKey(const RTGProtocol &protocol, const RTGShare &share,
                        const std::vector<Poly> &crps, SwitchingKey &rotKey);
} // namespace latticpp
