package discovery

import "fmt"

type NoInstanceAvailableError struct {
	service string
}

func (e NoInstanceAvailableError) Error() string {
	return fmt.Sprintf("No instance is available for %s", e.service)
}
