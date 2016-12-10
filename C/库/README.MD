# 带中文名字的函数库

## 参考
* http://www.cprogramming.com/tutorial/shared-libraries-linux-gcc.html
* http://tldp.org/HOWTO/Program-Library-HOWTO/dl-libraries.html


## 编译
```bash
gcc -c -Wall -Werror -fpic mean.c
gcc -shared -o lib均.so mean.o
gcc -o 测试 动.c -ldl -L. -l均
```
