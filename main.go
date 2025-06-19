package main

import (
	"log/slog"
	"os"

	// This will fail, as there is no such package to import
	"demo-main/pkg/foo"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	foobar := foo.New("baz")
	logger.Info("Created FooBar", "content", foobar.Content)
}
