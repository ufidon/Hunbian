# 中英混编
有容乃大, 中英混合编程。

## 1. llvm开始

llvm是开发编译器的工具，支持多种源码语言和目标指令集。其核心是一套处理中间表示（IR）的工具、库和头文件。工具套件含：
* 汇编器和反汇编
* 点码分析器及优化器

类C语言(C、C++、Objective C 及Objective C++)采用clang编译成llvm点码(IR)，然后用llvm编译成目标文件。按以下流程获取并编译llvm源码。

```bash
# 1. 获取源码
git clone https://github.com/llvm/llvm-project.git
# 在Windows上禁止行尾符自动转换
git clone --config core.autocrlf=false https://github.com/llvm/llvm-project.git

# 2. 配置并编译llvm和clang
cd llvm-project && mkdir build && cd build

# 运行脚本/zao.sh进行配置
ninja # 开始编译
ninja check-clang # 测试
ninja install # 安装

```

## 2. 使用llvm
```bash
# 1. 编译
clang 源文件.c -S -c // 生成汇编文件及目标文件

clang 源文件.c -o 可执行程序名 // 生成可执行文件

# 2. 生成llvm自身汇编及位码
clang -O3 -emit-llvm 源文件.c -c -S -o 源文件.bc

# 3. 跑程序
./可执行程序名
# 或
lli 源文件.bc

# 4. 反汇编
llvm-dis < 源文件.bc

# 5. 把llvm汇编编译成目标汇编
llc 源文件.bc -o 源文件.s

# 6. 把目标汇编编译成可执行程序
gcc 源文件.s -o 可执行程序
```

## 3. 增加中文关键词
对于语言关键词，在 llvm-project/clang/include/clang/Basic/TokenKinds.def中增加中文别名即可。

但想增加预处理关键词，还需更新：
llvm-project/clang/lib/Basic/IdentifierTable.cpp

## 参考
* [llvm 开始](https://llvm.org/docs/GettingStarted.html)
* [LLVM MinGW](https://github.com/mstorsjo/llvm-mingw)


