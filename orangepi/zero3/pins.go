package zero3

import (
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/pin"
	"periph.io/x/host/v3/allwinner"
)

var (
	PA1_1  pin.Pin    = pin.V3_3 // VCC 3v3 Ext
	PA1_2  pin.Pin    = pin.V5
	PA1_3  gpio.PinIO = allwinner.PH5
	PA1_4  pin.Pin    = pin.V5
	PA1_5  gpio.PinIO = allwinner.PH4
	PA1_6  pin.Pin    = pin.GROUND
	PA1_7  gpio.PinIO = allwinner.PC9
	PA1_8  gpio.PinIO = allwinner.PH2
	PA1_9  pin.Pin    = pin.GROUND
	PA1_10 gpio.PinIO = allwinner.PH3
	PA1_11 gpio.PinIO = allwinner.PC6
	PA1_12 gpio.PinIO = allwinner.PC11
	PA1_13 gpio.PinIO = allwinner.PC5
	PA1_14 pin.Pin    = pin.GROUND

	PA1_15 gpio.PinIO = allwinner.PC8
	PA1_16 gpio.PinIO = allwinner.PC15
	PA1_17 pin.Pin    = pin.V3_3 // VCC 3v3 Ext
	PA1_18 gpio.PinIO = allwinner.PC15
	PA1_19 gpio.PinIO = allwinner.PH7
	PA1_20 pin.Pin    = pin.GROUND
	PA1_21 gpio.PinIO = allwinner.PH8
	PA1_22 gpio.PinIO = allwinner.PC7
	PA1_23 gpio.PinIO = allwinner.PH6
	PA1_24 gpio.PinIO = allwinner.PH9
	PA1_25 pin.Pin    = pin.GROUND
	PA1_26 gpio.PinIO = allwinner.PH10
)
