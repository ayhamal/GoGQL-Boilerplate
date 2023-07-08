package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ayhamal/gogql-boilerplate/api/middleware"
	"github.com/ayhamal/gogql-boilerplate/env"
	"github.com/ayhamal/gogql-boilerplate/graph"
	"github.com/ayhamal/gogql-boilerplate/graph/model"
	"github.com/ayhamal/gogql-boilerplate/pkg/pg"
	"github.com/gorilla/websocket"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	"go.uber.org/fx"
)

// Config Life Cycle definition Hook
func ConfigLifeCycleHooks(lc fx.Lifecycle, _ *pg.PgClient) {
	lc.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				log.Println("[INFO] Starting app configuration...")
				return nil
			},
			OnStop: func(_ context.Context) error {
				log.Println("[INFO] Closing database connection...")
				pgClient, err := pg.GetInstance()
				// Handle get instance error
				if err != nil {
					return err
				}
				// // TODO: Drop database tables only in test environment
				// log.Println("[WARN] Dropping database tables...")
				// pgClient.DropDatabaseTables()
				// Close database connection
				pgClient.CloseConnection()
				// Log desconection success
				log.Println("[INFO] Database connection closed...")
				return nil //db.CloseConnection()
			},
		},
	)
}

// Config GQL Life Cycle definition Hook
func ConfigGqlLifeCycleHooks(lc fx.Lifecycle, _ *pg.PgClient, env *env.Env) {
	lc.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				log.Println("[INFO] Starting GQL Server...")
				// Build cors middleware handler
				c := middleware.BuildCorsHandler()
				// Use New instead of NewDefaultServer in order to have full control over defining transports
				srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
					ChatMessages:  []*model.Message{},
					ChatObservers: map[string]chan []*model.Message{},
				}}))
				srv.AddTransport(transport.POST{})
				srv.AddTransport(transport.Websocket{
					KeepAlivePingInterval: 10 * time.Second,
					Upgrader: websocket.Upgrader{
						CheckOrigin: func(r *http.Request) bool {
							return true
						},
					},
				})
				srv.Use(extension.Introspection{})

				http.Handle("/", playground.Handler("GraphQL playground", "/query"))
				http.Handle("/query", c.Handler(srv))

				log.Printf("connect to http://localhost:%d/ for GraphQL playground", env.App.Port)
				go http.ListenAndServe(fmt.Sprintf(":%d", env.App.Port), nil)
				return nil
			},
			OnStop: func(_ context.Context) error {
				log.Println("[INFO] Stopping GQL Server...")
				return nil
			},
		},
	)
}
