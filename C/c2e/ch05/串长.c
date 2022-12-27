#include <stdio.h>

int 串长(char *串)
{
    int 长;
    for (长 = 0; *串 != '\0'; 串++)
        长++;

    return 长;
}

int 串长2(char *串)
{
    char *指 = 串;
    while (*指 != '\0')
    {
        指++;
    }
    return 指 - 串;
}

int main(int argc, char const *argv[])
{
    char 串[100] = "这是一个地球村.";
    char *名 = "张三";

    printf("串\"%s\"长%d个字节.\n", "学C语言有意思", 串长("学C语言有意思"));
    printf("串\"%s\"长%d个字节.\n", 串, 串长2(串));
    printf("串\"%s\"长%d个字节.\n", 名, 串长(名));

    // 子串
    printf("\"%s\"长%d个字节。\n", 串 + 12, 串长(串 + 12));

    // 负号
    char *尾 = 串 + 串长2(串);
    printf("\"%s\"长%d个字节。\n", &尾[-10], 串长(&尾[-10]));

    return 0;
}
