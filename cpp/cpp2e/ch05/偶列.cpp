#include <iostream>
#include <utility>
#include <tuple>
#include <string>
#include <typeinfo>
#include <memory>
#include <functional>
#include <future>

using namespace std;

template <int 号, int 长, typename... 列>
struct 印列
{
    static void 印(ostream &出, const tuple<列...> &某列)
    {
        出 << get<号>(某列) << (号 + 1 == 长 ? "" : ",");
        印列<号 + 1, 长, 列...>::印(出, 某列);
    }
};
template <int 长, typename... 列>
struct 印列<长, 长, 列...>
{
    static void 印(ostream &出, const tuple<列...> &某列) {}
};
template <typename... 列>
ostream &operator<<(ostream &出, const tuple<列...> &某列)
{
    出 << "[";
    印列<0, sizeof...(列), 列...>::印(出, 某列);
    return 出 << "]";
}

// https://florianjw.de/en/variadic_templates.html
template <typename 能, typename... 列>
void 逐用(能 某能, const 列 &...某列)
{
    (void)initializer_list<int>{
        [&](const auto &参)
        {某能(参); return 0; }(某列)...};
}

template <typename... 列>
void 印(ostream &出, const 列 &...某列)
{
    // (void)initializer_list<int>{(出 << 某列, 0)...};
    //    (void) initializer_list<int>{
    //     [&](const auto& 参){ 出 << 参; return 0;} (某列)...
    //    };
    逐用([&](const auto &参)
         { 出 << 参; },
         某列...);
}

class 复制
{
public:
    复制() { 值 = 0; }
    复制(复制 &复) { 值 = 复.值; }
    复制(const 复制 &复) { 值 = 复.值; }
    复制(复制 &&复) { 值 = 复.值; }

private:
    int 值;
};

class 变参类
{
public:
    变参类(tuple<int, double> &t)
    {
        cout << "以组为参:(" << get<0>(t) << "," << get<1>(t) << ")" << endl;
    }

    template <typename... 多参>
    变参类(多参... 参集)
    {
        cout << "始以多参" << endl;
        // cout << "始以多参:" << 参集 << endl;
        // 印(cout, 参集);
    }
};

// 隐转
void 能(pair<int, const char *> 对) { cout << "(" << 对.first << "," << 对.second << ")" << endl; }
void 功(pair<const int, string> 对) { cout << "(" << 对.first << "," << 对.second << ")" << endl; }

typedef pair<string, double> 名重;
int main()
{
    名重 张三("张三", 64.22);
    cout << get<0>(张三) << "重" << get<1>(张三) << "公斤" << endl;
    cout << "变量张三有" << tuple_size<名重>::value << "个值" << endl;
    cout << "其首值类型为" << typeid(tuple_element<0, 名重>::type).name() << endl;

    pair<int, const char *> 对(3, "张三");
    功(对);
    能(对);

    pair<复制, int> 偶;

    tuple<int, double> 偶2(3, 3.14);
    pair<int, 变参类> 偶3(4, 偶2);
    pair<int, 变参类> 偶4(piecewise_construct, make_tuple(5), 偶2);

    int 数 = 1;
    cout << "改变前:" << 数;
    auto 偶5 = make_pair(ref(数), ref(数));
    偶5.first++;
    偶5.second++;
    cout << " 改变后:" << 数 << endl;

    int 前 = 0;
    double 后 = 0;
    tie(ignore, 后) = 偶2;
    cout << "前:" << 前 << " 后:" << 后 << endl;
    auto [甲, 乙] = 偶2;
    cout << "甲:" << 甲 << " 乙:" << 乙 << endl;

    tuple<string, int, double> 列("李四", 7, 67.55);
    cout << 列 << endl;
}