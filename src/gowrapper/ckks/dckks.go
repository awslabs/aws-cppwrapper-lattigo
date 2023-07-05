// SPDX-License-Identifier: Apache-2.0

package ckks

/*
#include <stdint.h>
*/
import "C"

import (
	"lattigo-cpp/marshal"
	"lattigo-cpp/utils"
	"unsafe"

	"github.com/tuneinsight/lattigo/v4/dckks"
	"github.com/tuneinsight/lattigo/v4/drlwe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle13 = uint64

func getStoredCKGProtocol(protocolHandle Handle13) *drlwe.CKGProtocol {
	ref := marshal.CrossLangObjMap.Get(protocolHandle)
	return (*drlwe.CKGProtocol)(ref.Ptr)
}

func getStoredCKGShare(shareHandle Handle13) *drlwe.CKGShare {
	ref := marshal.CrossLangObjMap.Get(shareHandle)
	return (*drlwe.CKGShare)(ref.Ptr)
}

func getStoredCKGCRP(shareHandle Handle13) *drlwe.CKGCRP {
	ref := marshal.CrossLangObjMap.Get(shareHandle)
	return (*drlwe.CKGCRP)(ref.Ptr)
}

func getStoredRKGProtocol(protocolHandle Handle13) *drlwe.RKGProtocol {
	ref := marshal.CrossLangObjMap.Get(protocolHandle)
	return (*drlwe.RKGProtocol)(ref.Ptr)
}

func getStoredRKGShare(shareHandle Handle13) *drlwe.RKGShare {
	ref := marshal.CrossLangObjMap.Get(shareHandle)
	return (*drlwe.RKGShare)(ref.Ptr)
}

func getStoredRKGCRP(shareHandle Handle13) *drlwe.RKGCRP {
	ref := marshal.CrossLangObjMap.Get(shareHandle)
	return (*drlwe.RKGCRP)(ref.Ptr)
}

func getStoredCKSProtocol(protocolHandle Handle13) *drlwe.CKSProtocol {
	ref := marshal.CrossLangObjMap.Get(protocolHandle)
	return (*drlwe.CKSProtocol)(ref.Ptr)
}

func getStoredCKSShare(shareHandle Handle13) *drlwe.CKSShare {
	ref := marshal.CrossLangObjMap.Get(shareHandle)
	return (*drlwe.CKSShare)(ref.Ptr)
}

func getStoredRTGProtocol(protocolHandle Handle13) *drlwe.RTGProtocol {
	ref := marshal.CrossLangObjMap.Get(protocolHandle)
	return (*drlwe.RTGProtocol)(ref.Ptr)
}

func getStoredRTGShare(shareHandle Handle13) *drlwe.RTGShare {
	ref := marshal.CrossLangObjMap.Get(shareHandle)
	return (*drlwe.RTGShare)(ref.Ptr)
}

func getStoredRTGCRP(shareHandle Handle13) *drlwe.RTGCRP {
	ref := marshal.CrossLangObjMap.Get(shareHandle)
	return (*drlwe.RTGCRP)(ref.Ptr)
}

//export lattigo_newCKGProtocol
func lattigo_newCKGProtocol(paramHandle Handle13) Handle13 {
	param := getStoredParameters(paramHandle)
	protocol := dckks.NewCKGProtocol(*param)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(protocol))
}

//export lattigo_ckgAllocateShare
func lattigo_ckgAllocateShare(protocolHandle Handle13) Handle13 {
	ckg := getStoredCKGProtocol(protocolHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ckg.AllocateShare()))
}

//export lattigo_ckgSampleCRP
func lattigo_ckgSampleCRP(protocolHandle, prngHandle Handle13) Handle13 {
	ckg := getStoredCKGProtocol(protocolHandle)
	prng := utils.GetStoredKeyedPRNG(prngHandle)
	crp := ckg.SampleCRP(prng)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&crp))
}

//export lattigo_ckgGenShare
func lattigo_ckgGenShare(protocolHandle, skHandle, crpHandle, shareOutHandle Handle13) {
	ckg := getStoredCKGProtocol(protocolHandle)
	sk := getStoredSecretKey(skHandle)
	crp := getStoredCKGCRP(crpHandle)
	shareOut := getStoredCKGShare(shareOutHandle)
	ckg.GenShare(sk, *crp, shareOut)
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
	crp := getStoredCKGCRP(crpHandle)
	pk := getStoredPublicKey(pkHandle)
	ckg.GenPublicKey(roundShare, *crp, pk)
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

//export lattigo_rkgAllocateShare
func lattigo_rkgAllocateShare(protocolHandle, ephSkHandle, share1Handle, share2Handle Handle13) {
	ckg := getStoredRKGProtocol(protocolHandle)
	ephSk := getStoredSecretKey(ephSkHandle)
	share1 := getStoredRKGShare(share1Handle)
	share2 := getStoredRKGShare(share2Handle)

	ephSkTmp, share1Tmp, share2Tmp := ckg.AllocateShare()

	*ephSk = *ephSkTmp
	*share1 = *share1Tmp
	*share2 = *share2Tmp
}

//export lattigo_rkgSampleCRP
func lattigo_rkgSampleCRP(protocolHandle, prngHandle Handle13) Handle13 {
	rkg := getStoredRKGProtocol(protocolHandle)
	prng := utils.GetStoredKeyedPRNG(prngHandle)
	crp := rkg.SampleCRP(prng)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&crp))
}

//export lattigo_rkgGenShareRoundOne
func lattigo_rkgGenShareRoundOne(protocolHandle, skHandle, crpHandle, ephSkOutHandle, shareOutHandle Handle13) {
	protocol := getStoredRKGProtocol(protocolHandle)
	sk := getStoredSecretKey(skHandle)
	crp := getStoredRKGCRP(crpHandle)
	ephSkOut := getStoredSecretKey(ephSkOutHandle)
	shareOut := getStoredRKGShare(shareOutHandle)

	protocol.GenShareRoundOne(sk, *crp, ephSkOut, shareOut)
}

//export lattigo_rkgGenShareRoundTwo
func lattigo_rkgGenShareRoundTwo(protocolHandle, ephSkHandle, skHandle, round1Handle, shareOutHandle Handle13) {
	protocol := getStoredRKGProtocol(protocolHandle)
	ephSk := getStoredSecretKey(ephSkHandle)
	sk := getStoredSecretKey(skHandle)
	round1 := getStoredRKGShare(round1Handle)
	shareOut := getStoredRKGShare(shareOutHandle)

	protocol.GenShareRoundTwo(ephSk, sk, round1, shareOut)
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
	protocol.GenShare(skInput, skOutput, ct, shareOut)
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
func lattigo_cksKeySwitch(protocolHandle, ctHandle, combinedHandle, ctOutHandle Handle13) {
	protocol := getStoredCKSProtocol(protocolHandle)
	ct := getStoredCiphertext(ctHandle)
	combined := getStoredCKSShare(combinedHandle)
	ctOut := getStoredCiphertext(ctOutHandle)
	protocol.KeySwitch(ct, combined, ctOut)
}

//export lattigo_newRTGProtocol
func lattigo_newRTGProtocol(paramHandle Handle13) Handle13 {
	param := getStoredParameters(paramHandle)
	protocol := dckks.NewRTGProtocol(*param)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(protocol))
}

//export lattigo_rtgAllocateShare
func lattigo_rtgAllocateShare(protocolHandle Handle13) Handle13 {
	protocol := getStoredRTGProtocol(protocolHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(protocol.AllocateShare()))
}

//export lattigo_rtgSampleCRP
func lattigo_rtgSampleCRP(protocolHandle, prngHandle Handle13) Handle13 {
	rtg := getStoredRTGProtocol(protocolHandle)
	prng := utils.GetStoredKeyedPRNG(prngHandle)
	crp := rtg.SampleCRP(prng)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&crp))
}

//export lattigo_rtgGenShare
func lattigo_rtgGenShare(protocolHandle, skHandle Handle13, galEl uint64, crpHandle, shareOutHandle Handle13) {
	protocol := getStoredRTGProtocol(protocolHandle)
	sk := getStoredSecretKey(skHandle)
	crp := getStoredRTGCRP(crpHandle)
	shareOut := getStoredRTGShare(shareOutHandle)
	protocol.GenShare(sk, galEl, *crp, shareOut)
}

//export lattigo_rtgAggregateShares
func lattigo_rtgAggregateShares(protocolHandle, share1Handle, share2Handle, shareOutHandle Handle13) {
	protocol := getStoredRTGProtocol(protocolHandle)
	share1 := getStoredRTGShare(share1Handle)
	share2 := getStoredRTGShare(share2Handle)
	shareOut := getStoredRTGShare(shareOutHandle)
	protocol.AggregateShares(share1, share2, shareOut)
}

//export lattigo_rtgGenRotationKey
func lattigo_rtgGenRotationKey(protocolHandle, shareHandle Handle13, crpHandle, rotKeyHandle Handle13) {
	protocol := getStoredRTGProtocol(protocolHandle)
	share := getStoredRTGShare(shareHandle)
	crp := getStoredRTGCRP(crpHandle)
	rotKey := getStoredSwitchingKey(rotKeyHandle)
	protocol.GenRotationKey(share, *crp, rotKey)
}
