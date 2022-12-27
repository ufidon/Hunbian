#include <stdio.h>
#include <ctype.h>

#define 行数上限 100

double 串转数(char 串[])
{
    double 数, 幂;
    int 号, 正负;

    for (号 = 0; isspace(串[号]); 号++)
        ; // 略去串首空白

    正负 = (串[号] == '-') ? -1 : 1;
    if (串[号] == '+' || 串[号] == '-')
        号++;

    for(数=0.0; isdigit(串[号]); 号++)
        数 = 10.0 * 数 + (串[号] - '0');

    if(串[号] == '.') 号++;

    for ( 幂 = 1.0; isdigit(串[号]); 号++)
    {
        数 = 10.0 * 数 + (串[号] - '0');
        幂 *= 10.0;
    }
    return 正负 * 数 / 幂;
}

int main(){
    // 注意函数声明
    double 和, 串转数(char []);
    char 行[行数上限];
    int 取行(char 行[], int 行限);

    和 = 0;
    while(取行(行, 行数上限) > 0)
        printf("\t%g\n", 和+= 串转数(行));

    return 0;
}

int 串转整(char 串[])
{
    return (int)串转数(串);
}

int 取行(char 行[], int 行限)
{
    int 字, 号;
    号 = 0;
    while (行限-- && (字 = getchar()) != EOF && 字 != '\n')
    {
        行[号++] = 字;
    }
    if (字 == '\n')
    {
        行[号++] = 字;
    }

    行[号] = '\0';
    return 号;
}

/*练习
2. 拓展串转整支持科学计数法,如 2.99796e+8, 1.602176634E-19。
*/