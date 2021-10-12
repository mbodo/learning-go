#!/bin/bash

# Run from vscode terminal

if [[ -z "${REMOTE_CONTAINERS}" ]]; then
    printf "Please run it from vscode terminal\n"
    exit 0
fi

CHAPTER=$1

mkdir -p ../$CHAPTER
cp -R chapter.go go.mod Makefile ../$CHAPTER

pushd ../$CHAPTER
sed -i "s/ch0X/$CHAPTER/g" go.mod
git init
popd

pushd ../
code -a $CHAPTER
code $CHAPTER/chapter.go $CHAPTER/go.mod $CHAPTER/Makefile
popd
