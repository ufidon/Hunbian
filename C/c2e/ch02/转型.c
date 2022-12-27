#include <stdio.h>
#include <string.h>

// 整数串转整数
int 串转数(char 串[])
{
    int 号, 数 = 0;

    for (号 = 0; 串[号] >= '0' && 串[号] <= '9'; 号++)
    {
        数 = 数 * 10 + (串[号] - '0');
    }

    return 数;
}

int 转小写(int 字母)
{
    if(字母 >= 'A' && 字母 <= 'Z')
        return 字母 + 'a' - 'A';
    else
        return 字母;
}

// 造随机数
unsigned long int 翌 = 1;
void 设随机数种(unsigned int 种)
{
    翌 = 种;
}

int 随机数(void)
{
    翌 = 翌 * 1103515245 + 12345;
    return (unsigned int)(翌/65536) % 32768;
}

int main(int argc, char const *argv[])
{
    int 号;
    char 串[] = "314159";
    printf("串'%s'转整数为%d\n", 串, 串转数(串));

    char 拼音[] = "Yu Xue FeiFei"; // 长13个字符
    char 小写[14];
    for ( 号 = 0; 号 <  13; 号++)
    {
        小写[号] = 转小写(拼音[号]);
    }
    小写[13] = '\0';
    
    printf("'%s'转小写成'%s'\n", 拼音, 小写);

    int 数目=0;
    设随机数种(31415);
    printf("十个随机数: ");
    while (数目++ < 10)
    {
        printf("%d ", 随机数());
    }
    printf("\n");

    return 0;
}

/*练习
3. 写程序转十六进制串为对应整数. 如"0x1234"->4660.
3'.写程序转八进制串为对应整数. 如"01234"->668.
*/