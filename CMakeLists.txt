# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

cmake_minimum_required(VERSION 3.12)

project(lattigo-cpp VERSION 0.0.1 LANGUAGES CXX)

#################
# CMAKE OPTIONS #
#################

# Define global variables.
# Build in Release mode by default; otherwise use selected option
if (NOT CMAKE_BUILD_TYPE)
    set(CMAKE_BUILD_TYPE "Release" CACHE STRING "Build type" FORCE)
    set_property(CACHE CMAKE_BUILD_TYPE PROPERTY STRINGS "Release" "Debug" "MinSizeRel" "RelWithDebInfo")
endif ()
message(STATUS "Build type (CMAKE_BUILD_TYPE): ${CMAKE_BUILD_TYPE}")

#########
# Paths #
#########

set(LATTICPP_SOURCE_DIR ${CMAKE_CURRENT_SOURCE_DIR}/src/latticpp)

# Required files and directories
include(GNUInstallDirs)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/lib/${CMAKE_BUILD_TYPE})
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/lib/${CMAKE_BUILD_TYPE})
set(CMAKE_LIBRARY_RUNTIME_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/bin/${CMAKE_BUILD_TYPE})
set(LATTICPP_TARGETS_FILENAME ${CMAKE_CURRENT_BINARY_DIR}/cmake/latticppTargets.cmake)
set(LATTICPP_CONFIG_INSTALL_DIR ${CMAKE_INSTALL_LIBDIR}/cmake/latticpp)
set(LATTICPP_INCLUDES_INSTALL_DIR ${CMAKE_INSTALL_INCLUDE_DIR}/latticpp)
set(LATTICPP_INCLUDES_BUILD_DIR ${CMAKE_CURRENT_SOURCE_DIR}/src)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/bin/${CMAKE_BUILD_TYPE})

if(NOT LATTICPP_ROOT_DIR)
  set(LATTICPP_ROOT_DIR "${PROJECT_SOURCE_DIR}/")
endif()

########################
# Custom Cmake Scripts #
########################

# Include custom cmake module path
list(APPEND CMAKE_MODULE_PATH ${CMAKE_CURRENT_SOURCE_DIR}/cmake)
# And custom modules
include(set-common-flags)

###############
# lattigo-cpp Library #
###############

set(CGO_BUILD_ROOT ${CMAKE_CURRENT_BINARY_DIR}/cgo_output)
set(CGO_INCLUDE_PATH ${CGO_BUILD_ROOT}/cgo_output/include)
set(CGO_HEADER_DST ${CGO_INCLUDE_PATH}/cgo)
set(PRNG_LIB_NAME liblattigo.so)
set(LIB_FULL_PATH ${CGO_BUILD_ROOT}/${PRNG_LIB_NAME})
file(MAKE_DIRECTORY ${CGO_BUILD_ROOT})
