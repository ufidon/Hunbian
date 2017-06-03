#include <stdio.h>
#include <io.h>
#include <fcntl.h>
#include <wchar.h>

int main()
{
	//char16_t 名 = "你好！";

	//printf("%s\n", 名);
	wchar_t* a = L"千山鸟飞绝，万径人踪灭！";
	fflush(stdout); //_setmode 前必须清输出缓存
	_setmode(_fileno(stdout), _O_U16TEXT); // 将控制台模式设置为UTF-16
	wprintf(L"中文字符串及索引:\n\t%s %c\n", a, a[3]);      // 正常输出中文

	wchar_t 名字[16];
	fflush(stdin);
	_setmode(_fileno(stdin), _O_U16TEXT);
	wscanf(L"%s", &名字);
	wprintf(L"%s, 欢迎你!\n", 名字);

	return 0;
}

