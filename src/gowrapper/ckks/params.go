package ckks

import "C"

import (
	"github.com/ldsec/lattigo/v2/ckks"
	"lattigo-cpp/marshal"
	"unsafe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle6 = uint64

func getStoredParameters(paramHandle Handle6) *ckks.Parameters {
	ref := marshal.CrossLangObjMap.Get(paramHandle)
	return (*ckks.Parameters)(ref.Ptr)
}

//export lattigo_getParams
func lattigo_getParams(paramEnum uint8) Handle6 {
	var params *ckks.Parameters
	params = ckks.DefaultParams[paramEnum]
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(params))
}

//export lattigo_newParametersFromLogModuli
func lattigo_newParametersFromLogModuli(logN uint64, logQi *C.uchar, numQi uint8, logPi *C.uchar, numPi uint8, logScale uint8) Handle6 {
	size := unsafe.Sizeof(uint8(0))

	LogQi := make([]uint64, numQi)
	qiPtr := uintptr(unsafe.Pointer(logQi))
	for i := range LogQi {
		LogQi[i] = uint64(*(*uint8)(unsafe.Pointer(qiPtr + size*uintptr(i))))
	}

	LogPi := make([]uint64, numPi)
	piPtr := uintptr(unsafe.Pointer(logPi))
	for i := range LogPi {
		LogPi[i] = uint64(*(*uint8)(unsafe.Pointer(piPtr + size*uintptr(i))))
	}

	var lm ckks.LogModuli
	lm = ckks.LogModuli{LogQi, LogPi}

	var params *ckks.Parameters
	var err error
	params, err = ckks.NewParametersFromLogModuli(logN, &lm)
	if err != nil {
		panic(err)
	}
	params.SetLogSlots(logN - 1)
	params.SetScale(float64(uint64(1) << uint64(logScale)))
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(params))
}

//export lattigo_numSlots
func lattigo_numSlots(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return params.Slots()
}

//export lattigo_logN
func lattigo_logN(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return params.LogN()
}

//export lattigo_logQP
func lattigo_logQP(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return params.LogQP()
}

//export lattigo_maxLevel
func lattigo_maxLevel(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return params.MaxLevel()
}

//export lattigo_paramsScale
func lattigo_paramsScale(paramHandle Handle6) float64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return params.Scale()
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
	return params.Qi()[i]
}

//export lattigo_getPi
func lattigo_getPi(paramHandle Handle6, i uint64) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return params.Pi()[i]
}

//export lattigo_qiCount
func lattigo_qiCount(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return params.QiCount()
}

//export lattigo_piCount
func lattigo_piCount(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return params.PiCount()
}

//export lattigo_logQLvl
func lattigo_logQLvl(paramHandle Handle6, i uint64) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return params.LogQLvl(i)
}

//export lattigo_logSlots
func lattigo_logSlots(paramHandle Handle6) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)
	return params.LogSlots()
}
