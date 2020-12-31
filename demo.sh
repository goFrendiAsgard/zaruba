#!/bin/sh
set -e
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
../zaruba please setKwarg key=dockerRepo value=stalchmst

echo "=== ADD SUBREPOS"
../zaruba please addSubrepo url="https://github.com/state-alchemists/fibonacci-clock" prefix="fibo"
../zaruba please initSubrepos
../zaruba please pullSubrepos

echo "=== ADD FIBO SERVICE"
../zaruba please makeServiceTask location=fibo

echo "=== ADD DOCKER SERVICE"
../zaruba please makeDockerTask image=rabbitmq

echo "=== CREATE FASTAPI SERVICE"
../zaruba please makeFastService location=myservice

echo "=== CREATE FASTAPI MODULE"
../zaruba please makeFastModule location=myservice module=mymodule

echo "=== CREATE FASTAPI ROUTE, EVENT, AND RPC HANDLER"
../zaruba please makeFastRoute location=myservice module=mymodule url=/hello
../zaruba please makeFastEventHandler location=myservice module=mymodule event=myEvent
../zaruba please makeFastRPCHandler location=myservice module=mymodule event=myRPC

echo "=== CREATE FASTAPI CRUD"
../zaruba please makeFastCRUD location=myservice module=mymodule entity=book field=title,author,synopsis

echo "=== RUN AND AUTOSTOP"
../zaruba please run autostop

echo "=== SHOW FASTAPI LOG"
../zaruba please showLog task=myservice

echo "=== CLEAR LOG"
../zaruba please clearLog

echo "=== BUILD DOCKER IMAGE"
git add . -A && git commit -m 'first commit'
../zaruba please buildImage

echo "=== BUILD DOCKER IMAGE (TAGGED)"
git tag -a v0.0.0 -m 'version 0.0.0'
../zaruba please buildImage

echo "=== PUSH IMAGE"
../zaruba please pushImage

echo "=== EXPLAIN START"
../zaruba please explain start

echo "=== DONE!!!"
echo "You can re-run the services by moving to playground directory and run: zaruba please run"
