for i in `find "/vendor" -name "*thermal*.json" -o -name "*thermal*.conf" -type f`; do
    echo "$i" | fgrep -iq 'android.' && continue
    cp -Rf  $i /data/vendor/thermal/config/
done
