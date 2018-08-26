package common

import "log"

var (
	// ErrorCheck is a wrapper for error checking
	ErrorCheck = func(v ...interface{}) {
		for _, e := range v {
			switch e.(type) {
			case error:
				if e != nil {
					log.Fatalln(v...)
				}
			}
		}
	}
)
