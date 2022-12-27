#include <stdio.h>
#include <string.h>
#include <stdlib.h>

struct 散列
{
    struct 散列 *翌;
    char *名;
    char *换;
};

#define 散列表夵 101
static struct 散列 *散列表[散列表夵];

unsigned 打散(char *串)
{
    unsigned 散;
    for (散 = 0; *串 != '\0'; 串++)
        散 = *串 + 31 * 散;
    return 散 % 散列表夵;
}

struct 散列 *查(char *串)
{
    struct 散列 *指;

    for (指 = 散列表[打散(串)]; 指 != NULL; 指 = 指->翌)
        if (strcmp(串, 指->名) == 0)
            return 指;

    return NULL;
}

char *复串(char *串)
{
    char *新;
    新 = (char *)malloc(strlen(串) + 1);
    if (新 != NULL)
        strcpy(新, 串);

    return 新;
}

struct 散列 *插表(char *名, char *换)
{
    struct 散列 *指;
    unsigned 散;

    if ((指 = 查(名)) == NULL)
    {
        指 = (struct 散列 *)malloc(sizeof(*指)); // 指空没有问题?
        if (指 == NULL || (指->名 = 复串(名)) == NULL)
            return NULL;
        散 = 打散(名);
        指->翌 = 散列表[散];
        散列表[散] = 指;
    }
    else
        free((void *)指->换);

    if ((指->换 = 复串(换)) == NULL)
        return NULL;

    return 指;
}

struct
{
    char *名;
    unsigned int 旗;
    int 类;
    union
    {
        int 整;
        float 浮;
        char *串;
    } 值;

} 表[101];

struct
{
    unsigned int 读 : 1;
    unsigned int 写 : 1;
    unsigned int 跑 : 1;
    unsigned int : 29; // 无名域
} 旗;

int main(int argc, char const *argv[])
{
    struct 散列 *指;
    if((指 = 查("π")) != NULL)
        printf("1. π=%s\n", 指->换);
    else
        插表("π", "3.14159");

    if((指 = 查("π")) != NULL)
        printf("2. π=%s\n", 指->换);

    旗.写 = 0;
    旗.读 = 1;
    旗.跑 = 1;
    printf("%s、%s、%s\n", 旗.写?"可写":"不可写", 旗.读?"可读":"不可读", 旗.跑?"可跑":"不可跑");

    表[0].名 = "收入";
    表[0].旗 = *(unsigned int*)&旗;
    表[0].类 = 3;
    表[0].值.整 = 123;
    return 0;
}
