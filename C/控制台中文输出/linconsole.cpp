#include <stdio.h>
#include <fcntl.h>
#include <wchar.h>
#include <stdlib.h>
#include <locale.h>

int main()
{
	//char16_t 名 = "你好！";

	//printf("%s\n", 名);
	const	wchar_t* a = L"千山鸟飞绝，万径人踪灭！";

	setlocale(LC_CTYPE, "");
	wprintf(L"中文字符串及索引:\n\t%ls %lc\n", a, a[3]);      // 中文输出正常

  return 0;
}

