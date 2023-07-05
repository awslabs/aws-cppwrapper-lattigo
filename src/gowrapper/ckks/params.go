// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

/*
#include "stdint.h"
typedef const uint8_t constUChar;
typedef const uint64_t constULong;
*/
import "C"

import (
	"errors"
	"lattigo-cpp/marshal"
	"unsafe"

	"github.com/tuneinsight/lattigo/v4/ckks"
	"github.com/tuneinsight/lattigo/v4/ring"
	"github.com/tuneinsight/lattigo/v4/rlwe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle6 = uint64

func getStoredParameters(paramHandle Handle6) *ckks.Parameters {
	ref := marshal.CrossLangObjMap.Get(paramHandle)
	return (*ckks.Parameters)(ref.Ptr)
}

//export lattigo_getDefaultClassicalParams
func lattigo_getDefaultClassicalParams(paramEnum uint8) Handle6 {
	if int(paramEnum) >= len(ckks.DefaultParams) {
		panic(errors.New("classical parameter enum index out of bounds"))
	}

	var paramsLit ckks.ParametersLiteral
	paramsLit = ckks.DefaultParams[paramEnum]

	var params ckks.Parameters
	var err error
	params, err = ckks.NewParametersFromLiteral(paramsLit)
	if err != nil {
		panic(err)
	}

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&params))
}

//export lattigo_getDefaultPQParams
func lattigo_getDefaultPQParams(paramEnum uint8) Handle6 {
	if int(paramEnum) >= len(ckks.DefaultPostQuantumParams) {
		panic(errors.New("quantum parameter enum index out of bounds"))
	}

	var paramsLit ckks.ParametersLiteral
	paramsLit = ckks.DefaultPostQuantumParams[paramEnum]

	var params ckks.Parameters
	var err error
	params, err = ckks.NewParametersFromLiteral(paramsLit)
	if err != nil {
		panic(err)
	}

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&params))
}

//export lattigo_newParameters
func lattigo_newParameters(logN uint64, qi *C.constULong, numQi uint8, pi *C.constULong, numPi uint8, logScale uint8) Handle6 {
	size := unsafe.Sizeof(uint64(0))

	Qi := make([]uint64, numQi)
	qiPtr := uintptr(unsafe.Pointer(qi))
	for i := range Qi {
		Qi[i] = *(*uint64)(unsafe.Pointer(qiPtr + size*uintptr(i)))
	}

	Pi := make([]uint64, numPi)
	piPtr := uintptr(unsafe.Pointer(pi))
	for i := range Pi {
		// https://stackoverflow.com/a/32701024/925978
		Pi[i] = *(*uint64)(unsafe.Pointer(piPtr + size*uintptr(i)))
	}

	var rlweParams rlwe.Parameters
	var err error
	scale := rlwe.NewScale(float64(uint64(1) << uint64(logScale)))
	rlweParams, err = rlwe.NewParameters(int(logN), Qi, Pi, 0, 192, rlwe.DefaultSigma, ring.Standard, scale, ckks.DefaultNTTFlag)
	if err != nil {
		panic(err)
	}

	var params ckks.Parameters
	params, err = ckks.NewParameters(rlweParams, int(logN-1))
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&params))
}

//export lattigo_newParametersFromLogModuli
func lattigo_newParametersFromLogModuli(logN uint64, logQi *C.constUChar, numQi uint8, logPi *C.constUChar, numPi uint8, logScale uint8) Handle6 {
	size := unsafe.Sizeof(uint8(0))

	LogQi := make([]int, numQi)
	qiPtr := uintptr(unsafe.Pointer(logQi))
	for i := range LogQi {
		LogQi[i] = int(*(*uint8)(unsafe.Pointer(qiPtr + size*uintptr(i))))
	}

	LogPi := make([]int, numPi)
	piPtr := uintptr(unsafe.Pointer(logPi))
	for i := range LogPi {
		LogPi[i] = int(*(*uint8)(unsafe.Pointer(piPtr + size*uintptr(i))))
	}

	var paramsLit ckks.ParametersLiteral
	paramsLit = ckks.ParametersLiteral{
		LogN:         int(logN),
		LogQ:         LogQi,
		LogP:         LogPi,
		Sigma:        rlwe.DefaultSigma,
		LogSlots:     int(logN) - 1,
		DefaultScale: float64(uint64(1) << uint64(logScale)),
	}

	var params ckks.Parameters
	var err error
	params, err = ckks.NewParametersFromLiteral(paramsLit)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&params))
}

//export lattigo_numSlots
func lattigo_numSlots(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return uint64(params.Slots())
}

//export lattigo_logN
func lattigo_logN(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return uint64(params.LogN())
}

//export lattigo_ringQ
func lattigo_ringQ(paramHandle Handle6) Handle6 {
	params := getStoredParameters(paramHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(params.RingQ()))
}

//export lattigo_noiseBound
func lattigo_noiseBound(paramHandle Handle6) uint64 {
	params := getStoredParameters(paramHandle)
	return uint64(params.NoiseBound())
}

//export lattigo_ringP
func lattigo_ringP(paramHandle Handle6) Handle6 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(params.RingP()))
}

//export lattigo_ringQP
func lattigo_ringQP(paramHandle Handle6) Handle6 {
	params := getStoredParameters(paramHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(params.RingQP()))
}

//export lattigo_logQP
func lattigo_logQP(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return uint64(params.LogQP())
}

//export lattigo_maxLevel
func lattigo_maxLevel(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return uint64(params.MaxLevel())
}

//export lattigo_paramsScale
func lattigo_paramsScale(paramHandle Handle6) float64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	scale := params.DefaultScale()
	return scale.Float64()
}

//export lattigo_sigma
func lattigo_sigma(paramHandle Handle6) float64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return params.Sigma()
}

//export lattigo_getQi
func lattigo_getQi(paramHandle Handle6, i uint64) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return params.Q()[i]
}

//export lattigo_getPi
func lattigo_getPi(paramHandle Handle6, i uint64) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return params.P()[i]
}

//export lattigo_qiCount
func lattigo_qiCount(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return uint64(params.QCount())
}

//export lattigo_piCount
func lattigo_piCount(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return uint64(params.PCount())
}

//export lattigo_logQLvl
func lattigo_logQLvl(paramHandle Handle6, i uint64) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return uint64(params.LogQLvl(int(i)))
}

//export lattigo_logSlots
func lattigo_logSlots(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return uint64(params.LogSlots())
}

//export lattigo_galoisElementForRowRotation
func lattigo_galoisElementForRowRotation(paramHandle Handle6) uint64 {
	params := getStoredParameters(paramHandle)
	return params.GaloisElementForRowRotation()
}

//export lattigo_galoisElementsForRowInnerSum
func lattigo_galoisElementsForRowInnerSum(paramHandle Handle6, outValues *C.uint64_t) {
	params := getStoredParameters(paramHandle)
	res := params.GaloisElementsForRowInnerSum()

	size := unsafe.Sizeof(uint64(0))
	basePtr := uintptr(unsafe.Pointer(outValues))
	for i := range res {
		*(*uint64)(unsafe.Pointer(basePtr + size*uintptr(i))) = res[i]
	}
}

//export lattigo_inverseGaloisElement
func lattigo_inverseGaloisElement(paramHandle Handle6, galEl uint64) uint64 {
	params := getStoredParameters(paramHandle)
	return params.InverseGaloisElement(galEl)
}

//export lattigo_rotationFromGaloisElement
func lattigo_rotationFromGaloisElement(paramHandle Handle6, galEl uint64) uint64 {
	params := getStoredParameters(paramHandle)
	return params.RotationFromGaloisElement(galEl)
}
