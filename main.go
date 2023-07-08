package main

import (
	"context"

	"github.com/ayhamal/gogql-boilerplate/api"
	"github.com/ayhamal/gogql-boilerplate/env"
	"github.com/ayhamal/gogql-boilerplate/pkg/pg"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			env.New,
			context.Background,
			pg.New,
		),
		fx.Decorate(),
		fx.Invoke(
			api.ConfigLifeCycleHooks,
			api.ConfigGqlLifeCycleHooks,
		),
	)

	app.Run()
}
