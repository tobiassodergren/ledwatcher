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
		currentState := value < treshold
		if verbose {
			fmt.Printf("Value: %d, CurrentState: %t, Treshold: %d\n", value, currentState, treshold)
		}
		if firstRead || currentState != state {
			if verbose {
				fmt.Printf("Acting because firstread: %t, state: %t\n", firstRead, state)
			}
			state = currentState
			c <- Status{state, value, treshold}
			firstRead = false
		}
		time.Sleep(5 * time.Second)
	}
}
