package metagenerator

import (
	"encoding/json"
	"fmt"
)

type Format int

const (
	SD Format = iota
	HD
	UHD
)

func parseFormat(f string) (Format, error) {
	value, ok := toFormatID[f]
	if !ok {
		return value, fmt.Errorf("%q is not a valid format", f)
	}
	return value, nil
}

func (f Format) String() string {
	return toFormatString[f]
}

func (f Format) Length() int {
	return len(toFormatString)
}

var (
	toFormatString = map[Format]string{
		SD:  "SD",
		HD:  "HD",
		UHD: "4K",
	}
	toFormatID = map[string]Format{
		"SD": SD,
		"HD": HD,
		"4K": UHD,
	}
)

func (f Format) marshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

func (f *Format) UnmarshalJSON(data []byte) (err error) {
	var format string
	if err := json.Unmarshal(data, &format); err != nil {
		return err
	}
	if *f, err = parseFormat(format); err != nil {
		return err
	}
	return nil
}
