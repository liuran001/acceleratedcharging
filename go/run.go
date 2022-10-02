package main

import (
    "os/exec"
    "os"
    "fmt"
    "time"
    "strings"
    "strconv"
)

var (
    temperature string
    temperaturewall string
    speed string
    timesleep string
    times int
    ion int
    uon int
)

func shell(command string) string{ //调用shell执行命令(root权限)
    var output []byte
    var err error
    var cmd *exec.Cmd
    cmd = exec.Command("su", "-c", command)
    if output, err = cmd.CombinedOutput(); err == nil {
    }
    return string(output)
}

func uninstalltemperaturecontrol() { //删除温控
    temperature = string(shell("cat /sys/class/power_supply/battery/temp")) //读取电池温度
    shell("chmod 771 -R /data/vendor/thermal/config") //重新给予权限
    //开始写入云端温控配置
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
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-tgame.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-video.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-videochat.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-yuanshen.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-india-phone.conf")
    shell("echo $( base64 -d <<< \"DPrretij42fIjnrsdemBUwo=\" )>/data/vendor/thermal/config/thermal-chg-only.conf")
    shell("chmod 644 /data/vendor/thermal/config/*") //给予权限
    if temperature > temperaturewall {
    shell("umount /sys/class/power_supply/battery/constant_charge_current") //解除挂载充电电流速度挂载
    shell(". /data/adb/modules/acceleratedcharging/restore.sh") //恢复温控
    } else {
        shell("echo '" + speed + "' > /data/adb/modules/acceleratedcharging/constant_charge_current") //写入充电电流到模块缓存文件
        shell("umount /sys/class/power_supply/battery/constant_charge_current")
        shell("mount -o rw /data/adb/modules/acceleratedcharging/constant_charge_current /sys/class/power_supply/battery/constant_charge_current")
    }
    //通过mount命令挂载充电电流速度
    shell("kill -9 $(pidof mi_thermald)") //重启云控
}

func installtemperaturecontrol() {
    shell(". /data/adb/modules/acceleratedcharging/restore.sh")
    shell("umount /sys/class/power_supply/battery/constant_charge_current")
    shell("kill -9 $(pidof mi_thermald)")
}

func sleeps(times int) { //硬核休眠
    sum := 1
    for ; sum <= times; {
        sum = sum + 1
        time.Sleep(time.Second)
    }
}
func main() {
    //读取命令行参数
    speed = os.Args[1]
    temperaturewall = os.Args[2]
    timesleep =os.Args[3]
    timesleep, err := strconv.Atoi(timesleep) //将string类型转为int类型
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    //初始化变量
    uon = 0 
    ion = 0
    for true { //循环
            var batterystatus = shell("dumpsys battery")
            var dl = strings.Contains(batterystatus,"status: 2")
            if dl { //判断是否在充电
                if uon == 0 {
                    uninstalltemperaturecontrol() //删除温控 修改充电速度
                    uon = 1
                    ion = 0
                }
            } else {
                if ion == 0 {
                    installtemperaturecontrol()//恢复
                    ion = 1
                    uon = 0
                }
            }
        sleeps(timesleep) //休眠
    }
}
