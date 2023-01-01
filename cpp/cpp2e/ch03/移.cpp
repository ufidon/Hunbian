#include <iostream>
#include <vector>
#include <algorithm>
#include <string>
#include <complex>
#include <cstddef>
using namespace std;

class 类
{
public:
    类() { 串 = "复制或移动"; }
    类(const 类 &左值){串 = 左值.串;} // 复构器
    类(类 &&右值){串 = 右值.串;}      // 移构器

    类 &operator=(const 类 &左值){串 = 左值.串;} // 复赋子
    类 &operator=(类 &&右值){串 = 右值.串;}      // 移赋子
private:
    string 串;
};

void 复或移(类 &例){}
void 复或移(const 类 &例){}
void 复或移(类 &&例){}

类 复或移()
{
    类 例;
    return 例;
};

// 类&& 复或移(){ // 不能编译
//     类 例;
//     return 例;
// };

int main()
{
}