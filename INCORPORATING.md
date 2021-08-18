# Incorporating latticpp into a CMake project

## Including latticpp source
The easiest way to consume latticpp with your project's build system is to include the latticpp source directly in your repository (say in `third-party/aws-latticpp`). Then in your project's CMakeLists.txt, just include:

```
# Add 'third-party/aws-latticpp/src', which defines 'aws-latticpp' target.
add_subdirectory(third-party/aws-latticpp)

# Define your project target.
add_executable(my_homom_app main.cpp)

# Link the project target against aws-latticpp.
target_link_libraries(my_homom_app aws-latticpp)
```

## Using latticpp as a dependency without source
A more complex way to include latticpp in your CMake project is to have CMake download the latticpp source from Github. This avoids the need to include the latticpp source code directly in your project.

### 1. Under `third-party/aws-latticpp` directory, create a file `CMakeLists.txt`, which will be used to download latticpp GitHub code.
```cmake
cmake_minimum_required(VERSION 3.12)

project(AWS_LATTICPP_DOWNLOAD 0.1.2) # Change the version number "0.1.2" to whichever version you want

include(ExternalProject)
ExternalProject_Add(EP_AWS_LATTICPP
    SOURCE_DIR           ${CMAKE_CURRENT_LIST_DIR}/src
    BINARY_DIR           ${CMAKE_CURRENT_LIST_DIR}/build
    GIT_REPOSITORY       https://github.com/awslabs/aws-cppwrapper-lattigo.git
    GIT_TAG              v${PROJECT_VERSION}
    GIT_CONFIG           advice.detachedHead=false
    CONFIGURE_COMMAND    ""
    BUILD_COMMAND        ""
    INSTALL_COMMAND      ""
    TEST_COMMAND         ""
)
```

### 2. In `CMakeLists.txt` of your project, link `aws-latticpp` target.
```cmake
# Define a external project download method.
function(download_external_project project_dir)
    message(STATUS "Downloading ${project_dir}.")
    set(EXTERNAL_PROJECT_CMAKE_CACHE_FILE ${project_dir}/CMakeCache.txt)
    if(EXISTS ${EXTERNAL_PROJECT_CMAKE_CACHE_FILE})
        message(STATUS "Removing old ${EXTERNAL_PROJECT_CMAKE_CACHE_FILE}")
        file(REMOVE ${EXTERNAL_PROJECT_CMAKE_CACHE_FILE})
    endif()
    set(COMMAND_WORK_DIR ${project_dir})
    execute_process(
            COMMAND ${CMAKE_COMMAND} -G "${CMAKE_GENERATOR}" .
            RESULT_VARIABLE result
            WORKING_DIRECTORY ${COMMAND_WORK_DIR})
    if(result)
        message(FATAL_ERROR "Failed to download (${result}).")
    endif()
    execute_process(COMMAND ${CMAKE_COMMAND} --build .
            RESULT_VARIABLE result
            WORKING_DIRECTORY ${COMMAND_WORK_DIR})
    if(result)
        message(FATAL_ERROR "Failed to build (${result}).")
    endif()
endfunction()

# Download AWS HIT.
download_external_project(third-party/aws-latticpp)
# Add 'third-party/aws-latticpp/src', which defines 'aws-latticpp' target.
add_subdirectory(third-party/aws-latticpp/src)
# Define your project target.
add_executable(my_homom_app main.cpp)
# Link the project target against aws-latticpp as needed.
target_link_libraries(my_homom_app aws-latticpp)
```

### 3. HIT header files are now available in `main.cpp` of your project.

```c++
#include "latticpp/latticpp.h"

using namespace latticpp;

int main(int, char **argv) {
    Parameters params = getDefaultClassicalParams(PN14QP438);
    // Ready to use `params` to generate CKKS keys.
}
```
