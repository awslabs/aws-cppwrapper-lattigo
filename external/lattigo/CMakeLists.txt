# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

cmake_minimum_required(VERSION 3.12)

project(LATTIGO_DOWNLOAD VERSION 2.1.1)

# suppress a CMake warning about an unused variable
set(SUPPRESS_VAR_WARNING ${3P_INSTALL_DIR})

if (EXISTS ${CMAKE_CURRENT_LIST_DIR}/src)
    set(DOWNLOAD_URL)
    message(STATUS "Found source code for Lattigo in ${CMAKE_CURRENT_LIST_DIR}.")
else()
    message(STATUS "No source code found for Lattigo; downloading in ${CMAKE_CURRENT_LIST_DIR}.")
    include(ExternalProject)
    ExternalProject_Add(EP_Lattigo
        SOURCE_DIR           ${CMAKE_CURRENT_LIST_DIR}/src
        GIT_REPOSITORY       https://github.com/ldsec/lattigo.git
        GIT_TAG              v${PROJECT_VERSION}
        GIT_CONFIG           advice.detachedHead=false
        UPDATE_COMMAND       ""
        CONFIGURE_COMMAND    ""
        BUILD_COMMAND        ""
        INSTALL_COMMAND      ""
        TEST_COMMAND         ""
        TEST_COMMAND         ""
    )
endif()
