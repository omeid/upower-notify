//go:generate stringer -type=State

// A minimal binding for UPower over DBUS.
// it is designed to be as simple as possible.

package upower

import (
	"errors"
	"time"

	"github.com/godbus/dbus"
)

var NoUpower = errors.New("Couldn't get org.freedesktop.UPower")

type State int

const (
	//This order is incosistant it seems, this is listed according to:
	// But it is different, at least in my HP laptop, thus thus they have
	// been set to variables. Feel free to change them.
	Unknown State = iota
	Charging
	Discharging
	Empty
	FullCharged
	PendingCharge
	PendingDischarge
)

type Update struct {
	Capacity         float64
	Energy           float64
	EnergyEmpty      float64
	EnergyFull       float64
	EnergyFullDesign float64
	EnergyRate       float64
	HasHistory       bool
	HasStatistics    bool
	IconName         string
	IsPresent        bool
	IsRechargeable   bool
	Luminosity       float64
	Model            string
	NativePath       string
	Online           bool
	Percentage       float64
	PowerSupply      bool
	Serial           string
	State            State
	Technology       uint32
	Temperature      float64
	TimeToEmpty      time.Duration
	TimeToFull       time.Duration
	Type             uint32
	UpdateTime       uint64
	Vendor           string
	Voltage          float64
	WarningLevel     uint32
}

func (s *Update) Changed(old Update) bool {
	if s.Capacity != old.Capacity {
		return true
	}

	if s.Energy != old.Energy {
		return true
	}

	if s.EnergyEmpty != old.EnergyEmpty {
		return true
	}

	if s.EnergyFull != old.EnergyFull {
		return true
	}

	if s.EnergyFullDesign != old.EnergyFullDesign {
		return true
	}

	if s.EnergyRate != old.EnergyRate {
		return true
	}

	if s.HasHistory != old.HasHistory {
		return true
	}

	if s.HasStatistics != old.HasStatistics {
		return true
	}

	if s.IconName != old.IconName {
		return true
	}

	if s.IsPresent != old.IsPresent {
		return true
	}

	if s.IsRechargeable != old.IsRechargeable {
		return true
	}

	if s.Luminosity != old.Luminosity {
		return true
	}

	if s.Model != old.Model {
		return true
	}

	if s.NativePath != old.NativePath {
		return true
	}

	if s.Online != old.Online {
		return true
	}

	if s.Percentage != old.Percentage {
		return true
	}

	if s.PowerSupply != old.PowerSupply {
		return true
	}

	if s.Serial != old.Serial {
		return true
	}

	if s.State != old.State {
		return true
	}

	if s.Technology != old.Technology {
		return true
	}

	if s.Temperature != old.Temperature {
		return true
	}

	if s.TimeToEmpty != old.TimeToEmpty {
		return true
	}

	if s.TimeToFull != old.TimeToFull {
		return true
	}

	if s.Type != old.Type {
		return true
	}

	if s.UpdateTime != old.UpdateTime {
		return true
	}

	if s.Vendor != old.Vendor {
		return true
	}

	if s.Voltage != old.Voltage {
		return true
	}

	if s.WarningLevel != old.WarningLevel {
		return true
	}

	return false

}

func New(device string) (*UPower, error) {

	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}

	up := conn.Object("org.freedesktop.UPower", dbus.ObjectPath(device))
	if up == nil {
		return nil, NoUpower
	}

	return &UPower{dbus: up}, nil
}

type UPower struct {
	dbus dbus.BusObject
}

func (u *UPower) Get() (Update, error) {

	probs := map[string]dbus.Variant{}
	update := Update{}
	err := u.dbus.Call("org.freedesktop.DBus.Properties.GetAll", 0, "org.freedesktop.UPower.Device").Store(&probs)

	if err != nil {
		return update, err
	}

	update.Capacity = probs["Capacity"].Value().(float64)
	update.Energy = probs["Energy"].Value().(float64)
	update.EnergyEmpty = probs["EnergyEmpty"].Value().(float64)
	update.EnergyFull = probs["EnergyFull"].Value().(float64)
	update.EnergyFullDesign = probs["EnergyFullDesign"].Value().(float64)
	update.EnergyRate = probs["EnergyRate"].Value().(float64)
	update.HasHistory = probs["HasHistory"].Value().(bool)
	update.HasStatistics = probs["HasStatistics"].Value().(bool)
	update.IconName = probs["IconName"].Value().(string)
	update.IsPresent = probs["IsPresent"].Value().(bool)
	update.IsRechargeable = probs["IsRechargeable"].Value().(bool)
	update.Luminosity = probs["Luminosity"].Value().(float64)
	update.Model = probs["Model"].Value().(string)
	update.NativePath = probs["NativePath"].Value().(string)
	update.Online = probs["Online"].Value().(bool)
	update.Percentage = probs["Percentage"].Value().(float64)
	update.PowerSupply = probs["PowerSupply"].Value().(bool)
	update.Serial = probs["Serial"].Value().(string)
	update.State = State(probs["State"].Value().(uint32))
	update.Technology = probs["Technology"].Value().(uint32)
	update.Temperature = probs["Temperature"].Value().(float64)
	update.TimeToEmpty = time.Duration(time.Duration(probs["TimeToEmpty"].Value().(int64)) * time.Second)
	update.TimeToFull = time.Duration(time.Duration(probs["TimeToFull"].Value().(int64)) * time.Second)
	update.Type = probs["Type"].Value().(uint32)
	update.UpdateTime = probs["UpdateTime"].Value().(uint64)
	update.Vendor = probs["Vendor"].Value().(string)
	update.Voltage = probs["Voltage"].Value().(float64)
	update.WarningLevel = probs["WarningLevel"].Value().(uint32)

	return update, err
}
