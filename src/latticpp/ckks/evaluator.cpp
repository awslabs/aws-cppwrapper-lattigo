// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "evaluator.h"

using namespace std;

namespace latticpp {

    Evaluator newEvaluator(const Parameters &params, const EvaluationKey &evalKey) {
        return Evaluator(lattigo_newEvaluator(params.getRawHandle(), evalKey.getRawHandle()));
    }

    Evaluator evaluatorWithKey(const Evaluator &eval,
                               const EvaluationKey &evalKey) {
      return Evaluator(lattigo_evaluatorWithKey(eval.getRawHandle(),
                                                evalKey.getRawHandle()));
    }

    void rotate(const Evaluator &eval, const Ciphertext &ctIn, uint64_t k, Ciphertext &ctOut) {
        lattigo_rotate(eval.getRawHandle(), ctIn.getRawHandle(), k, ctOut.getRawHandle());
    }

    vector<Ciphertext> rotateHoisted(const Evaluator &eval, const Ciphertext &ctIn, vector<uint64_t> ks) {
        vector<uint64_t> outputHandles(ks.size());
        lattigo_rotateHoisted(eval.getRawHandle(), ctIn.getRawHandle(), ks.data(), ks.size(), outputHandles.data());
        vector<Ciphertext> outputCts(ks.size());
        for (int i = 0; i < ks.size(); i++) {
            outputCts[i] = Ciphertext(outputHandles[i]);
        }
        return outputCts;
    }

    void multByConst(const Evaluator &eval, const Ciphertext &ctIn, double constant, Ciphertext &ctOut) {
        lattigo_multByConst(eval.getRawHandle(), ctIn.getRawHandle(), constant, ctOut.getRawHandle());
    }

    void addConst(const Evaluator &eval, const Ciphertext &ctIn, double constant, Ciphertext &ctOut) {
        lattigo_addConst(eval.getRawHandle(), ctIn.getRawHandle(), constant, ctOut.getRawHandle());
    }

    void rescale(const Evaluator &eval, const Ciphertext &ctIn, double scale, Ciphertext &ctOut) {
        lattigo_rescale(eval.getRawHandle(), ctIn.getRawHandle(), scale, ctOut.getRawHandle());
    }

    Ciphertext mulRelinNew(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1) {
        return Ciphertext(lattigo_mulRelinNew(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle()));
    }

    void mulRelin(const Evaluator &eval, const Ciphertext &ct0, const Ciphertext &ct1, Ciphertext &ctOut) {
        lattigo_mulRelin(eval.getRawHandle(), ct0.getRawHandle(), ct1.getRawHandle(), ctOut.getRawHandle());
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

    void relinearize(const Evaluator &eval, const Ciphertext &ctIn, Ciphertext &ctOut) {
        lattigo_relinearize(eval.getRawHandle(), ctIn.getRawHandle(), ctOut.getRawHandle());
    }

    void switchKeys(const Evaluator &eval, const Ciphertext &ctxIn, const SwitchingKey &swk, const Ciphertext &ctxOut) {
        lattigo_switchKeys(eval.getRawHandle(), ctxIn.getRawHandle(), swk.getRawHandle(), ctxOut.getRawHandle());
    }
} // namespace latticpp