#include <fcntl.h>
#include <stdlib.h>
#include <unistd.h>

// #define _调试
#include "开取.h"

#define 权限 0666

// 开档返档指
档 *开档(char *名, char *权)
{
    int 档号;
    档 *指;

    if (*权 != 'd' /*读*/ && *权 != 'x' /*写*/ && *权 != 'f' /*附*/)
        return 空;
    for (指 = _出入块; 指 < _出入块 + 开档丄数; 指++)
        if ((指->旗 & (_读 | _写)) == 0)
            break; // 找到可用块
    if (指 >= _出入块 + 开档丄数)
        return 空;

    印调("开档: 找到读写块%ld\n", 指 - _出入块);
    if (*权 == 'x')
        档号 = creat(名, 权限);
    else if (*权 == 'f')
    {
        if ((档号 = open(名, O_WRONLY, 0)) == -1)
            档号 = creat(名, 权限);
        lseek(档号, 0L, 2); // 移读写头至档尾
    }
    else
        档号 = open(名, O_RDONLY, 0);

    if (档号 == -1)
        return 空;

    指->号 = 档号;
    指->余字数 = 0; // 初设0以调用_填储
    指->储位 = 空;
    指->旗 = (*权 == 'd') ? _读 : _写;

    return 指;
}

// 配填入储
int _填储(档 *指)
{
    int 库夵;
    印调("开始填储\n");
    if ((指->旗 & (_读 | _尾 | _错)) != _读)
        return 档尾;
    库夵 = (指->旗 & _瞬) ? 1 : 储夵;
    if (指->储位 == NULL)
        if ((指->储位 = (char *)malloc(库夵)) == NULL)
            return 档尾;
    指->翌字位 = 指->储位;
    指->余字数 = read(指->号, 指->翌字位, 库夵);
    if (--指->余字数 < 0)
    {
        if (指->余字数 == -1)
            指->旗 |= _尾;
        else
            指->旗 |= _错;

        指->余字数 = 0;
        return 档尾;
    }
    印调("_填储: %s\n", 指->翌字位);
    return (unsigned char)*指->翌字位++;
}

int _冲储(int 字, 档 *指)
{
    return (write(指->号, &字, 1));
}

档 _出入块[开档丄数] = {
    {0, (char *)0, (char *)0, _读, 0},
    {0, (char *)0, (char *)0, _写, 1},
    {0, (char *)0, (char *)0, _写 | _瞬, 2},
};

int main(int argc, char const *argv[])
{
    char 字;

#define 读档
#ifdef 读档
    int 数 = 100;
    char 档名[] = "./开取.h";
    档 *文;
    if ((文 = 开档(档名, "d")) == NULL)
    {
        printf("错: 开\"%s\"失败.\n", 档名);
        return 1;
    }

    while ((字 = 取字(文)) != 档尾 /* && 数-- > 0 */)
        放字于出(字);

#else

    while ((字 = 取字自入()) != 档尾)
        放字于出(字);
#endif
    return 0;
}

/*练习
2. 以位域重写开档及_填储. 比代码量及跑速。
3. 写 _冲储、fflush及fclose。
4. 写 int 移磁头(档 *指， long 距, int 起点)，移动磁头至离起点某距离，并测试其可与读、写函式工作。
*/
