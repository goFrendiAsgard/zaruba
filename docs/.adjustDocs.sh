if [ ! -f "${HOME}/tocer/tocer.py" ]
then
    echo "Clone tocer"
    git clone git@github.com:state-alchemists/tocer.git ${HOME}/tocer
fi

echo "Update tocer"
_PWD=$(pwd)
cd "${HOME}/tocer"
git pull origin main
cd "${_PWD}"

echo "Remove examples/playground"
rm -Rf examples/playground

echo "Run tocer"
python "${HOME}/tocer/tocer.py" README.md 'source .adjustDocs.bashrc'