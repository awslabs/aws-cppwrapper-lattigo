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
	return uint64(bootParams.H)
}

//export lattigo_bootstrap_depth
func lattigo_bootstrap_depth(bootParamHandle Handle11) uint64 {
	var bootParams *ckks.BootstrappingParameters
	bootParams = getStoredBootstrappingParameters(bootParamHandle)
	// len(bootParams.ResidualModuli) is the number of moduli available
	// post-bootstrapping, which is one more than the ciphertext level
	// after bootstrapping. Thus the difference, plus one, is the depth of
	// the bootstrapping circuit. For example, if the highest ciphertext
	// level is 10 and the post-bootstrapping *level* is 10, then the
	// length of the residual moduli vector is 6, so 5 = 10 - 6 + 1.
	return uint64(bootParams.MaxLevel() - len(bootParams.ResidualModuli) + 1)
}
