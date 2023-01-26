// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

import "C"

import (
	"lattigo-cpp/marshal"

	"github.com/ldsec/lattigo/v2/ckks"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle7 = uint64

func getStoredPlaintext(ptHandle Handle7) *ckks.Plaintext {
	ref := marshal.CrossLangObjMap.Get(ptHandle)
	return (*ckks.Plaintext)(ref.Ptr)
}
