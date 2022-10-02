#!/system/bin/sh
SKIPUNZIP=0
rm -rf /data/thermal/
ui_print "检测温控是否存在...."
for i in `find "/vendor" -name "*thermal*.json" -o -name "*thermal*.conf" -type f`; do
    ui_print "$i" | fgrep -iq 'android.' && continue
    if [ ! -f $i ];then
        ui_print "温控:$i不存在！"
        abort 1
    fi
done
for o in `find "/vendor" -name "*thermal*" ! -name "*thermal*.json" ! -name "*thermal*.conf" -type f`; do
    ui_print "$o" | fgrep -iq 'android.' && continue
    if [ ! -f $o ];then
        ui_print "温控:$k不存在！"
        abort 1
    fi
done
for a in `find "/vendor/etc/perf" -name "*.xml" -type f`; do
    ui_print "$a" | fgrep -iq 'android.' && continue
    if [ ! -f $a ];then
        ui_print "温控:$a不存在！"
        abort 1
    fi
done
for l in `find "/vendor/etc/perf" -name "perf*" -type f`; do
    ui_print "$l" | fgrep -iq 'android.' && continue
    if [ ! -f $l ];then
        ui_print "温控:$l不存在！"
        abort 1
    fi
done
for b in `find "/vendor/etc/perf" -name "*conf" -type f`; do
    ui_print "$b" | fgrep -iq 'android.' && continue
    if [ ! -f $b ];then
        ui_print "温控:$b不存在！"
        abort 1
    fi
done
ui_print "温控存在，正在进行下一步安装..."
set_perm_recursive $MODPATH 0 0 0777 0777
set_perm $MODPATH/go/charge-current 0 0 0777 0777
set_perm $MODPATH/restore.sh 0 0 0777 0777
