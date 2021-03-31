// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#pragma once

#include "latticpp/marshall/gohandle.h"
#include "cgo/evaluator.h"

namespace latticpp {

    GoHandle<Evaluator> newEvaluator(const GoHandle<Parameters> &params);

    void multByConst(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ctIn, double constant, const GoHandle<Ciphertext> &ctOut);

    void addConst(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ctIn, double constant, const GoHandle<Ciphertext> &ctOut);

    void rescale(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ctIn, double threshold, const GoHandle<Ciphertext> &ctOut);

    GoHandle<Ciphertext> mulRelinNew(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ct0, const GoHandle<Ciphertext> &ct1, const GoHandle<RelinKey> &evakey);

    void mulRelin(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ct0, const GoHandle<Ciphertext> &ct1, const GoHandle<RelinKey> &evakey, const GoHandle<Ciphertext> &ctOut);

    void add(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ct0, const GoHandle<Ciphertext> &ct1, const GoHandle<Ciphertext> &ctOut);

    void sub(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ct0, const GoHandle<Ciphertext> &ct1, const GoHandle<Ciphertext> &ctOut);

    void multByGaussianIntegerAndAdd(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ctIn, uint64_t cReal, uint64_t cImag, const GoHandle<Ciphertext> &ctOut);

    void dropLevel(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ct, uint64_t levels);
}  // namespace latticpp
