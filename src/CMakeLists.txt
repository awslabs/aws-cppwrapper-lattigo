# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

# Create an object library to compile sources only once
# Sources are added by the CMakeLists in the source tree via `target_sources`
add_library(lattigo_cpp_obj OBJECT)
# Declare a dependency on lattigo_cpp_gowrapper so that the Go library and header files
# are generated prior to compiling lattigo_cpp_obj, which imports the generated headers
add_dependencies(lattigo_cpp_obj lattigo_cpp_gowrapper)

add_subdirectory(gowrapper)
add_subdirectory(latticpp)

target_include_directories(lattigo_cpp_obj
    PRIVATE
        $<BUILD_INTERFACE:${LATTICPP_INCLUDES_BUILD_DIR}>
        $<INSTALL_INTERFACE:${LATTICPP_INCLUDES_INSTALL_DIR}>
        ${LATTICPP_SOURCE_DIR}
        ${CGO_INCLUDE_PATH}
    )

set_common_flags(lattigo_cpp_obj)
