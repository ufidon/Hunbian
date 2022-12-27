#include <stdio.h>

int main(int argc, char const *argv[])
{
    FILE *档;
    void 复档(FILE *, FILE *);
    char *程名 = argv[0];

    if (argc == 1)
        复档(stdin, stdout);
    else
        while (--argc > 0)
            if ((档 = fopen(*++argv, "r")) == NULL)
            {
                frintf(stderr, "%s开%s失败.\n", 程名, *argv);
                exit(1); 
            }
            else
            {
                复档(档, stdout);
                fclose(档);
            }
    if (ferror(stdout))
    {
        fprintf(stderr, "%s写stdout出错.\n", 程名);
        exit(2);
    }
    // exit调fclose关闭打开的输出档, 并写缓冲至输出
    // 在main内 return 式 等价于 exit(式). 而exit(式)可在任何函式内被调用.
    exit(0);
}

void 复档(FILE *入, FILE *出)
{
    int 字;
    while ((字 = getc(入)) != EOF)
        putc(字, 出);
}