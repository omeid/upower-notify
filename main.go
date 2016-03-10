package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/omeid/upower-notify/notify"
	"github.com/omeid/upower-notify/upower"
)

func init() {
	//See the notes in "github.com/omeid/upower-notify/upower"
	//This setting is for HP Envy series late 2012.
}

var (
	tick     time.Duration
	warn     time.Duration
	critical time.Duration
	device   string
	report   bool
)

func main() {

	flag.DurationVar(&tick,     "tick",     10*time.Second,  "Update rate")
	flag.DurationVar(&warn,     "warn",     20*time.Minute,  "Time to start warning. (Warn)")
	flag.DurationVar(&critical, "critical", 10*time.Minute,  "Time to start warning. (Critical)")
	flag.StringVar(  &device,   "device",   "DisplayDevice", "DBus device name for the battery")
	flag.BoolVar(    &report,   "report",   false,           "Print out updates to stdout.")

	flag.Parse()

	up, err := upower.New(device)

	if err != nil {
		log.Fatal(err)
	}

	update, err := up.Get()
	if err != nil {
		log.Fatal(err)
	}

	notifier, err := notify.New("Upower Agent")
	if err != nil {
		log.Fatal(err)
	}

	err = notifier.Low("Ol Correct.", "Everything seems okay, I will keep you posted.", 100)
	if err != nil {
		log.Fatal(err)
	}

	var old upower.Update

	for _ = range time.Tick(tick) {
		update, err = up.Get()
		if err != nil {
			notifier.Critical("Oh Noosss!", fmt.Sprintf("Something went wrong: %s", err), 40)
			fmt.Printf("ERROR!!")
		}
		if update.Changed(old) {
			Notify(update, notifier, old.State != update.State)
			if report {
				Print(update, notifier)
			}
		}
		old = update
	}
}

func Print(battery upower.Update, notifier *notify.Notifier) {
	switch battery.State {
	case upower.Charging:
		fmt.Printf("C(%v%%):%v\n", battery.Percentage, battery.TimeToFull)
	case upower.Discharging:
		fmt.Printf("D(%v%%):%v\n", battery.Percentage, battery.TimeToEmpty)
	case upower.Empty:
		fmt.Printf("DEAD!\n")
	case upower.FullCharged:
		fmt.Printf("F:%v%%\n", battery.Percentage)
	case upower.PendingCharge:
		fmt.Printf("PC\n")
	case upower.PendingDischarge:
		fmt.Printf("PD\n")
	default:
		fmt.Printf("UNKN(%v)", battery.State)
	}
}

func Notify(battery upower.Update, notifier *notify.Notifier, changed bool) {
	if changed {
		notifier.Normal("Power Change.", fmt.Sprintf("Heads up!! We are now %s.", battery.State), 120)
	}

	switch battery.State {
	case upower.Charging:
		//Do nothing.
	case upower.Discharging:
		switch {
		case battery.TimeToEmpty < critical:
			notifier.Critical("BATTERY LOW!", fmt.Sprintf("Things are getting critical here. %s to go.", battery.TimeToEmpty), 120)
			time.Sleep(critical / 10)
		case battery.TimeToEmpty < warn:
			notifier.Normal("Heads up!!", fmt.Sprintf("We only got %s of juice. Any powerpoints around?", battery.TimeToEmpty), 120)
			time.Sleep(warn / 10)
		default:
			//Do nothing. Everything seems good.
		}
	case upower.Empty:
		notifier.Critical("BATTERY DEAD!", fmt.Sprintf("Things are pretty bad. Battery is dead. %s to go.", battery.TimeToEmpty), 120)
	case upower.FullCharged, upower.PendingCharge, upower.PendingDischarge:
		//Do nothing.
	default:
		notifier.Critical("Oh Noosss!", "I can't figure out battery state!", 40)
	}
}
