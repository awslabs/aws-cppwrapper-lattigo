// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

import "C"

import (
	"lattigo-cpp/marshal"
	"unsafe"

	"github.com/ldsec/lattigo/v2/ckks"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle8 = uint64

func getStoredCiphertext(ctHandle Handle8) *ckks.Ciphertext {
	ref := marshal.CrossLangObjMap.Get(ctHandle)
	return (*ckks.Ciphertext)(ref.Ptr)
}

//export lattigo_level
func lattigo_level(ctHandle Handle8) uint64 {
	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctHandle)
	return uint64(ctIn.Level())
}

//export lattigo_ciphertextScale
func lattigo_ciphertextScale(ctHandle Handle8) float64 {
	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctHandle)
	return ctIn.ScalingFactor()
}

//export lattigo_ciphertextDegree
func lattigo_ciphertextDegree(ctHandle Handle8) int {
	ct := getStoredCiphertext(ctHandle)
	return ct.Degree()
}

//export lattigo_copyNew
func lattigo_copyNew(ctHandle Handle8) Handle8 {
	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctHandle)

	var ctClone *ckks.Ciphertext
	ctClone = ctIn.CopyNew()
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ctClone))
}

//export lattigo_newCiphertext
func lattigo_newCiphertext(paramsHandle Handle8, degree uint64, level uint64, scale float64) Handle8 {
	var params *ckks.Parameters
	params = getStoredParameters(paramsHandle)

	var newCt *ckks.Ciphertext
	newCt = ckks.NewCiphertext(*params, int(degree), int(level), scale)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(newCt))
}
