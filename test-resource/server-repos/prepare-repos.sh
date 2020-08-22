#!/bin/sh
set -e

for INDEX in init push pull
do
    rm  -Rf parent-${INDEX} parent-${INDEX}.git
    mkdir -p parent-${INDEX}
    cd parent-${INDEX}
    git init --shared=true
    cd ..
    git clone --bare parent-${INDEX} parent-${INDEX}.git

    rm  -Rf child-${INDEX} child-${INDEX}.git
    mkdir -p child-${INDEX}
    cd child-${INDEX}
    git init --shared=true
    cd ..
    git clone --bare child-${INDEX} child-${INDEX}.git
done