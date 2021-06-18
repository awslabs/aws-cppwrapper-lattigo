# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

add_executable(sigmoidexample ${CMAKE_CURRENT_SOURCE_DIR}/sigmoid.cpp)
target_link_libraries(sigmoidexample aws-lattigo-cpp)
add_custom_target(
  run_sigmoidexample
  COMMAND bin/${CMAKE_BUILD_TYPE}/sigmoidexample
  WORKING_DIRECTORY ${LATTICPP_ROOT_DIR}
  DEPENDS sigmoidexample)

add_executable(bootstrapexample ${CMAKE_CURRENT_SOURCE_DIR}/bootstrapping.cpp)
target_link_libraries(bootstrapexample aws-lattigo-cpp)
add_custom_target(
  run_bootstrapexample
  COMMAND bin/${CMAKE_BUILD_TYPE}/bootstrapexample
  WORKING_DIRECTORY ${LATTICPP_ROOT_DIR}
  DEPENDS bootstrapexample)
