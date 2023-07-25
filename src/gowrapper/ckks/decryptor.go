// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

import "C"

import (
	"lattigo-cpp/marshal"
	"unsafe"

	"github.com/tuneinsight/lattigo/v4/ckks"
	"github.com/tuneinsight/lattigo/v4/rlwe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle1 = uint64

func getStoredDecryptor(decryptorHandle Handle1) *rlwe.Decryptor {
	ref := marshal.CrossLangObjMap.Get(decryptorHandle)
	return (*rlwe.Decryptor)(ref.Ptr)
}

//export lattigo_newDecryptor
func lattigo_newDecryptor(paramHandle Handle1, skHandle Handle1) Handle1 {
	params := getStoredParameters(paramHandle)
	sk := getStoredSecretKey(skHandle)
	var decryptor rlwe.Decryptor
	decryptor = ckks.NewDecryptor(*params, sk)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&decryptor))
}

//export lattigo_decryptNew
func lattigo_decryptNew(decryptorHandle Handle1, ctHandle Handle1) Handle1 {
	var dec *rlwe.Decryptor
	dec = getStoredDecryptor(decryptorHandle)

	var ct *rlwe.Ciphertext
	ct = getStoredCiphertext(ctHandle)

	var pt *rlwe.Plaintext
	pt = (*dec).DecryptNew(ct)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(pt))
}
