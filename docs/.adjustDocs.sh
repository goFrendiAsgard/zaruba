if [ -f "${HOME}/tocer/tocer.py" ]
then 
    echo "Run tocer"
    python "${HOME}/tocer/tocer.py" README.md 'source .adjustDocs.bashrc'
fi