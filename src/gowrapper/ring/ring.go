// SPDX-License-Identifier: Apache-2.0

package ring

/*
#include <stdint.h>
*/
import "C"

import (
	"lattigo-cpp/marshal"
	"lattigo-cpp/utils"
	"unsafe"

	"github.com/ldsec/lattigo/v2/ring"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle14 = uint64

func getStoredRing(ringHandle Handle14) *ring.Ring {
	ref := marshal.CrossLangObjMap.Get(ringHandle)
	return (*ring.Ring)(ref.Ptr)
}

func GetStoredPoly(polyHandle Handle14) *ring.Poly {
	ref := marshal.CrossLangObjMap.Get(polyHandle)
	return (*ring.Poly)(ref.Ptr)
}

func getStoredUniformSampler(samplerHandle Handle14) *ring.UniformSampler {
	ref := marshal.CrossLangObjMap.Get(samplerHandle)
	return (*ring.UniformSampler)(ref.Ptr)
}

//export lattigo_newRing
func lattigo_newRing(n int, moduli *C.uint64_t, moduliLen uint64) Handle14 {
	moduliTmp := make([]uint64, moduliLen)
	size := unsafe.Sizeof(uint64(0))
	basePtrIn := uintptr(unsafe.Pointer(&moduli))
	for i := range moduliTmp {
		moduliTmp[i] = *(*uint64)(unsafe.Pointer(basePtrIn + size*uintptr(i)))
	}

	r, err := ring.NewRing(n, moduliTmp)

	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(r))
}

//export lattigo_newPoly
func lattigo_newPoly(ringHandle Handle14) Handle14 {
	r := getStoredRing(ringHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(r.NewPoly()))
}

//export lattigo_ringAdd
func lattigo_ringAdd(ringHandle, poly1Handle, poly2Handle, poly3Handle Handle14) {
	r := getStoredRing(ringHandle)
	p1 := GetStoredPoly(poly1Handle)
	p2 := GetStoredPoly(poly2Handle)
	p3 := GetStoredPoly(poly3Handle)
	r.Add(p1, p2, p3)
}

//export lattigo_polyCopy
func lattigo_polyCopy(polyTargetHandle, polySrcHandle Handle14) {
	pTarget := GetStoredPoly(polyTargetHandle)
	pSrc := GetStoredPoly(polySrcHandle)
	pTarget.Copy(pSrc)
}

//export lattigo_newUniformSampler
func lattigo_newUniformSampler(prngHandle, baseRingHandle Handle14) Handle14 {
	prng := utils.GetStoredKeyedPRNG(prngHandle)
	r := getStoredRing(baseRingHandle)
	sampler := ring.NewUniformSampler(prng, r)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(sampler))
}

//export lattigo_readNewFromSampler
func lattigo_readNewFromSampler(samplerHandle Handle14) Handle14 {
	sampler := getStoredUniformSampler(samplerHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(sampler.ReadNew()))
}
