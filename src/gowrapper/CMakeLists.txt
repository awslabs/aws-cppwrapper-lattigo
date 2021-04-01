# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

# We have to generate header files for each go file separately
# see https://github.com/golang/go/issues/35715
#
# The `go tool` command also doesn't like creating directories, so we have to create them externally
#
# Buidling a static library (-buildmode=c-archive) results in threading linker errors.
#
# I can't figure out how to build the go headers and library in-place without creating temporary files in the source directory.
# As a work-around, I copy the source files into the build directory, and then build them there.

add_custom_command(
  OUTPUT
    ${LIB_FULL_PATH}
    ${CGO_HEADER_DST}/params.h
    ${CGO_HEADER_DST}/storage.h
  COMMAND cp -r ${CMAKE_CURRENT_SOURCE_DIR}/. .
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/params.h ckks/params.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/storage.h marshall/storage.go
  COMMAND go build -buildmode=c-shared -o ${LIB_FULL_PATH}
  WORKING_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}
  DEPENDS
    latticppMain.go
    ckks/params.go
    marshall/storage.go
)

add_library(lattigo_cpp_gowrapper STATIC
    ${CGO_HEADER_DST}/params.h
    ${CGO_HEADER_DST}/storage.h)
target_include_directories(lattigo_cpp_gowrapper PUBLIC ${CMAKE_BINARY_DIR})
set_target_properties(lattigo_cpp_gowrapper PROPERTIES LINKER_LANGUAGE CXX)
