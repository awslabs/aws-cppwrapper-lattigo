# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

# We have to generate header files for each go file separately
# see https://github.com/golang/go/issues/35715
#
# The `go tool` command also doesn't like creating directories, so we have to create them externally
#
# Buidling a static library (-buildmode=c-archive) results in threading linker errors.
add_custom_command(
  OUTPUT
    ${LIB_FULL_PATH}
    ${CGO_BUILD_ROOT}/params.h
    ${CGO_BUILD_ROOT}/storage.h
  COMMAND go tool cgo -exportheader ${CGO_BUILD_ROOT}/params.h ckks/params.go
  COMMAND go tool cgo -exportheader ${CGO_BUILD_ROOT}/storage.h marshall/storage.go
  COMMAND go build -buildmode=c-shared -o ${LIB_FULL_PATH}
  COMMAND rm -rf _obj
  WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
  DEPENDS
    latticppMain.go
    ckks/params.go
    marshall/storage.go
)

add_library(lattigo_cpp_gowrapper STATIC
    ${CGO_BUILD_ROOT}/params.h
    ${CGO_BUILD_ROOT}/storage.h)
target_include_directories(lattigo_cpp_gowrapper PUBLIC ${CMAKE_BINARY_DIR})
set_target_properties(lattigo_cpp_gowrapper PROPERTIES LINKER_LANGUAGE CXX)
