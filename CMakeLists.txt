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
set(LATTICPP_THIRD_PARTY_DIR ${CMAKE_CURRENT_SOURCE_DIR}/external)

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
set(3P_INSTALL_DIR ${CMAKE_CURRENT_BINARY_DIR}/external)

# This is where most third-party dependencies are installed; adding the path
# to CMAKE_PREFIX_PATH allows `find_package()` to succeed on these dependencies
list(APPEND CMAKE_PREFIX_PATH ${3P_INSTALL_DIR})

if(NOT LATTICPP_ROOT_DIR)
  set(LATTICPP_ROOT_DIR "${PROJECT_SOURCE_DIR}/")
endif()

########################
# Custom Cmake Scripts #
########################

# Include custom cmake module path
list(APPEND CMAKE_MODULE_PATH ${CMAKE_CURRENT_SOURCE_DIR}/cmake)
list(APPEND CMAKE_MODULE_PATH ${CMAKE_CURRENT_SOURCE_DIR}/cmake/modules)
# And custom modules
include(set-common-flags)
include(download-external-project)

############################
# Third party dependencies #
############################

# Using third-party dependencies is a multi-step process.
# 1. The process starts in cmake/modules/FindFOOLib.cmake
# 2. For some dependencies, we check if the dependency is available on the system with `find_package`.
#    If the dependency is found, stop here.
# 3. Otherwise, we invoke `download_external_project()`. In turn, this calls `cmake` on the make_directory
#    third-party/FOO/, which includes a CMakeLists.txt with download and build instructions. The
#    invocation of cmake generates the build system for the dependency.
# 4. Next, `download_external_project()` calls `cmake --build`, which actually builds the dependency
#    we downloaded and configured in the previous step.
# 5. Finally, FindFOOLib.cmake applies any one-off fixups to make the dependency usable by lattigo-cpp.
# Boost is an exception to the process above, since it is a large header-only library and we only
# need a tiny fraction of it. The required headers are zipped and included in the repository,
# so if Boost is not installed on the system, we just extract the headers.

find_package(LattigoLib REQUIRED)

###############
# lattigo-cpp Library #
###############



set(CGO_INCLUDE_PATH ${CMAKE_CURRENT_BINARY_DIR})
set(CGO_BUILD_ROOT ${CGO_INCLUDE_PATH}/cgo)
set(PRNG_LIB_NAME liblattigo.so)
set(LIB_FULL_PATH ${CGO_BUILD_ROOT}/${PRNG_LIB_NAME})
file(MAKE_DIRECTORY ${CGO_BUILD_ROOT})

# Add the main source tree object library
add_subdirectory(src)

# Always build the static library
add_library(aws-lattigo-cpp STATIC $<TARGET_OBJECTS:lattigo_cpp_obj>)
target_compile_features(aws-lattigo-cpp PUBLIC cxx_std_17)
target_include_directories(aws-lattigo-cpp
    PUBLIC
        $<BUILD_INTERFACE:${LATTICPP_INCLUDES_BUILD_DIR}>
        $<INSTALL_INTERFACE:${LATTICPP_INCLUDES_INSTALL_DIR}>
    PRIVATE
        ${CGO_INCLUDE_PATH}
        ${LATTICPP_SOURCE_DIR}
)
set_common_flags(aws-lattigo-cpp)

# Link dependencies
target_link_libraries(aws-lattigo-cpp ${LIB_FULL_PATH} lattigo_cpp_gowrapper)

install(TARGETS aws-lattigo-cpp EXPORT LatticppTargets
        ARCHIVE DESTINATION ${CMAKE_INSTALL_LIBDIR}
        LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
        RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR})

# We export LatticppTargets from the build tree so it can be used by other projects
# without requiring an install.
export(
    EXPORT LatticppTargets
    NAMESPACE latticpp::
    FILE ${LATTICPP_TARGETS_FILENAME})

export(TARGETS aws-lattigo-cpp lattigo_cpp_gowrapper FILE ${LATTICPP_TARGETS_FILENAME})

############
# Examples #
############

# Build wrapped examples
option(LATTICPP_BUILD_EXAMPLES "Build the wrapped Lattigo examples." OFF)
if (LATTICPP_BUILD_EXAMPLES)
    add_subdirectory(examples)
endif ()
