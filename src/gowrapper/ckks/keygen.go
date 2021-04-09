package ckks

// cgo will automatically generate a struct for functions which return multiple values,
// but the auto-generated struct with generated names loses its semantic value. We opt
// to define our own struct here.

// #include "stdint.h"
// struct Lattigo_KeyPairHandle {
//   uint64_t sk;
//   uint64_t pk;
// };
import "C"

import (
	"github.com/ldsec/lattigo/v2/ckks"
	"lattigo-cpp/marshall"
	"unsafe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle5 = uint64

func getStoredKeyGenerator(keygenHandle Handle5) *ckks.KeyGenerator {
	ref := marshall.CrossLangObjMap.Get(keygenHandle)
	return (*ckks.KeyGenerator)(ref.Ptr)
}

func getStoredSecretKey(skHandle Handle5) *ckks.SecretKey {
	ref := marshall.CrossLangObjMap.Get(skHandle)
	return (*ckks.SecretKey)(ref.Ptr)
}

func getStoredPublicKey(pkHandle Handle5) *ckks.PublicKey {
	ref := marshall.CrossLangObjMap.Get(pkHandle)
	return (*ckks.PublicKey)(ref.Ptr)
}

func getStoredEvalKey(evalKeyHandle Handle4) *ckks.EvaluationKey {
	ref := marshall.CrossLangObjMap.Get(evalKeyHandle)
	return (*ckks.EvaluationKey)(ref.Ptr)
}

//export lattigo_newKeyGenerator
func lattigo_newKeyGenerator(paramHandle Handle5) Handle5 {
	paramPtr := getStoredParameters(paramHandle)
	var keyGenerator ckks.KeyGenerator
	keyGenerator = ckks.NewKeyGenerator(paramPtr)
	return marshall.CrossLangObjMap.Add(unsafe.Pointer(&keyGenerator))
}

//export lattigo_genKeyPair
func lattigo_genKeyPair(keygenHandle Handle5) C.struct_Lattigo_KeyPairHandle {
	keygen := getStoredKeyGenerator(keygenHandle)
	var sk *ckks.SecretKey
	var pk *ckks.PublicKey
	sk, pk = (*keygen).GenKeyPair()
	var kpHandle C.struct_Lattigo_KeyPairHandle
	kpHandle.sk = C.uint64_t(marshall.CrossLangObjMap.Add(unsafe.Pointer(sk)))
	kpHandle.pk = C.uint64_t(marshall.CrossLangObjMap.Add(unsafe.Pointer(pk)))
	return kpHandle
}

//export lattigo_genRelinKey
func lattigo_genRelinKey(keygenHandle Handle5, skHandle Handle5) Handle5 {
	keygen := getStoredKeyGenerator(keygenHandle)
	var sk *ckks.SecretKey
	sk = getStoredSecretKey(skHandle)
	return marshall.CrossLangObjMap.Add(unsafe.Pointer((*keygen).GenRelinKey(sk)))
}
