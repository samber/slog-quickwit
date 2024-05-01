package main

import (
	"fmt"
	"time"

	"log/slog"

	"github.com/samber/go-quickwit"
	slogquickwit "github.com/samber/slog-quickwit"
)

func main() {
	// docker-compose up -d
	// curl -X POST \
	//     'http://localhost:7280/api/v1/indexes' \
	//     -H 'Content-Type: application/yaml' \
	//     --data-binary @test-config.yaml

	client := quickwit.NewWithDefault("http://localhost:7280")
	defer client.Stop() // flush and stop

	logger := slog.New(slogquickwit.Option{Level: slog.LevelDebug, Client: client}.NewQuickwitHandler())
	logger = logger.With("release", "v1.0.0")

	logger.
		With(
			slog.Group("user",
				slog.String("id", "user-123"),
				slog.Time("created_at", time.Now().AddDate(0, 0, -1)),
			),
		).
		With("environment", "dev").
		With("error", fmt.Errorf("an error")).
		Error("A message")
}
