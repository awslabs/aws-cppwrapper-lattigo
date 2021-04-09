package ckks

import "C"

import (
	"github.com/ldsec/lattigo/v2/ckks"
	"lattigo-cpp/marshall"
	"unsafe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle6 = uint64

func getStoredParameters(paramHandle Handle6) *ckks.Parameters {
	ref := marshall.CrossLangObjMap.Get(paramHandle)
	return (*ckks.Parameters)(ref.Ptr)
}

//export lattigo_getParams
func lattigo_getParams(paramEnum uint8) Handle6 {
	var params *ckks.Parameters
	params = ckks.DefaultParams[paramEnum]
	return marshall.CrossLangObjMap.Add(unsafe.Pointer(params))
}

//export lattigo_numSlots
func lattigo_numSlots(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return (*params).Slots()
}

//export lattigo_logN
func lattigo_logN(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return (*params).LogN()
}

//export lattigo_logQP
func lattigo_logQP(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return (*params).LogQP()
}

//export lattigo_maxLevel
func lattigo_maxLevel(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return (*params).MaxLevel()
}

//export lattigo_paramsScale
func lattigo_paramsScale(paramHandle Handle6) float64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return (*params).Scale()
}

//export lattigo_sigma
func lattigo_sigma(paramHandle Handle6) float64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return (*params).Sigma()
}

//export lattigo_getQi
func lattigo_getQi(paramHandle Handle6, i uint64) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return params.Moduli().Qi[i]
}

//export lattigo_logQLvl
func lattigo_logQLvl(paramHandle Handle6, i uint64) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)

	return (*params).LogQLvl(i)
}

//export lattigo_logSlots
func lattigo_logSlots(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)

	return (*params).LogSlots()
}
