#include <iostream>
#include <bitset>
using namespace std;

void 印(){}

// template <typename 参型>
// void 印(const 参型& 参)
// {
//     cout << 参 << endl;
// }


template <typename 首型, typename... 余型>
void 印(const 首型& 首, const 余型&... 余)
{
    // cout << 首 << endl;
    // cout << "余参数=" << sizeof...(余) << endl;
    印(余...);
    cout << 首 << endl;
    cout << "余参数=" << sizeof...(余) << endl;    
}

int main(){
    印(7.5, "你好!", bitset<16>(618), 32, 3.14);
}