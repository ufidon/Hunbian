/* 采用无限精度数学库:
1. GNU 多精度算术库 GNU Multiple Precision Arithmetic Library
2. GNU 多精度浮点可靠库 GNU Multiple Precision Floating-Point Reliable Library (GNU MPFR)
3. GMP 算法:
https://gmplib.org/manual/Algorithms
https://www.mpfr.org/algorithms.pdf
https://www.boost.org/doc/libs/1_79_0/libs/multiprecision/doc/html/index.html
安装库:
sudo apt install libmpfrc++-dev libgmp-dev libmpfr-dev libgmpxx4ldbl libboost-all-dev
编译:
clang++ 数论.cpp -lmpfr -lgmp
*/
#include <iostream>
#include <mpreal.h>
#include <gmpxx.h>
#include <string>
#include <vector>
using namespace mpfr;
using namespace std;

// typedef long long 长整数;
// typedef mpreal 长整数;
typedef mpz_class 长整数;

template <typename T>
class 数论
{
public:
    // 课本480页算法11.1
    static T 最大公约数(T 甲, T 乙)
    {
        if (乙 == 0)
            return 甲;
        else
            return 最大公约数(乙, mod(甲, 乙));
    }

    // 课本484页算法11.2
    // 丄约 = 甲*甲系 + 乙*乙系
    static void 最大公约数(T 甲, T 乙, T &丄约, T &甲系, T &乙系)
    {
        if (乙 == 0)
            丄约 = 甲, 甲系 = 1, 乙系 = 0;
        else
        {
            T 丄约1, 甲系1, 乙系1;
            最大公约数(乙, mod(甲, 乙), 丄约1, 甲系1, 乙系1);
            丄约 = 丄约1;
            甲系 = 乙系1;
            乙系 = 甲系1 - floor(甲 / 乙) * 乙系1;
        }
    }

    // 课本506页算法11.3
    // 解同余方程[乙]_甲・未知量 = [丙]_甲
    static void 解同余(T 甲, T 乙, T 丙, vector<T> &解)
    {
        int 号;
        T 甲系, 乙系, 丄约;

        最大公约数(甲, 乙, 丄约, 甲系, 乙系);
        if (mod(丙, 丄约) == 0)
            for (号 = 0; 号 < 丄约; 号++)
                解.push_back(乙系 * 丙 / 丄约 + 号 * 甲 / 丄约);
    }

    // 课本508页算法11.4
    // 求 ([乙]_甲)^丙
    static void 求幂(T 甲, T 乙, T 丙, T &幂)
    {
        int 号;
        幂 = 1;
        string 贰丙 = 丙.get_str(2);
        for (号 = 贰丙.size() - 1; 号 >= 0; 号--)
        {
            幂 = 幂 * 幂;
            if (贰丙[号] == '1')
                幂 *= 乙;
        }
    }

    static string 印列(const vector<T> &列)
    {
        stringstream ss;
        ss << "[";
        for (int 号 = 0; 号 < 列.size(); 号++)
        {
            ss << 号 << ":" << 列[号] << (号 == 列.size() - 1 ? "" : ",");
        }
        ss << "]";
        return ss.str();
    }

public:
    数论(/* args */){};
};

int main(int argc, char const *argv[])
{
    mpreal::set_default_prec(1024);
    cout.precision(1024);

    mpreal a("123456789012345678901234567890"), b("987654321098765432109876543210");
    cout << "和" << a + b << "\n积" << a * b << endl;
    cout << 数论<mpreal>::最大公约数(a, b) << endl;

    mpreal 丄约, 甲系, 乙系;
    数论<mpreal>::最大公约数(a, b, 丄约, 甲系, 乙系);
    cout << 丄约 << "," << 甲系 << "," << 乙系 << "甲*甲系 + 乙*乙系=" << a * 甲系 + b * 乙系 << endl;

    mpreal 甲 = 8, 乙 = 6, 丙 = 4;
    vector<mpreal> 解;
    数论<mpreal>::解同余(甲, 乙, 丙, 解);
    cout << 数论<mpreal>::印列(解) << endl;


    mpz_class x("123456789");
    std::cout << x.get_str(2) << std::endl; // base 2 representatio

    mpz_class n("257"), m("5"), k("45"), s("0");
    数论<mpz_class>::求幂(n, m, k, s);
    cout << s.get_str() << endl;


    return 0;
}
