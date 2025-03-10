package compute

import (
	"context"
	"errors"
	"io"
)

type (
	Command   string
	Arguments []string

	Compute interface {
		Parse(ctx context.Context, in io.Reader, out io.Writer) error
	}
)

var (
	ErrEmptyQuerry           = errors.New("empty querry")
	ErrInvalidCommand        = errors.New("invalid command")
	ErrInvalidArgumentsCount = errors.New("invalid arguments count")
)
