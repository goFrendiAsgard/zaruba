#!/bin/sh

# USAGE
# /bin/sh get_main_python_script.sh

if [ -f __main__.py ]
then
    echo "."
elif [ -f start.py ]
    echo "start.py"
elif [ -f index.py ]
    echo "index.py"
elif [ -f main.py ]
    echo "main.py"
else
    echo $(ls *.py | head -1)
fi
