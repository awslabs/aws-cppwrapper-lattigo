// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

import "C"

import (
	"errors"
	"github.com/ldsec/lattigo/v2/ckks"
	"github.com/ldsec/lattigo/v2/ckks/bootstrapping"
	"lattigo-cpp/marshal"
	"unsafe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle11 = uint64

func getStoredBootstrappingParameters(bootParamHandle Handle11) *bootstrapping.Parameters {
	ref := marshal.CrossLangObjMap.Get(bootParamHandle)
	return (*bootstrapping.Parameters)(ref.Ptr)
}

//export lattigo_getBootstrappingParams
func lattigo_getBootstrappingParams(bootParamEnum uint8) Handle11 {
	if int(bootParamEnum) >= len(bootstrapping.DefaultParameters) {
		panic(errors.New("bootstrapping parameter enum index out of bounds"))
	}

	var bootParams *bootstrapping.Parameters
	bootParams = &bootstrapping.DefaultParameters[bootParamEnum]

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(bootParams))
}

//export lattigo_bootstrap_h
func lattigo_bootstrap_h(bootParamHandle Handle11) uint64 {
	var bootParams *bootstrapping.Parameters
	bootParams = getStoredBootstrappingParameters(bootParamHandle)
	return uint64(bootParams.H)
}

//export lattigo_bootstrap_depth
func lattigo_bootstrap_depth(ckksParamHandle Handle11, bootParamHandle Handle11) uint64 {
	var ckksParams *ckks.Parameters
	ckksParams = getStoredParameters(ckksParamHandle)

	var bootParams *bootstrapping.Parameters
	bootParams = getStoredBootstrappingParameters(bootParamHandle)
	// SlotsToCoeffsParameters.LevelStart is the starting level of the last step of bootstrapping,
	// and len(SlotsToCoeffsParameters.ScalingFactor) is the number of levels consumed by this step.
	// Thus the post-bootstrapping level is `SlotsToCoeffsParameters.LevelStart - len(SlotsToCoeffsParameters.ScalingFactor)`.
	// The difference between ckksParams.MaxLevel() and the post-bootstrapping level is therefore the number of levels
	// consumed by bootstrapping.
	// For example, if the highest ciphertext level is 10 and the SlotsToCoeffsParams.LevelStart is 7 and the length
	// of the ScalingFactor array is 2, then the post-bootstrapping *level* is 5, so the depth of the bootstrapping
	// circuit is 10 - 7 + 2 = 5.
	return uint64(ckksParams.MaxLevel() - bootParams.SlotsToCoeffsParameters.LevelStart + len(bootParams.SlotsToCoeffsParameters.ScalingFactor))
}

//export lattigo_getCKKSParamsForBootstrapping
func lattigo_getCKKSParamsForBootstrapping(bootParamEnum uint8) Handle11 {
	if int(bootParamEnum) >= len(bootstrapping.DefaultCKKSParameters) {
		panic(errors.New("bootstrapping parameter enum index out of bounds"))
	}
	var paramsLit ckks.ParametersLiteral
	paramsLit = bootstrapping.DefaultCKKSParameters[bootParamEnum]

	var params ckks.Parameters
	var err error
	params, err = ckks.NewParametersFromLiteral(paramsLit)

	if err != nil {
		panic(err)
	}

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&params))
}
