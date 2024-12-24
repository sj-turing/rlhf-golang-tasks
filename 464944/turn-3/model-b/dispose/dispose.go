package dispose

import (
	"fmt"
	"sync"
)

type Disposable interface {
	Close() error
}

type Disposer struct {
	disposables []Disposable
	sync.Mutex
}

func (d *Disposer) Add(ds ...Disposable) {
	d.Lock()
	d.disposables = append(d.disposables, ds...)
	d.Unlock()
}

func (d *Disposer) Dispose() error {
	d.Lock()
	defer d.Unlock()

	var errs []error
	for _, d := range d.disposables {
		if err := d.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	d.disposables = nil // Clear the disposables

	if len(errs) > 0 {
		return fmt.Errorf("%d error(s) occurred during disposal", len(errs))
	}
	return nil
}
