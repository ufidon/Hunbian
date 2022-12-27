#include <stdio.h>
#include <ctype.h>
#include <string.h>

#ifdef 写法1
#define 表键数 20
struct 键
{
    char *词;
    int 数;
} 键表[表键数];
#else
struct 键
{
    char *词;
    int 数;
} 键表[] = {
    {"auto", 0},
    {"break", 0},
    {"case", 0},
    {"char", 0},
    {"const", 0},
    {"continue", 0},
    {"default", 0},
    {"unsigned", 0},
    {"void", 0},
    {"volatile", 0},
    /* ... */
    {"while", 0}};
#define 表键数 (sizeof 键表 / sizeof 键表[0])

#endif // 写法1

#define 丄词长 100
int 取词(char *, int);
struct 键 *折半寻(char *, struct 键 *, int);

int main(int argc, char const *argv[])
{
    char 词[丄词长];
    struct 键 *指;

    while (取词(词, 丄词长) != EOF)
        if (isalpha(词[0]))
            if ((指 = 折半寻(词, 键表, 表键数)) != NULL)
                指->数++;

    for (指 = 键表; 指 < 键表 + 表键数; 指++)
        if (指->数 > 0)
            printf("%4d %s\n", 指->数, 指->词);
    return 0;
}

struct 键 *折半寻(char *词, struct 键 *表, int 数)
{
    // 左闭右开
    int 比;
    struct 键 *左 = &表[0], *右 = &表[数], *中;

    while (左 < 右)
    {
        中 = 左 + (右 - 左) / 2;
        if ((比 = strcmp(词, 中->词)) < 0)
            右 = 中;
        else if (比 > 0)
            左 = 中 + 1;
        else
            return 中;
    }
    return NULL;
}

#define 段长 100
char 工段[段长];
int 段位 = 0;

int 取字(void)
{
    return (段位 > 0) ? 工段[--段位] : getchar();
}

void 放回(int 字)
{
    if (段位 >= 段长)
        printf("放回:工段已满.\n");
    else
        工段[段位++] = 字;
}

// 从输入取翌词
int 取词(char *词, int 词长限)
{
    int 字;
    char *备词 = 词;

    while (isspace(字 = 取字()))
        ;

    if (字 != EOF)
        *备词++ = 字;
    if (!isalpha(字))
    {
        *备词 = '\0';
        return 字;
    }
    for (; --词长限 > 0; 备词++)
        if (!isalnum(*备词 = 取字()))
        {
            放回(*备词);
            break;
        }
    *备词 = '\0';
    return 词[0];
}

/* 练习
1. 改进取词, 支持下划线、有名恒量、注释和预处理键词。
*/