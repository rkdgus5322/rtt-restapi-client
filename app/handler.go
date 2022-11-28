package app

import (
	"context"
	"encoding/json"
	"fmt"
    "io/ioutil"
    "log"
    "net/http"

	"openscrap/domain"
)

type Handler struct {
	base string
	repo *Repository
	key  string
}

func NewHandler(base string, repo *Repository, serviceKey string) *Handler {
	return &Handler{
		base: base,
		repo: repo,
		key:  serviceKey,
	}
}

func (h *Handler) Scrap(ctx context.Context) error {
    // api path 추가

    response, err := http.Get(h.base + h.key)
    if err != nil {
    fmt.Print(err)

    }
    //	responseData, err := ioutil.ReadAll(response.Body)
    responseData, err := ioutil.ReadAll(response.Body)
    //responseData, err := ioutil.ReadAll(response.Body)

    if err != nil {
    log.Fatal(err)
    }
    fmt.Println(string(responseData))

    //    str, _ := json.Marshal(string(responseData))
    //    _ = str
    //    fmt.Printf("json: %v", str)

    //json 파싱을 진행
    var resp *domain.Response
    if err := json.Unmarshal(responseData, &resp); err != nil {
        fmt.Printf("handler %v", err)
        return err
    }
    //    fmt.Println(string(responseData))
    fmt.Printf("%#v\n", resp)
//    fmt.Printf("%s", resp)
    //    if err != nil {
    //        fmt.Print(err)
    //    }
     return h.repo.Create(ctx, resp.Body[0])
   // return nil
    //return err
    //파싱 이후 데이터 저장
    //return h.repo.Create(ctx, resp.Response.Body.Item....)
}
