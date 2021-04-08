package calendar

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 0 {
		return errors.New("Invalid error")
	}

	d.year = year

	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid month")
	}

	d.month = month

	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid day")
	}

	d.day = day

	return nil
}

func (d Date) GetYear() int {
	return d.year
}

func (d Date) GetMonth() int {
	return d.month
}

func (d Date) GetDay() int {
	return d.day
}
