#!/bin/bash

CGO_CXXFLAGS="-std=c++11"

go build -ldflags "-lc" -x
