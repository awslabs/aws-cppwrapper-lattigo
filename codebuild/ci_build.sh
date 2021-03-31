#!/bin/bash -ex
# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

cmake -Bbuild -GNinja .
# build all targets
ninja -Cbuild

# This script defines -ex on line 1, -e exits on any error, +e turns that off due to the known issue with this test
# set +e
