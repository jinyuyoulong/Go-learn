#!/usr/bin/env bash
# 预处理(也叫预编译)
#gcc -E  hello.c  -o hello.i
## 编译
#gcc  -S  hello.i   -o  hello.s
#
## 汇编
#gcc  -c  hello.s  -o  hello.o #或者 as  hello.s -o  hello.o

#gcc -c -fPIC -o hi.o hi.c
#gcc -shared -o libhi.so hi.o

gcc -c -fPIC -o hello.o hello.c
gcc -shared -o libhello.so hello.o