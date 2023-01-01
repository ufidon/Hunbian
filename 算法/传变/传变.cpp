#include <iostream>
#include <iomanip>
#include <string>
#include <sstream>
#include <vector>
#include <map>
#include <algorithm>
#include <iterator>
#include <functional>
#include <random>
#include <numeric>
#include <utility>
#include <cmath>

using namespace std;

typedef int 基因;
typedef vector<基因> 基因链;
typedef vector<基因链> 代;
typedef double 适度;

#define 印个体 印基链
#define 印种群 印代
#define 𝛑 3.14159265359

typedef int 号型;
typedef double 键型;
typedef vector<键型> 列型;

class 传变
{
public:
    // 课本445页算法
    /*
    1. 造初代, 随机造之
    2. 适应力(适度)为函数值
    3. 选种:
    4. 杂交变异
    5. 停
    */

    static void 传变选(
        代 &某代,
        代 &种,
        void (*造初代)(代 &初代),
        void (*算适度)(代 &某代, vector<适度> &代适度),
        void (*选种)(代 &某代, vector<适度> &代适度, 代 &种),
        void (*交变)(代 &父, 代 &子),
        bool (*停)(代 &某代))
    {
        vector<适度> 代适度;
        造初代(某代);

        do
        {
            选种(某代, 代适度, 种);
            交变(某代);
        } while (!停(某代));
    }

    // 课本445页第10.2.2节
    /* 求 f(x) = sin(𝛑x/256) 在区间[0,255]上之最大值
    x 为整数.
    1. 基因链: 8位整数, 每位即一基因
    2. 造初代, 随机造之
    3. 适应力(适度)为函数值
    4. 选种:
    5. 杂交变异
    6. 停
    */
    static void 造初代1(代 &初)
    {
        int 初代个数 = 8, 个值;
        uniform_int_distribution<int> 造个体(0, 255);

        while (初代个数--)
        {
            基因链 个体;
            int 基因链长 = 8;
            个值 = 造个体(随机源);
            // cout << "个值:" << hex << 个值 << endl;
            while (--基因链长 >= 0)
            {
                if (个值 & 0x80)
                    个体.push_back(1);
                else
                    个体.push_back(0);

                个值 <<= 1;
            }

            初.push_back(个体);
        }
#ifdef _调试
        cout << 印种群(初) << endl;
#endif        
    }

    static void 算个体适度1(基因链 &个体, int &x, 适度 &度)
    {
        int 基因链长 = 个体.size();
        int 位 = 基因链长;
        x = 0;
        while (--位 >= 0)
        {
            x += 个体[基因链长 - 1 - 位] << 位;
        }

        度 = sin(𝛑 * x / 256);
#ifdef _调试        
        cout << "x = " << x << ", 适度 = " << 度 << endl;
#endif        
    }

    static void 算适度1(代 &某代, vector<适度> &代适度)
    {
        int 个体数 = 某代.size();
        vector<int> x(个体数, 0);

        for (int 个 = 0; 个 < 个体数; 个++)
        {
            算个体适度1(某代[个], x[个], 代适度[个]);
        }

#ifdef _调试
        vector<适度> 归一适度(个体数, 0), 累积适度(个体数, 0);
        cout << 印列<适度>(代适度) << endl;
        适度 代适度和 = accumulate(代适度.begin(), 代适度.end(), 0.0);
        cout << "代适度和:" << 代适度和 << endl;
        transform(代适度.begin(), 代适度.end(), 归一适度.begin(),
                  [=](适度 元)
                  { return 元 / 代适度和; });
        partial_sum(归一适度.begin(), 归一适度.end(), 累积适度.begin());

        cout << left << setfill(' ') << setw(20) << "个体" << setw(8) << "x" << setw(13) << "适度" << setw(12) << "归一适度" << setw(13) << "累积适度" << endl;

        for (int 个 = 0; 个 < 个体数; 个++)
        {
            cout << left << setfill(' ') << setw(20) << 印个体(某代[个]) << setw(6) << x[个] << setw(10) << setprecision(3) << 代适度[个] << setw(10) << setprecision(3) << 归一适度[个] << setw(10) << setprecision(3) << 累积适度[个] << endl;
        }
#endif //  _调试
    }

    static void 选种(代 &某代, vector<适度> &代适度, 代 &种)
    {
        int 个体数 = 某代.size();
        uniform_real_distribution<> 挑种(0, 1);
        vector<适度> 归一适度(个体数, 0), 累积适度(个体数, 0);

        适度 代适度和 = accumulate(代适度.begin(), 代适度.end(), 0.0);
        transform(代适度.begin(), 代适度.end(), 归一适度.begin(),
                  [=](适度 元)
                  { return 元 / 代适度和; });
        partial_sum(归一适度.begin(), 归一适度.end(), 累积适度.begin());

        // 转轮盘选种: 良种可用多次且被选中概率大
        while (个体数--)
        {
            种.push_back(某代[折半寻(累积适度, 挑种(随机源))]);
        }
#ifdef _调试        
        cout << 印种群(种) << endl;
#endif        
    }


    static void 交换(const 基因链 &父, const 基因链 &母, 基因链 &子, const double 交换概率)
    {

    }

    static void 变异(const 基因链 &父, 基因链 &子, const double 变异概率)
    {
        
    }


    static 号型 折半寻(const 列型 &数, const 键型 &标)
    {
        号型 左, 中, 右, 位 = -1;
        左 = 0;
        右 = 数.size() - 1;

        while (左 <= 右 && 位 == -1)
        {
            中 = (左 + 右) / 2;

            if (标 == 数[中])
                return 中;
            else if (标 < 数[中])
                右 = 中 - 1;
            else
                左 = 中 + 1;
        }

        if (标 <= 数[中])
            位 = 中;
        else
            位 = 中 + 1;

        return 位;
    }

public:
    static default_random_engine 随机源;
    传变()
    {
        随机源.seed(31415926);
    };

    static string 印基因(const 基因 &基)
    {
        stringstream ss;
        ss << 基;
        return ss.str();
    }

    static string 印基链(const 基因链 &链)
    {
        stringstream ss;
        ss << "(";
        for (int 号 = 0; 号 < 链.size(); 号++)
        {
            ss << 印基因(链[号]) << (号 == 链.size() - 1 ? "" : ",");
        }
        ss << ")";
        return ss.str();
    }

    static string 印代(const 代 &某代)
    {
        stringstream ss;
        ss << "---------共" << 某代.size() << "个---------" << endl;
        int 个 = 1;
        for (auto 个体 : 某代)
            ss << setw(4) << 个++ << ":" << 印个体(个体) << endl;
        return ss.str();
    }

    template <typename T>
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
};

default_random_engine 传变::随机源 = default_random_engine();

int main(int argc, char const *argv[])
{
    代 初代, 种;
    传变::造初代1(初代);
    vector<适度> 代适度(初代.size(), 0);
    传变::算适度1(初代, 代适度);
    传变::选种(初代, 代适度, 种);

    return 0;
}
