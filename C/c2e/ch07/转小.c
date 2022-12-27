#include <stdio.h>
#include <ctype.h>

int main(int argc, char const *argv[])
{
    int 字;

    while ((字 = getchar()) != EOF)
    {
        putchar(tolower(字));
    }
    
    return 0;
}


/*练习
1. 根据程序名(argv[0])转, 转大则转大写, 转小则转小写.
*/
