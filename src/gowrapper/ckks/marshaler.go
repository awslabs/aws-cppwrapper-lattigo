// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

/*
#include <stdint.h>
typedef void (*streamWriter) (void*, void*, uint64_t);

// https://golang.org/cmd/cgo/#hdr-Go_references_to_C
// https://nelkinda.com/blog/suppress-warnings-in-gcc-and-clang/
__attribute__((unused)) static void callStreamWriter(streamWriter f, void* stream, void* data, uint64_t len) {
  f(stream, data, len);
}
*/
import "C"

import (
	"errors"
	"lattigo-cpp/marshal"
	"reflect"
	"unsafe"

	"github.com/tuneinsight/lattigo/v4/ckks"
	"github.com/tuneinsight/lattigo/v4/ckks/bootstrapping"
	"github.com/tuneinsight/lattigo/v4/rlwe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle9 = uint64

//export lattigo_marshalBinaryCiphertext
func lattigo_marshalBinaryCiphertext(ctHandle Handle9, callback C.streamWriter, stream *C.void) {
	var ct *rlwe.Ciphertext
	ct = getStoredCiphertext(ctHandle)

	data, err := ct.MarshalBinary()
	if err != nil {
		panic(err)
	}

	if len(data) > 0 {
		C.callStreamWriter(callback, unsafe.Pointer(stream), unsafe.Pointer(&data[0]), C.uint64_t(len(data)))
	}
}

//export lattigo_marshalBinaryParameters
func lattigo_marshalBinaryParameters(paramsHandle Handle9, callback C.streamWriter, stream *C.void) {
	var params *ckks.Parameters
	params = getStoredParameters(paramsHandle)

	data, err := params.MarshalBinary()
	if err != nil {
		panic(err)
	}

	if len(data) > 0 {
		C.callStreamWriter(callback, unsafe.Pointer(stream), unsafe.Pointer(&data[0]), C.uint64_t(len(data)))
	}
}

//export lattigo_marshalBinaryBootstrapParameters
func lattigo_marshalBinaryBootstrapParameters(paramsHandle Handle9, callback C.streamWriter, stream *C.void) {
	var params *bootstrapping.Parameters
	params = getStoredBootstrappingParameters(paramsHandle)

	data, err := params.MarshalBinary()
	if err != nil {
		panic(err)
	}

	if len(data) > 0 {
		C.callStreamWriter(callback, unsafe.Pointer(stream), unsafe.Pointer(&data[0]), C.uint64_t(len(data)))
	}
}

//export lattigo_marshalBinarySecretKey
func lattigo_marshalBinarySecretKey(skHandle Handle9, callback C.streamWriter, stream *C.void) {
	var sk *rlwe.SecretKey
	sk = getStoredSecretKey(skHandle)

	data, err := sk.MarshalBinary()
	if err != nil {
		panic(err)
	}

	if len(data) > 0 {
		C.callStreamWriter(callback, unsafe.Pointer(stream), unsafe.Pointer(&data[0]), C.uint64_t(len(data)))
	}
}

//export lattigo_marshalBinaryPublicKey
func lattigo_marshalBinaryPublicKey(pkHandle Handle9, callback C.streamWriter, stream *C.void) {
	var pk *rlwe.PublicKey
	pk = getStoredPublicKey(pkHandle)

	data, err := pk.MarshalBinary()
	if err != nil {
		panic(err)
	}

	if len(data) > 0 {
		C.callStreamWriter(callback, unsafe.Pointer(stream), unsafe.Pointer(&data[0]), C.uint64_t(len(data)))
	}
}

//export lattigo_marshalBinaryRelinearizationKey
func lattigo_marshalBinaryRelinearizationKey(relinKeyHandle Handle9, callback C.streamWriter, stream *C.void) {
	var relinKey *rlwe.RelinearizationKey
	relinKey = getStoredRelinKey(relinKeyHandle)

	data, err := relinKey.MarshalBinary()
	if err != nil {
		panic(err)
	}

	if len(data) > 0 {
		C.callStreamWriter(callback, unsafe.Pointer(stream), unsafe.Pointer(&data[0]), C.uint64_t(len(data)))
	}
}

//export lattigo_marshalBinaryRotationKeys
func lattigo_marshalBinaryRotationKeys(rotkeyHandle Handle9, callback C.streamWriter, stream *C.void) {
	var rotkeys *rlwe.RotationKeySet
	rotkeys = getStoredRotationKeys(rotkeyHandle)

	data, err := rotkeys.MarshalBinary()
	if err != nil {
		panic(err)
	}

	if len(data) > 0 {
		C.callStreamWriter(callback, unsafe.Pointer(stream), unsafe.Pointer(&data[0]), C.uint64_t(len(data)))
	}
}

// We need a way to convert C-allocated memory into a Go Slice.
// One option is to use C.GoBytes. This is safe, but there are
// two problems. First, it copies the data, which is not great
// when we're talking about serialized keys which are 10s of GBs.
// More importantly, it takes the length of the C memory as a C.int,
// which is a signed 32-bit type on my 64-bit system. However,
// rotation keys handily exceed 2^31 bytes, resulting in a runtime
// error when I try to use the memory allocated this way.
//
// Instead, we can create an unallocated slice, and then use
// the reflection package to make it point to the C-allocated memory.
// This gets around both problems with C.GoBytes: it does not copy the
// C memory, and the length of the slize is a Go `int`, which is
// 32 bits on a 32-bit system and 64-bits on a 64-bit system.
// The caveat is that it's not very "safe". First, Go could modify
// the C memory, but we ensure this function is only used to create
// slices used for read-only functions. Second, the Go garbage collector
// could GC the data. But that can't happen since the memory was
// allocated by C, and therefore is unknown to the garbage collector.
// Finally, and most importantly, the reflect.SliceHeader data type
// could change without notice in future Go releases.
// Overall, this is the best solution I've come up with. See
// https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
func unsafeCPtrToSlice(buf *C.char, len uint64) []byte {
	// Check if the uint64-length of the buffer is larger than the system `int` type.
	// Technically, this check assumes that `len` is < 2^63. However, this is a
	// reasonable assumption since 2^63 bytes is over 9000 PB, which is large even
	// by HE-key standards.
	if int64(len) != int64(int(len)) {
		panic(errors.New("Unable to make a slice of the requested size"))
	}
	// declare an unallocated slice
	var serializedBytes []byte
	// get the SliceHeader of the new slice using the `reflect` API
	// Modifying `hdr` implicitly modifies `serializedBytes` via reflection.
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&serializedBytes))
	// set the data, length, and capacity of the slice
	hdr.Data = uintptr(unsafe.Pointer(buf))
	hdr.Cap = int(len)
	hdr.Len = int(len)
	// return the "allocated" slice
	return serializedBytes
}

//export lattigo_unmarshalBinaryCiphertext
func lattigo_unmarshalBinaryCiphertext(buf *C.char, len uint64) Handle9 {
	var serializedBytes []byte = unsafeCPtrToSlice(buf, len)

	ct := new(rlwe.Ciphertext)
	err := ct.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ct))
}

//export lattigo_unmarshalBinaryParameters
func lattigo_unmarshalBinaryParameters(buf *C.char, len uint64) Handle9 {
	var serializedBytes []byte = unsafeCPtrToSlice(buf, len)

	params := new(ckks.Parameters)
	err := params.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(params))
}

//export lattigo_unmarshalBinaryBootstrapParameters
func lattigo_unmarshalBinaryBootstrapParameters(buf *C.char, len uint64) Handle9 {
	var serializedBytes []byte = unsafeCPtrToSlice(buf, len)

	params := new(bootstrapping.Parameters)
	err := params.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(params))
}

//export lattigo_unmarshalBinarySecretKey
func lattigo_unmarshalBinarySecretKey(buf *C.char, len uint64) Handle9 {
	var serializedBytes []byte = unsafeCPtrToSlice(buf, len)

	sk := new(rlwe.SecretKey)
	err := sk.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(sk))
}

//export lattigo_unmarshalBinaryPublicKey
func lattigo_unmarshalBinaryPublicKey(buf *C.char, len uint64) Handle9 {
	var serializedBytes []byte = unsafeCPtrToSlice(buf, len)

	pk := new(rlwe.PublicKey)
	err := pk.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(pk))
}

//export lattigo_unmarshalBinaryRelinearizationKey
func lattigo_unmarshalBinaryRelinearizationKey(buf *C.char, len uint64) Handle9 {
	var serializedBytes []byte = unsafeCPtrToSlice(buf, len)

	relinKey := new(rlwe.RelinearizationKey)
	err := relinKey.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(relinKey))
}

//export lattigo_unmarshalBinaryRotationKeys
func lattigo_unmarshalBinaryRotationKeys(buf *C.char, len uint64) Handle9 {
	var serializedBytes []byte = unsafeCPtrToSlice(buf, len)

	rotkeys := new(rlwe.RotationKeySet)
	err := rotkeys.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(rotkeys))
}

//export lattigo_marshalBinarySizeCiphertext
func lattigo_marshalBinarySizeCiphertext(ctHandle Handle9) uint64 {
	var ct *rlwe.Ciphertext
	ct = getStoredCiphertext(ctHandle)
	return uint64(ct.MarshalBinarySize())
}

//export lattigo_marshalBinarySizeParameters
func lattigo_marshalBinarySizeParameters(paramsHandle Handle9) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramsHandle)
	paramBytes, err := params.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return uint64(len(paramBytes))
}

//export lattigo_marshalBinarySizeSecretKey
func lattigo_marshalBinarySizeSecretKey(skHandle Handle9) uint64 {
	var sk *rlwe.SecretKey
	sk = getStoredSecretKey(skHandle)
	return uint64(sk.MarshalBinarySize())
}

//export lattigo_marshalBinarySizePublicKey
func lattigo_marshalBinarySizePublicKey(pkHandle Handle9) uint64 {
	var pk *rlwe.PublicKey
	pk = getStoredPublicKey(pkHandle)
	return uint64(pk.MarshalBinarySize())
}

//export lattigo_marshalBinarySizeRelinearizationKey
func lattigo_marshalBinarySizeRelinearizationKey(relinKeyHandle Handle9) uint64 {
	var relinKey *rlwe.RelinearizationKey
	relinKey = getStoredRelinKey(relinKeyHandle)
	return uint64(relinKey.MarshalBinarySize())
}

//export lattigo_marshalBinarySizeRotationKeys
func lattigo_marshalBinarySizeRotationKeys(rotkeysHandle Handle9) uint64 {
	var rotkeys *rlwe.RotationKeySet
	rotkeys = getStoredRotationKeys(rotkeysHandle)
	return uint64(rotkeys.MarshalBinarySize())
}
