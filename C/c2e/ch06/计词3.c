#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include <stdlib.h>

struct 节
{
    char *词;
    int 数;
    struct 节 *左;
    struct 节 *右;
};

struct 节 *挂树(struct 节 *, char *);
void 印树(struct 节 *);

#define 丄词长 100
int 取词(char *, int);

int main(int argc, char const *argv[])
{
    struct 节 *根;
    char 词[丄词长];

    根 = NULL;
    while (取词(词, 丄词长) != EOF)
        if (isalpha(词[0]))
            根 = 挂树(根, 词);

    印树(根);
    return 0;
}

struct 节 *配节(void)
{
    return (struct 节 *)malloc(sizeof(struct 节));
}
char *复串(char *串)
{
    char *新;
    新 = (char *)malloc(strlen(串) + 1);
    if (新 != NULL)
        strcpy(新, 串);

    return 新;
}

struct 节 *挂树(struct 节 *指, char *词)
{
    int 比;
    if (指 == NULL)
    {
        指 = 配节();
        指->词 = 复串(词);
        指->数 = 1;
        指->左 = 指->右 = NULL;
    }
    else if ((比 = strcmp(词, 指->词)) == 0)
        指->数++;
    else if (比 < 0)
        指->左 = 挂树(指->左, 词);
    else
        指->右 = 挂树(指->右, 词);
    return 指;
}

void 印树(struct 节 *指)
{
    if (指 != NULL)
    {
        印树(指->左);
        printf("%4d %s\n", 指->数, 指->词);
        印树(指->右);
    }
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

/* 练习写程序
2. 读C程序, 将前几个字母相同而其余字母不同的变量名划为一组, 并按字母序印. 
    子串及注释不能算. 组首字母个数可设以令参.
3. 写交叉参考, 印文章所有词及每个词在文章出现的全部行数. 略虚词如冠词、介词、助动词、连词等等。
4. 印输入不同词，按词频降序。每行: 词频 词。
*/