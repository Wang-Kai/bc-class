#!/bin/sh

# start redis
redis-server &

# start app
/app
