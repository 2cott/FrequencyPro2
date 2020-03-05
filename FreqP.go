package main

import (
	"fmt"
	"strings"

	"github.com/rucuriousyet/monolog"
)

func main() {
	notes := map[string]float32{
		"a":       27.5,
		"a#":      29.14,
		"a sharp": 29.14,
		"bb":      29.14,
		"b flat":  29.14,
		"b":       30.87,
		"b#":      32.7,
		"b sharp": 32.7,
		"cb":      30.87,
		"c flat":  30.87,
		"c":       32.7,
		"c#":      34.65,
		"c sharp": 34.65,
		"db":      34.65,
		"d flat":  34.65,
		"d":       36.71,
		"d#":      38.89,
		"d sharp": 38.89,
		"eb":      38.89,
		"e flat":  38.89,
		"e":       41.2,
		"e#":      43.65,
		"e sharp": 43.65,
		"fb":      41.2,
		"f flat":  41.2,
		"f":       43.65,
		"f#":      46.25,
		"f sharp": 46.25,
		"gb":      46.25,
		"g flat":  46.25,
		"g":       49,
		"g#":      51.91,
		"g sharp": 51.91,
		"ab":      51.91,
		"a flat":  51.91,
	}

	callbackFunc := func(p *monolog.Prompter) monolog.Cmd {
		p.Write("Type a note name (result will print in Hertz): ")
		input := p.Read()
		letter := strings.ToLower(input)

		fmt.Println(notes[letter])
		return monolog.Continue
	}

	err := monolog.New(nil, nil).Add(callbackFunc).Do()
	if err != nil {
		panic(err)
	}

}

/* C0 to B0 would be position 1 (P1), C1 to B1 would be position 2 (P2)
and so on up to C8 (P9). Frequencies in Position 1 would be the base frequencies for
multiplication all the way to Position 9 (P9). for ex: P2 frequencies
would be the result of P1 frequencies * 2. P3 frequencies
would be P1 frequencies * 4.
the following shows the breakdown
P1 = Base Frequencies
P2 = P1 * 2
P3 = P1 * 4
P4 = P1 * 8
P5 = P1 * 16
P6 = P1 * 32
P7 = P1 * 64
P8 = P1 * 128
P9 = P1 * 256

*/
//oct := []int{2, 6, 7, 8, 9, 0, 1}

//for index, multi := range oct {
//	fmt.Printf("E%d => %d\n", index, tones["e"] * multi)
// }

//fmt.Println("\n", oct)

//oct[3] = 1001

//fmt.Println(oct)

/*oct[2] = 999

fmt.Println(oct)*/
