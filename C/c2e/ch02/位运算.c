#include <stdio.h>

// 从数之某位号起向右取数位
// 位编号始于0, 从右向左编号
unsigned 取片(unsigned 数, int 位号, int 位数)
{
    return (数 >> (位号 + 1 - 位数)) & ~(~0 << 位数);
}

// 计数中1位
int 计1位(unsigned 数)
{
    int 位数 = 0;
    for (; 数 != 0; 数 >>= 1)
        if (数 & 01)
            位数++;
    return 位数;
}

// 以二进制印数, 四位一组, 隔以|
void 贰印(unsigned 数)
{
    int 位数 = sizeof(数) * 8;
    int 组数 = 位数 / 4;

    unsigned int 最高位掩码 = 1 << (位数 - 1);

    printf("(");

    while (位数-- > 0)
    {

        printf("%d", 最高位掩码 & 数 ? 1 : 0);
        数 <<= 1;
        if (位数 % 4 == 0 && 位数 != 0)
            printf("|");
    }

    printf(")\n");
}

int main(int argc, char const *argv[])
{
    unsigned 数 = 0x55aa55aa;
    int 位号 = 7, 位数 = 4;

    printf("0x%x二进制是", 数);
    贰印(数);

    printf("在0x%x内从第%d位起向右取%d位得0x%x\n", 数, 位号, 位数, 取片(数, 位号, 位数));

    printf("0x%x中为1的位数是%d\n", 数, 计1位(数));

    printf("3的二进制是");
    贰印(3);

    return 0;
}

/*练习
6. 写函数unsigned 设片(unsigned 数, int 位号, int 位数, int 值) 将数从位号起位数位设成值, 其余位不变并返结果.
7. 写函数unsigned 反片(unsigned 数, int 位号, int 位数) 将数从位号起位数位取反, 其余位不变并返结果.
8. 写函数unsigned 右转(unsigned 数, int 位数) 将数右转位数位(即转出位填至数左)并返结果.
9. 在二补码数系(参考:https://zh.wikipedia.org/wiki/二補數)中, 数&=(数-1) 删除其最右为1位. 阐述其理并用之改写上程序.
10. 用条件句改写上述函数计1位.
11. 写函数 void 简贰印(unsigned 数), 不印左侧全〇位, 如3只印(11), 26只印(1|1010).
*/
