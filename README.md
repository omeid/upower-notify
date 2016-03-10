# UPower-Notify


A simple tool to give you Desktop Notifications about your battery, requires UPower.


# Usage


```sh
$ upower-notify --help

Usage of upower-notify:
  -critical duration
    	Time to start warning. (Critical) (default 10m0s)
  -device string
    	DBus address of the battery device (default "/org/freedesktop/UPower/devices/battery_BAT0")
  -report
    	Print out updates to stdout.
  -tick duration
    	Update rate (default 10s)
  -warn duration
    	Time to start warning. (Warn) (default 20m0s)
```

if you're using `i3wm` or relatives, just chuck this or it's equalent in your startup script:


`exec --no-startup-id "upower-notify"`
