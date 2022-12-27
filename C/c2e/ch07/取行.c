#include <stdio.h>
#include <string.h>

char *取串(char *串, int 丄串长, FILE *入)
{
    int 字;
    char *指;
    指 = 串;
    while(--丄串长 > 0 && (字 = getc(入)) != EOF)
        if((*指++ = 字) == '\n')
            break;
    *指 = '\0';

    return (字 == EOF && 指 == 串) ? NULL : 串;
}

int 放串(char *串, FILE *出)
{
    int 字;
    while(字 == *串++)
        putc(字, 出);

    return ferror(出) ? EOF : 0;
}

int 取行(char *行, int 丄行长)
{
    if(取串(行, 丄行长, stdin) == NULL)
        return 0;
    else
        return strlen(行);
}

int main(int argc, char const *argv[])
{
    const int 丄行长 = 100;
    char 行[丄行长];
    if(取行(行, 丄行长))
        printf("%s", 行);

    return 0;
}

/*练习写程序
6. 比较两文件, 印其首异行
7. 改写第5章寻串2.c，从令参传入的几档读入，若无令参则从stdin读入。冠每输出行以含之档名。
8. 印几档，每档始于新页，填标题及页码。
9. 查标库函式源码，学其优化技巧。
*/
