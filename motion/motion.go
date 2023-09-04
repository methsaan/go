package main

import (
	"fmt"
	"strconv"
	"math"
)

// vf = vi + a*t
// d = vi*t + 0.5*a*math.Pow(t, 2)
// d = 0.5*(vi + vf)*t
// math.Pow(vf, 2) = math.Pow(vi, 2) + 2*a*d
// d = vf*t - 0.5*a*math.Pow(t, 2)

func getVariables() map[string]float64 {
	keys := [5]string{"displacement", "velocityInit", "velocityFinal", "acceleration", "time"}
	varNames := [5]string{"displacement", "initial velocity", "final velocity", "acceleration", "time"}
	units := [5]string{"m", "m/s", "m/s", "m/s^2", "s"}
	obj := make(map[string]float64)
	for x := 0; x < 5; x++ {
		var temp string
		fmt.Printf("Enter %s (%s), '-' to leave blank: ", varNames[x], units[x]);
		fmt.Scanln(&temp)
		if temp != "-" {
			if s, err := strconv.ParseFloat(temp, 64); err == nil {
				obj[keys[x]] = s
			}
		}
	}
	return obj
}

func missingVariables(object map[string]float64) []string {
	keys := [5]string{"displacement", "velocityInit", "velocityFinal", "acceleration", "time"}
	var missing []string
	for x := 0; x < 5; x++ {
		if _, exists := object[keys[x]]; !exists {
			missing = append(missing, keys[x])
		}
	}
	return missing
}

// vf = vi + a*t
func firstEquation(object map[string]float64, missing []string) (map[string]float64, []string) {
	inEquation := []string{"velocityInit", "velocityFinal", "acceleration", "time"}
	var newObj = make(map[string]float64)
	for k, v := range object {
		newObj[k] = v
	}
	var newMissing []string
	for i := range missing {
		newMissing = append(newMissing, missing[i])
	}
	unknownCnt := 0
	var unknown string
	for i := range missing {
		if contains(inEquation, missing[i]) {
			unknownCnt += 1
			unknown = missing[i]
		}
	}
	if unknownCnt > 1 {
		return newObj, newMissing
	}
	var solved float64
	vi := object["velocityInit"]
	vf := object["velocityFinal"]
	t := object["time"]
	a := object["acceleration"]
	if contains(missing, "acceleration") {
		a = (vf - vi) / t
		solved = a
	} else if contains(missing, "time") {
		t = (vf - vi) / a
		solved = t
	} else if contains(missing, "velocityFinal") {
		vf = vi + a * t
		solved = vf
	} else {
		vi = vf - a * t
		solved = vi
	}
	newObj[unknown] = solved
	newMissing = remove(newMissing, unknown)
	return newObj, newMissing
}

// d = vi*t + 0.5*a*math.Pow(t, 2)
func secondEquation(object map[string]float64, missing []string) (map[string]float64, []string) {
	inEquation := []string{"velocityInit", "displacement", "acceleration", "time"}
	var newObj = make(map[string]float64)
	for k, v := range object {
		newObj[k] = v
	}
	var newMissing []string
	for i := range missing {
		newMissing = append(newMissing, missing[i])
	}
	unknownCnt := 0
	var unknown string
	for i := range missing {
		if contains(inEquation, missing[i]) {
			unknownCnt += 1
			unknown = missing[i]
		}
	}
	if unknownCnt > 1 {
		return newObj, newMissing
	}
	var solved float64
	vi := object["velocityInit"]
	d := object["displacement"]
	t := object["time"]
	a := object["acceleration"]
	if contains(missing, "acceleration") {
		a = 2 * (d - vi * t) / math.Pow(t, 2)
		solved = a
	} else if contains(missing, "time") {
		t1 := (-vi + math.Sqrt(math.Pow(vi, 2) + 2 * a * d)) / a
		t2 := (-vi - math.Sqrt(math.Pow(vi, 2) + 2 * a * d)) / a
		if t1 > 0 {
			t = t1
		} else {
			t = t2
		}
		solved = t
	} else if contains(missing, "velocityInit") {
		vi = (d - 0.5 * a * math.Pow(t, 2)) / t
		solved = vi
	} else {
		d = vi * t + 0.5 * a * math.Pow(t, 2)
		solved = d
	}
	newObj[unknown] = solved
	newMissing = remove(newMissing, unknown)
	return newObj, newMissing
}

// d = 0.5*(vi + vf)*t
func thirdEquation(object map[string]float64, missing []string) (map[string]float64, []string) {
	inEquation := []string{"displacement", "velocityFinal", "velocityInit", "time"}
	var newObj = make(map[string]float64)
	for k, v := range object {
		newObj[k] = v
	}
	var newMissing []string
	for i := range missing {
		newMissing = append(newMissing, missing[i])
	}
	unknownCnt := 0
	var unknown string
	for i := range missing {
		if contains(inEquation, missing[i]) {
			unknownCnt += 1
			unknown = missing[i]
		}
	}
	if unknownCnt > 1 {
		return newObj, newMissing
	}
	var solved float64
	vi := object["velocityInit"]
	vf := object["velocityFinal"]
	t := object["time"]
	d := object["displacement"]
	if contains(missing, "displacement") {
		d = 0.5 * (vi + vf) * t
		solved = d
	} else if contains(missing, "time") {
		t = (2 * d) / (vi + vf)
		solved = t
	} else if contains(missing, "velocityFinal") {
		vf = (2 * d) / t - vi
		solved = vf
	} else {
		vi = (2 * d) / t - vf
		solved = vi
	}
	newObj[unknown] = solved
	newMissing = remove(newMissing, unknown)
	return newObj, newMissing
}

// math.Pow(vf, 2) = math.Pow(vi, 2) + 2*a*d
func fourthEquation(object map[string]float64, missing []string) (map[string]float64, []string) {
	inEquation := []string{"displacement", "velocityFinal", "velocityInit", "acceleration"}
	var newObj = make(map[string]float64)
	for k, v := range object {
		newObj[k] = v
	}
	var newMissing []string
	for i := range missing {
		newMissing = append(newMissing, missing[i])
	}
	unknownCnt := 0
	var unknown string
	for i := range missing {
		if contains(inEquation, missing[i]) {
			unknownCnt += 1
			unknown = missing[i]
		}
	}
	if unknownCnt > 1 {
		return newObj, newMissing
	}
	var solved float64
	vi := object["velocityInit"]
	vf := object["velocityFinal"]
	a := object["acceleration"]
	d := object["displacement"]
	if contains(missing, "displacement") {
		d = (math.Pow(vf, 2) - math.Pow(vi, 2)) / (2 * a)
		solved = d
	} else if contains(missing, "acceleration") {
		a = (math.Pow(vf, 2) - math.Pow(vi, 2)) / (2 * d)
		solved = a
	} else if contains(missing, "velocityFinal") {
		vf = math.Sqrt(math.Pow(vi, 2) + 2 * a * d)
		solved = vf
	} else {
		vi = math.Sqrt(math.Pow(vf, 2) - 2 * a * d)
		solved = vi
	}
	newObj[unknown] = solved
	newMissing = remove(newMissing, unknown)
	return newObj, newMissing
}

// d = vf*t - 0.5*a*math.Pow(t, 2)
func fifthEquation(object map[string]float64, missing []string) (map[string]float64, []string) {
	inEquation := []string{"displacement", "velocityFinal", "time", "acceleration"}
	var newObj = make(map[string]float64)
	for k, v := range object {
		newObj[k] = v
	}
	var newMissing []string
	for i := range missing {
		newMissing = append(newMissing, missing[i])
	}
	unknownCnt := 0
	var unknown string
	for i := range missing {
		if contains(inEquation, missing[i]) {
			unknownCnt += 1
			unknown = missing[i]
		}
	}
	if unknownCnt > 1 {
		return newObj, newMissing
	}
	var solved float64
	t := object["time"]
	vf := object["velocityFinal"]
	a := object["acceleration"]
	d := object["displacement"]
	if contains(missing, "displacement") {
		d = vf * t - 0.5 * a * math.Pow(t, 2)
		solved = d
	} else if contains(missing, "acceleration") {
		a = 2 * (d - vf * t) / math.Pow(t, 2)
		solved = a
	} else if contains(missing, "velocityFinal") {
		vf = (d + 0.5 * a * math.Pow(t, 2)) / 2
		solved = vf
	} else {
		t1 := (-vf + math.Sqrt(math.Pow(vf, 2) - 2 * a * d)) / a
		t2 := (-vf - math.Sqrt(math.Pow(vf, 2) - 2 * a * d)) / a
		fmt.Println(t1, t2)
		if t1 > 0 {
			t = t1
		} else {
			t = t2
		}
		solved = t
	}
	newObj[unknown] = solved
	newMissing = remove(newMissing, unknown)
	return newObj, newMissing
}

func contains(stringarr []string, str string) bool {
	for _, value := range stringarr {
		if value == str {
			return true
		}
	}
	return false
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func main() {
	x := getVariables()
	y := missingVariables(x)
	for i := range y {
		x[y[i]] = 0
	}
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(fifthEquation(x, y))
}
