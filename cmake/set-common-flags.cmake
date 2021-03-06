# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

function(set_common_flags target_name)
    target_compile_options(${target_name} PRIVATE -std=c++17 -Wall -Werror -Wformat=2 -Wwrite-strings -Wvla
            -fvisibility=hidden -fno-common -funsigned-char -Wextra -Wunused -Wcomment -Wchar-subscripts -Wuninitialized
            -Wunused-result -Wfatal-errors)
    if (CMAKE_CXX_COMPILER_ID STREQUAL "Clang")
        target_compile_options(${target_name} PRIVATE -Wmissing-declarations -Wmissing-field-initializers -Wshadow
                -Wpedantic -Wno-c99-extensions)
    endif()
    #-Wextra turns on sign-compare which is strict on comparing loop indexes (int) with size_t from vector length
    target_compile_options(${target_name} PRIVATE -Wno-sign-compare)
endfunction()
