REM !/usr/bin/env bash

REM setup MSVC environment
REM "C:\Program Files (x86)\Microsoft Visual Studio\2019\Community\VC\Auxiliary\Build\vcvarsall.bat" x64

REM Ref: https://llvm.org/docs/GettingStarted.html
REM Ref: https://stackoverflow.com/questions/15036909/clang-how-to-list-supported-target-architectures

REM projects: clang, clang-tools-extra, libcxx, libcxxabi, libunwind, lldb, compiler-rt, lld, polly, or debuginfo-tests
REM ref: https://llvm.org/OpenProjects.html
set projects=clang;libcxx;libcxxabi;lld

REM Specify for directory the full pathname of where you want the LLVM tools and libraries to be installed (default /usr/local).
set dir=d:\llvm

REM Build type: Debug, Release, RelWithDebInfo, and MinSizeRel. Default is Debug
set btype=Release

REM Compilers
set cc=c:\devtools\llvm\bin\clang.exe
set cxx=c:\devtools\llvm\bin\clang++.exe

REM Targets: AArch64, AMDGPU, ARM, AVR, BPF, Hexagon, Lanai, Mips, MSP430, NVPTX, PowerPC, RISCV, Sparc, SystemZ, WebAssembly, X86, XCore
set targets=AArch64;X86;RISCV;ARM

REM Ninja ref: https://ninja-build.org/manual.html
cmake -G Ninja -DCMAKE_C_COMPILER=%cc% -DCMAKE_CXX_COMPILER=%cxx% -DLLVM_TARGETS_TO_BUILD=%targets% -DLLVM_ENABLE_PROJECTS=%projects% -DCMAKE_INSTALL_PREFIX=%dir% -DCMAKE_BUILD_TYPE=%btype% ../llvm
