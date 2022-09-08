package main

import (
	"context"

	"github.com/benthosdev/benthos/v4/public/service"

	_ "github.com/benthosdev/benthos/v4/public/components/all"

	// import our custom processor
	_ "github.com/juandunbar/immunity/processors/rules"
)

func main() {
	service.RunCLI(context.Background())
}
