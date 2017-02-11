#!/bin/bash

# WORK_DIR="/Users/baidu/GoglandProjects/filemerger/file"
WORK_DIR="./file"

cd ${WORK_DIR}

for ((i=0;i<=9;i++))
do
    echo "Create $i file.."

    if ((i%2==0));then
        echo "hello" > $i.file
    else
        echo "world" > $i.file
    fi

    echo "Done"
done