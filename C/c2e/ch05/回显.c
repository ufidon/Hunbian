#include <stdio.h>

#define 号法
#define 指法

int main(int 矢长, char const *参矢[])
{
    int  参数 = 矢长;
#ifdef 号法
    int 号;

    for (号 = 1; 号 < 参数; 号++)
        printf("%s%s", 参矢[号], (号 < 参数 - 1) ? " " : "\n");
#endif // 号法

#ifdef 指法
    while (--参数 > 0)
    {
        printf("%s%s", *++参矢, (参数 == 1) ? "\n" : " ");
    }
#endif

    // 简法
    参矢  -= 矢长-1;
    while (*++参矢)
        printf("%s%s", *参矢, " ");
    printf("\n");
    
    return 0;
}
