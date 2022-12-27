#include <stdio.h>
#include <string.h>

#define 丄行数 1000
char *行集[丄行数];

int 读行(char *行集[], int 最大行数);
void 写行(char *行集[], int 行数);
void 快排(char *行集[], int 首行号, int 末行号);

int main(int argc, char const *argv[])
{
    int 行数;

    if ((行数 = 读行(行集, 丄行数)) >= 0)
    {
        快排(行集, 0, 行数 - 1);
        写行(行集, 行数);
        return 0;
    }
    else
    {
        printf("出错: 输入太多行.\n");
        return 1;
    }
}

#define 丄行长  100
int 取行(char 行集[], int 最大行长);
char *配(int);

int 读行(char *行集[], int 最大行数)
{
    int 行长, 行数;
    char *新行, 暂行[丄行长];

    行数 = 0;
    while((行长 = 取行(暂行, 丄行长)) > 0)
        if(行数 >= 丄行数 || (新行 = 配(行长)) == NULL)
            return -1;
        else
        {
            暂行[行长-1] = '\0';
            strcpy(新行, 暂行);
            行集[行数++] = 新行;
        }
    return 行数;    
}

void 写行(char *行集[], int 行数)
{
#ifdef 写法1
    int 行号;
    for(行号 = 0; 行号 < 行数; 行号++)
        printf("%s\n", 行集[行号]);
#endif //  写法1
    while(行数--)
        printf("%s\n", *行集++);
}

void 换行(char *行集[], int 行号1, int 行号2)
{
    char *暂行;
    暂行 = 行集[行号1];
    行集[行号1] = 行集[行号2];
    行集[行号2] = 暂行;
}

void 快排(char *行集[], int 首行号, int 末行号)
{
    int 行号, 桩号;

    if(首行号 >= 末行号)
        return;
    换行(行集, 首行号, (首行号+末行号)/2);
    桩号 = 首行号;
    for(行号 = 首行号+1; 行号 <= 末行号; 行号++)
        if(strcmp(行集[行号], 行集[首行号]) < 0)
            换行(行集, ++桩号, 行号);
    换行(行集, 首行号, 桩号);
    快排(行集, 首行号, 桩号-1);
    快排(行集, 桩号+1, 末行号);
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
7. 重写读行，存行于main传入之行集。
*/

