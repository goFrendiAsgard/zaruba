#!/bin/sh

echo 💀 Cloning Zaruba 
git clone -–depth 1 https://github.com/state-alchemists/zaruba "${HOME}/.zaruba"

echo 💀 Build Zaruba 
echo "Build Zaruba"
cd "${HOME}/.zaruba"
go build

echo 💀 Injecting '${HOME}/.zaruba' to PATH
PATH=$PATH:"${HOME}/.zaruba"
echo 'PATH=$PATH:"${HOME}/.zaruba"' >> "${HOME}/.bashrc"
echo 'PATH=$PATH:"${HOME}/.zaruba"' >> "${HOME}/.zshrc"
