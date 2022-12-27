#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
#include <fcntl.h>
#include <unistd.h>

#define 权限 0666
#define 储夵 1000

void 出错(char *, ...);

int main(int argc, char const *argv[])
{
    int 档1, 档2, 数;
    char 储[储夵];

    if (argc != 3)
        出错("用法: 复 从 到");
    if ((档1 = open(argv[1], O_RDONLY)) == -1)
        出错("复: 不能开%s", argv[1]);
    if ((档2 = creat(argv[2], 权限)) == -1)
        出错("复:不能造%s以%03o权限", argv[2], 权限);

    while ((数 = read(档1, 储, 储夵)) > 0)
        if (write(档2, 储, 数) != 数)
            出错("复:写档%s出错", argv[2]);

    return 0;
}

void 出错(char *格式, ...)
{
    va_list 参;
    va_start(参, 格式);
    fprintf(stderr, "错:");
    vfprintf(stderr, 格式, 参);
    fprintf(stderr, "\n");
    va_end(参);

    exit(1);
}

/*练习写程序
1. 用系统函式read, write, open及close重写第7章并印2.c，并比较二者性能。
*/