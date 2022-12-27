#include <stdio.h>

// 在序列中找数, 返其列号. 未找到返-1

int 折半查找(int 数, int 列[], int 列长)
{
    int 左号, 中号, 右号;
    左号 = 0;
    右号 = 列长 - 1;
    while (左号 <= 右号)
    {
        中号 = (左号 + 右号) / 2;
        if (数 < 列[中号]) // 数在左半列
            右号 = 中号 - 1;
        else if (数 > 列[中号]) // 数在右半列
            左号 = 中号 + 1;
        else
            return 中号;
    }
    return -1;
}

int main(int argc, char const *argv[])
{
    int 列[] = {3, 5, 6, 8, 11, 13, 15, 20, 29, 34};
    int 列长 = 10;
    int 数 = 12;
    int 列号 = 0;
    if ((列号 = 折半查找(数, 列, 列长)) == -1)
        printf("%d不在数列中.\n", 数);
    else
        printf("%d在列中第%d号.\n", 数, 列号);
    return 0;
}

/*练习
1. 折半查找中, 待找数与每个中点数比较了两次, 能否只比一次?
*/
