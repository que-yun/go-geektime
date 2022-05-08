package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go-geektime/graduation/internal/biz"
	"go-geektime/graduation/pkg/db"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
	db   *db.GormDB
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
		db:   db.NewGormDB("mysql", ""),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}
