#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "lib.h"

int
main(int argc, char **argv) {
    GoString str = {"Hello from C!", 13};
    char *err;
    char *result;
    GoSlice sli = {"slice data", 11, 11};
    StringCmp(sli,sli,sli);

	long long arr[3] = {2,3,4};
    GoSlice intSli = {arr, 3, 3};
    long long *intResult;
    intResult = IntSlicePrint(intSli);
    for(int i = 0;i < 4;i++) {
        printf("int result %lld\n", *(intResult+i));
    }

    result = StrConv(str);
    printf("cp %p\n", result);
    result = StrConv(str);
    printf("cp %p\n", result);
    result = StrConv(str);
	printf("cp %p\n", result);

    return 0;
}
