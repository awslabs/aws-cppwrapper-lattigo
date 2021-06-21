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
	"github.com/ldsec/lattigo/v2/ckks"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle12 = uint64

//export lattigo_precisionStats
func lattigo_precisionStats(paramHandle Handle12, expectedValues *C.constDouble, actualValues *C.constDouble, length uint64) *C.char {

	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)

	expectedComplexValues := CDoubleVecToGoComplex(expectedValues, length)
	actualComplexValues := CDoubleVecToGoComplex(actualValues, length)

	var prec ckks.PrecisionStats
	prec = ckks.GetPrecisionStats(params, nil, nil, expectedComplexValues, actualComplexValues)

	return C.CString(prec.String())
}
