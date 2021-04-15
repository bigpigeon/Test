#include <stdlib.h>
#include <stdint.h>
#include <stdio.h>
void printsth() {
    printf("cgo func running\n");
}
void panic() {
    printsth();
    int *a;
    *a=2;
}
