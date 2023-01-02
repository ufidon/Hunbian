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

#define _调试

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

class 函数最大值
{
public:
    // 课本445页算法
    /*
    1. 造初代, 随机造之
    2. 适应力(适度)为函数值
    3. 选种:
    4. 杂交变异
    5. 停
    确定每代个体数, 可节省重复分配释放内存时间, 改进性能.
    */

    void 逐代演化()
    {
        int 代数 = 0;

        造代(父代); // 初代
        cout << 印种群(父代, "初代") << endl;
        do
        {
            算代适度(父代, 代适度, 归一适度, 累积适度);
            选种(父代, 代适度, 归一适度, 累积适度, 种);
            群交(种, 子代);
            群变(子代);
            cout << 印种群(子代, "子代") << endl;

            父代 = 子代;
            代数++;
        } while (!停(代适度, p止演适度, 代数, p止演代数));
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
    void 造代(代 &初) const
    {
        int 个 = p每代个数, 个值;
        uniform_int_distribution<int> 造个体(0, 255);

        while (--个 >= 0)
        {
            基因链 个体(p基因链长, 0);
            int 基位 = p基因链长;
            个值 = 造个体(随机源);
            // cout << "个值:" << hex << 个值 << endl;
            while (--基位 >= 0)
            {
                if (个值 & 0x80)
                    个体[p基因链长 - 1 - 基位] = 1;
                else
                    个体[p基因链长 - 1 - 基位] = 0;

                个值 <<= 1;
            }

            初[个] = 个体;
        }
#ifdef _调试
        cout << "造代:\n"
             << 印种群(初) << endl;
#endif
    }

    void 算个适度(const 基因链 &个体,
                  int &基度, // 基因度量, 不同应用不同度量
                  适度 &度) const
    {
        int 位 = p基因链长;
        基度 = 0;
        while (--位 >= 0)
        {
            基度 += (个体[p基因链长 - 1 - 位] << 位);
        }

        度 = sin(𝛑 * 基度 / 256);
#ifdef _调试
        cout << "算个适度:\n"
             << "基度 = " << 基度 << ", 适度 = " << 度 << endl;
#endif
    }

    void 算代适度(const 代 &某代, vector<适度> &代适度, vector<适度> &归一适度, vector<适度> &累积适度) const
    {
        vector<int> 代基度(p每代个数, 0);

        for (int 个 = 0; 个 < p每代个数; 个++)
        {
            算个适度(某代[个], 代基度[个], 代适度[个]);
        }

        适度 代适度和 = accumulate(代适度.begin(), 代适度.end(), 0.0);
        transform(代适度.begin(), 代适度.end(), 归一适度.begin(),
                  [=](适度 元)
                  { return 元 / 代适度和; });
        partial_sum(归一适度.begin(), 归一适度.end(), 累积适度.begin());

#ifdef _调试
        cout << 印列<适度>(代适度) << endl;
        cout << "代适度和:" << 代适度和 << endl;

        cout << left << setfill(' ') << setw(20) << "个体" << setw(8) << "x" << setw(13) << "适度" << setw(12) << "归一适度" << setw(13) << "累积适度" << endl;

        for (int 个 = 0; 个 < p每代个数; 个++)
        {
            cout << left << setfill(' ') << setw(20) << 印个体(某代[个]) << setw(6) << 代基度[个] << setw(10) << setprecision(3) << 代适度[个] << setw(10) << setprecision(3) << 归一适度[个] << setw(10) << setprecision(3) << 累积适度[个] << endl;
        }
#endif //  _调试
    }

    void 选种(const 代 &某代, const vector<适度> &代适度, const vector<适度> &归一适度, const vector<适度> &累积适度, 代 &种) const
    {
        int 个 = p每代个数;
        uniform_real_distribution<> 挑种(0, 1);

        // 转轮盘选种: 良种可用多次且被选中概率大
        while (--个 >= 0)
        {
            种[个] = 某代[折半寻(累积适度, 挑种(随机源))];
        }
#ifdef _调试
        cout << 印种群(种) << endl;
#endif
    }

    void 杂交(const 基因链 &父, const 基因链 &母, 基因链 &子, 基因链 &女) const
    {
        // 定交换片段
        号型 起, 止;
        uniform_int_distribution<> 定点(0, 父.size() - 1);
        bernoulli_distribution 换(p杂交概率);
        起 = 定点(随机源), 止 = 定点(随机源);
        子 = 父, 女 = 母;

        if (换(随机源))
        {
            if (起 <= 止)
            {
                swap_ranges(子.begin() + 起, 子.begin() + 止 + 1, 女.begin() + 起);
            }
            else
            {
                swap_ranges(子.begin(), 子.begin() + 止, 女.begin());
                swap_ranges(子.begin() + 起, 子.end(), 女.begin() + 起);
            }
        }
#ifdef _调试
        cout << 印换(父, 起, 止, "父") << endl;
        cout << 印换(母, 起, 止, "母") << endl;
        cout << 印换(子, 起, 止, "子") << endl;
        cout << 印换(女, 起, 止, "女") << endl;
#endif // _调试
    }

    void 群交(代 &祖, 代 &后) const
    {
        shuffle(祖.begin(), 祖.end(), 随机源); // 随机配对

        for (号型 对 = 0; 对 < p每代个数 / 2; 对++)
        {
            杂交(祖[2 * 对], 祖[2 * 对 + 1], 后[2 * 对], 后[2 * 对 + 1]);
        }
    }

    void 变异(基因链 &某) const
    {
        bernoulli_distribution 变(p变异概率);
        基因链 原 = 某;
        号型 号 = 0;

        for (号 = 0; 号 < p基因链长; 号++)
            if (变(随机源))
                某[号] = (某[号] + 1) % 2;

#ifdef _调试
        cout << 印个体(原, "原") << endl;
        cout << 印个体(某, "变") << endl;
#endif // _调试
    }

    void 群变(代 &某) const
    {
        for (号型 个 = 0; 个 < p每代个数; 个++)
        {
            变异(某[个]);
        }
    }

    号型 折半寻(const 列型 &数, const 键型 &标) const
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

    bool 停(vector<适度> &代适度, const 适度 &止适度, const int 代数, const int &止代数) const
    {
        if (代数 >= 止代数)
        {
            cout << "代数已过!" << endl;
            return true;
        }

        // 取2/8超过止适度
        sort(代适度.begin(), 代适度.end());
        if (代适度[代适度.size() - 2] >= 止适度)
        {
            cout << "第" << 代数 << "代即成";
            cout << 印代(子代, "满意的后代") << endl;
            return true;
        }

        return false;
    }

public:
    static default_random_engine 随机源;

    // _演化配置
    int p每代个数; // 假设每代个体数一样
    int p基因链长;
    适度 p止演适度;
    int p止演代数;
    double p杂交概率;
    double p变异概率;

    // 种群
    代 父代;
    代 种;
    代 配代;
    代 子代;

    // 评估
    vector<适度> 代适度, 归一适度, 累积适度;

    函数最大值()
    {
        随机源.seed(314);

        // _演化配置
        p每代个数 = 8;
        p基因链长 = 8;
        p止演适度 = 0.99;
        p止演代数 = 200;
        p杂交概率 = 0.8;
        p变异概率 = 0.001;

        // 种群
        父代 = 种 = 子代 = 配代 = 代(p每代个数, 基因链(p基因链长, 0));

        // 评估
        代适度 = 归一适度 = 累积适度 = vector<适度>(p每代个数, 0);
    };

    string 印基因(const 基因 &基) const
    {
        stringstream ss;
        ss << 基;
        return ss.str();
    }

    string 印基链(const 基因链 &链, const string &啥 = "") const
    {
        stringstream ss;
        ss << 啥 << "(";
        for (int 号 = 0; 号 < 链.size(); 号++)
        {
            ss << 印基因(链[号]) << (号 == 链.size() - 1 ? "" : ",");
        }
        ss << ")";
        return ss.str();
    }

    string 印换(const 基因链 &链, const int 起, const int 止, const string &啥 = "") const
    {
        stringstream ss;
        ss << 啥;
        if (起 <= 止)
        {
            ss << 起 << "->" << 止 << "(";
            for (号型 号 = 0; 号 < p基因链长; 号++)
            {
                if (号 == 起)
                    ss << "|";
                ss << 链[号];
                if (号 == 止)
                    ss << "||";
            }
        }
        else
        {
            ss << 0 << "->" << 止 << "|" << 起 << "->" << p基因链长 - 1 << "(";
            for (号型 号 = 0; 号 < p基因链长; 号++)
            {
                if (号 == 止)
                    ss << "||";

                if (号 == 起)
                    ss << "|";

                ss << 链[号];
            }
        }
        ss << ")";
        return ss.str();
    }

    string 印代(const 代 &某代, const string &啥 = "") const
    {
        stringstream ss;
        ss << 啥 << endl;
        ss << "---------共" << 某代.size() << "个---------" << endl;
        int 个 = 1;
        for (auto 个体 : 某代)
            ss << setw(4) << 个++ << ":" << 印个体(个体) << endl;
        return ss.str();
    }

    template <typename T>
    string 印列(const vector<T> &列) const
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

default_random_engine 函数最大值::随机源 = default_random_engine();

int main(int argc, char const *argv[])
{
    函数最大值 传1;
    传1.逐代演化();

    return 0;
}
