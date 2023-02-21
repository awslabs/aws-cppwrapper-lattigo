// SPDX-License-Identifier: Apache-2.0

package ckks

/*
#include <stdint.h>
*/
import "C"

import (
	"lattigo-cpp/marshal"
	"lattigo-cpp/ring"
	"unsafe"

	"github.com/ldsec/lattigo/v2/dckks"
	"github.com/ldsec/lattigo/v2/drlwe"
	lattigoring "github.com/ldsec/lattigo/v2/ring"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle13 = uint64

func getStoredCKGProtocol(protocolHandle Handle13) *dckks.CKGProtocol {
	ref := marshal.CrossLangObjMap.Get(protocolHandle)
	return (*dckks.CKGProtocol)(ref.Ptr)
}

func getStoredCKGShare(shareHandle Handle13) *drlwe.CKGShare {
	ref := marshal.CrossLangObjMap.Get(shareHandle)
	return (*drlwe.CKGShare)(ref.Ptr)
}

func getStoredRKGProtocol(protocolHandle Handle13) *dckks.RKGProtocol {
	ref := marshal.CrossLangObjMap.Get(protocolHandle)
	return (*dckks.RKGProtocol)(ref.Ptr)
}

func getStoredRKGShare(shareHandle Handle13) *drlwe.RKGShare {
	ref := marshal.CrossLangObjMap.Get(shareHandle)
	return (*drlwe.RKGShare)(ref.Ptr)
}

func getStoredCKSProtocol(protocolHandle Handle13) *dckks.CKSProtocol {
	ref := marshal.CrossLangObjMap.Get(protocolHandle)
	return (*dckks.CKSProtocol)(ref.Ptr)
}

func getStoredCKSShare(shareHandle Handle13) *drlwe.CKSShare {
	ref := marshal.CrossLangObjMap.Get(shareHandle)
	return (*drlwe.CKSShare)(ref.Ptr)
}

func getStoredRTGProtocol(protocolHandle Handle13) *dckks.RTGProtocol {
	ref := marshal.CrossLangObjMap.Get(protocolHandle)
	return (*dckks.RTGProtocol)(ref.Ptr)
}

func getStoredRTGShare(shareHandle Handle13) *drlwe.RTGShare {
	ref := marshal.CrossLangObjMap.Get(shareHandle)
	return (*drlwe.RTGShare)(ref.Ptr)
}

func getStoredCrps(crpsHandles *C.uint64_t, crpsHandlesLen uint64) []*lattigoring.Poly {
	size := unsafe.Sizeof(Handle13(0))
	basePtrIn := uintptr(unsafe.Pointer(crpsHandles))
	crps := make([]*lattigoring.Poly, crpsHandlesLen)
	for i := range crps {
		crps[i] = ring.GetStoredPoly(*(*Handle13)(unsafe.Pointer(basePtrIn + size*uintptr(i))))
	}
	return crps
}

//export lattigo_newCKGProtocol
func lattigo_newCKGProtocol(paramHandle Handle13) Handle13 {
	param := getStoredParameters(paramHandle)
	protocol := dckks.NewCKGProtocol(*param)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(protocol))
}

//export lattigo_ckgAllocateShares
func lattigo_ckgAllocateShares(protocolHandle Handle13) Handle13 {
	ckg := getStoredCKGProtocol(protocolHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ckg.AllocateShares()))
}

//export lattigo_ckgGenShare
func lattigo_ckgGenShare(protocolHandle, skHandle, crpHandle, shareOutHandle Handle13) {
	ckg := getStoredCKGProtocol(protocolHandle)
	sk := getStoredSecretKey(skHandle)
	crp := ring.GetStoredPoly(crpHandle)
	shareOut := getStoredCKGShare(shareOutHandle)
	ckg.GenShare(sk, crp, shareOut)
}

//export lattigo_ckgAggregateShares
func lattigo_ckgAggregateShares(protocolHandle, share1Handle, share2Handle, shareOutHandle Handle13) {
	ckg := getStoredCKGProtocol(protocolHandle)
	share1 := getStoredCKGShare(share1Handle)
	share2 := getStoredCKGShare(share2Handle)
	shareOut := getStoredCKGShare(shareOutHandle)
	ckg.AggregateShares(share1, share2, shareOut)
}

//export lattigo_ckgGenPublicKey
func lattigo_ckgGenPublicKey(protocolHandle, roundShareHandle, crpHandle, pkHandle Handle13) {
	ckg := getStoredCKGProtocol(protocolHandle)
	roundShare := getStoredCKGShare(roundShareHandle)
	crp := ring.GetStoredPoly(crpHandle)
	pk := getStoredPublicKey(pkHandle)
	ckg.GenPublicKey(roundShare, crp, pk)
}

//export lattigo_newRKGProtocol
func lattigo_newRKGProtocol(paramHandle Handle13) Handle13 {
	param := getStoredParameters(paramHandle)
	protocol := dckks.NewRKGProtocol(*param)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(protocol))
}

//export lattigo_newRKGShare
func lattigo_newRKGShare() Handle13 {
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(new(drlwe.RKGShare)))
}

//export lattigo_rkgAllocateShares
func lattigo_rkgAllocateShares(protocolHandle, ephSkHandle, share1Handle, share2Handle Handle13) {
	ckg := getStoredRKGProtocol(protocolHandle)
	ephSk := getStoredSecretKey(ephSkHandle)
	share1 := getStoredRKGShare(share1Handle)
	share2 := getStoredRKGShare(share2Handle)

	ephSkTmp, share1Tmp, share2Tmp := ckg.AllocateShares()

	*ephSk = *ephSkTmp
	*share1 = *share1Tmp
	*share2 = *share2Tmp
}

//export lattigo_rkgGenShareRoundOne
func lattigo_rkgGenShareRoundOne(protocolHandle, skHandle Handle13, crpsHandles *C.uint64_t, crpsHandlesLen uint64, ephSkOutHandle, shareOutHandle Handle13) {
	protocol := getStoredRKGProtocol(protocolHandle)
	sk := getStoredSecretKey(skHandle)
	crps := getStoredCrps(crpsHandles, crpsHandlesLen)
	ephSkOut := getStoredSecretKey(ephSkOutHandle)
	shareOut := getStoredRKGShare(shareOutHandle)

	protocol.GenShareRoundOne(sk, crps, ephSkOut, shareOut)
}

//export lattigo_rkgGenShareRoundTwo
func lattigo_rkgGenShareRoundTwo(protocolHandle, ephSkHandle, skHandle, round1Handle Handle13, crpsHandles *C.uint64_t, crpsHandlesLen uint64, shareOutHandle Handle13) {
	protocol := getStoredRKGProtocol(protocolHandle)
	ephSk := getStoredSecretKey(ephSkHandle)
	sk := getStoredSecretKey(skHandle)
	round1 := getStoredRKGShare(round1Handle)
	crps := getStoredCrps(crpsHandles, crpsHandlesLen)
	shareOut := getStoredRKGShare(shareOutHandle)

	protocol.GenShareRoundTwo(ephSk, sk, round1, crps, shareOut)
}

//export lattigo_rkgAggregateShares
func lattigo_rkgAggregateShares(protocolHandle, share1Handle, share2Handle, shareOutHandle Handle13) {
	protocol := getStoredRKGProtocol(protocolHandle)
	share1 := getStoredRKGShare(share1Handle)
	share2 := getStoredRKGShare(share2Handle)
	shareOut := getStoredRKGShare(shareOutHandle)
	protocol.AggregateShares(share1, share2, shareOut)
}

//export lattigo_rkgGenRelinearizationKey
func lattigo_rkgGenRelinearizationKey(protocolHandle, round1Handle, round2Handle, rlnKeyOutHandle Handle13) {
	protocol := getStoredRKGProtocol(protocolHandle)
	round1 := getStoredRKGShare(round1Handle)
	round2 := getStoredRKGShare(round2Handle)
	rlnKeyOut := getStoredRelinKey(rlnKeyOutHandle)
	protocol.GenRelinearizationKey(round1, round2, rlnKeyOut)
}

//export lattigo_newCKSProtocol
func lattigo_newCKSProtocol(paramHandle Handle13, sigmaSmudging float64) Handle13 {
	param := getStoredParameters(paramHandle)
	protocol := dckks.NewCKSProtocol(*param, sigmaSmudging)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(protocol))
}

//export lattigo_cksAllocateShare
func lattigo_cksAllocateShare(protocolHandle Handle13, level uint64) Handle13 {
	cks := getStoredCKSProtocol(protocolHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(cks.AllocateShare(int(level))))
}

//export lattigo_cksGenShare
func lattigo_cksGenShare(protocolHandle, skInputHandle, skOutputHandle, ctHandle, shareOutHandle Handle13) {
	protocol := getStoredCKSProtocol(protocolHandle)
	skInput := getStoredSecretKey(skInputHandle)
	skOutput := getStoredSecretKey(skOutputHandle)
	ct := getStoredCiphertext(ctHandle)
	shareOut := getStoredCKSShare(shareOutHandle)
	protocol.GenShare(skInput, skOutput, ct.Ciphertext, shareOut)
}

//export lattigo_cksAggregateShares
func lattigo_cksAggregateShares(protocolHandle, share1Handle, share2Handle, shareOutHandle Handle13) {
	protocol := getStoredCKSProtocol(protocolHandle)
	share1 := getStoredCKSShare(share1Handle)
	share2 := getStoredCKSShare(share2Handle)
	shareOut := getStoredCKSShare(shareOutHandle)
	protocol.AggregateShares(share1, share2, shareOut)
}

//export lattigo_cksKeySwitch
func lattigo_cksKeySwitch(protocolHandle, combinedHandle, ctHandle, ctOutHandle Handle13) {
	protocol := getStoredCKSProtocol(protocolHandle)
	combined := getStoredCKSShare(combinedHandle)
	ct := getStoredCiphertext(ctHandle)
	ctOut := getStoredCiphertext(ctOutHandle)
	protocol.KeySwitchCKKS(combined, ct, ctOut)
}

//export lattigo_newRotKGProtocol
func lattigo_newRotKGProtocol(paramHandle Handle13) Handle13 {
	param := getStoredParameters(paramHandle)
	protocol := dckks.NewRotKGProtocol(*param)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(protocol))
}

//export lattigo_rtgAllocateShares
func lattigo_rtgAllocateShares(protocolHandle Handle13) Handle13 {
	protocol := getStoredRTGProtocol(protocolHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(protocol.AllocateShares()))
}

//export lattigo_rtgGenShare
func lattigo_rtgGenShare(protocolHandle, skHandle Handle13, galEl uint64, crpsHandles *C.uint64_t, crpsHandlesLen uint64, shareOutHandle Handle13) {
	protocol := getStoredRTGProtocol(protocolHandle)
	sk := getStoredSecretKey(skHandle)
	crps := getStoredCrps(crpsHandles, crpsHandlesLen)
	shareOut := getStoredRTGShare(shareOutHandle)
	protocol.GenShare(sk, galEl, crps, shareOut)
}

//export lattigo_rtgAggregate
func lattigo_rtgAggregate(protocolHandle, share1Handle, share2Handle, shareOutHandle Handle13) {
	protocol := getStoredRTGProtocol(protocolHandle)
	share1 := getStoredRTGShare(share1Handle)
	share2 := getStoredRTGShare(share2Handle)
	shareOut := getStoredRTGShare(shareOutHandle)
	protocol.Aggregate(share1, share2, shareOut)
}

//export lattigo_rtgGenRotationKey
func lattigo_rtgGenRotationKey(protocolHandle, shareHandle Handle13, crpsHandles *C.uint64_t, crpsHandlesLen uint64, rotKeyHandle Handle13) {
	protocol := getStoredRTGProtocol(protocolHandle)
	share := getStoredRTGShare(shareHandle)
	crps := getStoredCrps(crpsHandles, crpsHandlesLen)
	rotKey := getStoredSwitchingKey(rotKeyHandle)
	protocol.GenRotationKey(share, crps, rotKey)
}
