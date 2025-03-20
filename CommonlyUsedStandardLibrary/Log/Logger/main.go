package main

import "log"

// logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。
func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条普通的日志")
	v := "很普通的"
	log.SetPrefix("[pprof]")
	log.Printf("这是一条%s日志\n", v)
	log.Fatalln("这是一条会触发fatal的日志")
	//log.Panicln("这是一条会触发panic的日志")
}
