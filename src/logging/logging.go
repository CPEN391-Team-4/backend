package logging

import "log"

// LogError Helper function to log and return error
func LogError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

