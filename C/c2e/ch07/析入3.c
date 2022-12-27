#include <stdio.h>
#include <ctype.h>
#include <string.h>

int 取行(char 行[], int 最大行长);

int main(int argc, char const *argv[])
{
    int 年, 月, 日;
    char 月名[10], 行[100];

    while (取行(行, sizeof(行)) > 0)
    {
        if (sscanf(行, "%d %s %d", &日, 月名, &年) == 3)
            printf("%d年%s月%d日\n", 年, 月名, 日);
        else if (sscanf(行, "%d/%d/%d", &日, &月, &年) == 3)
            printf("%d年%d月%d日\n", 年, 月, 日);
        else if (sscanf(行, "%d-%d-%d", &日, &月, &年) == 3)
            printf("%d年%d月%d日\n", 年, 月, 日);
        else
            printf("未知日期格式:%s\n", 行);
    }

    return 0;
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

/*练习写程序
4. 参考变参.c, 写简版scanf。
5. 用scanf及sscanf重写第四章算器.c之输入与转换。
*/