#include <iostream>
#include <vector>
#include <set>
#include <map>
#include <algorithm>
#include <functional>
#include <string>
#include <complex>
#include <cstddef>
using namespace std;

int main()
{
    map<map<string, double>, double> 国;
    decltype(国)::value_type 中国总产值;
    decltype(国)::key_type 中国;
    cout << typeid(decltype(国)::value_type).name() << endl;
    cout << typeid(decltype(国)::key_type).name() << endl;

    set<string, vector<string>> 战士;
    //decltype(战士)::value_type 装备 {"板斧", "腰刀", "金盔"};
    decltype(战士)::key_type 名 = "程咬金";
}