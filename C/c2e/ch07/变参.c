#include <stdarg.h>
#include <stdlib.h>
#include <ctype.h>
#include <stdio.h>

void 印(char *格式, ...)
{
    va_list 参;
    char *格, *串值;
    int 整值;
    double 浮值;

    va_start(参, 格式); // 令参指向首个无名参
    for (格 = 格式; *格; 格++)
    {
        if (*格 != '%')
        {
            putchar(*格);
            continue;
        }
        switch (*++格)
        {
        case 'd':
            整值 = va_arg(参, int);
            printf("%d", 整值);
            break;
        case 'f':
            浮值 = va_arg(参, double);
            printf("%f", 浮值);
            break;
        case 's':
            for (串值 = va_arg(参, char *); *串值; 串值++)
                putchar(*串值);
            break;
        default:
            putchar(*格);
            break;
        }
    }
    va_end(参); // 清场
}

int main(int argc, char const *argv[])
{
    印("%s %d %f\n", "羊肉串", 314, 3.14159);
    return 0;
}

/*练习
3. 改进印, 处理更多printf格式.
*/