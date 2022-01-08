# SWIG Go and C++ Example

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

## Compiling go bindings - Auto

Using the `gen_lib` script you can easily generate the swig bindings, generate the c++ binary file and the library file.
