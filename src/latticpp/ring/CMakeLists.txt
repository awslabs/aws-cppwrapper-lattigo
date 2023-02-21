# SPDX-License-Identifier: Apache-2.0

target_sources(latticpp_obj
    PRIVATE
        ${CMAKE_CURRENT_LIST_DIR}/ring.cpp
)

install(
    FILES
        ${CMAKE_CURRENT_LIST_DIR}/ring.h
    DESTINATION
        ${LATTICPP_INCLUDES_INSTALL_DIR}/ring
)
