// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

import "C"

import (
	"lattigo-cpp/marshal"
	"unsafe"

	"github.com/tuneinsight/lattigo/v4/ckks"
	"github.com/tuneinsight/lattigo/v4/ckks/bootstrapping"
	"github.com/tuneinsight/lattigo/v4/rlwe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle10 = uint64

func getStoredBootstrapper(btpHandle Handle10) *bootstrapping.Bootstrapper {
	ref := marshal.CrossLangObjMap.Get(btpHandle)
	return (*bootstrapping.Bootstrapper)(ref.Ptr)
}

//export lattigo_newBootstrapper
func lattigo_newBootstrapper(paramHandle Handle10, btpParamHandle Handle10, btpKeyHandle Handle10) Handle10 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)

	var btpParams *bootstrapping.Parameters
	btpParams = getStoredBootstrappingParameters(btpParamHandle)

	var btpKey *bootstrapping.EvaluationKeys
	btpKey = getStoredBootstrappingKey(btpKeyHandle)

	var btp *bootstrapping.Bootstrapper
	var err error
	btp, err = bootstrapping.NewBootstrapper(*params, *btpParams, *btpKey)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(btp))
}

//export lattigo_bootstrap
func lattigo_bootstrap(btpHandle Handle10, ctHandle Handle10) Handle10 {
	var btp *bootstrapping.Bootstrapper
	btp = getStoredBootstrapper(btpHandle)

	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctHandle)

	var ctOut *rlwe.Ciphertext
	ctOut = btp.Bootstrap(ctIn)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ctOut))
}
