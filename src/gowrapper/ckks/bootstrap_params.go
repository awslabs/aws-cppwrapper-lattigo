// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

import "C"

import (
	"errors"
	"lattigo-cpp/marshal"
	"unsafe"

	"github.com/tuneinsight/lattigo/v4/ckks"
	"github.com/tuneinsight/lattigo/v4/ckks/bootstrapping"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle11 = uint64

func getStoredBootstrappingParameters(bootParamHandle Handle11) *bootstrapping.Parameters {
	ref := marshal.CrossLangObjMap.Get(bootParamHandle)
	return (*bootstrapping.Parameters)(ref.Ptr)
}

//export lattigo_getBootstrappingParams
func lattigo_getBootstrappingParams(bootParamEnum uint8) Handle11 {
	defaultParameters := bootstrapping.DefaultParametersDense

	if int(bootParamEnum) >= len(defaultParameters) {
		panic(errors.New("bootstrapping parameter enum index out of bounds"))
	}

	var bootParams bootstrapping.Parameters
	bootParams = defaultParameters[bootParamEnum].BootstrappingParams
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&bootParams))
}

//export lattigo_ephemeralSecretWeight
func lattigo_ephemeralSecretWeight(bootParamHandle Handle11) uint64 {
	var bootParams *bootstrapping.Parameters
	bootParams = getStoredBootstrappingParameters(bootParamHandle)
	return uint64(bootParams.EphemeralSecretWeight)
}

//export lattigo_params
func lattigo_params(bootParamEnum uint8) Handle11 {
	defaultParameters := bootstrapping.DefaultParametersDense

	if int(bootParamEnum) >= len(defaultParameters) {
		panic(errors.New("bootstrapping parameter enum index out of bounds"))
	}

	paramsSet := defaultParameters[bootParamEnum]
	var params ckks.Parameters
	var err error
	params, err = ckks.NewParametersFromLiteral(paramsSet.SchemeParams)
	if err != nil {
		panic(err)
	}

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&params))
}
