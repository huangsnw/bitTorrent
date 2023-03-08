package utils

import "log"

func Check(err error) {
	if err != nil {
		log.Fatalf("Something is wrong : %v", err)
	}
}
