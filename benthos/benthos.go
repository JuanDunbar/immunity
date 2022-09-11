package benthos

import (
	"context"
	"os"

	_ "github.com/benthosdev/benthos/v4/public/components/all"
	"github.com/benthosdev/benthos/v4/public/service"
)

func RunStream(ctx context.Context) error {
	builder := service.NewStreamBuilder()
	// get our yaml config
	config, err := os.ReadFile("./benthos/benthos.yaml")
	if err != nil {
		return err
	}
	// load yaml into builder
	err = builder.SetYAML(string(config))
	if err != nil {
		return err
	}
	// build our data stream
	stream, err := builder.Build()
	if err != nil {
		return err
	}
	// start our stream
	return stream.Run(ctx)
}
