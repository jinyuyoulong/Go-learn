#!/usr/bin/env bash
# 编辑 .so 动态链接库
gcc hello.c -fPIC -shared -o libhello.so