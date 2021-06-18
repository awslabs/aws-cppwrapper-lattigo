package ckks

import "C"

import (
	"github.com/ldsec/lattigo/v2/ckks"
	"lattigo-cpp/marshal"
	"unsafe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle11 = uint64

func getStoredBootstrappingParameters(bootParamHandle Handle11) *ckks.BootstrappingParameters {
	ref := marshal.CrossLangObjMap.Get(bootParamHandle)
	return (*ckks.BootstrappingParameters)(ref.Ptr)
}

//export lattigo_getBootstrappingParams
func lattigo_getBootstrappingParams(bootParamEnum uint8) Handle11 {
	var bootParams *ckks.BootstrappingParameters

	bootParams = ckks.DefaultBootstrapParams[bootParamEnum]

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(bootParams))
}

//export lattigo_bootstrap_h
func lattigo_bootstrap_h(bootParamHandle Handle11) uint64 {
	var bootParams *ckks.BootstrappingParameters
	bootParams = getStoredBootstrappingParameters(bootParamHandle)
	return bootParams.H
}
