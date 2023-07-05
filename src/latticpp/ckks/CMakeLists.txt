# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

target_sources(latticpp_obj
    PRIVATE
        ${CMAKE_CURRENT_LIST_DIR}/bootstrap.cpp
        ${CMAKE_CURRENT_LIST_DIR}/bootstrap_params.cpp
        ${CMAKE_CURRENT_LIST_DIR}/ciphertext.cpp
        ${CMAKE_CURRENT_LIST_DIR}/dckks.cpp
        ${CMAKE_CURRENT_LIST_DIR}/decryptor.cpp
        ${CMAKE_CURRENT_LIST_DIR}/encoder.cpp
        ${CMAKE_CURRENT_LIST_DIR}/encryptor.cpp
        ${CMAKE_CURRENT_LIST_DIR}/evaluator.cpp
        ${CMAKE_CURRENT_LIST_DIR}/keygen.cpp
        ${CMAKE_CURRENT_LIST_DIR}/marshaler.cpp
        ${CMAKE_CURRENT_LIST_DIR}/params.cpp
        ${CMAKE_CURRENT_LIST_DIR}/plaintext.cpp
        ${CMAKE_CURRENT_LIST_DIR}/precision.cpp
)

install(
    FILES
        ${CMAKE_CURRENT_LIST_DIR}/bootstrap.h
        ${CMAKE_CURRENT_LIST_DIR}/bootstrap_params.h
        ${CMAKE_CURRENT_LIST_DIR}/ciphertext.h
        ${CMAKE_CURRENT_LIST_DIR}/dckks.cpp
        ${CMAKE_CURRENT_LIST_DIR}/decryptor.h
        ${CMAKE_CURRENT_LIST_DIR}/encoder.h
        ${CMAKE_CURRENT_LIST_DIR}/encryptor.h
        ${CMAKE_CURRENT_LIST_DIR}/evaluator.h
        ${CMAKE_CURRENT_LIST_DIR}/keygen.h
        ${CMAKE_CURRENT_LIST_DIR}/marshaler.h
        ${CMAKE_CURRENT_LIST_DIR}/params.h
        ${CMAKE_CURRENT_LIST_DIR}/precision.h
    DESTINATION
        ${LATTICPP_INCLUDES_INSTALL_DIR}/ckks
)
