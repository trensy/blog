package utils

import "log"

type Base struct {

}

func (base *Base) Log(v interface{}) {
	log.Println(v)
}
