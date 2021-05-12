// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "evaluator.h"

namespace latticpp {

    Evaluator newEvaluator(const Parameters &params) {
        return Evaluator(lattigo_newEvaluator(params.getRawHandle()));
    }

    void rotate(const Evaluator &eval, const Ciphertext &ctIn, uint64_t k, const RotationKeys &rotKeys, Ciphertext &ctOut) {
        lattigo_rotate(eval.getRawHandle(), ctIn.getRawHandle(), k, rotKeys.getRawHandle(), ctOut.getRawHandle());
    }

    void multByConst(const Evaluator &eval, const Ciphertext &ctIn, double constant, Ciphertext &ctOut) {
        lattigo_multByConst(eval.getRawHandle(), ctIn.getRawHandle(), constant, ctOut.getRawHandle());
    }

    void addConst(const Evaluator &eval, const Ciphertext &ctIn, double constant, Ciphertext &ctOut) {
        lattigo_addConst(eval.getRawHandle(), ctIn.getRawHandle(), constant, ctOut.getRawHandle());
    }

    void rescale(const Evaluator &eval, const Ciphertext &ctIn, double threshold, Ciphertext &ctOut) {
        lattigo_rescale(eval.getRawHandle(), ctIn.getRawHandle(), threshold, ctOut.getRawHandle());
    }

    void rescaleMany(const Evaluator &eval, const Ciphertext &ctIn, uint64_t numRescales, Ciphertext &ctOut) {
        lattigo_rescaleMany(eval.getRawHandle(), ctIn.getRawHandle(), numRescales, ctOut.getRawHandle());
    }

    Ciphertext mulRelinNew(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1, const EvaluationKey &evakey) {
        return Ciphertext(lattigo_mulRelinNew(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle(), evakey.getRawHandle()));
    }

    void mulRelin(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1, const EvaluationKey &evakey, Ciphertext &ctOut) {
        lattigo_mulRelin(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle(), evakey.getRawHandle(), ctOut.getRawHandle());
    }

    void mul(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1, Ciphertext &ctOut) {
        lattigo_mul(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle(), ctOut.getRawHandle());
    }

    void mulPlain(const Evaluator &eval, const Ciphertext &ctIn, const Plaintext &pt, Ciphertext &ctOut) {
        lattigo_mulPlain(eval.getRawHandle(), ctIn.getRawHandle(), pt.getRawHandle(), ctOut.getRawHandle());
    }

    void add(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1, Ciphertext &ctOut) {
        lattigo_add(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle(), ctOut.getRawHandle());
    }

    void addPlain(const Evaluator &eval, const Ciphertext &ctIn, const Plaintext &pt, Ciphertext &ctOut) {
        lattigo_addPlain(eval.getRawHandle(), ctIn.getRawHandle(), pt.getRawHandle(), ctOut.getRawHandle());
    }

    void neg(const Evaluator &eval, const Ciphertext &ctIn, Ciphertext &ctOut) {
        lattigo_neg(eval.getRawHandle(), ctIn.getRawHandle(), ctOut.getRawHandle());
    }

    void sub(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1, Ciphertext &ctOut) {
        lattigo_sub(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle(), ctOut.getRawHandle());
    }

    void subPlain(const Evaluator &eval, const Ciphertext &ctIn, const Plaintext &pt, Ciphertext &ctOut) {
        lattigo_subPlain(eval.getRawHandle(), ctIn.getRawHandle(), pt.getRawHandle(), ctOut.getRawHandle());
    }

    void multByGaussianIntegerAndAdd(const Evaluator &eval, const Ciphertext &ctIn, uint64_t cReal, uint64_t cImag, Ciphertext &ctOut) {
        lattigo_multByGaussianIntegerAndAdd(eval.getRawHandle(), ctIn.getRawHandle(), cReal, cImag, ctOut.getRawHandle());
    }

    void dropLevel(const Evaluator &eval, Ciphertext &ct, uint64_t levels) {
        lattigo_dropLevel(eval.getRawHandle(), ct.getRawHandle(), levels);
    }

    void relinearize(const Evaluator &eval, const Ciphertext &ctIn, const EvaluationKey &evakey, Ciphertext &ctOut) {
        lattigo_relinearize(eval.getRawHandle(), ctIn.getRawHandle(), evakey.getRawHandle(), ctOut.getRawHandle());
    }
}  // namespace latticpp
