// SPDX-License-Identifier: Apache-2.0

package utils

import "C"

import (
	"lattigo-cpp/marshal"
	"unsafe"

	"github.com/ldsec/lattigo/v2/utils"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle15 = uint64

func GetStoredKeyedPRNG(prngHandle Handle15) *utils.KeyedPRNG {
	ref := marshal.CrossLangObjMap.Get(prngHandle)
	return (*utils.KeyedPRNG)(ref.Ptr)
}

//export lattigo_newPRNG
func lattigo_newPRNG() Handle15 {
	// prng is of type KeyedPRNG
	prng, err := utils.NewPRNG()

	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(prng))
}

//export lattigo_newKeyedPRNG
func lattigo_newKeyedPRNG(key *C.char, keyLen uint64) Handle15 {
	keyTmp := make([]byte, keyLen)
	size := unsafe.Sizeof(byte(0))
	basePtrIn := uintptr(unsafe.Pointer(&key))
	for i := range keyTmp {
		keyTmp[i] = *(*byte)(unsafe.Pointer(basePtrIn + size*uintptr(i)))
	}

	prng, err := utils.NewKeyedPRNG(keyTmp)

	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(prng))
}
