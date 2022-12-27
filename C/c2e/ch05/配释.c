#include <stdio.h>

#define 总容量 1000

static char 库[总容量];
static char *块位 = 库;
static char *库首 = 库;
static char *库尾 = 库 + 总容量 - 1;

char *配(int 量)
{
    if (库尾 - 块位 + 1 >= 量)
    {
        块位 += 量;
        return 块位 - 量;
    }
    else
        return 0;
}

void 释(char *块首)
{
    if (块首 >= 库首 && 块首 <= 库尾)
        块位 = 块首;
}

int main(int argc, char const *argv[])
{
    // 分库以存四整数
    int *整 = (int *)配(sizeof(int) * 4);
    for (int 号 = 0; 号 < 4; 号++)
    {
        *(整 + 号) = 100 + 号 * 3;
    }

    for (int 号 = 0; 号 < 4; 号++)
    {
        printf("%d ", *(整 + 号));
    }
    printf("\n");

    释((char *)整);

    return 0;
}
