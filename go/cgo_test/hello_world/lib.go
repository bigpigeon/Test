/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */
package main

// #include<stdio.h>
// #include<stdbool.h>
// void hello_world()
// {
//	    printf("hello world\n");
// }
import "C"

func printFalse() {
	C.hello_world()
}
