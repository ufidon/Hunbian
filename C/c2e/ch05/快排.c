#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define 丄行数 1000
char *行集[丄行数];

int 读行(char *行集[], int 最大行数);
void 写行(char *行集[], int 行数);
void 快排(void *行集[], int 首行号, int 末行号, int (*比)(const void *, const void *));
int 比数(const char *, const char *);

int main(int argc, char const *argv[])
{
    int 行数;
    int 按数排 = 0;

    if (argc > 1 && strcmp(argv[1], "-s") == 0)
        按数排 = 1;

    if ((行数 = 读行(行集, 丄行数)) >= 0)
    {
        快排((void **)行集, 0, 行数 - 1,
             (int (*)(const void *, const void *))(按数排 ? 比数 : strcmp));
        写行(行集, 行数);
        return 0;
    }
    else
    {
        printf("出错: 输入太多行.\n");
        return 1;
    }
}

#define 丄行长 100
int 取行(char 行集[], int 最大行长);
char *配(int);

int 读行(char *行集[], int 最大行数)
{
    int 行长, 行数;
    char *新行, 暂行[丄行长];

    行数 = 0;
    while ((行长 = 取行(暂行, 丄行长)) > 0)
        if (行数 >= 丄行数 || (新行 = 配(行长)) == NULL)
            return -1;
        else
        {
            暂行[行长 - 1] = '\0';
            strcpy(新行, 暂行);
            行集[行数++] = 新行;
        }
    return 行数;
}

void 写行(char *行集[], int 行数)
{
#ifdef 写法1
    int 行号;
    for (行号 = 0; 行号 < 行数; 行号++)
        printf("%s\n", 行集[行号]);
#endif //  写法1
    while (行数--)
        printf("%s\n", *行集++);
}

void 换行(void *行集[], int 行号1, int 行号2)
{
    void *暂行;
    暂行 = 行集[行号1];
    行集[行号1] = 行集[行号2];
    行集[行号2] = 暂行;
}

int 比数(const char *串1, const char *串2)
{
    double 数1, 数2;
    数1 = atof(串1);
    数2 = atof(串2);
    if (数1 < 数2)
        return -1;
    else if (数1 > 数2)
        return 1;
    else
        return 0;
}

void 快排(void *行集[], int 首行号, int 末行号, int (*比)(const void *, const void *))
{
    int 行号, 桩号;

    if (首行号 >= 末行号)
        return;
    换行(行集, 首行号, (首行号 + 末行号) / 2);
    桩号 = 首行号;
    for (行号 = 首行号 + 1; 行号 <= 末行号; 行号++)
        if ((*比)(行集[行号], 行集[首行号]) < 0)
            换行(行集, ++桩号, 行号);
    换行(行集, 首行号, 桩号);
    快排(行集, 首行号, 桩号 - 1, 比);
    快排(行集, 桩号 + 1, 末行号, 比);
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

#define 总容量 100000

static char 库[总容量];
static char *块位 = 库;
static char *库首 = 库;
static char *库尾 = 库 + 总容量 - 1;

char *配(int 量)
{
    if (库尾 - 块位 + 1 >= 量)
    {
        块位 += 量;
        return 块位 - 量;
    }
    else
        return 0;
}

void 释(char *块首)
{
    if (块首 >= 库首 && 块首 <= 库尾)
        块位 = 块首;
}

/*练习
14. 快排默认升序, 传令参-n支持降序, 确保-n与-s可一起工作. 参考寻串2.c。
15. 传令参-w支持大小写无关的串排序。
16. 传令参-z支持字典排序，确保-z可与-w一起工作。
17. 表格按列多重排序，每列可用独立令参集。
    例： 几个家庭成员收入排序，先按姓排家庭，然后每个家庭内按收入排序。
*/
