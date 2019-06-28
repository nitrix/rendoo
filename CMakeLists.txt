CMAKE_MINIMUM_REQUIRED(VERSION 3.0)
PROJECT(rendoo)

INCLUDE_DIRECTORIES(
    include
)

ADD_DEFINITIONS(
    -std=c11
    -pedantic
    -Wall
    -Werror
)

ADD_EXECUTABLE(rendoo src/main.c)