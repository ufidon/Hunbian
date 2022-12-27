#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/dir.h>
#include <dirent.h>
#include <errno.h>

void 档夵(const char *);
void 遍录(const char *, void (*函式)(const char *));

int main(int argc, char const *argv[])
{
    if (argc == 1)
        档夵(".");
    else
        while (--argc > 0)
            档夵(*++argv);

    return 0;
}

void 档夵(const char *档名)
{
    struct stat 态;

    if (stat(档名, &态) == -1)
    {
        fprintf(stderr, "档夵: 查不了%s\n", 档名);
        return;
    }

    if (S_ISDIR(态.st_mode))
        遍录(档名, 档夵);

    printf("%8ld %s\n", 态.st_size, 档名);
}

void 遍录(const char *录名, void (*函式)(const char *))
{
    char 全名[256];
    struct dirent *指;
    DIR *d;

    if ((d = opendir(录名)) == NULL)
    {
        fprintf(stderr, "遍录: 开不了%s\n", 录名);
        return;
    }

    // 参考: https://man7.org/linux/man-pages/man3/readdir.3.html
    while ((指 = readdir(d)) != NULL)//每次读返回下一录项
    {
        if (strcmp(指->d_name, ".") == 0 || strcmp(指->d_name, "..") == 0)
            continue;
        if (strlen(录名) + strlen(指->d_name) + 2 > sizeof(全名))
            fprintf(stderr, "遍录: 路径%s/%s太长.\n", 录名, 指->d_name);
        else
        {
            sprintf(全名, "%s/%s", 录名, 指->d_name);
            (*函式)(全名);
        }
    }

    closedir(d);
}

/*练习
5. 改进档夵加印节点内其它信息.
*/