package dto

import "reflect"

type Form struct {
	Sites []Site `form:"sites[]"`
}

type Site struct {
	Name string `form:"name"`
	Status string `form:"status"`
	Prefix string `form:"prefix"`
}

func (m Form ) IsEmpty() bool {
	return reflect.DeepEqual(Form{}, m)
}
