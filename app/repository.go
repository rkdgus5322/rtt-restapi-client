package app

import (
	"context"

	"gorm.io/gorm"

    "openscrap/domain"
)

// Repository 데이터 베이스 저장을 위한 레포지토리
type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Create 모델이 들어올경우 레코드 생성
// 한개가 아닐수 있으니 배열 형태로 수신
func (r *Repository) Create(ctx context.Context, model ...domain.Body) error {
	return r.db.WithContext(ctx).Create(model).Error
}
//func (r *Repository) Create(ctx context.Context, model ...domain.Body) error {
//    return r.db.WithContext(ctx).Create(model).Error
//}
