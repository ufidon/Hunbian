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

    // explicit 类(const string &串) // 禁串默转类
    // {
    //     this->串 = 串;
    // }
    类(const string &串) // 允串默转类
    {
        this->串 = 串;
    }
    friend ostream& operator<<(ostream& 出, const 类& 例);
private:
    string 串;
};
ostream& operator<<(ostream& 出, const 类& 例)
{
    出 << 例.串 << endl;
    return 出;
}

int main()
{
    for (int 元 : {3, 1, 4, 1, 5, 9})
    {
        cout << 元 << " ";
    }
    cout << endl;

    for (int 元 : vector<int>{2, 9, 9, 7, 9, 6})
    {
        cout << 元 << " ";
    }
    cout << endl;

    vector<int> 列{6, 6, 7, 3, 8, 4};
    for (auto 步 = begin(列); 步 != end(列); ++步)
    {
        cout << *步 << " ";
    }
    cout << endl;

    vector<string> 串列{"一念天堂", "一念地狱"};
    for (const 类 &元 : 串列) // 类构器须支持串默转类
    {
        cout << 元 << endl;
    }
}