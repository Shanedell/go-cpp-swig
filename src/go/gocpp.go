/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.1.0
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: src/bindings/gocpp.i

package gocpp

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


extern void _wrap_Swig_free_gocpp_87c7e8ed7687eccb(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_gocpp_87c7e8ed7687eccb(swig_intgo arg1);
extern swig_intgo _wrap_Math_Sum_gocpp_87c7e8ed7687eccb(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3);
extern swig_intgo _wrap_Math_diff_gocpp_87c7e8ed7687eccb(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3);
extern swig_intgo _wrap_Math_product_gocpp_87c7e8ed7687eccb(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3);
extern swig_intgo _wrap_Math_quotient_gocpp_87c7e8ed7687eccb(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3);
extern swig_intgo _wrap_Math_modulus_gocpp_87c7e8ed7687eccb(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3);
extern uintptr_t _wrap_new_Math_gocpp_87c7e8ed7687eccb(void);
extern void _wrap_delete_Math_gocpp_87c7e8ed7687eccb(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_gocpp_87c7e8ed7687eccb(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_gocpp_87c7e8ed7687eccb(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrMath uintptr

func (p SwigcptrMath) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMath) SwigIsMath() {
}

func (arg1 SwigcptrMath) Sum(arg2 int, arg3 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (int)(C._wrap_Math_Sum_gocpp_87c7e8ed7687eccb(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2)))
	return swig_r
}

func (arg1 SwigcptrMath) Diff(arg2 int, arg3 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (int)(C._wrap_Math_diff_gocpp_87c7e8ed7687eccb(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2)))
	return swig_r
}

func (arg1 SwigcptrMath) Product(arg2 int, arg3 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (int)(C._wrap_Math_product_gocpp_87c7e8ed7687eccb(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2)))
	return swig_r
}

func (arg1 SwigcptrMath) Quotient(arg2 int, arg3 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (int)(C._wrap_Math_quotient_gocpp_87c7e8ed7687eccb(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2)))
	return swig_r
}

func (arg1 SwigcptrMath) Modulus(arg2 int, arg3 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (int)(C._wrap_Math_modulus_gocpp_87c7e8ed7687eccb(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2)))
	return swig_r
}

func NewMath() (_swig_ret Math) {
	var swig_r Math
	swig_r = (Math)(SwigcptrMath(C._wrap_new_Math_gocpp_87c7e8ed7687eccb()))
	return swig_r
}

func DeleteMath(arg1 Math) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Math_gocpp_87c7e8ed7687eccb(C.uintptr_t(_swig_i_0))
}

type Math interface {
	Swigcptr() uintptr
	SwigIsMath()
	Sum(arg2 int, arg3 int) (_swig_ret int)
	Diff(arg2 int, arg3 int) (_swig_ret int)
	Product(arg2 int, arg3 int) (_swig_ret int)
	Quotient(arg2 int, arg3 int) (_swig_ret int)
	Modulus(arg2 int, arg3 int) (_swig_ret int)
}

