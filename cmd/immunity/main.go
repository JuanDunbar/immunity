package main

import (
	"context"

	"github.com/benthosdev/benthos/v4/public/service"

	_ "github.com/benthosdev/benthos/v4/public/components/all"

	_ "github.com/juandunbar/immunity/processors/rules"
)

func main() {
	service.RunCLI(context.Background())
}
