// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "precision.h"

using namespace std;

namespace latticpp {


    std::string precisionStats(const Parameters &params, const Encoder &encoder, const std::vector<double> &expectedValues, const std::vector<double> &actualValues) {
        int len = expectedValues.size();
        if (len != actualValues.size()) {
            throw invalid_argument("Inputs to precisionStats do not have the same length.");
        }

        char* strData = lattigo_precisionStats(params.getRawHandle(), encoder.getRawHandle(), expectedValues.data(), actualValues.data(), len);
        string str = string(strData);
        free(strData);
        return str;
    }
}  // namespace latticpp
