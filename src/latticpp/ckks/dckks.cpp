// SPDX-License-Identifier: Apache-2.0

#include "dckks.h"

using namespace std;

namespace latticpp {

    vector<uint64_t> getCrpsRawHandles(const vector<Poly> &crps) {
        vector<uint64_t> crpsHandles(crps.size());
        for (size_t i = 0; i < crps.size(); i++) {
            crpsHandles.at(i) = crps.at(i).getRawHandle();
        }
        return crpsHandles;
    }

    CKGProtocol newCKGProtocol(const Parameters &params) {
        return CKGProtocol(lattigo_newCKGProtocol(params.getRawHandle()));
    }

    CKGShare ckgAllocateShares(const CKGProtocol &protocol) {
        return CKGShare(lattigo_ckgAllocateShares(protocol.getRawHandle()));
    }

    void ckgGenShare(const CKGProtocol &protocol, const SecretKey &sk,
                    const Poly &crp, CKGShare &shareOut) {
        lattigo_ckgGenShare(protocol.getRawHandle(), sk.getRawHandle(),
                            crp.getRawHandle(), shareOut.getRawHandle());
    }

    void ckgAggregateShares(const CKGProtocol &protocol, const CKGShare &share1,
                            const CKGShare &share2, CKGShare &shareOut) {
        lattigo_ckgAggregateShares(protocol.getRawHandle(), share1.getRawHandle(),
                                    share2.getRawHandle(), shareOut.getRawHandle());
    }

    void ckgGenPublicKey(const CKGProtocol &protocol, const CKGShare &roundShare,
                        const Poly &crp, PublicKey &pk) {
        lattigo_ckgGenPublicKey(protocol.getRawHandle(), roundShare.getRawHandle(),
                                crp.getRawHandle(), pk.getRawHandle());
    }

    RKGProtocol newRKGProtocol(const Parameters &params) {
        return RKGProtocol(lattigo_newRKGProtocol(params.getRawHandle()));
    }

    RKGShare newRKGShare() { return RKGShare(lattigo_newRKGShare()); }

    void rkgAllocateShares(const RKGProtocol &protocol, SecretKey &ephSk,
                        RKGShare &share1, RKGShare &share2) {
        lattigo_rkgAllocateShares(protocol.getRawHandle(), ephSk.getRawHandle(),
                                    share1.getRawHandle(), share2.getRawHandle());
    }

    void rkgGenShareRoundOne(const RKGProtocol &protocol, const SecretKey &sk,
                            const vector<Poly> &crps, SecretKey &ephSkOut,
                            RKGShare &shareOut) {
        vector<uint64_t> crpsHandles = getCrpsRawHandles(crps);
        lattigo_rkgGenShareRoundOne(protocol.getRawHandle(), sk.getRawHandle(),
                                    crpsHandles.data(), crpsHandles.size(),
                                    ephSkOut.getRawHandle(), shareOut.getRawHandle());
    }

    void rkgGenShareRoundTwo(const RKGProtocol &protocol, const SecretKey &ephSk,
                            const SecretKey &sk, const RKGShare &round1,
                            const vector<Poly> &crps, RKGShare &shareOut) {
        vector<uint64_t> crpsHandles = getCrpsRawHandles(crps);
        lattigo_rkgGenShareRoundTwo(protocol.getRawHandle(), ephSk.getRawHandle(),
                                    sk.getRawHandle(), round1.getRawHandle(),
                                    crpsHandles.data(), crpsHandles.size(),
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

    void cksKeySwitch(const CKSProtocol &protocol, const CKSShare &combined,
                    const Ciphertext &ct, Ciphertext &ctOut) {
        lattigo_cksKeySwitch(protocol.getRawHandle(), combined.getRawHandle(),
                            ct.getRawHandle(), ctOut.getRawHandle());
    }

    RTGProtocol newRTGProtocol(const Parameters &params) {
        return RTGProtocol(lattigo_newRotKGProtocol(params.getRawHandle()));
    }

    RTGShare rtgAllocateShares(const RTGProtocol &protocol) {
        return RTGShare(lattigo_rtgAllocateShares(protocol.getRawHandle()));
    }

    void rtgGenShare(const RTGProtocol &protocol, const SecretKey &sk,
                    uint64_t galEl, const vector<Poly> &crps, RTGShare &shareOut) {
        vector<uint64_t> crpsHandles = getCrpsRawHandles(crps);
        lattigo_rtgGenShare(protocol.getRawHandle(), sk.getRawHandle(), galEl,
                            crpsHandles.data(), crpsHandles.size(),
                            shareOut.getRawHandle());
    }

    void rtgAggregate(const RTGProtocol &protocol, const RTGShare &share1,
                    const RTGShare &share2, RTGShare &shareOut) {
        lattigo_rtgAggregate(protocol.getRawHandle(), share1.getRawHandle(),
                            share2.getRawHandle(), shareOut.getRawHandle());
    }

    void rtgGenRotationKey(const RTGProtocol &protocol, const RTGShare &share,
                        const vector<Poly> &crps, SwitchingKey &rotKey) {
        vector<uint64_t> crpsHandles = getCrpsRawHandles(crps);
        lattigo_rtgGenRotationKey(protocol.getRawHandle(), share.getRawHandle(),
                                    crpsHandles.data(), crpsHandles.size(),
                                    rotKey.getRawHandle());
    }
} // namespace latticpp
