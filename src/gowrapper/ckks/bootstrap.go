// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

import "C"

import (
	"lattigo-cpp/marshal"
	"unsafe"

	"github.com/ldsec/lattigo/v2/ckks"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle10 = uint64

func getStoredBootstrapper(btpHandle Handle10) *ckks.Bootstrapper {
	ref := marshal.CrossLangObjMap.Get(btpHandle)
	return (*ckks.Bootstrapper)(ref.Ptr)
}

//export lattigo_newBootstrapper
func lattigo_newBootstrapper(paramHandle Handle10, btpParamHandle Handle10, btpKeyHandle Handle10) Handle10 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)

	var btpParams *ckks.BootstrappingParameters
	btpParams = getStoredBootstrappingParameters(btpParamHandle)

	var btpKey *ckks.BootstrappingKey
	btpKey = getStoredBootstrappingKey(btpKeyHandle)

	var btp *ckks.Bootstrapper
	var err error
	btp, err = ckks.NewBootstrapper(*params, btpParams, *btpKey)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(btp))
}

//export lattigo_bootstrap
func lattigo_bootstrap(btpHandle Handle10, ctHandle Handle10) Handle10 {
	var btp *ckks.Bootstrapper
	btp = getStoredBootstrapper(btpHandle)

	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctHandle)

	var ctOut *ckks.Ciphertext
	ctOut = btp.Bootstrapp(ctIn)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ctOut))
}
