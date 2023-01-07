hello() {
    if [ -z "$1"]
    then
        echo "Hello world"
    else
        echo "Hello $1"
    fi
}

hello
hello universe