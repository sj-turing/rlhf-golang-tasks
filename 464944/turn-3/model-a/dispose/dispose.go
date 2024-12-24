// Dispose package
package dispose

import (
	"fmt"
	"sync"
)

// Define a custom cleanup protocol
type Cleanup interface {
	Close() error
}

type Disposer struct {
	mu      sync.RWMutex
	closers []Cleanup
}

func New() *Disposer {
	return &Disposer{}
}

func (d *Disposer) Register(closer Cleanup) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.closers = append(d.closers, closer)
}

func (d *Disposer) Deregister(closer Cleanup) {
	d.mu.Lock()
	defer d.mu.Unlock()
	for i, c := range d.closers {
		if c == closer {
			d.closers = append(d.closers[:i], d.closers[i+1:]...)
			return
		}
	}
}

func (d *Disposer) Dispose() error {
	var errs []error
	d.mu.RLock()
	defer d.mu.RUnlock()
	for _, c := range d.closers {
		err := c.Close()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return multiError(errs)
}

func multiError(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	return fmt.Errorf("multiple errors: %v", errs)
}
