#include <stdio.h>

// 从串中删除某字
void 删字(char 串[], int 字)
{
    int 步, 余步;
    for (步 = 余步 = 0; 串[步] != '\0'; 步++)
        if (串[步] != 字)
            串[余步++] = 串[步];
    串[余步] = '\0';
}

// 将自串于至之尾
void 串串(char 至[], char 自[]) // 至需够长
{
    int 至步, 自步;
    至步 = 自步 = 0;
    while (至[至步] != '\0') // 步至尾
        至步++;

    while ((至[至步++] = 自[自步++]) != '\0') // 串自于至尾
        ;
}

int main(int argc, char const *argv[])
{
    char 登鹳雀楼[100] = "dēng guàn què lóu"; // 100足够串其下两串
    char 登鹳雀楼1[] = "bái rì yī shān jǐn, huáng hé rù hǎi liú。";
    char 登鹳雀楼2[] = "yù qióng qiān lǐ mù, gèng shàng yī céng lóu。";

    printf("\"%s\"删除h", 登鹳雀楼1);
    删字(登鹳雀楼1, 'h');
    printf("得\"%s\"\n", 登鹳雀楼1);

    串串(登鹳雀楼, 登鹳雀楼1);
    串串(登鹳雀楼, 登鹳雀楼2);

    printf("%s\n", 登鹳雀楼);

    return 0;
}

/*练习
1. 改进上程序, 给诗歌标题及句子分行.
4. 写函数 void 删字(char 原[], char 集[]), 从原串中去除所有集串所含字母.
5. 写函数 int 找(char 串[], char 集[]), 从左至右在串中找首先出现集中任一字母的位置并返此位, 若皆无则返-1.
*/