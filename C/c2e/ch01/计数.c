#include <stdio.h>

int main(int argc, char const *argv[])
{
    // 计数字、空白及其余
    int 字, 号, 空白数, 其余字数;
    int 数字数目[10];

    空白数 = 其余字数 = 0;
    for (号 = 0; 号 < 10; 号++)
    {
        数字数目[号] = 0;
    }

    while ((字 = getchar()) != EOF)
    {
        if (字 >= '0' && 字 <= '9')
        {
            ++数字数目[字 - '0'];
        }
        else if (字 == ' ' || 字 == '\t' || 字 == '\n')
        {
            ++空白数;
        }
        else
            ++其余字数;
    }

    printf("数字: ");
    for (号 = 0; 号 <= 9; 号++)
    {
        printf("%2d ", 号);
    }
    printf("\n");

    printf("数目: ");
    for (号 = 0; 号 <= 9; 号++)
    {
        printf("%2d ", 数字数目[号]);
    }
    printf("\n");

    printf("空白数=%d, 其余字数=%d\n", 空白数, 其余字数);

    return 0;
}

/* 练习：
11. 怎样测试上程序？并揭露其缺陷？
12. 写程序，将每字或每词单独输出为一行。
13. 写程序，以输入词之长度画横向直方图及纵向直方图。
14. 写程序，画词频直方图及字频直方图。
*/