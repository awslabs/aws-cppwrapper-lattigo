package ckks

import "C"

import (
    "unsafe"
    "github.com/ldsec/lattigo/v2/ckks"
    "lattigo-cpp/marshall"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle3 = uint64

func getStoredEncrypter(encryptorHandle Handle3) *ckks.Encryptor {
    ref := marshall.CrossLangObjMap.Get(encryptorHandle)
    return (*ckks.Encryptor)(ref.Ptr)
}

//export lattigo_newEncryptorFromPk
func lattigo_newEncryptorFromPk(paramHandle Handle3, pkHandle Handle3) Handle3 {
    params := getStoredParameters(paramHandle)
    pk := getStoredPublicKey(pkHandle)
    var encryptor ckks.Encryptor
    encryptor = ckks.NewEncryptorFromPk(params, pk)
    return marshall.CrossLangObjMap.Add(unsafe.Pointer(&encryptor))
}

//export lattigo_encryptNew
func lattigo_encryptNew(encryptorHandle Handle3, ptHandle Handle3) Handle3 {
    encryptorPtr := getStoredEncrypter(encryptorHandle)
    ptPtr := getStoredPlaintext(ptHandle)
    var ct *ckks.Ciphertext
    ct = (*encryptorPtr).EncryptNew(ptPtr)
    return marshall.CrossLangObjMap.Add(unsafe.Pointer(ct))
}
