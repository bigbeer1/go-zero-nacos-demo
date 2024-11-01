package global

import (
	"sync"
	"time"
)

/*
全局变量
*/

var (
	ShangHaiTime, _ = time.LoadLocation("Asia/Shanghai") //上海
	Ulock           sync.Mutex
)
