package roadie

import "io"

type Entity interface {
	Write(writer io.Writer) (err error)
}
