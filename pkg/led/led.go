package led

import (
	"fmt"
	"os"

	"golang.org/x/exp/io/i2c"

	"time"
)

const CHAN0 = byte(0x0C)
const CHAN1 = byte(0x0E)

const TSL_ADDR = int(0x39)
const TSL_CMD = byte(0x80)

const CMD_ON = byte(0x03)
const CMD_OFF = byte(0x00)

const LOW_SHORT = byte(0x00) // x1 Gain 13.7 miliseconds
const LOW_MED = byte(0x01)   // x1 Gain 101 miliseconds
const LOW_LONG = byte(0x02)  // x1 Gain 402 miliseconds

const HIGH_SHORT = byte(0x10) // LowLight x16 Gain 13.7 miliseconds
const HIGH_MED = byte(0x11)   // LowLight x16 Gain 100 miliseconds
const HIGH_LONG = byte(0x12)  // LowLight x16 Gain 402 miliseconds

func getGain(level int) byte {

	switch level {
	case 1:
		return LOW_SHORT
	case 2:
		return LOW_MED
	case 3:
		return LOW_LONG
	case 4:
		return HIGH_SHORT
	case 5:
		return HIGH_MED
	default:
		return HIGH_LONG
	}
}

func errState(message string) {
	fmt.Println(message)
	os.Exit(2)
}

func assertError(err error) {
	if err != nil {
		errState(err.Error())
	}
}

func ReadLight(gainLevel int, dev string) uint64 {
	device, err := i2c.Open(&i2c.Devfs{Dev: dev}, TSL_ADDR)
	assertError(err)
	defer device.Close()

	gain := getGain(gainLevel)

	powerOn(device)
	setGain(device, gain)
	defer powerOff(device)

	return readSensorValue(device)
}

func powerOn(device *i2c.Device) {
	device.WriteReg(TSL_CMD, []byte{CMD_ON})
	time.Sleep(10 * time.Millisecond)
}

func powerOff(device *i2c.Device) {
	device.WriteReg(TSL_CMD, []byte{CMD_OFF})
	time.Sleep(10 * time.Millisecond)
}

func setGain(device *i2c.Device, value byte) {
	device.WriteReg(0x01|TSL_CMD, []byte{value})
	time.Sleep(20 * time.Millisecond)
}

func allocWord() []byte {
	return make([]byte, 2)
}

func readSensorValue(device *i2c.Device) uint64 {

	var value uint64
	for i := 0; i < 5; i++ {

		result0 := allocWord()
		result1 := allocWord()

		device.ReadReg(CHAN0|TSL_CMD, result0)
		time.Sleep(10 * time.Millisecond)
		device.ReadReg(CHAN1|TSL_CMD, result1)

		ch0 := uint64(result0[1])*256 + uint64(result0[0])
		ch1 := uint64(result1[1])*256 + uint64(result1[0])

		vResult := ch0 - ch1

		value += vResult
	}

	return value / 5
}
