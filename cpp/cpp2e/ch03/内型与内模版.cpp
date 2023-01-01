#include <iostream>
#include <string>
using namespace std;



class 型集{
public:   
    typedef int 子型;
};

template <typename 型>
class 内型
{
public:    
    typename 型::子型 * 指;
};

template <typename 型>
class 内模{
public:
    型 值;     
public:
    template <typename 型1>
    void 赋值(const 内模<型1>& 物){ 值 = 物.值;}
};


int main()
{
    内型<型集> 内1;

    内模<double> d;
    内模<int> i;
    d.赋值(d);
    d.赋值(i);
}