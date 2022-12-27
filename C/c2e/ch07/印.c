#include <stdio.h>
#include <string.h>

int main(int argc, char const *argv[])
{
    int 宽 = 8;
    char *长 = "1234567890123456";
    char *短 = "123456";
    // %对齐下限.上限s
    printf("0|12345678901234567890|\n");
    printf("1|%.*s|\n", 宽, 长);
    printf("2|%s|\n", 长);
    printf("3|%15s|\n", 长);
    printf("4|%.15s|\n", 长);
    printf("5|%.20s|\n", 长);
    printf("6|%-20s|\n", 长);
    printf("7|%15.10s|\n", 长);
    printf("8|%-15.10s|\n", 长);

    printf("\n0|12345678901234567890|\n");
    printf("1|%.*s|\n", 宽, 短);
    printf("2|%s|\n", 短);
    printf("3|%15s|\n", 短);
    printf("4|%.15s|\n", 短);
    printf("5|%.20s|\n", 短);
    printf("6|%-20s|\n", 短);
    printf("7|%15.10s|\n", 短);
    printf("8|%-15.10s|\n", 短);    
    return 0;
}

/*练习写程序
2. 印任意输入. 以八进制或十六进制印控制符, 自动折长行.
*/