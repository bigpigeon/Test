/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s\n", s);
}
char* toLower(char* s) {
    char * ret = malloc(128*sizeof(char));
    int i = 0;
	while(s[i] != '\0'){
		if(s[i] <='Z' && s[i] >= 'A'){
			ret[i] = (s[i] - 'A' + 'a');
		}else {
			ret[i] = s[i];
		}
		i++;
    }
	return ret;
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	s := "Hello from stdio"
	cs := C.CString(s)
	C.myprint(cs)
	cStr := C.toLower(cs)

	C.free(unsafe.Pointer(cs))
	goStr := C.GoString(cStr)
	fmt.Println(goStr)
	C.free(unsafe.Pointer(cStr))
}
