#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

#define 丄式长 100
#define 数哉 '0' // 自式中析出数
#define 筒深 100
int 项位 = 0;
double 筒[筒深];

int 取项(char[]);
void 入筒(double);
double 出筒(void);
int 取字(void);
void 放回(int);

// 算附后置式算器

int main(int argc, char const *argv[])
{
    int 项型;
    double 右数;
    char 式[丄式长];

    while ((项型 = 取项(式)) != EOF)
    {
        switch (项型)
        {
        case 数哉:
            入筒(atof(式));
            break;
        case '+':
            入筒(出筒() + 出筒());
            break;
        case '-':
            右数 = 出筒();
            入筒(出筒() - 右数);
            break;
        case '*':
            入筒(出筒() * 出筒());
            break;
        case '/':
            右数 = 出筒();
            if (右数 != 0.0)
                入筒(出筒() / 右数);
            else
                printf("障: 不能除以〇\n");
            break;
        case '\n':
            printf("\t%.8g\n", 出筒());
            break;
        default:
            printf("障: 未知算式 %s\n", 式);
            break;
        }
    }

    return 0;
}

void 入筒(double 数)
{
    if (项位 < 筒深)
        筒[项位++] = 数;
    else
        printf("障: 筒满了装不下%g\n", 数);
}

double 出筒(void)
{
    if (项位 > 0)
        return 筒[--项位];
    else
    {
        printf("障: 空筒!\n");
        return 0.0;
    }
}

// 从式中取下一个算符或数
int 取项(char 式[])
{
    int 号, 字;
    while ((式[0] = 字 = 取字()) == ' ' || 字 == '\t')
        ;

    式[1] = '\0';

    if (!isdigit(字) && 字 != '.')
        return 字;

    号 = 0;
    if(isdigit(字)) // 取数项整数部分
        while(isdigit(式[++号] = 字 = 取字()));

    if(字 == '.') // 取小数点右部
        while(isdigit(式[++号] = 字 = 取字()));

    式[号] = '\0';

    if(字 != EOF)
        放回(字);

    return 数哉;
}

#define 段长    100
char 工段[段长];
int 段位 = 0;

int 取字(void)
{
    return (段位 > 0) ? 工段[--段位] : getchar();
}

void 放回(int 字)
{
    if(段位 >= 段长)
        printf("放回: 工段已满.\n");
    else
        工段[段位++] = 字;
}

/*练习: 拓展算器功能
3. 求余%及支持负数.
4. 1)不出筒, 印筒顶项和复制筒顶项; 2) 交换筒顶两项; 3) 清筒.
5. 调库函数如 pow, exp, log, sin 等等
6. 1) 引入变量并处理变量; 2) 增加一变量存最近印值
7. 写函数放串(char 串)放串回输入
8. 若放回多于一个字的情况不存在, 如何简化函数放回及取字.
9. 程序中的放回及取字函数不能处理放回的EOF, 如何改进?
10. 一次读入一行后再处理, 则无需取字及放回函数. 请实现此功能.
11. 用内部静变量改写取项函数以无需放回函数.
*/