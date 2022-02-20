//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go-geektime/fourth-week/internal/biz"
	"go-geektime/fourth-week/internal/conf"
	"go-geektime/fourth-week/internal/data"
	"go-geektime/fourth-week/internal/server"
	"go-geektime/fourth-week/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
