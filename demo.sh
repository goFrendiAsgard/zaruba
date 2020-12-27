#!/bin/sh
go build
ZARUBA_HOME=$(realpath $(pwd))
sudo rm -Rf playground
mkdir -p playground
cd playground

echo "=== SORRY"
../zaruba sorry

echo "=== THANKS"
../zaruba thanks

echo "=== SETUP UBUNTU"
sudo -E ../zaruba please setupUbuntu

echo "=== INIT PROJECT"
../zaruba please initProject

echo "=== ADD SUBREPOS"
../zaruba please addSubrepo url="https://github.com/therealvasanth/fibonacci-clock" prefix="fibo"
../zaruba please initSubrepos
../zaruba please pullSubrepos

echo "=== ADD FIBO SERVICE"
../zaruba please makeServiceTask location=fibo

echo "=== ADD DOCKER SERVICE"
../zaruba please makeDockerTask image=rabbitmq

echo "=== CREATE FASTAPI SERVICE"
../zaruba please makeFastService location=oreno

echo "=== RUN AND AUTOSTOP"
../zaruba please run autostop

echo "=== SHOW FASTAPI LOG"
../zaruba please showLog task=oreno

echo "=== CLEAR LOG"
../zaruba please clearLog

echo "=== EXPLAIN START"
../zaruba please explain start

echo "=== DONE!!!"
