// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

/*
#include <stdint.h>
*/
import "C"

import (
	"lattigo-cpp/marshal"
	"unsafe"

	"github.com/tuneinsight/lattigo/v4/ckks"
	"github.com/tuneinsight/lattigo/v4/rlwe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle4 = uint64

func getStoredEvaluator(evalHandle Handle4) *ckks.Evaluator {
	ref := marshal.CrossLangObjMap.Get(evalHandle)
	return (*ckks.Evaluator)(ref.Ptr)
}

//export lattigo_newEvaluator
func lattigo_newEvaluator(paramsHandle Handle4, evalkeyHandle Handle4) Handle4 {
	var params *ckks.Parameters
	params = getStoredParameters(paramsHandle)

	var evalKeys *rlwe.EvaluationKey
	evalKeys = getStoredEvaluationKey(evalkeyHandle)

	var evaluator ckks.Evaluator
	evaluator = ckks.NewEvaluator(*params, *evalKeys)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&evaluator))
}

//export lattigo_evaluatorWithKey
func lattigo_evaluatorWithKey(evalHandle, evalKeyHandle Handle4) Handle4 {
	eval := getStoredEvaluator(evalHandle)
	evalKey := getStoredEvaluationKey(evalKeyHandle)
	evalWithKey := (*eval).WithKey(*evalKey)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&evalWithKey))
}

//export lattigo_rotate
func lattigo_rotate(evalHandle Handle4, ctInHandle Handle4, k uint64, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var ctOut *rlwe.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).Rotate(ctIn, int(k), ctOut)
}

//export lattigo_rotateHoisted
func lattigo_rotateHoisted(evalHandle Handle4, ctInHandle Handle4, ks *C.uint64_t, ksLen uint64, outHandles *C.uint64_t) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	rotations := make([]int, ksLen)
	size := unsafe.Sizeof(uint64(0))
	basePtrIn := uintptr(unsafe.Pointer(ks))
	for i := range rotations {
		rotations[i] = *(*int)(unsafe.Pointer(basePtrIn + size*uintptr(i)))
	}

	var rotatedCts map[int]*rlwe.Ciphertext
	(*eval).RotateHoisted(ctIn, rotations, rotatedCts)

	basePtrOut := uintptr(unsafe.Pointer(outHandles))
	for i := range rotations {
		*(*uint64)(unsafe.Pointer(basePtrOut + size*uintptr(i))) = marshal.CrossLangObjMap.Add(unsafe.Pointer(rotatedCts[rotations[i]]))
	}
}

//export lattigo_multByConst
func lattigo_multByConst(evalHandle Handle4, ctInHandle Handle4, constant float64, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var ctOut *rlwe.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).MultByConst(ctIn, constant, ctOut)
}

//export lattigo_addConst
func lattigo_addConst(evalHandle Handle4, ctInHandle Handle4, constant float64, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var ctOut *rlwe.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).AddConst(ctIn, constant, ctOut)
}

//export lattigo_rescale
func lattigo_rescale(evalHandle Handle4, ctInHandle Handle4, scale float64, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var ctOut *rlwe.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	err := (*eval).Rescale(ctIn, rlwe.NewScale(scale), ctOut)
	if err != nil {
		panic(err)
	}
}

//export lattigo_mulRelinNew
func lattigo_mulRelinNew(evalHandle Handle4, op0Handle Handle4, op1Handle Handle4) Handle4 {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ct0 *rlwe.Ciphertext
	ct0 = getStoredCiphertext(op0Handle)

	var ct1 *rlwe.Ciphertext
	ct1 = getStoredCiphertext(op1Handle)

	var ctOut *rlwe.Ciphertext
	ctOut = (*eval).MulRelinNew(ct0, ct1)

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ctOut))
}

// multiply two ciphertexts and relinearize the result
//
//export lattigo_mulRelin
func lattigo_mulRelin(evalHandle Handle4, op0Handle Handle4, op1Handle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ct0 *rlwe.Ciphertext
	ct0 = getStoredCiphertext(op0Handle)

	var ct1 *rlwe.Ciphertext
	ct1 = getStoredCiphertext(op1Handle)

	var ctOut *rlwe.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).MulRelin(ct0, ct1, ctOut)
}

// multiply two ciphertexts without relinearization
//
//export lattigo_mul
func lattigo_mul(evalHandle Handle4, op0Handle Handle4, op1Handle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ct0 *rlwe.Ciphertext
	ct0 = getStoredCiphertext(op0Handle)

	var ct1 *rlwe.Ciphertext
	ct1 = getStoredCiphertext(op1Handle)

	var ctOut *rlwe.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).Mul(ct0, ct1, ctOut)
}

// multiply a ciphertext by a plaintext
//
//export lattigo_mulPlain
func lattigo_mulPlain(evalHandle Handle4, ctInHandle Handle4, ptHandle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var pt *rlwe.Plaintext
	pt = getStoredPlaintext(ptHandle)

	var ctOut *rlwe.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).Mul(ctIn, pt, ctOut)
}

//export lattigo_add
func lattigo_add(evalHandle Handle4, op0Handle Handle4, op1Handle Handle4, outHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ct0 *rlwe.Ciphertext
	ct0 = getStoredCiphertext(op0Handle)

	var ct1 *rlwe.Ciphertext
	ct1 = getStoredCiphertext(op1Handle)

	var ctOut *rlwe.Ciphertext
	ctOut = getStoredCiphertext(outHandle)

	(*eval).Add(ct0, ct1, ctOut)
}

//export lattigo_addPlain
func lattigo_addPlain(evalHandle Handle4, ctInHandle Handle4, ptHandle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var pt *rlwe.Plaintext
	pt = getStoredPlaintext(ptHandle)

	var ctOut *rlwe.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).Add(ctIn, pt, ctOut)
}

//export lattigo_neg
func lattigo_neg(evalHandle Handle4, ctInHandle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var ctOut *rlwe.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).Neg(ctIn, ctOut)
}

//export lattigo_sub
func lattigo_sub(evalHandle Handle4, op0Handle Handle4, op1Handle Handle4, outHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ct0 *rlwe.Ciphertext
	ct0 = getStoredCiphertext(op0Handle)

	var ct1 *rlwe.Ciphertext
	ct1 = getStoredCiphertext(op1Handle)

	var ctOut *rlwe.Ciphertext
	ctOut = getStoredCiphertext(outHandle)

	(*eval).Sub(ct0, ct1, ctOut)
}

//export lattigo_subPlain
func lattigo_subPlain(evalHandle Handle4, ctInHandle Handle4, ptHandle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var pt *rlwe.Plaintext
	pt = getStoredPlaintext(ptHandle)

	var ctOut *rlwe.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).Sub(ctIn, pt, ctOut)
}

//export lattigo_dropLevel
func lattigo_dropLevel(evalHandle Handle4, ctHandle Handle4, levels uint64) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ct *rlwe.Ciphertext
	ct = getStoredCiphertext(ctHandle)

	(*eval).DropLevel(ct, int(levels))
}

//export lattigo_multByGaussianIntegerAndAdd
func lattigo_multByGaussianIntegerAndAdd(evalHandle Handle4, ct0Handle Handle4, cReal int64, cImag int64, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ct0 *rlwe.Ciphertext
	ct0 = getStoredCiphertext(ct0Handle)

	var ctOut *rlwe.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).MultByGaussianIntegerAndAdd(ct0, cReal, cImag, ctOut)
}

//export lattigo_relinearize
func lattigo_relinearize(evalHandle Handle4, ctInHandle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *rlwe.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var ctOut *rlwe.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).Relinearize(ctIn, ctOut)
}

//export lattigo_switchKeys
func lattigo_switchKeys(evalHandle, ctxInHandle, swkHandle, ctxOutHandle Handle4) {
	eval := getStoredEvaluator(evalHandle)
	swk := getStoredSwitchingKey(swkHandle)
	ctxIn := getStoredCiphertext(ctxInHandle)
	ctxOut := getStoredCiphertext(ctxOutHandle)
	(*eval).SwitchKeys(ctxIn, swk, ctxOut)
}
