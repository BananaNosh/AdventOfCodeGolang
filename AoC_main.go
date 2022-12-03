package main

import (
	AoC "AoC/AoC2022"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var dayStr string
	if len(os.Args) > 1 {
		dayStr = os.Args[1]
	} else {
		dayStr = strconv.Itoa(time.Now().Day())
	}
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch d := day; d {
	case 1:
		AoC.AoC1()
	case 2:
		AoC.AoC2()
	case 3:
		AoC.AoC3()
	case 4:
		AoC.AoC4()
		//		//	case 5:
		//		//		AoC.AoC5()
		//		//	case 6:
		//		//		AoC.AoC6()
		//		//	case 7:
		//		//		AoC.AoC7()
		//		//	case 8:
		//		//		AoC.AoC8()
		//		//	case 9:
		//		//		AoC.AoC9()
		//		//	case 10:
		//		//		AoC.AoC10()
		//		//	case 11:
		//		//		AoC.AoC11()
		//		//	case 12:
		//		//		AoC.AoC12()
		//		//	case 13:
		//		//		AoC.AoC13()
		//		//	case 14:
		//		//		AoC.AoC14()
		//		//	case 15:
		//		//		AoC.AoC15()
		//		//	case 16:
		//		//		AoC.AoC16()
		//		//	case 17:
		//		//		AoC.AoC17()
		//		//	case 18:
		//		//		AoC.AoC18()
		//		//	case 19:
		//		//		AoC.AoC19()
		//		//	case 20:
		//		//		AoC.AoC20()
		//		//	case 21:
		//		//		AoC.AoC21()
		//		//	case 22:
		//		//		AoC.AoC22()
		//		//	case 23:
		//		//		AoC.AoC23()
		//		//	case 24:
		//		//		AoC.AoC24()
		//		//	case 25:
		//		//		AoC.AoC25()
	default:
		AoC.AoC1()
	}
}
