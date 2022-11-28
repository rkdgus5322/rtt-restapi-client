package main

import (
	"context"
	"log"
	"openscrap/app"
	"os/signal"
	"syscall"

	"github.com/go-resty/resty/v2"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)

	serviceKey := ""

	// api 호출을 위한 클라이언트 생성
	cli := resty.New()
	cli.SetBaseURL("https://api.csi.go.kr")

	//데이터 베이스 접속
	db, err := gorm.Open(mysql.Open("kgh:1234@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=true&multiStatements=true"))
	if err != nil {
		panic(err)
	}

	repo := app.NewRepository(db)
	h := app.NewHandler(cli, repo, serviceKey)

	// 크론잡 실행  크론 표현식 참고
	c := cron.New()
	c.AddFunc("* * * * *", func() {
		if err := h.Scrap(ctx); err != nil {
			log.Printf("err: %v\n", err)
		}
	})

	c.Start()

	<-ctx.Done()
	cancel()
	c.Stop()

}
