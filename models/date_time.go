package models

import "time"

type DateTime struct {
	time.Time
}

func (d *DateTime) MarshalCSV() (string, error) {
	return d.Time.Format(time.RFC3339), nil
}

func (d *DateTime) String() string {
	return d.String()
}

func (d *DateTime) UnmarshalCSV(csv string) (err error) {
	d.Time, err = time.Parse(time.RFC3339, csv)
	return err
}
