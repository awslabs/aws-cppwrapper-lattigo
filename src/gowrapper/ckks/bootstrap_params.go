// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

import "C"

import (
	"errors"
	"lattigo-cpp/marshal"
	"unsafe"

	"github.com/ldsec/lattigo/v2/ckks"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle11 = uint64

func getStoredBootstrappingParameters(bootParamHandle Handle11) *ckks.BootstrappingParameters {
	ref := marshal.CrossLangObjMap.Get(bootParamHandle)
	return (*ckks.BootstrappingParameters)(ref.Ptr)
}

//export lattigo_getBootstrappingParams
func lattigo_getBootstrappingParams(bootParamEnum uint8) Handle11 {
	if int(bootParamEnum) >= len(ckks.DefaultBootstrapParams) {
		panic(errors.New("bootstrapping parameter enum index out of bounds"))
	}

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
	// level is 10 and the post-bootstrapping *level* is 5, then the
	// length of the residual moduli vector is 6, so the depth of the bootstrapping
	// circuit is 10 - 6 + 1 = 5.
	return uint64(bootParams.MaxLevel() - len(bootParams.ResidualModuli) + 1)
}

//export lattigo_params
func lattigo_params(bootParamHandle Handle11) Handle11 {
	var bootParams *ckks.BootstrappingParameters
	bootParams = getStoredBootstrappingParameters(bootParamHandle)

	var params ckks.Parameters
	var err error
	params, err = bootParams.Params()

	if err != nil {
		panic(err)
	}

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&params))
}
