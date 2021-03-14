package logging

import "log"

func LogError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

