package ckks

import "C"

import (
    "unsafe"
    "github.com/ldsec/lattigo/v2/ckks"
    "lattigo-cpp/marshall"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle1 = uint64

func getStoredDecryptor(decryptoHandle Handle8) *ckks.Decryptor {
    ref := marshall.CrossLangObjMap.Get(decryptoHandle)
    return (*ckks.Decryptor)(ref.Ptr)
}

//export lattigo_newDecryptor
func lattigo_newDecryptor(paramHandle Handle1, skHandle Handle1) Handle1 {
    params := getStoredParameters(paramHandle)
    sk := getStoredSecretKey(skHandle)
    var decryptor ckks.Decryptor
    decryptor = ckks.NewDecryptor(params, sk)
    return marshall.CrossLangObjMap.Add(unsafe.Pointer(&decryptor))
}

//export lattigo_decryptNew
func lattigo_decryptNew(decryptoHandle Handle1, ctHandle Handle1) Handle1 {
	var dec *ckks.Decryptor
    dec = getStoredDecryptor(decryptoHandle)

    var ct *ckks.Ciphertext
    ct = getStoredCiphertext(ctHandle)

    var pt *ckks.Plaintext
    pt = (*dec).DecryptNew(ct)
    return marshall.CrossLangObjMap.Add(unsafe.Pointer(pt))
}
