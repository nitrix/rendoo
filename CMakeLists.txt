CMAKE_MINIMUM_REQUIRED(VERSION 3.0)
PROJECT(rendoo)

ADD_SUBDIRECTORY(lib/toolkit)

ADD_DEFINITIONS(
    -std=c11
    -pedantic
    -Wall
    -Werror
)

ADD_EXECUTABLE(
    rendoo
    src/main.c
    src/rasterizer.c
    src/scene.c
)

TARGET_LINK_LIBRARIES(rendoo toolkit)

TARGET_INCLUDE_DIRECTORIES(rendoo PRIVATE include)