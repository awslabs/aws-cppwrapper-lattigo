#!/bin/bash -ex
# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

cmake -Bbuild -GNinja . -DLATTICPP_BUILD_EXAMPLES=ON
# build all targets
ninja -Cbuild
ninja -Cbuild run_bootstrapexample
ninja -Cbuild run_eulerexample
