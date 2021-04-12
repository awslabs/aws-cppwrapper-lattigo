# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

add_executable(gobindingtest ${CMAKE_CURRENT_SOURCE_DIR}/sigmoid.cpp)
target_link_libraries(gobindingtest aws-lattigo-cpp)
add_custom_target(
  run_gobindingtest
  COMMAND bin/${CMAKE_BUILD_TYPE}/gobindingtest
  WORKING_DIRECTORY ${LATTICPP_ROOT_DIR}
  DEPENDS gobindingtest)
