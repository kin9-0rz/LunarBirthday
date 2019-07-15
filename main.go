package main

import (
	"flag"
	"fmt"
	"github.com/nosixtools/solarlunar"
	"os"
	"strconv"
	"strings"
	"time"
)

var VERSION = "0.0.1"

var (
	subject  string
	birthday string
	lunar    bool
	times    int
	help     bool
	version  bool
)

func init() {
	flag.StringVar(&subject, "s", "", "标题")
	flag.StringVar(&birthday, "b", "", "输入日期，1990-01-01")
	flag.BoolVar(&lunar, "l", false, "是否为农历")
	flag.IntVar(&times, "n", 31, "次数，默认31")
	flag.BoolVar(&help, "h", false, "显示用法")
	flag.BoolVar(&version, "v", false, "show version")
}

func main() {
	flag.Parse()

	if help || len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}

	if version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	days := strings.Split(birthday, "-")
	if len(days) < 2 {
		fmt.Println("生日格式为：1900-01-01")
		os.Exit(1)
	}

	year := time.Now().Year()
	fmt.Println("Subject, Start Date, Start Time, End Date, End Time, Private, All Day Event, Location")
	for a := 0; a < times; a++ {
		day := strconv.Itoa(year) + "-" + days[1] + "-" + days[2]
		if lunar {
			day = solarlunar.LunarToSolar(day, false)
		}
		fmt.Printf("%s, %s, 8:00 AM, %s, , FALSE, TRUE, HOME\n", subject, day, day)
		year++
		if year > 2049 {
			// 这个库不支持超过2049年的农历
			break
		}
	}

}
