// SPDX-License-Identifier: Apache-2.0

#include "dckks.h"

using namespace std;

namespace latticpp {

    CKGProtocol newCKGProtocol(const Parameters &params) {
        return CKGProtocol(lattigo_newCKGProtocol(params.getRawHandle()));
    }

    CKGShare ckgAllocateShare(const CKGProtocol &protocol) {
        return CKGShare(lattigo_ckgAllocateShare(protocol.getRawHandle()));
    }

    CKGCRP ckgSampleCRP(const CKGProtocol &protocol, const PRNG &prng) {
        return CKGCRP(lattigo_ckgSampleCRP(protocol.getRawHandle(), prng.getRawHandle()));
    }    

    void ckgGenShare(const CKGProtocol &protocol, const SecretKey &sk,
                    const CKGCRP &crp, CKGShare &shareOut) {
        lattigo_ckgGenShare(protocol.getRawHandle(), sk.getRawHandle(),
                            crp.getRawHandle(), shareOut.getRawHandle());
    }

    void ckgAggregateShares(const CKGProtocol &protocol, const CKGShare &share1,
                            const CKGShare &share2, CKGShare &shareOut) {
        lattigo_ckgAggregateShares(protocol.getRawHandle(), share1.getRawHandle(),
                                    share2.getRawHandle(), shareOut.getRawHandle());
    }

    void ckgGenPublicKey(const CKGProtocol &protocol, const CKGShare &roundShare,
                        const CKGCRP &crp, PublicKey &pk) {
        lattigo_ckgGenPublicKey(protocol.getRawHandle(), roundShare.getRawHandle(),
                                crp.getRawHandle(), pk.getRawHandle());
    }

    RKGProtocol newRKGProtocol(const Parameters &params) {
        return RKGProtocol(lattigo_newRKGProtocol(params.getRawHandle()));
    }

    RKGShare newRKGShare() { return RKGShare(lattigo_newRKGShare()); }

    void rkgAllocateShare(const RKGProtocol &protocol, SecretKey &ephSk,
                        RKGShare &share1, RKGShare &share2) {
        lattigo_rkgAllocateShare(protocol.getRawHandle(), ephSk.getRawHandle(),
                                    share1.getRawHandle(), share2.getRawHandle());
    }

    RKGCRP rkgSampleCRP(const RKGProtocol &protocol, const PRNG &prng) {
        return RKGCRP(lattigo_rkgSampleCRP(protocol.getRawHandle(), prng.getRawHandle()));
    }

    void rkgGenShareRoundOne(const RKGProtocol &protocol, const SecretKey &sk,
                            const RKGCRP &crp, SecretKey &ephSkOut,
                            RKGShare &shareOut) {
        lattigo_rkgGenShareRoundOne(protocol.getRawHandle(), sk.getRawHandle(),
                                    crp.getRawHandle(), ephSkOut.getRawHandle(),
                                    shareOut.getRawHandle());
    }

    void rkgGenShareRoundTwo(const RKGProtocol &protocol, const SecretKey &ephSk,
                            const SecretKey &sk, const RKGShare &round1,
                            RKGShare &shareOut) {
        lattigo_rkgGenShareRoundTwo(protocol.getRawHandle(), ephSk.getRawHandle(),
                                    sk.getRawHandle(), round1.getRawHandle(),
                                    shareOut.getRawHandle());
    }

    void rkgAggregateShares(const RKGProtocol &protocol, const RKGShare &share1,
                            const RKGShare &share2, RKGShare &shareOut) {
        lattigo_rkgAggregateShares(protocol.getRawHandle(), share1.getRawHandle(),
                                    share2.getRawHandle(), shareOut.getRawHandle());
    }

    void rkgGenRelinearizationKey(const RKGProtocol &protocol,
                                const RKGShare &round1, const RKGShare &round2,
                                RelinearizationKey &rlnKeyOut) {
        lattigo_rkgGenRelinearizationKey(protocol.getRawHandle(),
                                        round1.getRawHandle(), round2.getRawHandle(),
                                        rlnKeyOut.getRawHandle());
    }

    CKSProtocol newCKSProtocol(const Parameters &params, double sigmaSmudging) {
        return CKSProtocol(
            lattigo_newCKSProtocol(params.getRawHandle(), sigmaSmudging));
    }

    CKSShare cksAllocateShare(const CKSProtocol &protocol, uint64_t level) {
        return CKSShare(lattigo_cksAllocateShare(protocol.getRawHandle(), level));
    }

    void cksGenShare(const CKSProtocol &protocol, const SecretKey &skInput,
                    const SecretKey &skOutput, const Ciphertext &ct,
                    CKSShare &shareOut) {
        lattigo_cksGenShare(protocol.getRawHandle(), skInput.getRawHandle(),
                            skOutput.getRawHandle(), ct.getRawHandle(),
                            shareOut.getRawHandle());
    }

    void cksAggregateShares(const CKSProtocol &protocol, const CKSShare &share1,
                            const CKSShare &share2, CKSShare &shareOut) {
        lattigo_cksAggregateShares(protocol.getRawHandle(), share1.getRawHandle(),
                                    share2.getRawHandle(), shareOut.getRawHandle());
    }

    void cksKeySwitch(const CKSProtocol &protocol, const Ciphertext &ct,
                    const CKSShare &combined, Ciphertext &ctOut) {
        lattigo_cksKeySwitch(protocol.getRawHandle(), ct.getRawHandle(),
                            combined.getRawHandle(), ctOut.getRawHandle());
    }

    RTGProtocol newRTGProtocol(const Parameters &params) {
        return RTGProtocol(lattigo_newRTGProtocol(params.getRawHandle()));
    }

    RTGShare rtgAllocateShare(const RTGProtocol &protocol) {
        return RTGShare(lattigo_rtgAllocateShare(protocol.getRawHandle()));
    }

    RTGCRP rtgSampleCRP(const RTGProtocol &protocol, const PRNG &prng) {
        return RTGCRP(lattigo_rtgSampleCRP(protocol.getRawHandle(), prng.getRawHandle()));
    }    

    void rtgGenShare(const RTGProtocol &protocol, const SecretKey &sk,
                    uint64_t galEl, const RTGCRP &crp, RTGShare &shareOut) {
        lattigo_rtgGenShare(protocol.getRawHandle(), sk.getRawHandle(), galEl,
                            crp.getRawHandle(), shareOut.getRawHandle());
    }

    void rtgAggregateShares(const RTGProtocol &protocol, const RTGShare &share1,
                    const RTGShare &share2, RTGShare &shareOut) {
        lattigo_rtgAggregateShares(protocol.getRawHandle(), share1.getRawHandle(),
                            share2.getRawHandle(), shareOut.getRawHandle());
    }

    void rtgGenRotationKey(const RTGProtocol &protocol, const RTGShare &share,
                        const RTGCRP &crp, SwitchingKey &rotKey) {
        lattigo_rtgGenRotationKey(protocol.getRawHandle(), share.getRawHandle(),
                                    crp.getRawHandle(), rotKey.getRawHandle());
    }
} // namespace latticpp