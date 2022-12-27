#include <stdio.h>

int 闰年(int 年)
{
    // 四百年一大闰四年一小闰
    if (年 % 400 == 0 || (年 % 4 == 0 && 年 % 100 != 0))
    {
        printf("%d是闰年.\n", 年);
        return 1;
    }
    else
    {
        printf("%d不是闰年.\n", 年);
        return 0;
    }
}

int main(int argc, char const *argv[])
{
    int 年 = 1995;
    while (年 <= 2006)
    {
        闰年(年++);
    }

    return 0;
}

/*练习
2. 改写第五行, 不用||及&&.
3. 编程演示||及&&之短路(得手即止)效应.
*/