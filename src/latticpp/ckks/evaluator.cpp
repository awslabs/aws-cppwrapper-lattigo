// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "evaluator.h"

namespace latticpp {

    GoHandle<Evaluator> newEvaluator(const GoHandle<Parameters> &params) {
        return GoHandle<Evaluator>(lattigo_newEvaluator(params.getRawHandle()));
    }

    void multByConst(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ctIn, double constant, const GoHandle<Ciphertext> &ctOut) {
        lattigo_multByConst(eval.getRawHandle(), ctIn.getRawHandle(), constant, ctOut.getRawHandle());
    }

    void addConst(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ctIn, double constant, const GoHandle<Ciphertext> &ctOut) {
        lattigo_addConst(eval.getRawHandle(), ctIn.getRawHandle(), constant, ctOut.getRawHandle());
    }

    void rescale(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ctIn, double threshold, const GoHandle<Ciphertext> &ctOut) {
        lattigo_rescale(eval.getRawHandle(), ctIn.getRawHandle(), threshold, ctOut.getRawHandle());
    }

    GoHandle<Ciphertext> mulRelinNew(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ct0, const GoHandle<Ciphertext> &ct1, const GoHandle<RelinKey> &evakey) {
        return GoHandle<Ciphertext>(lattigo_mulRelinNew(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle(), evakey.getRawHandle()));
    }

    void mulRelin(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ct0, const GoHandle<Ciphertext> &ct1, const GoHandle<RelinKey> &evakey, const GoHandle<Ciphertext> &ctOut) {
        lattigo_mulRelin(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle(), evakey.getRawHandle(), ctOut.getRawHandle());
    }

    void add(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ct0, const GoHandle<Ciphertext> &ct1, const GoHandle<Ciphertext> &ctOut) {
        lattigo_add(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle(), ctOut.getRawHandle());
    }

    void sub(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ct0, const GoHandle<Ciphertext> &ct1, const GoHandle<Ciphertext> &ctOut) {
        lattigo_sub(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle(), ctOut.getRawHandle());
    }

    void multByGaussianIntegerAndAdd(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ctIn, uint64_t cReal, uint64_t cImag, const GoHandle<Ciphertext> &ctOut) {
        lattigo_multByGaussianIntegerAndAdd(eval.getRawHandle(), ctIn.getRawHandle(), cReal, cImag, ctOut.getRawHandle());
    }

    void dropLevel(const GoHandle<Evaluator> &eval, const GoHandle<Ciphertext> &ct, uint64_t levels) {
        lattigo_dropLevel(eval.getRawHandle(), ct.getRawHandle(), levels);
    }
}  // namespace latticpp
