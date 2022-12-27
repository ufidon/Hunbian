#include <stdio.h>
#include <string.h>
#define 丄行长 1000

int 取行(char 行[], int 最大行长);

int main(int argc, char const *argv[])
{
    char 行[丄行长];
    long 行号 = 0;
    int 字, 反 = 0, 填行号 = 0, 寻得 = 0;

    while (--argc > 0 && (*++argv)[0] == '-')
        while ((字 = *++argv[0]))
        {
            switch (字)
            {
            case 'f':
                反 = 1;
                break;
            case 'h':
                填行号 = 1;
                break;
            default:
                printf("寻串2: 未知配置%c\n", 字);
                argc = 0;
                寻得 = -1;
                break;
            }
        }

    if (argc != 1)
        printf("用法: 寻串 -f -h 串\n");
    else
    {
        while (取行(行, 丄行长) > 0)
        {
            行号++;
            if ((strstr(行, *argv) != NULL) != 反)
            {
                if (填行号)
                    printf("%ld:", 行号);
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

/*练习
10. 写程序算式, 求令行算符后置式, 各数符单独为参. 
    例: 算式 3 1 4 - / 求 3/(1-4)
11. 以令参设每跳格空格数, 重写跳格空格互换程序. 无此令参则取一默认值.
12. 用字画菱形: 画菱形 宽 高 字. 
    例: 画菱形 3 3 * 得
       *   
      *  *
    *      *
      *  *    
        * 
13. 写程序印尾行: 印尾行 行数. 印输入行集末尾指定行数.
*/