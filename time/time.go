package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UTC().Format("2006-01-02 15:04:05"), " ", time.Now().UTC().Local().Unix()) //输出英国伦敦时区时间 - 0时区 - UTC±0
	fmt.Println(time.Now().Local().Format("2006-01-02 15:04:05"), " ", time.Now().Local().Unix())     //输出中国上海时区 - 中国时间
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " ", time.Now().Unix())                     //输出中国上海时区 - 中国时间

	location, err := time.LoadLocation("Asia/Shanghai") //"America/New_York"
	if err == nil {
		time.Local = location
	}
	fmt.Println(time.Now().UTC().Format("2006-01-02 15:04:05"), " ", time.Now().UTC().Local().Unix()) //输出英国伦敦时区时间 - 0时区 - UTC±0
	fmt.Println(time.Now().Local().Format("2006-01-02 15:04:05"), " ", time.Now().Local().Unix())     //输出中国上海时区 - 中国时间
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " ", time.Now().Unix())                     //输出中国上海时区 - 中国时间
}

/*
2021-04-22 14:09:49   1619100589
2021-04-22 22:09:49   1619100589
2021-04-22 22:09:49   1619100589
2021-04-22 14:09:49   1619100589
2021-04-22 22:09:49   1619100589
2021-04-22 22:09:49   1619100589

https://blog.csdn.net/wschq/article/details/80114036
*/

/*
使用 time.LoadLocation 和 time.ParseInLocation 可以得到本地时间而不是UTC时间
time.Now().UTC().Unix() 和 time.Now().Unix() 得到的时间是一样的，说明unix时间戳不区分是否是UTC
*/
