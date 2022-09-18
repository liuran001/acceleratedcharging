for i in `find "/vendor" -name "*thermal*.json" -o -name "*thermal*.conf" -type f`; do
    echo "$i" | fgrep -iq 'android.' && continue
    echo "$i" | fgrep -iq 'thermald' && continue
    echo "$i" | fgrep -iq 'perf' && continue
    # cp -Rf $i /data/vendor/thermal/config/
    echo $i
done
