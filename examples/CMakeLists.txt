# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

add_executable(eulerexample ${CMAKE_CURRENT_SOURCE_DIR}/euler.cpp)
target_link_libraries(eulerexample aws-lattigo-cpp)
add_custom_target(
  run_eulerexample
  COMMAND bin/${CMAKE_BUILD_TYPE}/eulerexample
  WORKING_DIRECTORY ${LATTICPP_ROOT_DIR}
  DEPENDS eulerexample)

add_executable(bootstrapexample ${CMAKE_CURRENT_SOURCE_DIR}/bootstrapping.cpp)
target_link_libraries(bootstrapexample aws-lattigo-cpp)
add_custom_target(
  run_bootstrapexample
  COMMAND bin/${CMAKE_BUILD_TYPE}/bootstrapexample
  WORKING_DIRECTORY ${LATTICPP_ROOT_DIR}
  DEPENDS bootstrapexample)

add_executable(multikeyexample ${CMAKE_CURRENT_SOURCE_DIR}/multikey.cpp)
target_link_libraries(multikeyexample aws-lattigo-cpp)
add_custom_target(
  run_multikeyexample
  COMMAND bin/${CMAKE_BUILD_TYPE}/multikeyexample
  WORKING_DIRECTORY ${LATTICPP_ROOT_DIR}
  DEPENDS multikeyexample)