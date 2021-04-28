#include <stdlib.h>
#include <stdint.h>
#include <stdio.h>
int cpuHogCSalt1 = 0;
int cpuHogCSalt2 = 0;
void add100WTimesSub(int foo) {
    int i;
    for (i = 0; i < 100000; i++) {
        if (foo > 0) {
            foo *= foo;
        } else {
            foo *= foo + 1;
        }
        cpuHogCSalt2 = foo;
    }
}

void add100WTimes() {
    add100WTimesSub(cpuHogCSalt1);
}

void alloc10MHeap() {
    char * data = malloc(sizeof(char) * 30<<20);
    int count = 0;
    for(int i=0; i< (30<<20)-1 ;i++) {
        *(data+i) = i%64;
        count++;
    }
    printf("data alloc %d\n", count);
}
