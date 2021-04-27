// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshal/gohandle.h"
#include "cgo/evaluator.h"

namespace latticpp {

    Evaluator newEvaluator(const Parameters &params);

    void multByConst(const Evaluator &eval, const Ciphertext &ctIn, double constant, const Ciphertext &ctOut);

    void addConst(const Evaluator &eval, const Ciphertext &ctIn, double constant, const Ciphertext &ctOut);

    void rescale(const Evaluator &eval, const Ciphertext &ctIn, double threshold, const Ciphertext &ctOut);

    Ciphertext mulRelinNew(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1, const RelinKeys &evakey);

    void mulRelin(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1, const RelinKeys &evakey, const Ciphertext &ctOut);

    void add(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1, const Ciphertext &ctOut);

    void sub(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1, const Ciphertext &ctOut);

    void multByGaussianIntegerAndAdd(const Evaluator &eval, const Ciphertext &ctIn, uint64_t cReal, uint64_t cImag, const Ciphertext &ctOut);

    void dropLevel(const Evaluator &eval, const Ciphertext &ct, uint64_t levels);
}  // namespace latticpp
