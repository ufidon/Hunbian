#include <stdio.h>
#include <ctype.h>
#include <string.h>

int main(int argc, char const *argv[])
{
    double 和, 项;
    int 年, 日;
    char 月[10];


    和 = 0;
    while (scanf("%lf", &项) == 1)
        printf("\t%.2f\n", 和 += 项);


    return 0;
}
