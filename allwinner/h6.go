// Copyright 2022 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// This file contains pin mapping information that is specific to the Allwinner
// H5 model.

package allwinner

import (
	"fmt"
	"strings"

	"periph.io/x/conn/v3/pin"
	"periph.io/x/host/v3/sysfs"
)

// mappingH6 describes the mapping of the H5 processor gpios to their
// alternate functions.
// According to https://linux-sunxi.org/H616 H616 and H618 are the same.
// allwinner marketing team seems to disagree.
//
// It omits the in & out functions which are available on all gpio.
//
// The mapping comes from the datasheet page 26:
// https://linux-sunxi.org/images/3/30/H616_Datasheet_V1.0.pdf
//
//   - The datasheet uses TWI instead of I2C but it is renamed here for
//     consistency.
//   - RGMII means Reduced gigabit media-independent interface.
//   - SDC means SDCard?
//   - NAND connects to a NAND flash controller.
//   - CSI and CCI are for video capture.
var mappingH6 = map[string][5]pin.Func{
	"PC0": {"NAND_WE", "SDC2_DS", "SPI0_CLK", "", "PA_EINT0"},
	"PC1": {"NAND_ALE", "SDC2_RST", "", "", "PA_EINT1"},
	"PC2": {"NAND_CLE", "", "SPI0_MOSI", "", "PA_EINT2"},
	"PC3": {"NAND_CE1", "", "SPI0_CS0", "BOOT_SEL1", "PA_EINT3"},
	"PC4": {"NAND_CE0", "", "SPI0_MIS0", "BOOT_SEL2", "PA_EINT4"},
	"PC5": {"NAND_RE", "SDC2_CLK", "", "BOOT_SEL3", "PA_EINT5"},
	"PC6": {"NAND_RB0", "SDC2_CMD", "", "BOOT_SEL4", "PA_EINT6"},
	"PC7": {"NAND_RB1", "", "SPI0_CS1", "", "PA_EINT7"},
	"PC8": {"NAND_DQ7", "SDC2_D3", "", "", "PA_EINT8"},

	"PC9":  {"NAND_DQ6", "SDC2_D4", "", "", "PA_EINT9"},
	"PC10": {"NAND_DQ5", "SDC2_D0", "", "", "PA_EINT10"},
	"PC11": {"NAND_DQ4", "SDC2_D5", "", "", "PA_EINT11"},
	"PC12": {"NAND_DQS", "", "", "", "PA_EINT12"},
	"PC13": {"NAND_DQ3", "SDC2_D1", "", "", "PA_EINT13"},
	"PC14": {"NAND_DQ2", "SDC2_D6", "", "", "PA_EINT14"},
	"PC15": {"NAND_DQ1", "SDC2_D1", "SPI0_WP", "", "PA_EINT15"},
	"PC16": {"NAND_DQ0", "SDC2_D7", "SPI0_HOLD", "", "PA_EINT16"},
}

// mapH6Pins uses mappingH6 to actually set the altFunc fields of all gpio
// and mark them as available.
//
// It is called by the generic allwinner processor code if an H6 is detected.
func mapH6Pins() error {
	fmt.Println("Initializing H6 pins.")
	for name, altFuncs := range mappingH6 {
		pin := cpupins[name]
		pin.altFunc = altFuncs
		pin.available = true
		if strings.Contains(string(altFuncs[4]), "_EINT") {
			pin.supportEdge = true
		}

		// Initializes the sysfs corresponding pin right away.
		pin.sysfsPin = sysfs.Pins[pin.Number()]
	}
	return nil
}
