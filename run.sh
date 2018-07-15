func=$1

compile(){
    GOOS=linux go build -v -o ./app .
}

$func
exit 0 