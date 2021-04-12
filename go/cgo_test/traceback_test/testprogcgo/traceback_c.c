// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The C definitions for traceback.go. That file uses //export so
// it can't put function definitions in the "C" import comment.

#include <stdint.h>
#include <stdio.h>
#include <inttypes.h>
char *p;

int crashInGo;
extern void h1(void);

int tracebackF3(void) {
	if (crashInGo)
		h1();
	else
		*p = 0;
	return 0;
}

int tracebackF2(void) {
	return tracebackF3();
}

int tracebackF1(void) {
	return tracebackF2();
}

struct cgoTracebackArg {
	uintptr_t  context;
	uintptr_t  sigContext;
	uintptr_t* buf;
	uintptr_t  max;
};

struct cgoSymbolizerArg {
	uintptr_t   pc;
	const char* file;
	uintptr_t   lineno;
	const char* func;
	uintptr_t   entry;
	uintptr_t   more;
	uintptr_t   data;
};

void cgoTraceback(void* parg) {
	struct cgoTracebackArg* arg = (struct cgoTracebackArg*)(parg);
	printf("buf 0x%" PRIXPTR " ctx 0x%" PRIXPTR " sig ctx 0x%" PRIXPTR " max 0x%" PRIXPTR "\n", arg->buf[0], arg->context, arg->sigContext, arg->max);
	arg->buf[0] = (uintptr_t)(tracebackF1);
	arg->buf[1] = (uintptr_t)(tracebackF2);
	arg->buf[2] = (uintptr_t)(tracebackF3);
	arg->buf[3] = 0;
}

void cgoSymbolizer(void* parg) {
	struct cgoSymbolizerArg* arg = (struct cgoSymbolizerArg*)(parg);
	printf("pc 0x%" PRIXPTR "\n", arg->pc);
	if (arg->pc != arg->data + 1) {
		arg->file = "unexpected data";
	} else {
		arg->file = "cgo symbolizer";
	}
	arg->lineno = arg->data + 1;
	arg->data++;
}
