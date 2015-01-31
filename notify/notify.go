//go:generate stringer -type=Urgency
package notify

import (
	"errors"

	"github.com/godbus/dbus"
)

var NoNotifications = errors.New("Couldn't get org.freedesktop.Notifications")

type Urgency byte

const (
	Low Urgency = iota
	Normal
	Critical
)

type Message struct {
	AppName       string
	ReplacesId    uint32
	AppIcon       string
	Summary       string
	Body          string
	Actions       []string
	Hints         map[string]dbus.Variant
	ExpireTimeout int32
}

type Notifier struct {
	dbus *dbus.Object

	app string
}

func New(app string) (*Notifier, error) {

	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}
	notification := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	if notification == nil {
		return nil, NoNotifications
	}

	return &Notifier{dbus: notification, app: app}, nil
}

func (n *Notifier) Low(Summary string, Body string, ExpireTimeout int32) error {
	return n.Send(Summary, Body, Low, ExpireTimeout)
}

func (n *Notifier) Normal(Summary string, Body string, ExpireTimeout int32) error {
	return n.Send(Summary, Body, Normal, ExpireTimeout)
}

func (n *Notifier) Critical(Summary string, Body string, ExpireTimeout int32) error {
	return n.Send(Summary, Body, Critical, ExpireTimeout)
}

func (n *Notifier) SendMessage(m *Message) error {

	return n.dbus.Call("org.freedesktop.Notifications.Notify", 0,
		m.AppName,
		m.ReplacesId,
		m.AppIcon,
		m.Summary,
		m.Body,
		m.Actions,
		m.Hints,
		m.ExpireTimeout,
	).Err
}

func (n *Notifier) Send(Summary string, Body string, urgency Urgency, ExpireTimeout int32) error {
	return n.SendMessage(&Message{
		n.app,
		0,
		"",
		Summary,
		Body,
		[]string{},
		map[string]dbus.Variant{"urgency": dbus.MakeVariant(urgency)},
		ExpireTimeout})
}
