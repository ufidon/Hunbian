#include <iostream>
#include <vector>
#include <string>
#include <complex>
#include <cstddef>
using namespace std;

template <typename 参型>
void 印(const 参型 &参)
{
    for (auto 元 : 参)
    {
        cout << 元 << " ";
    }
    cout << endl;
}

template <typename 参型>
void 印(initializer_list<参型> 集)
{
    for (auto 步 = 集.begin(); 步 != 集.end(); ++步)
    {
        cout << *步 << " ";
    }
    cout << endl;
}

typedef complex<double> 复数;
void 印(const 复数 &参)
{
    cout << 参 << endl;
}

template <typename 参>
void 显(const 参 &参1)
{
    cout << 参1 << endl;
}

template <typename 参>
void 显(const string &名, const 参 &参1)
{
    cout << 名 << "=" << 参1 << endl;
}

class 有理数
{
public:
    有理数(const int 子, const int 母) : 分子(子), 分母(母)
    {
        cout << "调有理数(const int 子, const int 母)" << endl;
    }
    有理数(initializer_list<int> 始列)
    {
        cout << "调有理数(initializer_list<int> 始列)" << endl;
        auto 步 = 始列.begin();
        分子 = *步++;
        分母 = *步++;
    }

public:
    int 分子, 分母;
};

class 默转
{
public:
    默转(double 横, double 纵) : 横(横), 纵(纵)
    {
        cout << "允默转型" << endl;
    }
    explicit 默转(double 横, double 纵, double 深) : 深(深)
    {
        this->横 = 横;
        this->纵 = 纵;
        cout << "禁默转型" << endl;
    }

public:
    double 横, 纵, 深;
};

void 用(const 默转 &参) {}

int main()
{
    int 列[]{8, 6, 3};
    印(列);
    vector<int> 行{2, 4, 6, 8};
    印(行);
    vector<string> 水果{"苹果", "桃", "李", "香蕉"};
    印(水果);
    complex<double> 复数1{8, 10};
    印(复数1);

    int 未定义;
    显("未定义", 未定义);
    int 始零{};
    显("始零", 始零);
    int *指哪;
    显("指哪", 指哪);
    int *空指{};
    显("空指", 空指);

    int 数1(1.1); // 被警告
    显("数1", 数1);
    int 数2 = 2.1; // 被警告
    显("数2", 数2);
    // int 数3{3.1}; // 始列中double 不能窄成 int
    // 显("数3", 数3);
    // int 数4 = {4.1};  // 始列中double 不能窄成 int
    // 显("数4", 数4);
    char 符1{30};
    显("符1", 符1);
    // char 符2{123456}; // 始列中整数123456不能窄成char
    // 显("符2", 符2);

    vector<int> 同可{2, 3, 5, 7, 11};
    印(同可);
    // vector<int> 异否{2, 4, 6.0, 8.0}; // 始列中double 不能窄成 int
    // 印(异否);

    印({3, 1, 4, 1, 5, 9});

    有理数 有1(1, 2);
    有理数 有2{3, 4};
    有理数 有3{5, 6, 7};
    有理数 有4 = {8, 9};

    默转 矢1(3, 4);
    默转 矢2{3, 4};
    默转 矢3(3, 4, 5);
    默转 矢4{3, 4, 5};
    默转 矢5 = {3, 4};
    // 默转 矢6 = {3, 4, 5}; // 禁默转型

    用({1, 2});
    // 用({1, 2, 3}); // 禁默转型
    用(默转{1, 2});
    用(默转{1, 2, 3});
}