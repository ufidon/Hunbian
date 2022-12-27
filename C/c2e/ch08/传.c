#include <unistd.h>
#include <stdio.h>

#define 储夵 1000

int 取字(void)
{
    char 字;
    return (read(0, &字, 1) == 1) ? (unsigned char)字 : EOF;
}

int 取字2(void)
{
    static char 储[储夵];
    static char *指 = 储;
    static int 数 = 0;

    if (数 == 0)
    {
        数 = read(0, 储, sizeof 储);
        指 = 储;
    }
    return (--数 >= 0) ? (unsigned char)*指++ : EOF;
}

int main(int argc, char const *argv[])
{
    char 储[储夵];
    int 数;

    while ((数 = read(0, 储, 储夵)) > 0)
        write(1, 储, 数);

    return 0;
}
