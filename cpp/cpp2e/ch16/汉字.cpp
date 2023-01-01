#include <string>
#include <iostream>
#include <codecvt>

using namespace std;

// 参考： http://zhiyi.live/2020/02/16/std-wcout输出乱码

int main(int argc, char const *argv[])
{
    std::ios::sync_with_stdio(false);
    std::wcout.imbue(std::locale(""));
    std::cout << "你好" << std::endl;
    std::wstring wstr = L"你能输出中文？";
    std::wcout << wstr << std::endl;

    cout << "各输出各自的，这来自cout。" << endl;
    wcout << L"各输出各自的，这来自wcout。" << endl;

    wstring 字;

    std::wcin.imbue(std::locale(""));
    std::wcerr.imbue(std::locale(""));
    while (wcin >> 字)
    {
        wcout << 字 << wcslen(字.data());
    }

    return 0;
}
