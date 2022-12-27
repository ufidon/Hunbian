#include <stdio.h>
#define 行长上限 1000

int 取行(char 行[], int 行数上限);
int 寻词(char 文[], char 词[]);

char 词[] = "春风";

// 求输入之最长行
int main(int argc, char const *argv[])
{
    char 文[行长上限];
    int 寻得 = 0;

    while (取行(文, 行长上限) > 0)
    {
        if (寻词(文, 词) >= 0)
        {
            printf("%s", 文);
            寻得++;
        }
    }

    return 寻得;
}

int 取行(char 行[], int 行数上限)
{
    int 字, 号;
    号 = 0;
    while (行数上限-- && (字 = getchar()) != EOF && 字 != '\n')
    {
        行[号++] = 字;
    }
    if (字 == '\n')
    {
        行[号++] = 字;
    }

    行[号] = '\0';
    return 号;
}

int 寻词(char 文[], char 词[])
{
    int 号, 文步, 词步;
    for (号 = 0; 文[号] != '\0'; 号++)
    {
        for (文步 = 号, 词步 = 0; 词[词步] != '\0' && 文[文步] == 词[词步]; 文步++, 词步++)
            ;
        if (词步 > 0 && 词[词步] == '\0')
            return 号;
    }
    return -1;
}

/*练习
1. 上寻词在文中从左往右寻, 只返首寻者. 改写函数 int 寻词(char 文[], char 词[]) 返最后出现者号, 无则返-1.
*/
