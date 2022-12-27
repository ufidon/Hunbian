/*
C语宣语法：
宣:     可选 * 直宣
直宣    : 名
        ｜(宣)
        ｜直宣()
        ｜直宣[可选夵]
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>

#define 丄词数 100
enum
{
    C名,
    C括,
    C方
};

int 取字();
void 放回(int 字);
void 析宣(void);
void 析直宣(void);

int 取词(void);
int 词类;
char 词[丄词数];
char 名[丄词数];
char 数类[丄词数];
char 出[1000];

int main(int argc, char const *argv[])
{
    while (取词() != EOF)
    {
        strcpy(数类, 词);
        出[0] = '\0';
        析宣();
        if (词类 != '\n')
            printf("语法错误\n");
        printf("%s: %s %s\n", 名, 出, 数类);
    }

    return 0;
}

int 取词(void)
{
    int 字;
    char *指 = 词;

    while ((字 = 取字()) == ' ' || 字 == '\t')
        ;

    if (字 == '(')
    {
        if ((字 = 取字()) == ')')
        {
            strcpy(词, "()");
            return 词类 = C括;
        }
        else
        {
            放回(字);
            return 词类 = '(';
        }
    }
    else if (字 == '[')
    {
        for (*指++ = 字; (*指++ = 取字()) != ']';)
            ;
        *指 = '\0';
        return 词类 = C方;
    }
    else if (isalpha(字))
    {
        for (*指++ = 字; isalnum(字 = 取字());)
            *指++ = 字;

        *指 = '\0';
        放回(字);
        return 词类 = C名;
    }
    else
        return 词类 = 字;
}

void 析宣(void) // 析宣为文
{
    int 星数;

    for (星数 = 0; 取词() == '*';)
        星数++;
    析直宣();
    while (星数-- > 0)
    {
        strcat(出, "指至");
    }
}

void 析直宣(void)
{
    int 类;

    if (词类 == '(')
    {
        析宣();
        if (词类 != ')')
            printf("错:缺)\n");
    }
    else if (词类 == C名)
        strcpy(名, 词);
    else
        printf("错:需名或(宣)\n");

    while ((类 = 取词()) == C括 || 类 == C方)
        if (类 == C括)
            strcat(出, "函返");
        else
        {
            strcat(出, "列");
            strcat(出, 词);
            strcat(出, "属");
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

/*练习
18. 改进析宣, 使其能从错中恢复
19. 改进反宣, 使其勿加冗余括号至宣
20. 改进析宣, 以析带参函式、束词如const等。
*/