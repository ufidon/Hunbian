#include <iostream>
#include <string>
using namespace std;

enum class 周 : char
{
    一 = '1',
    二,
    三,
    四,
    五,
    六,
    日
};

int main()
{
    cout << (char)(周::一) << endl;
}