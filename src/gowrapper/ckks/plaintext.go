package ckks

import "C"

import (
    "github.com/ldsec/lattigo/v2/ckks"
    "lattigo-cpp/marshall"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle7 = uint64

func getStoredPlaintext(ptHandle uint64) *ckks.Plaintext {
    ref := marshall.CrossLangObjMap.Get(ptHandle)
    return (*ckks.Plaintext)(ref.Ptr)
}
