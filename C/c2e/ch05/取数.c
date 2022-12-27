#include <stdio.h>
#include <ctype.h>

int 取字(void);
void 放回(int);

int 寻整数(int *整)
{
    int 字, 正负;

    while (isspace(字 = 取字()))
        ; // 略文首空白

    if (!isdigit(字) && 字 != EOF && 字 != '+' && 字 != '-')
    {
        放回(字); // 字非数字
        return 0;
    }

    正负 = (字 == '-') ? -1 : 1;
    if (字 == '+' || 字 == '-')
        字 = 取字();

    for (*整 = 0; isdigit(字); 字 = 取字())
        *整 = 10 * *整 + (字 - '0');

    *整 *= 正负;

    if (字 != EOF)
        放回(字);

    return 字;
}

int main(int argc, char const *argv[])
{
    int 整数 = 0;

    寻整数(&整数);
    printf("寻得整数%d\n", 整数);

    return 0;
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
        printf("放回: 工段已满.\n");
    else
        工段[段位++] = 字;
}

/*练习
1. 上寻整数处理+或-带一非数字为0, 请更正之.
2. 写功能寻浮点数.
*/