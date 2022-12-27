#include <stdio.h>

const double π = 3.14159265359;

int 串长(const char 串[])
{
    int 位;
    while (串[位] != 0) // 或 != '\0'
    {
        ++位;
    }
    return 位;
}

int main(int argc, char const *argv[])
{
    printf("反斜杠：\\ 退格\b- 问号\? ? 单引号: ' \' 双引号: \" \n");
    printf("反斜杠: %x 退格: %x 问号: %x ? 单引号: %x '  双引号: %x \n",
     '\\', '\b', '\?', '\'', '\"');

    printf("反斜杠：\x5c 退格\x08- 问号\x3f ? 单引号: ' \x27 双引号: \x22 \n");

    printf("\f");
    printf("\n");

    printf("铃：\a 喂纸: \f 行尾: \n 回车: \r 横跳格: \t 纵跳格: \v \n");
    printf("铃: %x 喂纸: %x 行尾: %x 回车: %x 横跳格: %x 纵跳格: %x \n",
        '\a', '\f','\n','\r','\t','\v');

    printf("回车不见了?\r");

    printf("空符: %x\n", '\0');

    char 串[] = "自古天下是一家";
    int 长度 = 串长(串);
    printf("'自古天下是一家'在内存中储以万国码-8, 占用%d个字节\n"
    "每个常用汉字占3个字节, 故其长%d个汉字.\n", 
    长度, 长度/3);

    double 半径 = 3;
    printf("半径为%.0f的圆面积为%.2f\n", 半径, π*半径*半径);
    
    return 0;
}

/*练习
1. 编程确定char, short和int值域。
*/
