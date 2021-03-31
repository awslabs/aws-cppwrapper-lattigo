# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

# Download and install dependency - Lattigo.
download_external_project("lattigo")
find_package(lattigo 2.1.1 REQUIRED)
