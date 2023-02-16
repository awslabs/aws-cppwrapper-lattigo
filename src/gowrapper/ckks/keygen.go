// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

// cgo will automatically generate a struct for functions which return multiple values,
// but the auto-generated struct with generated names loses its semantic value. We opt
// to define our own struct here.

/*
#include "stdint.h"
struct Lattigo_KeyPairHandle {
  uint64_t sk;
  uint64_t pk;
};
*/
import "C"

import (
	"errors"
	"lattigo-cpp/marshal"
	"unsafe"

	"github.com/ldsec/lattigo/v2/ckks"
	"github.com/ldsec/lattigo/v2/rlwe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle5 = uint64

func getStoredKeyGenerator(keygenHandle Handle5) *rlwe.KeyGenerator {
	ref := marshal.CrossLangObjMap.Get(keygenHandle)
	return (*rlwe.KeyGenerator)(ref.Ptr)
}

func getStoredSecretKey(skHandle Handle5) *rlwe.SecretKey {
	ref := marshal.CrossLangObjMap.Get(skHandle)
	return (*rlwe.SecretKey)(ref.Ptr)
}

func getStoredPublicKey(pkHandle Handle5) *rlwe.PublicKey {
	ref := marshal.CrossLangObjMap.Get(pkHandle)
	return (*rlwe.PublicKey)(ref.Ptr)
}

func getStoredEvaluationKey(evalKeyHandle Handle5) *rlwe.EvaluationKey {
	ref := marshal.CrossLangObjMap.Get(evalKeyHandle)
	return (*rlwe.EvaluationKey)(ref.Ptr)
}

func getStoredRelinKey(relinKeyHandle Handle5) *rlwe.RelinearizationKey {
	ref := marshal.CrossLangObjMap.Get(relinKeyHandle)
	return (*rlwe.RelinearizationKey)(ref.Ptr)
}

func getStoredRotationKeys(rotKeysHandle Handle5) *rlwe.RotationKeySet {
	ref := marshal.CrossLangObjMap.Get(rotKeysHandle)
	return (*rlwe.RotationKeySet)(ref.Ptr)
}

func getStoredSwitchingKey(switchingKeyHandle Handle5) *rlwe.SwitchingKey {
	ref := marshal.CrossLangObjMap.Get(switchingKeyHandle)
	return (*rlwe.SwitchingKey)(ref.Ptr)
}

func getStoredBootstrappingKey(bootKeyHandle Handle5) *ckks.BootstrappingKey {
	ref := marshal.CrossLangObjMap.Get(bootKeyHandle)
	return (*ckks.BootstrappingKey)(ref.Ptr)
}

//export lattigo_newKeyGenerator
func lattigo_newKeyGenerator(paramHandle Handle5) Handle5 {
	paramPtr := getStoredParameters(paramHandle)
	var keyGenerator rlwe.KeyGenerator
	keyGenerator = ckks.NewKeyGenerator(*paramPtr)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&keyGenerator))
}

//export lattigo_newSecretKey
func lattigo_newSecretKey(paramHandle Handle5) Handle5 {
	paramPtr := getStoredParameters(paramHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ckks.NewSecretKey(*paramPtr)))
}

//export lattigo_newPublicKey
func lattigo_newPublicKey(paramHandle Handle5) Handle5 {
	paramPtr := getStoredParameters(paramHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ckks.NewPublicKey(*paramPtr)))
}

//export lattigo_newRelinearizationKey
func lattigo_newRelinearizationKey(paramHandle Handle5) Handle5 {
	paramPtr := getStoredParameters(paramHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ckks.NewRelinearizationKey(*paramPtr)))
}

//export lattigo_newRotationKeys
func lattigo_newRotationKeys(paramHandle Handle5, galoisElements *C.uint64_t, galoisElementsLen uint64) Handle5 {
	paramPtr := getStoredParameters(paramHandle)

	galoisElementsTmp := make([]uint64, galoisElementsLen)
	size := unsafe.Sizeof(uint64(0))
	basePtrIn := uintptr(unsafe.Pointer(galoisElements))
	for i := range galoisElementsTmp {
		galoisElementsTmp[i] = *(*uint64)(unsafe.Pointer(basePtrIn + size*uintptr(i)))
	}

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ckks.NewRotationKeySet(*paramPtr, galoisElementsTmp)))
}

//export lattigo_genSecretKey
func lattigo_genSecretKey(keygenHandle Handle5) Handle5 {
	keygen := getStoredKeyGenerator(keygenHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer((*keygen).GenSecretKey()))
}

//export lattigo_genPublicKey
func lattigo_genPublicKey(keygenHandle Handle5, skHandle Handle5) Handle5 {
	keygen := getStoredKeyGenerator(keygenHandle)
	sk := getStoredSecretKey(skHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer((*keygen).GenPublicKey(sk)))
}

//export lattigo_genKeyPair
func lattigo_genKeyPair(keygenHandle Handle5) C.struct_Lattigo_KeyPairHandle {
	var keygen *rlwe.KeyGenerator
	keygen = getStoredKeyGenerator(keygenHandle)
	var sk *rlwe.SecretKey
	var pk *rlwe.PublicKey
	sk, pk = (*keygen).GenKeyPair()
	var kpHandle C.struct_Lattigo_KeyPairHandle
	kpHandle.sk = C.uint64_t(marshal.CrossLangObjMap.Add(unsafe.Pointer(sk)))
	kpHandle.pk = C.uint64_t(marshal.CrossLangObjMap.Add(unsafe.Pointer(pk)))
	return kpHandle
}

//export lattigo_genKeyPairSparse
func lattigo_genKeyPairSparse(keygenHandle Handle5, hw uint64) C.struct_Lattigo_KeyPairHandle {
	var keygen *rlwe.KeyGenerator
	keygen = getStoredKeyGenerator(keygenHandle)
	var sk *rlwe.SecretKey
	var pk *rlwe.PublicKey
	sk, pk = (*keygen).GenKeyPairSparse(int(hw))
	var kpHandle C.struct_Lattigo_KeyPairHandle
	kpHandle.sk = C.uint64_t(marshal.CrossLangObjMap.Add(unsafe.Pointer(sk)))
	kpHandle.pk = C.uint64_t(marshal.CrossLangObjMap.Add(unsafe.Pointer(pk)))
	return kpHandle
}

// only generates relinearization keys for ciphertexts up to degree 2
// (i.e., you must relinearize after each ct/ct multiplication)
//
//export lattigo_genRelinearizationKey
func lattigo_genRelinearizationKey(keygenHandle Handle5, skHandle Handle5) Handle5 {
	var keygen *rlwe.KeyGenerator
	keygen = getStoredKeyGenerator(keygenHandle)
	var sk *rlwe.SecretKey
	sk = getStoredSecretKey(skHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer((*keygen).GenRelinearizationKey(sk, 2)))
}

// Positive k is for left rotation by k positions
// Negative k is equivalent to a right rotation by k positions
//
//export lattigo_genRotationKeysForRotations
func lattigo_genRotationKeysForRotations(keygenHandle Handle5, skHandle Handle5, ks *C.int64_t, ksLen uint64) Handle5 {
	var keygen *rlwe.KeyGenerator
	keygen = getStoredKeyGenerator(keygenHandle)

	var sk *rlwe.SecretKey
	sk = getStoredSecretKey(skHandle)

	rotations := make([]int, ksLen)
	size := unsafe.Sizeof(uint64(0))
	basePtrIn := uintptr(unsafe.Pointer(ks))
	for i := range rotations {
		rotations[i] = int(*(*int64)(unsafe.Pointer(basePtrIn + size*uintptr(i))))
	}

	var rotKeys *rlwe.RotationKeySet
	// The second argument determines if conjugation keys are generated or not. This wrapper API does
	// not support generating a conjugation key.
	rotKeys = (*keygen).GenRotationKeysForRotations(rotations, false, sk)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(rotKeys))
}

//export lattigo_makeEvaluationKey
func lattigo_makeEvaluationKey(relinKeyHandle Handle5, rotKeyHandle Handle5) Handle5 {
	var relinKey *rlwe.RelinearizationKey
	relinKey = getStoredRelinKey(relinKeyHandle)

	var rotKeys *rlwe.RotationKeySet
	rotKeys = getStoredRotationKeys(rotKeyHandle)

	var evalKey rlwe.EvaluationKey
	evalKey = rlwe.EvaluationKey{Rlk: relinKey, Rtks: rotKeys}

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&evalKey))
}

//export lattigo_makeEmptyEvaluationKey
func lattigo_makeEmptyEvaluationKey() Handle5 {
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&rlwe.EvaluationKey{}))
}

//export lattigo_setRelinKeyForEvaluationKey
func lattigo_setRelinKeyForEvaluationKey(evalKeyHandle Handle5, relinKeyHandle Handle5) {
	evalKey := getStoredEvaluationKey(evalKeyHandle)
	evalKey.Rlk = getStoredRelinKey(relinKeyHandle)
}

//export lattigo_setRotKeysForEvaluationKey
func lattigo_setRotKeysForEvaluationKey(evalKeyHandle Handle5, rotKeysHandle Handle5) {
	evalKey := getStoredEvaluationKey(evalKeyHandle)
	evalKey.Rtks = getStoredRotationKeys(rotKeysHandle)
}

// Generates any missing Galois keys
//
//export lattigo_genBootstrappingKey
func lattigo_genBootstrappingKey(keygenHandle Handle5, paramHandle Handle5, btpParamsHandle Handle5, skHandle Handle5, relinKeyHandle Handle5, rotKeyHandle Handle5) Handle5 {
	var keygen *rlwe.KeyGenerator
	keygen = getStoredKeyGenerator(keygenHandle)

	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)

	var btpParams *ckks.BootstrappingParameters
	btpParams = getStoredBootstrappingParameters(btpParamsHandle)

	var sk *rlwe.SecretKey
	sk = getStoredSecretKey(skHandle)

	// get existing keys
	var relinKey *rlwe.RelinearizationKey
	relinKey = getStoredRelinKey(relinKeyHandle)

	var rotKeys *rlwe.RotationKeySet
	rotKeys = getStoredRotationKeys(rotKeyHandle)

	// generate the set of keys needed for bootstrapping
	btpRots := btpParams.RotationsForBootstrapping(params.LogSlots())

	// galois elements corresponding to the bootstrapping rotation indices
	var galEls []uint64
	for _, k := range btpRots {
		// generate the Galois index for this rotation
		var galoisIdx uint64
		galoisIdx = params.GaloisElementForColumnRotationBy(k)

		// test if this galoisIdx is already in the set of rotation keys.
		// if NOT, add it to galEls so that it will be generated
		if _, ok := rotKeys.Keys[galoisIdx]; !ok {
			galEls = append(galEls, galoisIdx)
		}
	}
	// include a conjugation key
	var conjIdx uint64
	conjIdx = params.GaloisElementForRowRotation()
	if _, ok := rotKeys.Keys[conjIdx]; !ok {
		galEls = append(galEls, conjIdx)
	}

	// galEls contains the Galois indices needed for bootstrapping, but which are not already in the set of rotation keys.
	// Generate a new set of rotation keys for the missing indices, then merge the two maps.
	var btpRotKeys *rlwe.RotationKeySet
	btpRotKeys = (*keygen).GenRotationKeys(galEls, sk)

	for k, v := range btpRotKeys.Keys {
		if _, ok := rotKeys.Keys[k]; ok {
			panic(errors.New("Internal error: Generated a bootstrapping key that is already in the map"))
		}
		rotKeys.Keys[k] = v
	}

	var btpKey ckks.BootstrappingKey
	btpKey = ckks.BootstrappingKey{Rlk: relinKey, Rtks: rotKeys}

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&btpKey))
}

//export lattigo_makeBootstrappingKey
func lattigo_makeBootstrappingKey(relinKeyHandle Handle5, rotKeyHandle Handle5) Handle5 {
	// get existing keys
	var relinKey *rlwe.RelinearizationKey
	relinKey = getStoredRelinKey(relinKeyHandle)

	var rotKeys *rlwe.RotationKeySet
	rotKeys = getStoredRotationKeys(rotKeyHandle)

	var btpKey ckks.BootstrappingKey
	btpKey = ckks.BootstrappingKey{Rlk: relinKey, Rtks: rotKeys}

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&btpKey))
}

//export lattigo_getSecretKeyValue
func lattigo_getSecretKeyValue(skHandle Handle5) Handle5 {
	sk := getStoredSecretKey(skHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(sk.Value))
}

//export lattigo_getSwitchingKey
func lattigo_getSwitchingKey(rotKeyHandle Handle5, galoisElement uint64) Handle5 {
	rotKeys := getStoredRotationKeys(rotKeyHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(rotKeys.Keys[galoisElement]))
}
