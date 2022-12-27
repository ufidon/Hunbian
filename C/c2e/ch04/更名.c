//  clang -E 更名.c

#include <stdio.h>
#include <string.h>
#include <math.h>

#define 死循环  while(1)
#define 大者(甲,乙) (((甲)>(乙))?(甲):(乙))

#if defined _MATH_H
    #define 平方(数)    pow(数,2)
#else
    #define 平方(数)    ((数)*(数))
#endif

// 更为引串: #串->"串"
#define 印障(式)    printf(#式 "=%g\n", 式)
// 黏: 甲##乙->甲乙
#define 并(属性,值)   属性##值

int main(int argc, char const *argv[])
{
    int 甲=20, 乙=30;
    char 姓[] = "孙", 名[] = "悟空";
    int 孙悟空年龄 = 7000;

    printf("%d的平方是%lf\n", 甲, 平方(甲));

    printf("%d与%d中大者是%d.\n", 甲, 乙, 大者(甲,乙));
    printf("%s%s已有%d岁.\n", 姓, 名, 并(孙悟空,年龄));
    印障(2.0/3);

    死循环;

    return 0;
}

/*练习
14. 定义更名交换(某型, 甲,乙)交换为某型之甲乙.
*/
