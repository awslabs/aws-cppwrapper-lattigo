# SPDX-License-Identifier: Apache-2.0

target_sources(latticpp_obj
    PRIVATE
        ${CMAKE_CURRENT_LIST_DIR}/utils.cpp
)

install(
    FILES
        ${CMAKE_CURRENT_LIST_DIR}/utils.h
    DESTINATION
        ${LATTICPP_INCLUDES_INSTALL_DIR}/utils
)
