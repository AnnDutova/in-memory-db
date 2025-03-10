package service

import "context"

type Servicer interface {
	Run(context.Context) error
}
