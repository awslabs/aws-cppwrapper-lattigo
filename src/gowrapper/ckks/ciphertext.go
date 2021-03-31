package ckks

import "C"

import (
    "unsafe"
    "github.com/ldsec/lattigo/v2/ckks"
    "lattigo-cpp/marshall"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle8 = uint64

func getStoredCiphertext(ctHandle Handle8) *ckks.Ciphertext {
    ref := marshall.CrossLangObjMap.Get(ctHandle)
    return (*ckks.Ciphertext)(ref.Ptr)
}

//export lattigo_level
func lattigo_level(ctHandle Handle8) uint64 {
	var ctIn *ckks.Ciphertext
    ctIn = getStoredCiphertext(ctHandle)
    return ctIn.Level()
}

//export lattigo_ciphertextScale
func lattigo_ciphertextScale(ctHandle Handle8) float64 {
    var ctIn *ckks.Ciphertext
    ctIn = getStoredCiphertext(ctHandle)
    return ctIn.Scale()
}

//export lattigo_copyNew
func lattigo_copyNew(ctHandle Handle8) Handle8 {
	var ctIn *ckks.Ciphertext
    ctIn = getStoredCiphertext(ctHandle)

    var ctClone *ckks.Ciphertext
    ctClone = ctIn.CopyNew().Ciphertext();
    return marshall.CrossLangObjMap.Add(unsafe.Pointer(ctClone))
}

//export lattigo_newCiphertext
func lattigo_newCiphertext(paramsHandle Handle8, degree uint64, level uint64, scale float64) Handle8 {
    var params *ckks.Parameters
    params = getStoredParameters(paramsHandle)

    var newCt *ckks.Ciphertext
    newCt = ckks.NewCiphertext(params, degree, level, scale)
    return marshall.CrossLangObjMap.Add(unsafe.Pointer(newCt))
}
