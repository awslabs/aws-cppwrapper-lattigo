// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

/*
typedef const double constDouble;
*/
import "C"

import (
	"lattigo-cpp/marshal"

	"math"
	"unsafe"

	"github.com/tuneinsight/lattigo/v4/ckks"
	"github.com/tuneinsight/lattigo/v4/rlwe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle2 = uint64

func getStoredEncoder(encoderHandle Handle2) *ckks.Encoder {
	ref := marshal.CrossLangObjMap.Get(encoderHandle)
	return (*ckks.Encoder)(ref.Ptr)
}

//export lattigo_newEncoder
func lattigo_newEncoder(paramHandle Handle2) Handle2 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)

	var encoder ckks.Encoder
	encoder = ckks.NewEncoder(*params)

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&encoder))
}

func CDoubleVecToGoComplex(realValues *C.constDouble, length uint64) []complex128 {
	complexValues := make([]complex128, length)
	size := unsafe.Sizeof(float64(0))
	basePtr := uintptr(unsafe.Pointer(realValues))
	for i := range complexValues {
		var x float64
		// https://stackoverflow.com/a/32701024/925978
		x = *(*float64)(unsafe.Pointer(basePtr + size*uintptr(i)))
		complexValues[i] = complex(x, 0)
	}
	return complexValues
}

// Take the encoder handle and an array of _real_ numbers of length 2^logLen (checked in C++-land)
// Converts these doubles to complex numbers where the imaginary component is 0, then encode with Lattigo
//
//export lattigo_encode
func lattigo_encode(encoderHandle Handle2, realValues *C.constDouble, logLen uint64, ptHandle Handle2) Handle2 {
	var encoder *ckks.Encoder
	encoder = getStoredEncoder(encoderHandle)

	complexValues := CDoubleVecToGoComplex(realValues, uint64(math.Pow(2, float64(logLen))))

	var plaintext *rlwe.Plaintext
	plaintext = getStoredPlaintext(ptHandle)

	(*encoder).Encode(complexValues, plaintext, int(logLen))
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(plaintext))
}

//export lattigo_encodeNew
func lattigo_encodeNew(encoderHandle Handle2, realValues *C.constDouble, level uint64, scale float64, logLen uint64) Handle2 {
	var encoder *ckks.Encoder
	encoder = getStoredEncoder(encoderHandle)

	complexValues := CDoubleVecToGoComplex(realValues, uint64(math.Pow(2, float64(logLen))))

	var plaintext *rlwe.Plaintext
	plaintext = (*encoder).EncodeNew(complexValues, int(level), rlwe.NewScale(scale), int(logLen))
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(plaintext))
}

//export lattigo_decode
func lattigo_decode(encoderHandle, ptHandle Handle2, logSlots uint64, outValues *C.double) {
	var enc *ckks.Encoder
	enc = getStoredEncoder(encoderHandle)

	var pt *rlwe.Plaintext
	pt = getStoredPlaintext(ptHandle)

	var res []complex128
	res = (*enc).Decode(pt, int(logSlots))

	size := unsafe.Sizeof(float64(0))
	basePtr := uintptr(unsafe.Pointer(outValues))
	for i := range res {
		var x complex128
		x = res[i]
		*(*float64)(unsafe.Pointer(basePtr + size*uintptr(i))) = real(x)
	}
}
