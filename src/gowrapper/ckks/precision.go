// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

/*
typedef const double constDouble;

#include "stdint.h"
struct Lattigo_String {
  char *data;
  uint64_t len;
};
*/
import "C"

import (
	"math"

	"github.com/tuneinsight/lattigo/v4/ckks"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle12 = uint64

//export lattigo_precisionStats
func lattigo_precisionStats(paramHandle Handle12, encoderHandle Handle12, expectedValues *C.constDouble, actualValues *C.constDouble, length uint64) *C.char {

	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)

	var encoder *ckks.Encoder
	encoder = getStoredEncoder(encoderHandle)

	expectedComplexValues := CDoubleVecToGoComplex(expectedValues, length)
	actualComplexValues := CDoubleVecToGoComplex(actualValues, length)

	var prec ckks.PrecisionStats
	prec = ckks.GetPrecisionStats(*params, *encoder, nil, expectedComplexValues, actualComplexValues, int(math.Log2(float64(length))), 0)

	return C.CString(prec.String())
}
