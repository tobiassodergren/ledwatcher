package led

import (
	"fmt"
	"time"
)

var lastState bool

func Read(verbose bool, gain int, treshold uint64, device string, c chan Status) {
	if verbose {
		fmt.Println("About to start reading values from led.")
	}
	initialize := true
	for {
		value := ReadLight(gain, device)
		isLit := value > treshold
		if verbose {
			fmt.Printf("Value: %d, Treshold: %d\n", value, treshold)
		}
		if initialize || isLit != lastState {
			if verbose {
				fmt.Printf("Acting because firstread: %t, lastState: %t, isLit: %t\n",
					initialize, lastState, isLit)
			}
			lastState = isLit
			c <- Status{lastState, value, treshold}
			initialize = false
		}
		time.Sleep(5 * time.Second)
	}
}
