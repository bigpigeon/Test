package main

// #include<stdio.h>
// #include<stdbool.h>
// bool boolFunc()
// {
//	    return false;
// }
import "C"
import "fmt"

func main() {

	fmt.Println(bool(C.boolFunc()))
	// Output: 42
}
