// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "cgo/dckks.h"
#include "latticpp/marshal/gohandle.h"
#include <vector>

namespace latticpp {

    CKGProtocol newCKGProtocol(const Parameters &params);

    CKGShare ckgAllocateShare(const CKGProtocol &protocol);

    CKGCRP ckgSampleCRP(const CKGProtocol &protocol, const PRNG &prng);

    void ckgGenShare(const CKGProtocol &protocol, const SecretKey &sk,
                    const CKGCRP &crp, CKGShare &shareOut);

    void ckgAggregateShares(const CKGProtocol &protocol, const CKGShare &share1,
                            const CKGShare &share2, CKGShare &shareOut);

    void ckgGenPublicKey(const CKGProtocol &protocol, const CKGShare &roundShare,
                        const CKGCRP &crp, PublicKey &pk);

    RKGProtocol newRKGProtocol(const Parameters &params);

    RKGShare newRKGShare();

    void rkgAllocateShare(const RKGProtocol &protocol, SecretKey &ephSk,
                        RKGShare &share1, RKGShare &share2);

    RKGCRP rkgSampleCRP(const RKGProtocol &protocol, const PRNG &prng);

    void rkgGenShareRoundOne(const RKGProtocol &protocol, const SecretKey &sk,
                            const RKGCRP &crp, SecretKey &ephSkOut,
                            RKGShare &shareOut);

    void rkgGenShareRoundTwo(const RKGProtocol &protocol, const SecretKey &ephSk,
                            const SecretKey &sk, const RKGShare &round1,
                            RKGShare &shareOut);

    void rkgAggregateShares(const RKGProtocol &protocol, const RKGShare &share1,
                            const RKGShare &share2, RKGShare &shareOut);

    void rkgGenRelinearizationKey(const RKGProtocol &protocol,
                                const RKGShare &round1, const RKGShare &round2,
                                RelinearizationKey &rlnKeyOut);

    CKSProtocol newCKSProtocol(const Parameters &params, double sigmaSmudging);

    CKSShare cksAllocateShare(const CKSProtocol &protocol, uint64_t level);

    void cksGenShare(const CKSProtocol &protocol, const SecretKey &skInput,
                    const SecretKey &skOutput, const Ciphertext &ct,
                    CKSShare &shareOut);

    void cksAggregateShares(const CKSProtocol &protocol, const CKSShare &share1,
                            const CKSShare &share2, CKSShare &shareOut);

    void cksKeySwitch(const CKSProtocol &protocol, const Ciphertext &ct,
                    const CKSShare &combined, Ciphertext &ctOut);

    RTGProtocol newRTGProtocol(const Parameters &params);

    RTGShare rtgAllocateShare(const RTGProtocol &protocol);

    RTGCRP rtgSampleCRP(const RTGProtocol &protocol, const PRNG &prng);

    void rtgGenShare(const RTGProtocol &protocol, const SecretKey &sk,
                    uint64_t galEl, const RTGCRP &crp,
                    RTGShare &shareOut);

    void rtgAggregateShares(const RTGProtocol &protocol, const RTGShare &share1,
                    const RTGShare &share2, RTGShare &shareOut);

    void rtgGenRotationKey(const RTGProtocol &protocol, const RTGShare &share,
                        const RTGCRP &crp, SwitchingKey &rotKey);
} // namespace latticpp