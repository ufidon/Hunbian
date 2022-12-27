#include <stdio.h>


int main(int argc, char const *argv[])
{
    // 统计数字、空白及其余字符
    int 字, 号, 空白数, 其余字符数, 数字数[10];
    空白数 = 其余字符数 = 0;

    for(号=0; 号<10; 号++)
        数字数[号] = 0;

    while ((字 = getchar()) != EOF)
    {
        switch (字)
        {
        case '0': case '1': case '2': case '3': case '4': 
        case '5': case '6': case '7': case '8': case '9':
            数字数[字 - '0']++;
            break;
        case ' ':
        case '\n':
        case '\t':
            空白数++;
            break;
        default:
            其余字符数++;
            break;
        }
    }
    printf("数字: ");
    for(号=0; 号<=9; 号++)
        printf("%2d ", 号);
    printf("\n");
    printf("数目: ");
    for(号=0; 号<=9; 号++)
        printf("%2d ", 数字数[号]);
    printf("\n");

    printf("空白数 = %d\n其余字符数 = %d\n", 空白数, 其余字符数);
    return 0;
}

/*练习
2. 写函数 void 特变转义(char 自[], char 至[]), 将字从自串复制到至串过程中, 把特殊字符换成其转义, 如换行->\n, 横跳格->\t等等.
    再写一函数做相反转换.

*/