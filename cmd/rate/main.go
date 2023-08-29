package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gogf/gf/v2/os/gsession"
	"golang.org/x/time/rate"
)

func main() {
	// limit表示每秒产生token数，buret最多存令牌数
	l := rate.NewLimiter(5, 10)
	time.Sleep(time.Second * 3)
	log.Println(l.Limit(), l.Burst())

	for i := 0; i < 15; i++ {
		//这里是阻塞等待的，一直等到取到一个令牌为止
		go func(k int) {
			c, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*225))
			// Wait阻塞等待
			if err := l.Wait(c); err != nil {
				log.Println("limiter wait error : " + err.Error())
				return
			}
			// a = l.Allow()
			// log.Println("wait后, 当前是否可以取到令牌:", i, a)
			log.Println("得到新的令牌:", k, "now=", time.Now())
		}(i)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-quit
	gsession.DefaultStorageFileCryptoKey
}
