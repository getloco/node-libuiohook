#!/usr/bin/env sh

# Windows Cross Compilation
apt install mingw-w64 -y

# i386 Cross Compilation
apt install gcc-multilib -y

# Linux x11 Dependencies for i386
apt install \
    libx11-dev:i386 libxtst-dev:i386 libxcb1-dev:i386 \
    libxkbcommon-dev:i386 libxkbcommon-x11-dev:i386 \
    libx11-xcb-dev:i386 \
    -y

