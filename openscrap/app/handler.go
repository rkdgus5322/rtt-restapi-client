package app

import (
	"context"
	"encoding/json"
	"openscrap/domain"

	"github.com/go-resty/resty/v2"
)

type Handler struct {
	cli  *resty.Client
	repo *Repository
	key  string
}

func NewHandler(cli *resty.Client, repo *Repository, serviceKey string) *Handler {
	return &Handler{cli: cli, repo: repo, key: serviceKey}
}

func (h *Handler) Scrap(ctx context.Context) error {
	// api path 추가
	res, err := h.cli.R().Get("/api/service/acd/acdSagoCaseInfo?serviceKey=" + h.key)
	if err != nil {
		return err
	}

	//json 파싱을 진행
	var model domain.Model
	if err := json.Unmarshal(res.Body(), &model); err != nil {
		return err
	}

	//파싱 이후 데이터 저장
	return h.repo.Create(ctx, model.Response.Body.Items...)
}
