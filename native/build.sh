#!/usr/bin/env sh

rm -f ../native-artifacts/*.h
rm -f ../native-artifacts/*.so
rm -f ../native-artifacts/*.dll






# Linux -- amd64
echo "Linux -- amd64"

rm -f callback.o
gcc -c callback.c -o callback.o

env \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=1 \
    go build \
        -o ../native-artifacts/node-libuiohook-linux-amd64.so \
        -buildmode=c-shared \
        main.go






# Linux -- i386
echo "Linux -- i386"

rm -f callback.o

gcc -m32 -c callback.c -o callback.o

env \
    GOOS=linux \
    GOARCH=386 \
    CGO_ENABLED=1 \
    CC="gcc -m32" \
    go build \
        -o ../native-artifacts/node-libuiohook-linux-i386.so \
        -buildmode=c-shared \
        main.go





# Windows -- amd64
echo "Windows -- amd64"

rm -f callback.o
rm -f libcallback.a

x86_64-w64-mingw32-gcc -c callback.c -o callback.o
x86_64-w64-mingw32-ar rcs libcallback.a callback.o

env \
    GOOS=windows \
    GOARCH=amd64 \
    CGO_ENABLED=1 \
    CC=x86_64-w64-mingw32-gcc \
    \
    go build \
        -o ../native-artifacts/node-libuiohook-windows-amd64.dll \
        -buildmode=c-shared \
        main.go





# Windows -- i386
echo "Windows -- i386"

rm -f callback.o
rm -f libcallback.a

i686-w64-mingw32-gcc -c callback.c -o callback.o
i686-w64-mingw32-ar rcs libcallback.a callback.o

env \
    GOOS=windows \
    GOARCH=386 \
    CGO_ENABLED=1 \
    CC=i686-w64-mingw32-gcc \
    go build \
        -o ../native-artifacts/node-libuiohook-windows-i386.dll \
        -buildmode=c-shared \
        main.go





# CLEANUP   
rm -rf ./callback.o
rm -rf ./libcallback.a
