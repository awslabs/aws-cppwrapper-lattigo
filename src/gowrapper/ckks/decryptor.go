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
type Handle1 = uint64

func getStoredDecryptor(decryptorHandle Handle1) *ckks.Decryptor {
	ref := marshal.CrossLangObjMap.Get(decryptorHandle)
	return (*ckks.Decryptor)(ref.Ptr)
}

//export lattigo_newDecryptor
func lattigo_newDecryptor(paramHandle Handle1, skHandle Handle1) Handle1 {
	params := getStoredParameters(paramHandle)
	sk := getStoredSecretKey(skHandle)
	var decryptor ckks.Decryptor
	decryptor = ckks.NewDecryptor(*params, sk)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&decryptor))
}

//export lattigo_decryptNew
func lattigo_decryptNew(decryptorHandle Handle1, ctHandle Handle1) Handle1 {
	var dec *ckks.Decryptor
	dec = getStoredDecryptor(decryptorHandle)

	var ct *ckks.Ciphertext
	ct = getStoredCiphertext(ctHandle)

	var pt *ckks.Plaintext
	pt = (*dec).DecryptNew(ct)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(pt))
}
