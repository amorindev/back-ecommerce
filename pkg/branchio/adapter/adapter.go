package adapter

import "com.fernando/pkg/branchio/port"

var _ port.Adapter = &Adapter{}

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}
