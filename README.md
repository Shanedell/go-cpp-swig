# SWIG Go and C++ Example

![Build Status](https://github.com/Shanedell/go-cpp-swig/workflows/Test%20Gocpp%20Module/badge.svg)

This repo is an example of bindings created via SWIG between C++ example

## Requirements

- **swig** to generate language bindings (http://www.swig.org/svn.html)
- **goloang** for running the shared library
- **c++ compiler** for creating the shared libraried (use `g++`)

## Generate SWIG bindings

```bash
swig -go -cgo -c++ -intgosize 64 -outdir gocpp gocpp/gocpp.i
```

Updates `gocpp/gocpp.go` and `gocpp/gocpp_wrap.cxx`.

This needs to be done every time the code for `gocpp.cpp` and `gocpp.h` changes. 

## Compiling go bindings -- Manual

### Generate C++ Binary

Works for both MacOS and Linux.

```bash
cd gocpp
g++ -std=c++11 -fPIC -c gocpp.cpp gocpp_wrap.cxx
cd ../
```

### Generate lib file

Works for both MacOS and Linux.

```bash
g++ -shared -o gocpp/_gocpp.so gocpp/gocpp.o gocpp/gocpp_wrap.o
```

### Test library

```bash
go test -v
```

## Compiling python bindings - Auto

Two Options:

- 1.) You can run the `gen_lib` file locally and everything will be built
- 2.) Run `docker-compose up -d --build`
  - Creates a container that builds all the bindings, creates the c++ binary and creates the library file
  - Creates another container once the build one finishes that runs the tests