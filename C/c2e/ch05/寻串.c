#include <stdio.h>
#include <string.h>
#define 丄行长 1000

int 取行(char 行[], int 最大行长);

int main(int argc, char const *argv[])
{
    char 行[丄行长];
    int 寻得 = 0;

    if (argc != 2)
        printf("用法: 寻串 串\n");
    else
    {
        while (取行(行, 丄行长) > 0)
        {
            if (strstr(行, argv[1]) != NULL)
            {
                printf("%s", 行);
                寻得++;
            }
        }
    }

    return 寻得;
}

int 取行(char 行[], int 最大行长)
{
    int 字, 号;
    for (号 = 0; 号 < 最大行长 && (字 = getchar()) != EOF && 字 != '\n'; 号++)
    {
        行[号] = 字;
    }
    if (字 == '\n')
    {
        行[号] = 字;
        ++号;
    }

    行[号] = '\0';
    return 号;
}