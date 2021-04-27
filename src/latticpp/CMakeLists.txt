# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

add_subdirectory(ckks)

install(
    FILES
        ${CMAKE_CURRENT_LIST_DIR}/marshal/gohandle.h
    DESTINATION
        ${LATTICPP_INCLUDES_INSTALL_DIR}/marshal
)
