#include <stdio.h>

int main(int argc, char const *argv[])
{
    FILE *档;
    void 复档(FILE *, FILE *);

    if (argc == 1)
        复档(stdin, stdout);
    else
        while (--argc > 0)
            if ((档 = fopen(*++argv, "r")) == NULL)
            {
                printf("开%s失败.\n", *argv);
                return 1;
            }
            else
            {
                复档(档, stdout);
                fclose(档);
            }

    return 0;
}

void 复档(FILE *入, FILE *出)
{
    int 字;
    while ((字 = getc(入)) != EOF)
        putc(字, 出);
}