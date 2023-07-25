# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

# We have to generate header files for each go file separately
# see https://github.com/golang/go/issues/35715
#
# The `go tool` command also doesn't like creating directories, so we have to create them externally
#
# Building a static library (-buildmode=c-archive) results in threading linker errors.
#
# I can't figure out how to build the go headers and library in-place without creating temporary files in the source directory.
# As a work-around, I copy the source files into the build directory, and then build them there.

add_custom_command(
  OUTPUT
    ${LATTIGO_LIB_FULL_PATH}
    ${CGO_HEADER_DST}/bootstrap.h
    ${CGO_HEADER_DST}/bootstrap_params.h
    ${CGO_HEADER_DST}/ciphertext.h
    ${CGO_HEADER_DST}/decryptor.h
    ${CGO_HEADER_DST}/encoder.h
    ${CGO_HEADER_DST}/encryptor.h
    ${CGO_HEADER_DST}/evaluator.h
    ${CGO_HEADER_DST}/keygen.h
    ${CGO_HEADER_DST}/marshaler.h
    ${CGO_HEADER_DST}/params.h
    ${CGO_HEADER_DST}/plaintext.h
    ${CGO_HEADER_DST}/precision.h
    ${CGO_HEADER_DST}/dckks.h
    ${CGO_HEADER_DST}/ring.h
    ${CGO_HEADER_DST}/utils.h
    ${CGO_HEADER_DST}/storage.h
  COMMAND cp -r ${CMAKE_CURRENT_SOURCE_DIR}/. .

  COMMAND go mod download github.com/tuneinsight/lattigo/v4
  COMMAND go get github.com/tuneinsight/lattigo/v4/utils@v4.1.0

  COMMAND go vet
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/ckks/bootstrap.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/ckks/bootstrap_params.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/ckks/ciphertext.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/ckks/decryptor.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/ckks/encoder.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/ckks/encryptor.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/ckks/evaluator.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/ckks/keygen.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/ckks/marshaler.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/ckks/params.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/ckks/plaintext.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/ckks/precision.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/ckks/dckks.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/ring/ring.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/utils/utils.go
  COMMAND go fmt ${CMAKE_CURRENT_SOURCE_DIR}/marshal/storage.go

  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/bootstrap.h ckks/bootstrap.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/bootstrap_params.h ckks/bootstrap_params.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/ciphertext.h ckks/ciphertext.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/decryptor.h ckks/decryptor.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/encoder.h ckks/encoder.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/encryptor.h ckks/encryptor.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/evaluator.h ckks/evaluator.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/keygen.h ckks/keygen.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/marshaler.h ckks/marshaler.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/params.h ckks/params.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/plaintext.h ckks/plaintext.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/precision.h ckks/precision.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/dckks.h ckks/dckks.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/ring.h ring/ring.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/utils.h utils/utils.go
  COMMAND go tool cgo -exportheader ${CGO_HEADER_DST}/storage.h marshal/storage.go
  COMMAND go build -buildmode=c-shared -o ${LATTIGO_LIB_FULL_PATH}
  WORKING_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}
  DEPENDS
    latticppMain.go
    ckks/bootstrap.go
    ckks/bootstrap_params.go
    ckks/ciphertext.go
    ckks/decryptor.go
    ckks/encoder.go
    ckks/encryptor.go
    ckks/evaluator.go
    ckks/keygen.go
    ckks/marshaler.go
    ckks/params.go
    ckks/plaintext.go
    ckks/precision.go
    ckks/dckks.go
    ring/ring.go
    utils/utils.go    
    marshal/storage.go
    go.mod
)

add_library(latticpp_gowrapper STATIC
    ${CGO_HEADER_DST}/bootstrap.h
    ${CGO_HEADER_DST}/bootstrap_params.h
    ${CGO_HEADER_DST}/ciphertext.h
    ${CGO_HEADER_DST}/decryptor.h
    ${CGO_HEADER_DST}/encoder.h
    ${CGO_HEADER_DST}/encryptor.h
    ${CGO_HEADER_DST}/evaluator.h
    ${CGO_HEADER_DST}/keygen.h
    ${CGO_HEADER_DST}/marshaler.h
    ${CGO_HEADER_DST}/params.h
    ${CGO_HEADER_DST}/plaintext.h
    ${CGO_HEADER_DST}/precision.h
    ${CGO_HEADER_DST}/dckks.h
    ${CGO_HEADER_DST}/ring.h
    ${CGO_HEADER_DST}/utils.h    
    ${CGO_HEADER_DST}/storage.h)
target_include_directories(latticpp_gowrapper PUBLIC ${CMAKE_BINARY_DIR})
set_target_properties(latticpp_gowrapper PROPERTIES LINKER_LANGUAGE CXX)
