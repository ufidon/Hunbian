#include <stdio.h>

void 复串1(char *的, char *源)
{
    int 号 = 0;
    while ((的[号] = 源[号]) != '\0')
    {
        号++;
    }
}

void 复串2(char *的, char *源)
{
    while ((*的 = *源) != '\0')
    {
        的++;
        源++;
    }
}

void 复串3(char *的, char *源)
{
    while ((*的++ = *源++) != '\0')
        ;
}

void 复串4(char *的, char *源)
{
    while ((*的++ = *源++))
        ;
}

int 比串1(char *甲, char *乙)
{
    int 号;

    for (号 = 0; 甲[号] == 乙[号]; 号++)
        if (甲[号] == '\0')
            return 0;

    return 甲[号] - 乙[号];
}

int 比串2(char *甲, char *乙)
{
    for (; *甲 == *乙; 甲++, 乙++)
        if (*甲 == '\0')
            return 0;

    return *甲 - *乙;
}

int 串长2(char *串)
{
    char *指 = 串;
    while (*指 != '\0')
    {
        指++;
    }
    return 指 - 串;
}
int main(int argc, char const *argv[])
{
    char *串1 = "飞流直下三千尺,";
    char *串2 = "疑是银河落九天.";

    char 诗[100];
    复串1(诗, 串1);
    复串4(诗 + 串长2(串1), 串2);
    printf("%s\n", 诗);

    printf("%s>%s=%d\n", 串1, 串2, 比串2(串1, 串2));

    return 0;
}

/*练习用指针写
3. 函式黏串(char *左, char *右)黏右串首至左串尾.
4. 函式现尾(char *甲, char *乙)返1若乙现于甲尾否返0
5. 实现考库函式 strncpy, strncat, strncmp
6. 重写前章节函式取行,串转整,整转串,逆串,串号,取项
*/
