package helpers

import "log"

// CheckErr - Simple error checking helper to reduce boilerplate
func CheckErr(err error, msg string) {
	if err != nil {
		if msg != "" {
			log.Fatal(msg)
		}
		panic(err)
	}
}
