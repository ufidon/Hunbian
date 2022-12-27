#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
#include <fcntl.h>
#include <unistd.h>

int 读(int 档号, long 从, char *容, int 节数)
{
    if (lseek(档号, 从, 0) >= 0)
        return read(档号, 容, 节数);
    else
        return -1;
}

int main(int argc, char const *argv[])
{
    int 档号, 节数;
    long 从;
    char *容;

    if (argc != 4)
    {
        printf("用法: 读 档名 始位 节数\n");
        exit(1);
    }

    档号 = open(argv[1], O_RDONLY);
    从 = atol(argv[2]);
    节数 = atoi(argv[3]);
    容 = (char *)calloc(节数, 1);
    读(档号, 从, 容, 节数);
    printf("%s\n", 容);
    free(容);
    close(档号);

    return 0;
}
