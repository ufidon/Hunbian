#include "历录.h"

void 档夵(const char *);
void 历录(const char *, void (*函)(const char *));

int main(int argc, char const *argv[])
{
    if (argc == 1)
        档夵(".");
    else
        while (--argc > 0)
            档夵(*++argv);

    return 0;
}

void 档夵(const char *名)
{
    struct stat 态;

    if (stat(名, &态) == -1)
    {
        fprintf(stderr, "档夵: 读不了%s\n", 名);
        return;
    }

    if (S_ISDIR(态.st_mode) /* (态.st_mode & S_IFMT) == S_IFDIR */)
    {
        印调("档夵调历录查目录\"%s\"\n", 名);
        历录(名, 档夵);
    }
    printf("%8ld %s\n", 态.st_size, 名);
}

#define 径名长限 1024

void 历录(const char *录名, void (*函)(const char *))
{
    char 径名[径名长限];
    录项 *项指;
    录 *录指;

    if ((录指 = 开录(录名)) == NULL)
    {
        fprintf(stderr, "历录: 打不开%s\n", 录名);
        return;
    }
    while ((项指 = 读录(录指)) != NULL)
    {
        印调("历录: 项指->名=%s\n", 项指->名);
        if (strcmp(项指->名, ".") == 0 || strcmp(项指->名, "..") == 0)
            continue;
        if (strlen(录名) + strlen(项指->名) + 2 > sizeof(径名))
            fprintf(stderr, "历录: 路径%s/%s太长\n", 录名, 项指->名);
        else
        {
            sprintf(径名, "%s/%s", 录名, 项指->名);
            (*函)(径名);
        }
    }

    关录(录指);
}

录 *开录(const char *录名)
{
    int 档号;
    struct stat 态;
    录 *录指;

    if ((档号 = open(录名, O_RDONLY, 0)) == -1 
    || fstat(档号, &态) == -1 
    || !S_ISDIR(态.st_mode) /* (态.st_mode & S_IFMT) != S_IFDIR */
    || (录指 = (录 *)malloc(sizeof(录))) == NULL)
        return NULL;

    录指->号 = 档号;
    印调("开录: 录名=%s, 档号=%d\n", 录名, 录指->号);

    return 录指;
}

录项 *读录(录 *某录)
{
    struct direct *目指;
    static DIR *录指;
    static 录项 项;

#if 0
    // https://stackoverflow.com/questions/42154232/linux-c-how-to-open-a-directory-and-get-a-file-descriptor
    int 节数; 
    节数 = read(某录->号, (char *)&目, sizeof(目));
    印调("读录: 档号=%d 节数 = %d 错码: %d 错:%s\n", 某录->号, 节数, errno, strerror(errno));
#endif // _调试

    录指 = fdopendir(某录->号);
    while ((目指 = readdir(录指))
        /*read(某录->号, (char *)&目, sizeof(目)) == sizeof(目)*/)
    {
        if (目指->d_ino == 0)
            continue;
        项.节号 = 目指->d_ino;
        strncpy(项.名, 目指->d_name, 档名长限);
        印调("读录: %s\n", 项.名);
        项.名[档名长限] = '\0';
        return &项;
    }
    印调("读录: 失败: %s %d 错码:%d 错:%s\n", 某录->项.名, 某录->号, errno, strerror(errno));
    return NULL;
}

void 关录(录 *某录)
{
    if (某录)
    {
        close(某录->号);
        free(某录);
    }
}