package middlewares

import "time"

func ExecuteExample() error {
	// do something
	<-time.After(time.Millisecond * 10)

	return nil
}
