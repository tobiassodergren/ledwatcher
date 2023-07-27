package led

import (
	"fmt"
	"time"
)

var state bool

func Read(verbose bool, gain int, treshold uint64, device string, c chan Status) {
	if verbose {
		fmt.Println("About to start reading values from led.")
	}
	firstRead := true
	for {
		value := ReadLight(gain, device)
		currentLitState := value > treshold
		if verbose {
			fmt.Printf("Value: %d, Treshold: %d\n", value, treshold)
		}
		if firstRead || currentLitState != state {
			if verbose {
				fmt.Printf("Acting because firstread: %t, state: %t, newState: %t\n", firstRead, state, currentLitState)
			}
			state = currentLitState
			c <- Status{state, value, treshold}
			firstRead = false
		}
		time.Sleep(5 * time.Second)
	}
}
