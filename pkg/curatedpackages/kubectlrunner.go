package curatedpackages

import (
	"bytes"
	"context"
)

type KubectlRunner interface {
	ExecuteCommand(ctx context.Context, opts ...string) (bytes.Buffer, error)
	CreateFromYaml(ctx context.Context, yaml []byte, opts ...string) (bytes.Buffer, error)
}
