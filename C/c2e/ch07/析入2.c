#include <stdio.h>
#include <ctype.h>
#include <string.h>

int main(int argc, char const *argv[])
{
    int 年,月, 日;
    char 月名[10];

    if (scanf("%d %s %d", &日, 月名, &年) == 3)
        printf("%d年%s月%d日\n", 年, 月名, 日);
    if (scanf("%d/%d/%d", &日, &月, &年) == 3)
        printf("%d年%d月%d日\n", 年, 月, 日);
    if (scanf("%d-%d-%d", &日, &月, &年) == 3)
        printf("%d年%d月%d日\n", 年, 月, 日);

    return 0;
}
