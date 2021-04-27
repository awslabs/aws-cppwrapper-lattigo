package ckks

import "C"

import (
    "github.com/ldsec/lattigo/v2/ckks"
    "lattigo-cpp/marshal"
    "unsafe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle9 = uint64

func marshalBytes(input []byte, output *C.uchar) {
	size := unsafe.Sizeof(uint8(0))
	basePtr := uintptr(unsafe.Pointer(output))
	for i := range input {
		*(*uint8)(unsafe.Pointer(basePtr + size*uintptr(i))) = input[i]
	}
}

//export lattigo_marshalBinaryCiphertext
func lattigo_marshalBinaryCiphertext(ctHandle Handle9, output *C.uchar) {
	var ct *ckks.Ciphertext
	ct = getStoredCiphertext(ctHandle)

	data, err := ct.MarshalBinary()
	if err != nil {
		panic(err)
	}

	marshalBytes(data, output)
}

//export lattigo_marshalBinaryParams
func lattigo_marshalBinaryParams(paramsHandle Handle9, output *C.uchar) {
	var params *ckks.Parameters
	params = getStoredParameters(paramsHandle)

	data, err := params.MarshalBinary()
	if err != nil {
		panic(err)
	}

	marshalBytes(data, output)
}

//export lattigo_marshalBinarySecretKey
func lattigo_marshalBinarySecretKey(skHandle Handle9, output *C.uchar) {
	var sk *ckks.SecretKey
	sk = getStoredSecretKey(skHandle)

	data, err := sk.MarshalBinary()
	if err != nil {
		panic(err)
	}

	marshalBytes(data, output)
}

//export lattigo_marshalBinaryPublicKey
func lattigo_marshalBinaryPublicKey(pkHandle Handle9, output *C.uchar) {
	var pk *ckks.PublicKey
	pk = getStoredPublicKey(pkHandle)

	data, err := pk.MarshalBinary()
	if err != nil {
		panic(err)
	}

	marshalBytes(data, output)
}



//export lattigo_unmarshalBinaryCiphertext
func lattigo_unmarshalBinaryCiphertext(output *C.uchar, len uint64) Handle9 {
	var serializedBytes []byte
	serializedBytes = C.GoBytes(unsafe.Pointer(output), C.int(len))

	ct := new(ckks.Ciphertext)
	err := ct.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ct))
}

//export lattigo_unmarshalBinaryParams
func lattigo_unmarshalBinaryParams(output *C.uchar, len uint64) Handle9 {
	var serializedBytes []byte
	serializedBytes = C.GoBytes(unsafe.Pointer(output), C.int(len))

	params := new(ckks.Parameters)
	err := params.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(params))
}

//export lattigo_unmarshalBinarySecretKey
func lattigo_unmarshalBinarySecretKey(output *C.uchar, len uint64) Handle9 {
	var serializedBytes []byte
	serializedBytes = C.GoBytes(unsafe.Pointer(output), C.int(len))

	sk := new(ckks.SecretKey)
	err := sk.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(sk))
}

//export lattigo_unmarshalBinaryPublicKey
func lattigo_unmarshalBinaryPublicKey(output *C.uchar, len uint64) Handle9 {
	var serializedBytes []byte
	serializedBytes = C.GoBytes(unsafe.Pointer(output), C.int(len))

	pk := new(ckks.PublicKey)
	err := pk.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(pk))
}
