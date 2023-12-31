/* Code generated by cmd/cgo; DO NOT EDIT. */

/* package command-line-arguments */


#line 1 "cgo-builtin-export-prolog"

#include <stddef.h>

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#endif

#endif

/* Start of preamble from import "C" comments.  */


#line 7 "main.go"
 #include "../cloud.h"

#line 1 "cgo-generated-wrapper"


/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef size_t GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
#ifdef _MSC_VER
#include <complex.h>
typedef _Fcomplex GoComplex64;
typedef _Dcomplex GoComplex128;
#else
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;
#endif

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef _GoString_ GoString;
#endif
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif

extern __declspec(dllexport) component* Dial(char* conf);
extern __declspec(dllexport) char* Ping(component* comp);
extern __declspec(dllexport) char* Close(component* comp);
extern __declspec(dllexport) char* Set(component* comp, char* key, char* value, double expire);
extern __declspec(dllexport) char* Get(component* comp, char* key);
extern __declspec(dllexport) char* Del(component* comp, char* c);
extern __declspec(dllexport) char* SetNX(component* comp, char* key, char* value, double expire);
extern __declspec(dllexport) char* Do(component* comp, char* cmd);
extern __declspec(dllexport) char* Eval(component* comp, char* c);
extern __declspec(dllexport) char* HSet(component* comp, char* c);
extern __declspec(dllexport) char* HGet(component* comp, char* key, char* value);
extern __declspec(dllexport) char* HDel(component* comp, char* c);
extern __declspec(dllexport) char* LPush(component* comp, char* c);
extern __declspec(dllexport) char* RPush(component* comp, char* c);
extern __declspec(dllexport) char* LRange(component* comp, char* c);
extern __declspec(dllexport) char* Incr(component* comp, char* key);
extern __declspec(dllexport) char* Decr(component* comp, char* key);
extern __declspec(dllexport) char* LPop(component* comp, char* key);
extern __declspec(dllexport) char* RPop(component* comp, char* key);
extern __declspec(dllexport) char* Scan(component* comp, char* prefix, long long int count);
extern __declspec(dllexport) char* SScan(component* comp, char* key, char* prefix, long long int count);

#ifdef __cplusplus
}
#endif
