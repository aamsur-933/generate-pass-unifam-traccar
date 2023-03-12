package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"time"
)

func main() {
	paramDateRaw := flag.String("date", "", "a string for description")
	flag.Parse()

	paramDate := ""
	if *paramDateRaw == "" {
		// using time.Now() function
		// to get the current time
		currentTime := time.Now()
		paramDate = currentTime.Format("20060102")
	} else {
		paramDate = *paramDateRaw
	}

	// random logic
	// currentTime = currentTime.AddDate(0, 0, 99)

	r := GetMD5Hash(paramDate)
	fmt.Printf("Password : %s \n", r[0:10])
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
