package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	// using time.Now() function
	// to get the current time
	currentTime := time.Now()

        // random logic
        // currentTime = currentTime.AddDate(0, 0, 99)

	r := GetMD5Hash(currentTime.Format("20060102"))
	fmt.Printf("Password : %s \n", r[0:10])
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
