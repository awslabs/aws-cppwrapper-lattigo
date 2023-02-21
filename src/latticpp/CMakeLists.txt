# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

add_subdirectory(ckks)
add_subdirectory(utils)
add_subdirectory(ring)

install(
    FILES
        ${CMAKE_CURRENT_LIST_DIR}/marshal/gohandle.h
    DESTINATION
        ${LATTICPP_INCLUDES_INSTALL_DIR}/marshal
)

install(
    FILES
        ${CMAKE_CURRENT_LIST_DIR}/latticpp.h
    DESTINATION
        ${LATTICPP_INCLUDES_INSTALL_DIR}
)
