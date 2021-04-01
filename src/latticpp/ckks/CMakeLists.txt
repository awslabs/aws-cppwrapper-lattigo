# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

target_sources(lattigo_cpp_obj
    PRIVATE
        ${CMAKE_CURRENT_LIST_DIR}/params.cpp
)

install(
    FILES
        ${CMAKE_CURRENT_LIST_DIR}/params.h
    DESTINATION
        ${LATTICPP_INCLUDES_INSTALL_DIR}/ckks
)
