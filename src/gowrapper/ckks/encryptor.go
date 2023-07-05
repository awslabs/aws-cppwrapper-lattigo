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
type Handle3 = uint64

func getStoredEncrypter(encryptorHandle Handle3) *rlwe.Encryptor {
	ref := marshal.CrossLangObjMap.Get(encryptorHandle)
	return (*rlwe.Encryptor)(ref.Ptr)
}

//export lattigo_newEncryptorFromSk
func lattigo_newEncryptorFromSk(paramHandle Handle3, skHandle Handle3) Handle3 {
	params := getStoredParameters(paramHandle)
	sk := getStoredSecretKey(skHandle)
	var encryptor rlwe.Encryptor
	encryptor = ckks.NewEncryptor(*params, sk)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&encryptor))
}

//export lattigo_newEncryptorFromPk
func lattigo_newEncryptorFromPk(paramHandle Handle3, pkHandle Handle3) Handle3 {
	params := getStoredParameters(paramHandle)
	pk := getStoredPublicKey(pkHandle)
	var encryptor rlwe.Encryptor
	encryptor = ckks.NewEncryptor(*params, pk)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&encryptor))
}

//export lattigo_encryptNew
func lattigo_encryptNew(encryptorHandle Handle3, ptHandle Handle3) Handle3 {
	encryptorPtr := getStoredEncrypter(encryptorHandle)
	ptPtr := getStoredPlaintext(ptHandle)
	var ct *rlwe.Ciphertext
	ct = (*encryptorPtr).EncryptNew(ptPtr)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ct))
}

//export lattigo_encryptZeroQP
func lattigo_encryptZeroQP(paramHandle, skHandle, ctxQPHandle Handle3) {
	params := getStoredParameters(paramHandle)
	sk := getStoredSecretKey(skHandle)
	enc := rlwe.NewEncryptor(params.Parameters, sk)
	ctxQP := getStoredCiphertextQP(ctxQPHandle)
	enc.EncryptZero(ctxQP)
}
