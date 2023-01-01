#include <iostream>
#include <vector>
#include <algorithm>
#include <functional>
#include <string>
#include <complex>
#include <cstddef>
using namespace std;

function<double(double, double)> 求面积(int 形)
{
    switch (形)
    {
    case 0: // 圆
        return [](double r, double 未用)
        { return 3.1415926 * r * r; };
        break;
    case 1: // 长方形
        return [](double 长, double 宽)
        { return 长 * 宽; };
    case 2: // 直角三角形
        return [](double 直边1, double 直边2)
        { return 直边1 * 直边2 / 2; };
    default:
        return [](double, double) -> double
        { return 0; };
        break;
    }
}

int main()
{
    auto 介绍 = []
    {
        cout << "这是一个能" << endl;
    };
    介绍();

    []
    {
        cout << "直接调用能" << endl;
    }();

    auto 问候 = [](const string &名)
    {
        cout << "嘿" << 名 << ", 您好!" << endl;
    };
    问候("张三");

    auto 三角形周长 = [](double 边1, double 边2, double 边3) -> double
    {
        return (边1 + 边2 + 边3);
    };
    cout << 三角形周长(3, 4, 5) << endl;

    // 顺手牵羊
    double 高 = 3, 宽 = 4;
    cout << "(" << 高 << "," << 宽 << ")" << endl;
    auto 直角三角形周长 = [高, &宽]() -> double
    {
        // 高++; // 不能编译, 高为只读传值
        宽++;
        return (高 + 宽 + sqrt(高 * 高 + 宽 * 宽));
    };
    cout << 直角三角形周长() << endl;
    cout << "(" << 高 << "," << 宽 << ")" << endl;

    auto 面积 = [高, &宽]() mutable -> double
    {
        宽--, 高++; // 高可读写, 可写传值
        return 宽 * 高 / 2;
    };
    cout << 面积() << endl;
    cout << "(" << 高 << "," << 宽 << ")" << endl;

    double 半径 = 1;
    auto 求圆面积 = 求面积(0);
    cout << 求圆面积(半径,0) << endl;
}
