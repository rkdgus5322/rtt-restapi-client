package main

import (
    "context"
    "log"
    "os/signal"
    "syscall"

    //"github.com/go-resty/resty/v2"
    "github.com/robfig/cron/v3"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "openscrap/app"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)

	//	serviceKey := "get=message"
	//serviceKey := "get-message"
    serviceKey := "/3670000/ViewIssueLocalPay/ViewIssueLocalPayList?serviceKey=" +
        "KCv6Cl0rBxgnErm7aHRPoosg%2BEePQ%2BXCnPDwlLh%2Bz0ueYn7NomT5SPqVyL5VD22ysYpJPQ6A%2BK3%2F8huN9NkWww%3D%3D" +
        "&find_year=2021&find_month=01&find_area=%EA%B5%AC%EC%95%94&find_age=20&find_sex=%EB%82%A8%EC%84%B1&%20" +
        "find_sigu=%EC%9C%A0%EC%84%B1&find_city=%EB%8C%80%EC%A0%84%EA%B4%91%EC%97%AD%EC%8B%9C"


	// api 호출을 위한 클라이언트 생성
	//	cli := resty.New()
	////	cli.SetBaseURL("https://api.csi.go.kr")
	//    cli.SetBaseURL("http://192.168.100.56:7123")

	//데이터 베이스 접속
	db, err := gorm.Open(mysql.Open("kgh:1234@tcp(mariadb)/testdb?charset=utf8mb4&parseTime=true&multiStatements=true"))
	if err != nil {
		panic(err)
	}
    _=db

	repo := app.NewRepository(db)
    h := app.NewHandler("http://apis.data.go.kr", repo, serviceKey)

    if err := h.Scrap(ctx); err != nil {
        log.Printf("err: %v\n", err)
    }

//	 크론잡 실행  크론 표현식 참고
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
