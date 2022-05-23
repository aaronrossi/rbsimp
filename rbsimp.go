package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	currentTime := time.Now()

	rbsimpYear := (currentTime.Format("06"))
	rbsimpMonth := (currentTime.Format("02"))
	rbsimpDay := (currentTime.Format("01"))

	fmt.Println(rbsimpDay)
	fmt.Println(rbsimpMonth)
	fmt.Println(rbsimpYear)

	iRbsimpDay, err := strconv.Atoi(rbsimpDay)
	iRbsimpMonth, err := strconv.Atoi(rbsimpMonth)
	iRbsimpYear, err := strconv.Atoi(rbsimpYear)
	if err == nil {
		total := iRbsimpDay + iRbsimpMonth - iRbsimpYear
		totalFormatted := fmt.Sprintf("%02d", total)
		fmt.Println("Todays number is:", totalFormatted)
	}
}
