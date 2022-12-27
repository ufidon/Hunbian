#include <stdio.h>

//typedef int 温型;
typedef double 温型;

#define 低  0
#define 高  300
#define 步长    10

int main(int argc, char const *argv[])
{
    // 摄氏度 = 5/9*(华氏度-32)
    温型 摄氏度, 华氏度;

#ifdef 当循环
    温型 低 = 0, 高=300, 步长 = 10;
    华氏度 = 低;
    printf("%5s\t%5s\n", "华氏度", "摄氏度");
    while (华氏度 <= 高)
    {
        摄氏度 = 5*(华氏度 - 32)/9;
        //printf("%3d\t%3d\n", 华氏度, 摄氏度);
        printf("%5.2f\t%5.2f\n", 华氏度, 摄氏度);
        华氏度 += 步长;
    }
#endif // 当循环

#define 逐循环
#ifdef 逐循环
    printf("%5s\t%5s\n", "华氏度", "摄氏度");
    for ( 华氏度 = 低; 华氏度 <= 高; 华氏度+=步长)
    {
        printf("%5.2f\t%5.2f\n", 华氏度, 5*(华氏度 - 32)/9);
    }
#endif // 逐循环    
    

    return 0;
}

/* 练习
3. 给上程序输出增加表格线。
4. 写一个程序，输出摄氏度至开氏温度转换表：开氏温度 = 摄氏度 + 273.15。
5. 修改上程序，逆序输出， 即从华氏温度300度至0度。
*/
