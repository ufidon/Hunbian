#!/usr/bin/env bash

# Ref: https://llvm.org/docs/GettingStarted.html
# https://stackoverflow.com/questions/15036909/clang-how-to-list-supported-target-architectures

## projects: clang, clang-tools-extra, libcxx, libcxxabi, libunwind, lldb, compiler-rt, lld, polly, or debuginfo-tests
# ref: https://llvm.org/OpenProjects.html
projects="clang;libcxx;libcxxabi;lld"

## Specify for directory the full pathname of where you want the LLVM tools and libraries to be installed (default /usr/local).
dir="/opt/llvm"

## Build type: Debug, Release, RelWithDebInfo, and MinSizeRel. Default is Debug
btype=Release

## Compilers
cc=/bin/clang
cxx=/bin/clang++

## Targets: AArch64, AMDGPU, ARM, AVR, BPF, Hexagon, Lanai, Mips, MSP430, NVPTX, PowerPC, RISCV, Sparc, SystemZ, WebAssembly, X86, XCore
targets="AArch64;X86;RISCV;ARM"

# Ninja ref: https://ninja-build.org/manual.html
cmake -G Ninja -DCMAKE_C_COMPILER=$cc -DCMAKE_CXX_COMPILER=$cxx -DLLVM_TARGETS_TO_BUILD=$targets -DLLVM_ENABLE_PROJECTS=$projects -DCMAKE_INSTALL_PREFIX=$dir -DCMAKE_BUILD_TYPE=$btype ../llvm
