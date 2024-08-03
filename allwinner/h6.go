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
	"PC0":  {"NAND_WE", "SDC2_DS", "SPI0_CLK", "", "PC_EINT0"},
	"PC1":  {"NAND_ALE", "SDC2_RST", "", "", "PC_EINT1"},
	"PC2":  {"NAND_CLE", "", "SPI0_MOSI", "", "PC_EINT2"},
	"PC3":  {"NAND_CE1", "", "SPI0_CS0", "BOOT_SEL1", "PC_EINT3"},
	"PC4":  {"NAND_CE0", "", "SPI0_MIS0", "BOOT_SEL2", "PC_EINT4"},
	"PC5":  {"NAND_RE", "SDC2_CLK", "", "BOOT_SEL3", "PC_EINT5"},
	"PC6":  {"NAND_RB0", "SDC2_CMD", "", "BOOT_SEL4", "PC_EINT6"},
	"PC7":  {"NAND_RB1", "", "SPI0_CS1", "", "PC_EINT7"},
	"PC8":  {"NAND_DQ7", "SDC2_D3", "", "", "PC_EINT8"},
	"PC9":  {"NAND_DQ6", "SDC2_D4", "", "", "PC_EINT9"},
	"PC10": {"NAND_DQ5", "SDC2_D0", "", "", "PC_EINT10"},
	"PC11": {"NAND_DQ4", "SDC2_D5", "", "", "PC_EINT11"},
	"PC12": {"NAND_DQS", "", "", "", "PC_EINT12"},
	"PC13": {"NAND_DQ3", "SDC2_D1", "", "", "PC_EINT13"},
	"PC14": {"NAND_DQ2", "SDC2_D6", "", "", "PC_EINT14"},
	"PC15": {"NAND_DQ1", "SDC2_D1", "SPI0_WP", "", "PC_EINT15"},
	"PC16": {"NAND_DQ0", "SDC2_D7", "SPI0_HOLD", "", "PC_EINT16"},

	"PF0": {"SDC0_D1", "JTAG_MS", "", "", "PF_EINT0"},
	"PF1": {"SDC0_D0", "JTAG_DI", "", "", "PF_EINT1"},
	"PF2": {"SDC0_CLK", "UTAR0_TX", "", "", "PF_EINT2"},
	"PF3": {"SDC0_CMD", "JTAG_DO", "", "", "PF_EINT3"},
	"PF4": {"SDC0_D3", "UTAR0_RX", "", "", "PF_EINT4"},
	"PF5": {"SDC0_D2", "JTAG_CK", "", "", "PF_EINT5"},
	"PF6": {"", "", "", "", "PF_EINT6"},

	"PG0":  {"SDC1_CLK", "", "", "", "PG_EINT0"},
	"PG1":  {"SDC1_CMD", "", "", "", "PG_EINT1"},
	"PG2":  {"SDC1_D0", "", "", "", "PG_EINT2"},
	"PG3":  {"SDC1_D1", "", "", "", "PG_EINT3"},
	"PG4":  {"SDC1_D2", "", "", "", "PG_EINT4"},
	"PG5":  {"SDC1_D3", "", "", "", "PG_EINT5"},
	"PG6":  {"UTAR1_TX", "", "JTAG_MS", "", "PG_EINT6"},
	"PG7":  {"UTAR1_RX", "", "JTAG_CK", "", "PG_EINT7"},
	"PG8":  {"UTAR1_RTS", "PLL_LOCK_DBG", "JTAG_DO", "", "PG_EINT8"},
	"PG9":  {"UTAR1_CTS", "", "JTAG_DI", "", "PG_EINT9"},
	"PG10": {"H_I2S2_MCLK", "X32KFOUT", "", "", "PG_EINT10"},
	"PG11": {"H_I2S2_BCLK", "", "BIST_RESULT0", "", "PG_EINT11"},
	"PG12": {"H_I2S2_LRCK", "", "BIST_RESULT1", "", "PG_EINT12"},
	"PG13": {"H_I2S2_DOUT0", "H_I2S2_DOUT1", "BIST_RESULT2", "", "PG_EINT13"},
	"PG14": {"H_I2S2_DIN0", "H_I2S2_DOUT1", "BIST_RESULT3", "", "PG_EINT14"},
	"PG15": {"UTAR2_TX", "", "", "TWI4_SCK", "PG_EINT15"},
	"PG16": {"UTAR2_RX", "", "", "TWI4_SDA", "PG_EINT16"},
	"PG17": {"UTAR2_RTS", "", "", "TWI3_SCK", "PG_EINT17"},
	"PG18": {"UTAR2_CTS", "", "", "TWI3_SDA", "PG_EINT18"},
	"PG19": {"", "", "PWM1", "", "PG_EINT19"},
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
