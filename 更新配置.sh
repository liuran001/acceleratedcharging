#!/system/bin/sh

. ${0%/*}/config.ini

kill -9 $(pidof charge-current)
echo "等待3秒后手动退出"
nohup ${0%/*}/go/charge-current $speed $temperaturewall $timesleep &

