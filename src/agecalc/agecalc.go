package main

import(
	"fmt"
	"time"
	"strconv"
)

func indexOf(word string, data [12]string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func nextNum(str string, index int) int {
	runeStr := []rune(str)
	strSlice := make([]string, len(runeStr))
	var months [40]string
	for x := 0; x < len(strSlice); x++ {
		strSlice[x] = string(runeStr[x])
	}
	var tempStr string = ""
	var cnt int = 0
	for x := 0; x < len(strSlice); x++ {
		if isNumeric(strSlice[x]) {
			tempStr += strSlice[x]
		}else {
			months[cnt] = tempStr
			cnt += 1
			tempStr = ""
		}
	}
	i1, err := strconv.Atoi(months[index])
	if err == nil {
		return i1
	}
	return 0
}

func main() {
	var months [12]string = [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	var month string
	var day int
	var year int
	fmt.Println("Enter date of birth (m d y): ")
	fmt.Scan(&month, &day, &year)
	var monthNum int = indexOf(month, months)+1
	var currentYear int = nextNum(time.Now().String(), 0)
	var currentMonth int = nextNum(time.Now().String(), 1)
	var currentDay int = nextNum(time.Now().String(), 2)
	var yearDiff int
	var monthDiff int
	var dayDiff int
	var accurateAge float64
	var carried bool = false
	if currentDay - day >= 0 {
		dayDiff = currentDay-day
	}else {
		dayDiff = (currentDay+30)-day
		carried = true
	}
	if carried == false {
		if currentMonth - monthNum >= 0 {
			monthDiff = currentMonth - monthNum
		}else {
			monthDiff = (currentMonth+30) - monthNum
			carried = true
		}
	}else {
		if currentMonth - monthNum >= 0 {
			monthDiff = (currentMonth-1) - monthNum
			carried = false
		}else {
			monthDiff = ((currentMonth-1)+12) - monthNum
		}
	}
	if carried == true {
		yearDiff = (currentYear-1) - year
	}else {
		yearDiff = currentYear - year
	}
	accurateAge = float64(yearDiff) + (float64(monthDiff)/12) + (float64(dayDiff)/365)
	fmt.Println(accurateAge)
	fmt.Print(yearDiff, " years, ")
	fmt.Print(monthDiff, " months and ")
	fmt.Println(dayDiff, "days old")
}
