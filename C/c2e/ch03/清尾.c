#include <stdio.h>
#include <string.h>

// 清串尾空白符
int 清尾(char 串[])
{
    int 号;
    for(号=strlen(串)-1; 号 >= 0; 号--)
        if(串[号] != ' ' && 串[号] != '\t' && 串[号] != '\n')
            break;
    串[号+1] = '\0';
    return 号;
}

int main(int argc, char const *argv[])
{
    char 串[] = "此串尾有2空格1跳格     ";
    printf("清尾前: \"%s\"\n", 串);
    清尾(串);
    printf("清尾后: \"%s\"\n", 串);
    return 0;
}
