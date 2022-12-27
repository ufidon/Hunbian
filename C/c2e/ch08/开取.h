#ifndef _开取_H
#define _开取_H

#include <stdio.h>
#define 空  0
#define 档尾    (-1)
#define 储夵    1024
#define 开档丄数    20

typedef struct _出入储
{
    int 余字数;
    char *翌字位;
    char *储位;
    int 旗;
    int 号;
} 档;

extern 档 _出入块[开档丄数];
#define 入档  (&_出入块[0])
#define 出档  (&_出入块[1])
#define 错档  (&_出入块[2])

enum _档旗{
    _读 = 01,
    _写 = 02,
    _瞬 = 04,
    _尾 = 010,
    _错 = 020
};

int _填储(档 *);
int _冲储(int, 档 *);

#define 是尾(指)  (((指)->旗 & _尾) != 0)
#define 是错(指)  (((指)->旗 & _错) != 0)
#define 档号(指)  ((指)->号)

#define 取字(指) (--(指)->余字数 >= 0 \
                ? (unsigned char) *(指)->翌字位++ : _填储(指))
#define 放字(字,指) (--(指)->余字数 >=0 \
                ? *(指)->翌字位++ = (字) : _冲储((字), 指))                

#define 取字自入()  取字(入档)                
#define 放字于出(字) 放字((字), 出档)

void 印调(char *格式, ...)
{
#ifdef _调试    
    va_list 参;
    va_start(参, 格式);
    fprintf(stderr, "错:");
    vfprintf(stderr, 格式, 参);
    fprintf(stderr, "\n");
    va_end(参);
#endif
}

#endif // _开取_H