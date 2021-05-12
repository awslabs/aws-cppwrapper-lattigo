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
	"lattigo-cpp/marshal"
	"unsafe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle5 = uint64

func getStoredKeyGenerator(keygenHandle Handle5) *ckks.KeyGenerator {
	ref := marshal.CrossLangObjMap.Get(keygenHandle)
	return (*ckks.KeyGenerator)(ref.Ptr)
}

func getStoredSecretKey(skHandle Handle5) *ckks.SecretKey {
	ref := marshal.CrossLangObjMap.Get(skHandle)
	return (*ckks.SecretKey)(ref.Ptr)
}

func getStoredPublicKey(pkHandle Handle5) *ckks.PublicKey {
	ref := marshal.CrossLangObjMap.Get(pkHandle)
	return (*ckks.PublicKey)(ref.Ptr)
}

func getStoredEvalKey(evalKeyHandle Handle4) *ckks.EvaluationKey {
	ref := marshal.CrossLangObjMap.Get(evalKeyHandle)
	return (*ckks.EvaluationKey)(ref.Ptr)
}

func getStoredRotationKeys(rotKeysHandle Handle4) *ckks.RotationKeys {
	ref := marshal.CrossLangObjMap.Get(rotKeysHandle)
	return (*ckks.RotationKeys)(ref.Ptr)
}

//export lattigo_newKeyGenerator
func lattigo_newKeyGenerator(paramHandle Handle5) Handle5 {
	paramPtr := getStoredParameters(paramHandle)
	var keyGenerator ckks.KeyGenerator
	keyGenerator = ckks.NewKeyGenerator(paramPtr)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&keyGenerator))
}

//export lattigo_genKeyPair
func lattigo_genKeyPair(keygenHandle Handle5) C.struct_Lattigo_KeyPairHandle {
	var keygen *ckks.KeyGenerator
	keygen = getStoredKeyGenerator(keygenHandle)
	var sk *ckks.SecretKey
	var pk *ckks.PublicKey
	sk, pk = (*keygen).GenKeyPair()
	var kpHandle C.struct_Lattigo_KeyPairHandle
	kpHandle.sk = C.uint64_t(marshal.CrossLangObjMap.Add(unsafe.Pointer(sk)))
	kpHandle.pk = C.uint64_t(marshal.CrossLangObjMap.Add(unsafe.Pointer(pk)))
	return kpHandle
}

//export lattigo_genRelinKey
func lattigo_genRelinKey(keygenHandle Handle5, skHandle Handle5) Handle5 {
	var keygen *ckks.KeyGenerator
	keygen = getStoredKeyGenerator(keygenHandle)
	var sk *ckks.SecretKey
	sk = getStoredSecretKey(skHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer((*keygen).GenRelinKey(sk)))
}

//export lattigo_genRotationKeysPow2
func lattigo_genRotationKeysPow2(keygenHandle Handle5, skHandle Handle5) Handle5 {
	var keygen *ckks.KeyGenerator
	keygen = getStoredKeyGenerator(keygenHandle)
	var sk *ckks.SecretKey
	sk = getStoredSecretKey(skHandle)
	var rotKeys *ckks.RotationKeys
	rotKeys = (*keygen).GenRotationKeysPow2(sk)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(rotKeys))
}
