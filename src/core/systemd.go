package core

import (
	"context"
	"fmt"
	"github.com/coreos/go-systemd/dbus"
	"github.com/coreos/go-systemd/sdjournal"
	"io"
)

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

func ReadUnitJournal(name string) error  {
	journalReaderConfig := sdjournal.JournalReaderConfig{
		Matches: []sdjournal.Match{
			{
				Field: sdjournal.SD_JOURNAL_FIELD_SYSTEMD_UNIT,
				Value: name,
			},
		},
	}

	journalReader, err := sdjournal.NewJournalReader(journalReaderConfig)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 64 * 1 << (10)) // 64KB.
	for {
		content, err := journalReader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Print(string(buffer[:content]))
	}
}
