#include <iostream>
#include <vector>
#include <algorithm>
#include <string>
#include <complex>
#include <cstddef>
using namespace std;

void 障()
{
    cout << "能抛任何异常" << endl;
}
void 无障() noexcept
{
    cout << "不能抛任何异常" << endl;
}

int main()
{
    cout << "\\\\字"
         << endl;
    cout << R"(\\字)" << endl;
    cout << "井\\\n 底 \\之()\"蛙\n  " << endl;
    cout << R"|(井\
 底 \之()"蛙
            )|" << endl;

    // 需终端编码配合
    cout << u8"8位万国码" << endl;
    cout << u"16位万国码" << endl;
    cout << U"32位万国码" << endl;
    wcout << L"宽符万国码" << endl;

}