package main

import(
	"fmt"
	"time"
)

func indexOf(word string, data []string) (int) {
    for k, v := range data {
        if word == v {
            return k
        }
    }
    return -1
}

func main() {
	var months []string = string[12]{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	var strmonth string
	fmt.Print("Enter date of birth (y m d): ")
	fmt.Scanf("%d %s %d", &strmonth)
	int month = indexOf(strmonth, months)
	fmt.Println(string(year) + "-" + string(month) + "-" + string(date))
	fmt.Println(time.Now)
}
