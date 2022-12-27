#include <stdio.h>
#include <ctype.h>

// 串转整数
int 串转整(char 串[])
{
    int 号, 数, 正负号;
    // 掠过开头空白
    for(号=0; isspace(串[号]); 号++) ;
    // 录正负号
    正负号 = (串[号] == '-') ? -1:1;
    // 录完略过
    if (串[号] == '+' || 串[号] == '-') 号++;
    for(数=0; isdigit(串[号]); 号++)
        数 = 10*数 + (串[号] - '0');

    return 正负号*数;
}

int main(int argc, char const *argv[])
{
    char 串[] = "-667384";
    printf("\"%s\"转整数为%d\n", 串, 串转整(串));

    return 0;
}
