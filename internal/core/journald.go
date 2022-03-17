package core

import (
	"github.com/coreos/go-systemd/v22/sdjournal"
	"io/ioutil"
)

// ReadUnitJournal to read a unit journal.
func ReadUnitJournal(name string) ([]byte, error) {
	_, err := checkUnit(name)
	if err != nil {
		return nil, err
	}
	readerConfig := sdjournal.JournalReaderConfig{
		Matches: []sdjournal.Match{
			{
				Field: sdjournal.SD_JOURNAL_FIELD_SYSTEMD_UNIT,
				Value: name,
			},
		},
	}
	reader, err := sdjournal.NewJournalReader(readerConfig)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return content, err
}
