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

//export lattigo_bootstrap_depth
func lattigo_bootstrap_depth(bootParamHandle Handle11) uint64 {
	var bootParams *ckks.BootstrappingParameters
	bootParams = getStoredBootstrappingParameters(bootParamHandle)
	// bootParams.CtSLevel[0] is the highest ciphertext level (this is enforced at runtime in Lattigo)
	// and therefore the first level of the bootstrapping circuit.
	// By contrast, bootParams.StCLevel[len(bootParams.StCLevel)-1] is the last level of Slots-to-Coeffs
	// step, or the *second to last* bootstrapping level.
	// Thus, the difference plus one is the depth of the bootstrapping circuit. For example,
	// if the first level of bootstrapping is 10 and the last StC level is 6, then a ciphertext after
	// bootstrapping is at level 5, or 10 - 6 + 1.
	return bootParams.CtSLevel[0] - bootParams.StCLevel[len(bootParams.StCLevel)-1] + 1
}
