package main

import (
	"AoC/AoC2022/AoC_22_1"
	"AoC/AoC2022/AoC_22_10"
	"AoC/AoC2022/AoC_22_11"
	"AoC/AoC2022/AoC_22_12"
	"AoC/AoC2022/AoC_22_13"
	"AoC/AoC2022/AoC_22_14"
	"AoC/AoC2022/AoC_22_15"
	"AoC/AoC2022/AoC_22_16"
	"AoC/AoC2022/AoC_22_17"
	"AoC/AoC2022/AoC_22_18"
	"AoC/AoC2022/AoC_22_2"
	"AoC/AoC2022/AoC_22_3"
	"AoC/AoC2022/AoC_22_4"
	"AoC/AoC2022/AoC_22_5"
	"AoC/AoC2022/AoC_22_6"
	"AoC/AoC2022/AoC_22_7"
	"AoC/AoC2022/AoC_22_8"
	"AoC/AoC2022/AoC_22_9"
	//"AoC/AoC2022/AoC_22_19"
	//"AoC/AoC2022/AoC_22_20"
	//"AoC/AoC2022/AoC_22_21"
	//"AoC/AoC2022/AoC_22_22"
	//"AoC/AoC2022/AoC_22_23"
	//"AoC/AoC2022/AoC_22_24"
	//"AoC/AoC2022/AoC_22_25"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var dayStrs []string
	if len(os.Args) > 1 {
		dayStrs = os.Args[1:]
	} else {
		dayStrs = []string{strconv.Itoa(time.Now().Day())}
	}
	for _, dayStr := range dayStrs {
		day, err := strconv.Atoi(dayStr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		switch d := day; d {
		case 1:
			AoC_22_1.AoC1()
		case 2:
			AoC_22_2.AoC2()
		case 3:
			AoC_22_3.AoC3()
		case 4:
			AoC_22_4.AoC4()
		case 5:
			AoC_22_5.AoC5()
		case 6:
			AoC_22_6.AoC6()
		case 7:
			AoC_22_7.AoC7()
		case 8:
			AoC_22_8.AoC8()
		case 9:
			AoC_22_9.AoC9()
		case 10:
			AoC_22_10.AoC10()
		case 11:
			AoC_22_11.AoC11()
		case 12:
			AoC_22_12.AoC12()
		case 13:
			AoC_22_13.AoC13()
		case 14:
			AoC_22_14.AoC14()
		case 15:
			AoC_22_15.AoC15()
		case 16:
			AoC_22_16.AoC16()
		case 17:
			AoC_22_17.AoC17()
		case 18:
			AoC_22_18.AoC18()
		//case 19:
		//	AoC_22_19.AoC19()
		//case 20:
		//	AoC_22_20.AoC20()
		//case 21:
		//	AoC_22_21.AoC21()
		//case 22:
		//	AoC_22_22.AoC22()
		//case 23:
		//	AoC_22_23.AoC23()
		//case 24:
		//	AoC_22_24.AoC24()
		//case 25:
		//	AoC_22_25.AoC25()
		default:
			AoC_22_1.AoC1()
		}
	}
}
