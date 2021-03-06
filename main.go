package main

import (
	"Book/Crawler/biquge"
	"Book/Crawler/engine"
	"Book/Crawler/persist"
	"Book/Crawler/scheduler"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//Console.ActTable(reader)
	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:      "http://www.xbiquge.la/fenlei/1_1.html",
		ParseFun: biquge.ParseLinkUl,
	})

	//endless.DefaultReadTimeOut = Setting.ReadTimeout
	//endless.DefaultWriteTimeOut = Setting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", Setting.HTTPPort)
	//server := endless.NewServer(endPoint, Router.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
