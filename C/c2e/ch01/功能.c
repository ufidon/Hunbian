#include <stdio.h>

int 幂(int 底, int 指数);

int main(int argc, char const *argv[])
{
    int 指数;

    for ( 指数 = 0; 指数 < 10; 指数++)
    {
        printf("%d %d %d\n", 指数, 幂(2, 指数), 幂(-3, 指数));
    }
    
    return 0;
}

int 幂(int 底, int 指数)
{
    int 号,果;
    果 = 1;
    for ( 号 = 1; 号 <= 指数 ; 号++)
    {
        果 *= 底;
    }
    return 果;
}

/*练习
15. 以函数重写温度转换程序。
*/
