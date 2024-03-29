package core

import (
	"context"
	"github.com/coreos/go-systemd/v22/dbus"
)

// InvalidUnitName custom error.
type InvalidUnitName struct{}

func (e *InvalidUnitName) Error() string {
	return "Invalid systemd unit name"
}

func newDbusConnection() *dbus.Conn {
	connection, err := dbus.NewWithContext(context.Background())
	if err != nil {
		panic(err)
	}
	return connection
}

// GetUnits to list and return units.
func GetUnits() ([]dbus.UnitStatus, error) {
	conn := newDbusConnection()
	defer conn.Close()
	return conn.ListUnitsContext(context.Background())
}

func checkUnit(name string) (*string, error) {
	units, err := GetUnits()
	if err != nil {
		return nil, err
	}
	for index := range units {
		if units[index].Name == name {
			return &name, nil
		}
	}
	return nil, &InvalidUnitName{}
}

// GetUnit to get and return single unit.
func GetUnit(name string) (*dbus.UnitStatus, error) {
	conn := newDbusConnection()
	defer conn.Close()
	_, err := checkUnit(name)
	if err != nil {
		return nil, err
	}
	result, err := conn.ListUnitsByNamesContext(context.Background(), []string{name})
	return &result[0], err
}

// StartUnit to start a unit.
func StartUnit(name string, wait bool) error {
	conn := newDbusConnection()
	defer conn.Close()
	_, err := checkUnit(name)
	if err != nil {
		return err
	}
	result := make(chan string)
	_, err = conn.StartUnitContext(context.Background(), name, "replace", result)
	if wait {
		_ = <-result
	}
	return err
}

// StopUnit to stop a unit.
func StopUnit(name string, wait bool) error {
	conn := newDbusConnection()
	defer conn.Close()
	_, err := checkUnit(name)
	if err != nil {
		return err
	}
	result := make(chan string)
	_, err = conn.StopUnitContext(context.Background(), name, "replace", result)
	if wait {
		_ = <-result
	}
	return err
}
