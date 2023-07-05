// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

import "C"

import (
	"lattigo-cpp/marshal"
	"unsafe"

	"github.com/tuneinsight/lattigo/v4/ckks"
	"github.com/tuneinsight/lattigo/v4/rlwe"
	"github.com/tuneinsight/lattigo/v4/rlwe/ringqp"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle8 = uint64

func getStoredCiphertext(ctHandle Handle8) *rlwe.Ciphertext {
	ref := marshal.CrossLangObjMap.Get(ctHandle)
	return (*rlwe.Ciphertext)(ref.Ptr)
}

func getStoredCiphertextQP(ctHandle Handle8) *rlwe.CiphertextQP {
	ref := marshal.CrossLangObjMap.Get(ctHandle)
	return (*rlwe.CiphertextQP)(ref.Ptr)
}

func getStoredMetaData(metaDataHandle Handle8) *rlwe.MetaData {
	ref := marshal.CrossLangObjMap.Get(metaDataHandle)
	return (*rlwe.MetaData)(ref.Ptr)
}

//export lattigo_level
func lattigo_level(ctHandle Handle8) uint64 {
	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctHandle)
	return uint64(ctIn.Level())
}

//export lattigo_ciphertextGetScale
func lattigo_ciphertextGetScale(ctHandle Handle8) float64 {
	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctHandle)
	return ctIn.GetScale().Float64()
}

//export lattigo_ciphertextSetScale
func lattigo_ciphertextSetScale(ctHandle Handle8, scale float64) {
	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctHandle)
	ctIn.SetScale(rlwe.NewScale(scale))
}

//export lattigo_ciphertextDegree
func lattigo_ciphertextDegree(ctHandle Handle8) uint64 {
	ct := getStoredCiphertext(ctHandle)
	return uint64(ct.Degree())
}

//export lattigo_copyNew
func lattigo_copyNew(ctHandle Handle8) Handle8 {
	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctHandle)

	var ctClone *rlwe.Ciphertext
	ctClone = ctIn.CopyNew()
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ctClone))
}

//export lattigo_copyNewCiphertextQP
func lattigo_copyNewCiphertextQP(ctQPHandle Handle8) Handle8 {
	ctQPIn := getStoredCiphertextQP(ctQPHandle)
	ctQPOut := ctQPIn.CopyNew()
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ctQPOut))
}

//export lattigo_newCiphertext
func lattigo_newCiphertext(paramsHandle Handle8, degree uint64, level uint64) Handle8 {
	var params *ckks.Parameters
	params = getStoredParameters(paramsHandle)

	var newCt *rlwe.Ciphertext
	newCt = ckks.NewCiphertext(*params, int(degree), int(level))
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(newCt))
}

//export lattigo_newCiphertextQP
func lattigo_newCiphertextQP(paramsHandle Handle8) Handle8 {
	params := getStoredParameters(paramsHandle)
	ringQP := params.RingQP()
	newCtQp := rlwe.CiphertextQP{Value: [2]ringqp.Poly{ringQP.NewPoly(), ringQP.NewPoly()}, MetaData: rlwe.MetaData{Scale: params.DefaultScale(), IsNTT: true}}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&newCtQp))
}

//export lattigo_setCiphertextMetaData
func lattigo_setCiphertextMetaData(ctxHandle, metaDataHandle Handle8) {
	ctx := getStoredCiphertext(ctxHandle)
	metaData := getStoredMetaData(metaDataHandle)
	(*ctx).MetaData = *metaData
}

//export lattigo_setCiphertextQPMetaData
func lattigo_setCiphertextQPMetaData(ctxQPHandle, metaDataHandle Handle8) {
	ctxQP := getStoredCiphertextQP(ctxQPHandle)
	metaData := getStoredMetaData(metaDataHandle)
	(*ctxQP).MetaData = *metaData
}

//export lattigo_getCiphertextMetaData
func lattigo_getCiphertextMetaData(ctxHandle Handle8) Handle8 {
	ctx := getStoredCiphertext(ctxHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&ctx.MetaData))
}

//export lattigo_getCiphertextQPMetaData
func lattigo_getCiphertextQPMetaData(ctxHandle Handle8) Handle8 {
	ctxQP := getStoredCiphertextQP(ctxHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&ctxQP.MetaData))
}

//export lattigo_getCiphertextPoly
func lattigo_getCiphertextPoly(ctxHandle Handle8, i uint64) Handle8 {
	ctx := getStoredCiphertext(ctxHandle)
	if int(i) > ctx.Degree() {
		panic("index exceed ciphertext degree")
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ctx.Value[i]))
}

//export lattigo_getCiphertextPolyQP
func lattigo_getCiphertextPolyQP(ctQPHandle Handle8, i uint64) Handle8 {
	ctxQP := getStoredCiphertextQP(ctQPHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&ctxQP.Value[i]))
}
