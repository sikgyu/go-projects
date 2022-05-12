package details

import "os"

func GetHostname() (string, error) {
	hostname, err := os.Hostname()

	return hostname, err
}
