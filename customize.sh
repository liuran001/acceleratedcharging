#!/system/bin/sh
SKIPUNZIP=0
set_perm_recursive $MODPATH 0 0 0777 0777
set_perm $MODPATH/go/main 0 0 0777 0777