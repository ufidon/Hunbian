#include <stdio.h>
#include <string.h>

void 印数(int 数)
{
    if (数 < 0)
    {
        putchar('-');
        数 = -数;
    }

    // putchar(数 % 10 + '0'); // 得逆串

    if (数 / 10)
        印数(数 / 10);

    putchar(数 % 10 + '0'); // 得正串
}

static char 串[100];
void 数转串(int 数)
{
    // 静态变量只在第一次调用时始化
    static int 位 = 0;
    static int 自调次数 = 0;
    自调次数++;

    if (数 < 0)
    {
        数 = -数;
        串[位++] = '-';
    }

    // 串[位] = 数 % 10 + '0'; // 得逆串? 请分析

    if (数 / 10)
        数转串(数 / 10);

    串[位] = 数 % 10 + '0';
    printf("位: %d, %c\n", 位, 串[位]);

    位++;
}

void 快排(int 列[], int 左号, int 右号)
{
    int 号, 桩号;
    void 换(int 列[], int 甲, int 乙);

    if (左号 >= 右号)
        return;

    换(列, 左号, (左号 + 右号) / 2);
    桩号 = 左号; // 左首为桩

    for (号 = 左号 + 1; 号 <= 右号; 号++)
        if (列[号] < 列[左号])
            换(列, ++桩号, 号);
    换(列, 左号, 桩号);
    快排(列, 左号, 桩号 - 1);
    快排(列, 桩号 + 1, 右号);
}

void 换(int 列[], int 甲, int 乙)
{
    int 暂;
    暂 = 列[甲];
    列[甲] = 列[乙];
    列[乙] = 暂;
}

void 印列(int 列[], int 列长)
{
    int 号;
    printf("[");
    for (号 = 0; 号 < 列长; 号++)
    {
        printf("%d", 列[号]);
        printf("%s", 号 == 列长 - 1 ? "" : ",");
    }

    printf("]");
}

int main(int argc, char const *argv[])
{
    int 数 = -123;

    印数(数);
    putchar('\n');

    数转串(数);
    printf("%s\n", 串);

    int 列[] = {19, 3, 5, 17, 11, 23, 13, 2, 4};
    int 列长 = 9;

    printf("排序前: ");
    印列(列, 列长);
    printf("\n");

    快排(列, 0, 列长-1);

    printf("排序后: ");
    印列(列, 列长);
    printf("\n");

    return 0;
}

/*练习
12. 以调己法写函数整转串.
13. 以调己法实现在地逆串.
*/
