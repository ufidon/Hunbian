#ifndef _历录_H
#define _历录_H

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

//#define _调试
#define 档名长限    50

#undef S_IFMT 
#undef S_IFDIR
#undef S_IFCHR
#undef S_IFBLK
#undef S_IFREG

#define S_IFMT  0160000
#define S_IFDIR 0040000
#define S_IFCHR 0020000
#define S_IFBLK 0060000
#define S_IFREG 0100000


typedef struct _录项
{
    long 节号;
    char 名[档名长限+1];
}录项;

typedef struct _录
{
    int 号;
    录项 项;
} 录;

录 *开录(const char *录名);
录项 *读录(录 *某录);
void 关录(录 *某录);

void 印调(char *格式, ...)
{
#ifdef _调试    
    va_list 参;
    va_start(参, 格式);
    fprintf(stderr, "信:");
    vfprintf(stderr, 格式, 参);
    fprintf(stderr, "\n");
    va_end(参);
#endif
}

#endif // _历录_H