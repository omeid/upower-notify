# UPower-Notify


A simple tool to give you Desktop Notifications about your battery, requires UPower.


# Usage


```sh
$ upower-notify --help

Usage of upower-notify:
  -critical=10m0s: Time to start warning. (Critical)
  -report=false: Print out updates to stdout.
  -tick=10s: Update rate
  -warn=20m0s: Time to start warning. (Warn)
```

if you're using `i3wm` or relatives, just chuck this or it's equalent in your startup script:


`exec --no-startup-id "upower-notify"`
