package ckks

/*
#include <stdint.h>
*/
import "C"

import (
	"github.com/ldsec/lattigo/v2/ckks"
	"lattigo-cpp/marshal"
	"unsafe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle4 = uint64

func getStoredEvaluator(evalHandle Handle4) *ckks.Evaluator {
	ref := marshal.CrossLangObjMap.Get(evalHandle)
	return (*ckks.Evaluator)(ref.Ptr)
}

//export lattigo_newEvaluator
func lattigo_newEvaluator(paramsHandle Handle4) Handle4 {
	var params *ckks.Parameters
	params = getStoredParameters(paramsHandle)

	var evaluator ckks.Evaluator
	evaluator = ckks.NewEvaluator(params)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&evaluator))
}

//export lattigo_rotate
func lattigo_rotate(evalHandle Handle4, ctInHandle Handle4, k uint64, rotKeysHandle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var rotKeys *ckks.RotationKeys
	rotKeys = getStoredRotationKeys(rotKeysHandle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).Rotate(ctIn, k, rotKeys, ctOut)
}

//export lattigo_rotateHoisted
func lattigo_rotateHoisted(evalHandle Handle4, ctInHandle Handle4, ks *C.uint64_t, ksLen uint64, rotKeysHandle Handle4, outHandles *C.uint64_t) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var rotKeys *ckks.RotationKeys
	rotKeys = getStoredRotationKeys(rotKeysHandle)

	rotations := make([]uint64, ksLen)
	size := unsafe.Sizeof(uint64(0))
	basePtrIn := uintptr(unsafe.Pointer(ks))
	for i := range rotations {
		rotations[i] = *(*uint64)(unsafe.Pointer(basePtrIn + size*uintptr(i)))
	}

	var rotatedCts map[uint64]*ckks.Ciphertext
	rotatedCts = (*eval).RotateHoisted(ctIn, rotations, rotKeys)

	basePtrOut := uintptr(unsafe.Pointer(outHandles))
	for i := range rotations {
		*(*uint64)(unsafe.Pointer(basePtrOut + size*uintptr(i))) = marshal.CrossLangObjMap.Add(unsafe.Pointer(rotatedCts[rotations[i]]))
	}
}

//export lattigo_multByConst
func lattigo_multByConst(evalHandle Handle4, ctInHandle Handle4, constant float64, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).MultByConst(ctIn, constant, ctOut)
}

//export lattigo_addConst
func lattigo_addConst(evalHandle Handle4, ctInHandle Handle4, constant float64, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).AddConst(ctIn, constant, ctOut)
}

//export lattigo_rescale
func lattigo_rescale(evalHandle Handle4, ctInHandle Handle4, threshold float64, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	err := (*eval).Rescale(ctIn, threshold, ctOut)
	if err != nil {
		panic(err)
	}
}

//export lattigo_rescaleMany
func lattigo_rescaleMany(evalHandle Handle4, ctInHandle Handle4, numRescales uint64, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	err := (*eval).RescaleMany(ctIn, numRescales, ctOut)
	if err != nil {
		panic(err)
	}
}

//export lattigo_mulRelinNew
func lattigo_mulRelinNew(evalHandle Handle4, op0Handle Handle4, op1Handle Handle4, evakeyHandle Handle4) Handle4 {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ct0 *ckks.Ciphertext
	ct0 = getStoredCiphertext(op0Handle)

	var ct1 *ckks.Ciphertext
	ct1 = getStoredCiphertext(op1Handle)

	var evakey *ckks.EvaluationKey
	evakey = getStoredEvalKey(evakeyHandle)

	var ctOut *ckks.Ciphertext
	ctOut = (*eval).MulRelinNew(ct0, ct1, evakey)

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ctOut))
}

// multiply two ciphertexts and relinearize the result
//export lattigo_mulRelin
func lattigo_mulRelin(evalHandle Handle4, op0Handle Handle4, op1Handle Handle4, evakeyHandle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ct0 *ckks.Ciphertext
	ct0 = getStoredCiphertext(op0Handle)

	var ct1 *ckks.Ciphertext
	ct1 = getStoredCiphertext(op1Handle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	var evakey *ckks.EvaluationKey
	evakey = getStoredEvalKey(evakeyHandle)

	(*eval).MulRelin(ct0, ct1, evakey, ctOut)
}

// multiply two ciphertexts without relinearization
//export lattigo_mul
func lattigo_mul(evalHandle Handle4, op0Handle Handle4, op1Handle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ct0 *ckks.Ciphertext
	ct0 = getStoredCiphertext(op0Handle)

	var ct1 *ckks.Ciphertext
	ct1 = getStoredCiphertext(op1Handle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).MulRelin(ct0, ct1, nil, ctOut)
}

// multiply a ciphertext by a plaintext
//export lattigo_mulPlain
func lattigo_mulPlain(evalHandle Handle4, ctInHandle Handle4, ptHandle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var pt *ckks.Plaintext
	pt = getStoredPlaintext(ptHandle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).MulRelin(ctIn, pt, nil, ctOut)
}

//export lattigo_add
func lattigo_add(evalHandle Handle4, op0Handle Handle4, op1Handle Handle4, outHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ct0 *ckks.Ciphertext
	ct0 = getStoredCiphertext(op0Handle)

	var ct1 *ckks.Ciphertext
	ct1 = getStoredCiphertext(op1Handle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(outHandle)

	(*eval).Add(ct0, ct1, ctOut)
}

//export lattigo_addPlain
func lattigo_addPlain(evalHandle Handle4, ctInHandle Handle4, ptHandle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var pt *ckks.Ciphertext
	pt = getStoredCiphertext(ptHandle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).Add(ctIn, pt, ctOut)
}

//export lattigo_neg
func lattigo_neg(evalHandle Handle4, ctInHandle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).Neg(ctIn, ctOut)
}

//export lattigo_sub
func lattigo_sub(evalHandle Handle4, op0Handle Handle4, op1Handle Handle4, outHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ct0 *ckks.Ciphertext
	ct0 = getStoredCiphertext(op0Handle)

	var ct1 *ckks.Ciphertext
	ct1 = getStoredCiphertext(op1Handle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(outHandle)

	(*eval).Sub(ct0, ct1, ctOut)
}

//export lattigo_subPlain
func lattigo_subPlain(evalHandle Handle4, ctInHandle Handle4, ptHandle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var pt *ckks.Ciphertext
	pt = getStoredCiphertext(ptHandle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).Sub(ctIn, pt, ctOut)
}

//export lattigo_dropLevel
func lattigo_dropLevel(evalHandle Handle4, ctHandle Handle4, levels uint64) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ct *ckks.Ciphertext
	ct = getStoredCiphertext(ctHandle)

	(*eval).DropLevel(ct, levels)
}

//export lattigo_multByGaussianIntegerAndAdd
func lattigo_multByGaussianIntegerAndAdd(evalHandle Handle4, ct0Handle Handle4, cReal int64, cImag int64, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ct0 *ckks.Ciphertext
	ct0 = getStoredCiphertext(ct0Handle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	(*eval).MultByGaussianIntegerAndAdd(ct0, cReal, cImag, ctOut)
}

//export lattigo_relinearize
func lattigo_relinearize(evalHandle Handle4, ctInHandle Handle4, evakeyHandle Handle4, ctOutHandle Handle4) {
	var eval *ckks.Evaluator
	eval = getStoredEvaluator(evalHandle)

	var ctIn *ckks.Ciphertext
	ctIn = getStoredCiphertext(ctInHandle)

	var ctOut *ckks.Ciphertext
	ctOut = getStoredCiphertext(ctOutHandle)

	var evakey *ckks.EvaluationKey
	evakey = getStoredEvalKey(evakeyHandle)

	(*eval).Relinearize(ctIn, evakey, ctOut)
}
