#!/bin/sh

kill -9 $(pgrep webserver)
cd ~/newweb/
git pull https://github.com/pychango/newgo.git
cd webserver/
./webserver &