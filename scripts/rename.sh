#!/usr/bin/env bash

read -p "Enter your project name: " name
name=${name}

find . \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i "s/goecho-projectbase-api/${name}/g"
