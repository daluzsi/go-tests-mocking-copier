package copier

import "github.com/jinzhu/copier"

type Copier interface {
	Copy(dst, src interface{}) error
}

type Impl struct{}

func NewCopier() Copier {
	return &Impl{}
}

func (c *Impl) Copy(dst, src interface{}) error {
	return copier.Copy(dst, src)
}
