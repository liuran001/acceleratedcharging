package main

import (
    "os/exec"
    "fmt"
    "time"
    "strings"
)

  
func shell(command string) string{
    var output []byte
    var err error
    var cmd *exec.Cmd
    cmd = exec.Command("su", "-c", command)
    if output, err = cmd.CombinedOutput(); err == nil {
    }
    return string(output)
}

func uninstalltemperaturecontrol() {
    temperature := string(shell("cat /sys/class/power_supply/battery/temp"))
    temperaturewall := string(shell(". /data/adb/modules/acceleratedcharging/config.ini && echo ${temperaturewall}"))
    great := string(shell(". /data/adb/modules/acceleratedcharging/config.ini && echo ${great}"))
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-4k.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-arvr.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-camera.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-class0.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-hp-mgame.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-hp-normal.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-huanji.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-4k.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-camera.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-class0.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-huanji.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-mgame.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-navigation.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-nolimits.conf")
    shell("echo $( base64 -d <<< \"boFgrg4gvtcCzEOsftjlcg5OxGAYz62gCBthbPvh5Rh0IhF/vqRPtdlO9qRg1K+BNiEGrwiKJvqU Lmgww887XR6JnESBTI3DfKIvwR9WTn1CFp/vZm+F+hq+3PnKYdpAPhqXRbqrO0EuE8K6LShxrQ41 HxaOqFPHBWO47TL6eLygyUT0iQeZE3Nz+WrkZzYyEBjjmdmjitMxKkc9z/C9ToVIVviNkxatcK75 X9lfavA1fjUejqd87GCPlVCwSBmn+E1cfWyOn0nrPMc701jnFUGeEP6mTzClxiiQuAIi3vJSGLQ0 rSgGvNHFK7FPwuKnNCJveas3FVQJNuy5TG2kKB62dBOL/79GkAn8adbKZmWHXbbUcXObVsX72jjw LSb0lgYlkXmDKyUyFYircI7oNQMUfhVQYZjMs4/6rPjvIqcdV1Teoy3h+A81bcrt4rM40OYw7m6s CCIpmSeW06QL6GBs8bKuDzS24qmr0QSfwkkLykhUOQ5CKermH6CtHAFo3njcaP5iUVkcpJUxHmZI K/+GoOrFnOkfe4aUhg9geZz2tEW2WmBkm68ufxFoR1U2pkVQNasCQG1VFJTqkyHvH0vbVJjsp2XO 6vhw7ik54COpXBoI2y9TDLFUyXft1qpm3sDSt2OReRY/DEKSSN6xUy6tIYS7gmBjqlkwbVVs/go=\" )>/data/vendor/thermal/config/thermal-india-normal.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-per-camera.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-per-class0.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-per-navigation.conf")
    shell("echo $( base64 -d <<< \"boFgrg4gvtcCzEOsftjlcg5OxGAYz62gCBthbPvh5Rh0IhF/vqRPtdlO9qRg1K+BNiEGrwiKJvqU Lmgww887XR6JnESBTI3DfKIvwR9WTn1CFp/vZm+F+hq+3PnKYdpAPhqXRbqrO0EuE8K6LShxrQ41 HxaOqFPHBWO47TL6eLygyUT0iQeZE3Nz+WrkZzYyEBjjmdmjitMxKkc9z/C9ToVIVviNkxatcK75 X9lfavA1fjUejqd87GCPlVCwSBmn+E1cfWyOn0nrPMc701jnFUGeEP6mTzClxiiQuAIi3vJSGLQ0 rSgGvNHFK7FPwuKnNCJveas3FVQJNuy5TG2kKB62dBOL/79GkAn8adbKZmWHXbbUcXObVsX72jjw LSb0lgYlkXmDKyUyFYircI7oNQMUfhVQYZjMs4/6rPjvIqcdV1Teoy3h+A81bcrt4rM40OYw7m6s CCIpmSeW06QL6GBs8bKuDzS24qmr0QSfwkkLykhUOQ5CKermH6CtHAFo3njcaP5iUVkcpJUxHmZI K/+GoOrFnOkfe4aUhg9geZz2tEW2WmBkm68ufxFoR1U2pkVQNasCQG1VFJTqkyHvH0vbVJjsp2XO 6vhw7ik54COpXBoI2y9TDLFUyXft1qpm3sDSt2OReRY/DEKSSN6xUy6tIYS7gmBjqlkwbVVs/go=\" )>/data/vendor/thermal/config/thermal-india-per-normal.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-per-video.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-phon.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-tgame.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-video.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-videochat.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-yuanshen.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-mgame.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-navigation.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-nolimits.conf")
    shell("echo $( base64 -d <<< \"boFgrg4gvtcCzEOsftjlcg5OxGAYz62gCBthbPvh5Rh0IhF/vqRPtdlO9qRg1K+BNiEGrwiKJvqU Lmgww887XR6JnESBTI3DfKIvwR9WTn1CFp/vZm+F+hq+3PnKYdpAPhqXRbqrO0EuE8K6LShxrQ41 HxaOqFPHBWO47TL6eLygyUT0iQeZE3Nz+WrkZzYyEBjjmdmjitMxKkc9z/C9ToVIVviNkxatcK75 X9lfavA1fjUejqd87GCPlVCwSBmn+E1cfWyOn0nrPMc701jnFUGeEP6mTzClxiiQuAIi3vJSGLQ0 rSgGvNHFK7FPwuKnNCJveas3FVQJNuy5TG2kKB62dBOL/79GkAn8adbKZmWHXbbUcXObVsX72jjw LSb0lgYlkXmDKyUyFYircI7oNQMUfhVQYZjMs4/6rPjvIqcdV1Teoy3h+A81bcrt4rM40OYw7m6s CCIpmSeW06QL6GBs8bKuDzS24qmr0QSfwkkLykhUOQ5CKermH6CtHAFo3njcaP5iUVkcpJUxHmZI K/+GoOrFnOkfe4aUhg9geZz2tEW2WmBkm68ufxFoR1U2pkVQNasCQG1VFJTqkyHvH0vbVJjsp2XO 6vhw7ik54COpXBoI2y9TDLFUyXft1qpm3sDSt2OReRY/DEKSSN6xUy6tIYS7gmBjqlkwbVVs/go=\" )>/data/vendor/thermal/config/thermal-normal.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-per-camera.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-per-class0.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-per-navigation.conf")
    shell("echo $( base64 -d <<< \"boFgrg4gvtcCzEOsftjlcg5OxGAYz62gCBthbPvh5Rh0IhF/vqRPtdlO9qRg1K+BNiEGrwiKJvqU Lmgww887XR6JnESBTI3DfKIvwR9WTn1CFp/vZm+F+hq+3PnKYdpAPhqXRbqrO0EuE8K6LShxrQ41 HxaOqFPHBWO47TL6eLygyUT0iQeZE3Nz+WrkZzYyEBjjmdmjitMxKkc9z/C9ToVIVviNkxatcK75 X9lfavA1fjUejqd87GCPlVCwSBmn+E1cfWyOn0nrPMc701jnFUGeEP6mTzClxiiQuAIi3vJSGLQ0 rSgGvNHFK7FPwuKnNCJveas3FVQJNuy5TG2kKB62dBOL/79GkAn8adbKZmWHXbbUcXObVsX72jjw LSb0lgYlkXmDKyUyFYircI7oNQMUfhVQYZjMs4/6rPjvIqcdV1Teoy3h+A81bcrt4rM40OYw7m6s CCIpmSeW06QL6GBs8bKuDzS24qmr0QSfwkkLykhUOQ5CKermH6CtHAFo3njcaP5iUVkcpJUxHmZI K/+GoOrFnOkfe4aUhg9geZz2tEW2WmBkm68ufxFoR1U2pkVQNasCQG1VFJTqkyHvH0vbVJjsp2XO 6vhw7ik54COpXBoI2y9TDLFUyXft1qpm3sDSt2OReRY/DEKSSN6xUy6tIYS7gmBjqlkwbVVs/go=\" )>/data/vendor/thermal/config/thermal-per-normal.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-per-video.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-phone.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-phon.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-tgame.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-video.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-videochat.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-yuanshen.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-phone.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-chg-only.conf")
    if temperature > temperaturewall {
    shell("umount /sys/class/power_supply/battery/constant_charge_current")
    shell(". /data/adb/modules/acceleratedcharging/restore.sh")
    fmt.Println("达到温度墙"+ temperaturewall)
    } else {
        shell("echo '" + great + "' > /data/adb/modules/acceleratedcharging/constant_charge_current")
    }
    shell("umount /sys/class/power_supply/battery/constant_charge_current")
    shell("mount -o rw /data/adb/modules/acceleratedcharging/constant_charge_current /sys/class/power_supply/battery/constant_charge_current")
}

func installtemperaturecontrol() {
    shell(". /data/adb/modules/acceleratedcharging/restore.sh")
    shell("umount /sys/class/power_supply/battery/constant_charge_current")
}



func main() {
    uon := 0
    ion := 0
    for true {
            var batterystatus = shell("dumpsys battery")
            var dl = strings.Contains(batterystatus,"status: 2")
            if dl {
                if uon == 0 {
                    uninstalltemperaturecontrol()
                    fmt.Println("已开启满速充电")
                    uon = 1
                    ion = 0
                }
            } else {
                if ion == 0 {
                    installtemperaturecontrol()
                    fmt.Println("已恢复充电温控")
                    ion = 1
                    uon = 0
                }
            }
        time.Sleep(time.Duration(60)*time.Second)
    }
}
