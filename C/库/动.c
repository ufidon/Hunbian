#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>

typedef double 双型;
typedef int 整型;
typedef char 符型;
typedef void 空;

#define 若 if
#define 返 return
#define 结构 struct
#define 定型 typedef

整型 双乘(整型 甲, 整型 乙)
{
	return 甲*乙;
}

定型 结构 人{
	符型 名[20];
	整型 身高;
}构人;

#define 主函数 main

整型  主函数(整型  参目, 符型 **参列)
{
	空 *柄;
	双型 (*函数)(双型, 双型);

	符型 *错误;
	整型  甲=12, 乙=21, 丙;


	构人 小明={"小明", 176};
	printf("%s身高是%d\n", 小明.名, 小明.身高);
	柄 = dlopen("./lib均.so", RTLD_LAZY);
	若(!柄){
		fputs(dlerror(), stderr);
		exit(1);
	}
	函数=dlsym(柄, "均值");
	若((错误=dlerror()) != NULL)
	{
		fputs(错误, stderr);
		exit(1);
	}
	puts("这是一个共享库测试...\n");
	丙= 函数(甲,乙);
	printf("(%d+%d)/2=%d\n", 甲,乙,丙);
	dlclose(柄);

	丙 = 双乘(甲,乙);
	printf("%d乘%d之积为%d\n",甲,乙,丙);

	返 0;
}
