// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "evaluator.h"

namespace latticpp {

    Evaluator newEvaluator(const Parameters &params) {
        return Evaluator(lattigo_newEvaluator(params.getRawHandle()));
    }

    void multByConst(const Evaluator &eval, const Ciphertext &ctIn, double constant, const Ciphertext &ctOut) {
        lattigo_multByConst(eval.getRawHandle(), ctIn.getRawHandle(), constant, ctOut.getRawHandle());
    }

    void addConst(const Evaluator &eval, const Ciphertext &ctIn, double constant, const Ciphertext &ctOut) {
        lattigo_addConst(eval.getRawHandle(), ctIn.getRawHandle(), constant, ctOut.getRawHandle());
    }

    void rescale(const Evaluator &eval, const Ciphertext &ctIn, double threshold, const Ciphertext &ctOut) {
        lattigo_rescale(eval.getRawHandle(), ctIn.getRawHandle(), threshold, ctOut.getRawHandle());
    }

    Ciphertext mulRelinNew(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1, const EvaluationKey &evakey) {
        return Ciphertext(lattigo_mulRelinNew(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle(), evakey.getRawHandle()));
    }

    void mulRelin(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1, const EvaluationKey &evakey, const Ciphertext &ctOut) {
        lattigo_mulRelin(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle(), evakey.getRawHandle(), ctOut.getRawHandle());
    }

    void add(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1, const Ciphertext &ctOut) {
        lattigo_add(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle(), ctOut.getRawHandle());
    }

    void sub(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1, const Ciphertext &ctOut) {
        lattigo_sub(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle(), ctOut.getRawHandle());
    }

    void multByGaussianIntegerAndAdd(const Evaluator &eval, const Ciphertext &ctIn, uint64_t cReal, uint64_t cImag, const Ciphertext &ctOut) {
        lattigo_multByGaussianIntegerAndAdd(eval.getRawHandle(), ctIn.getRawHandle(), cReal, cImag, ctOut.getRawHandle());
    }

    void dropLevel(const Evaluator &eval, const Ciphertext &ct, uint64_t levels) {
        lattigo_dropLevel(eval.getRawHandle(), ct.getRawHandle(), levels);
    }
}  // namespace latticpp
