package service

import (
	"github.com/koopa0/blog-service/global"
	"github.com/koopa0/blog-service/internal/dao"
	"golang.org/x/net/context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
