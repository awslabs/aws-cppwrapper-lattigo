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

//export lattigo_marshalBinaryParameters
func lattigo_marshalBinaryParameters(paramsHandle Handle9, output *C.uchar) {
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

//export lattigo_marshalBinaryEvaluationKey
func lattigo_marshalBinaryEvaluationKey(evakeyHandle Handle9, output *C.uchar) {
	var evakey *ckks.EvaluationKey
	evakey = getStoredEvalKey(evakeyHandle)

	data, err := evakey.MarshalBinary()
	if err != nil {
		panic(err)
	}

	marshalBytes(data, output)
}

//export lattigo_marshalBinaryRotationKeys
func lattigo_marshalBinaryRotationKeys(rotkeyHandle Handle9, output *C.uchar) {
	var rotkeys *ckks.RotationKeys
	rotkeys = getStoredRotationKeys(rotkeyHandle)

	data, err := rotkeys.MarshalBinary()
	if err != nil {
		panic(err)
	}

	marshalBytes(data, output)
}

//export lattigo_unmarshalBinaryCiphertext
func lattigo_unmarshalBinaryCiphertext(buf *C.char, len uint64) Handle9 {
	var serializedBytes []byte
	serializedBytes = C.GoBytes(unsafe.Pointer(buf), C.int(len))

	ct := new(ckks.Ciphertext)
	err := ct.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(ct))
}

//export lattigo_unmarshalBinaryParameters
func lattigo_unmarshalBinaryParameters(buf *C.char, len uint64) Handle9 {
	var serializedBytes []byte
	serializedBytes = C.GoBytes(unsafe.Pointer(buf), C.int(len))

	params := new(ckks.Parameters)
	err := params.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(params))
}

//export lattigo_unmarshalBinarySecretKey
func lattigo_unmarshalBinarySecretKey(buf *C.char, len uint64) Handle9 {
	var serializedBytes []byte
	serializedBytes = C.GoBytes(unsafe.Pointer(buf), C.int(len))

	sk := new(ckks.SecretKey)
	err := sk.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(sk))
}

//export lattigo_unmarshalBinaryPublicKey
func lattigo_unmarshalBinaryPublicKey(buf *C.char, len uint64) Handle9 {
	var serializedBytes []byte
	serializedBytes = C.GoBytes(unsafe.Pointer(buf), C.int(len))

	pk := new(ckks.PublicKey)
	err := pk.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(pk))
}

//export lattigo_unmarshalBinaryEvaluationKey
func lattigo_unmarshalBinaryEvaluationKey(buf *C.char, len uint64) Handle9 {
	var serializedBytes []byte
	serializedBytes = C.GoBytes(unsafe.Pointer(buf), C.int(len))

	evakey := new(ckks.EvaluationKey)
	err := evakey.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(evakey))
}

//export lattigo_unmarshalBinaryRotationKeys
func lattigo_unmarshalBinaryRotationKeys(buf *C.char, len uint64) Handle9 {
	var serializedBytes []byte
	serializedBytes = C.GoBytes(unsafe.Pointer(buf), C.int(len))

	rotkeys := new(ckks.RotationKeys)
	err := rotkeys.UnmarshalBinary(serializedBytes)
	if err != nil {
		panic(err)
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(rotkeys))
}

//export lattigo_getDataLenCiphertext
func lattigo_getDataLenCiphertext(ctHandle Handle9, withMetaData bool) uint64 {
	var ct *ckks.Ciphertext
	ct = getStoredCiphertext(ctHandle)
	return ct.GetDataLen(withMetaData)
}

//export lattigo_getDataLenParameters
func lattigo_getDataLenParameters(paramsHandle Handle9, withMetaData bool) uint64 {
	var params *ckks.Parameters
	params = getStoredParameters(paramsHandle)
	// see https://github.com/ldsec/lattigo/issues/115
	// return params.GetDataLen(withMetaData)
	paramBytes, err := params.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return uint64(len(paramBytes))
}

//export lattigo_getDataLenSecretKey
func lattigo_getDataLenSecretKey(skHandle Handle9, withMetaData bool) uint64 {
	var sk *ckks.SecretKey
	sk = getStoredSecretKey(skHandle)
	return sk.GetDataLen(withMetaData)
}

//export lattigo_getDataLenPublicKey
func lattigo_getDataLenPublicKey(pkHandle Handle9, withMetaData bool) uint64 {
	var pk *ckks.PublicKey
	pk = getStoredPublicKey(pkHandle)
	return pk.GetDataLen(withMetaData)
}

//export lattigo_getDataLenEvaluationKey
func lattigo_getDataLenEvaluationKey(evakeyHandle Handle9, withMetaData bool) uint64 {
	var evakey *ckks.EvaluationKey
	evakey = getStoredEvalKey(evakeyHandle)
	return evakey.GetDataLen(withMetaData)
}

//export lattigo_getDataLenRotationKeys
func lattigo_getDataLenRotationKeys(rotkeysHandle Handle9, withMetaData bool) uint64 {
	var rotkeys *ckks.RotationKeys
	rotkeys = getStoredRotationKeys(rotkeysHandle)
	return rotkeys.GetDataLen(withMetaData)
}
