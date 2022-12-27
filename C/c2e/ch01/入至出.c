#include <stdio.h>

int main(int argc, char const *argv[])
{
    int 字;

#ifdef 写法1
    字 = getchar();
    while (字 != EOF)
    {
        putchar(字);
        字 = getchar();
    }
#endif // 写法1

#ifdef 写法2
    printf("EOF=%d\n", EOF);

    while ((字 = getchar()) != EOF)
    {
        putchar(字);
    }
#endif

#ifdef 字节计数1
    // 字节计数?
    long 字节数 = 0;
    while ((字 = getchar()) != EOF)
    {
        ++字节数;
        putchar(字);
    }
    printf("字节数=%ld\n", 字节数);
#endif // 字节计数1

#ifdef 字节计数2
    double 字节数;
    for (字节数 = 0; getchar() != EOF; ++字节数)
        ;
    printf("字节数=%.0f\n", 字节数);
#endif // 字节计数2

    // 计行
    int 行数;
    行数 = 0;
    while ((字 = getchar()) != EOF)
    {
        if (字 == '\n')
        {
            行数++;
        }
    }
    printf("共计%d行\n", 行数);

    return 0;
}

/* 练习：
6. 检测 getchar() != EOF 为0或1。
7. 写一程序，将字符以数字输出。
*/
