// Package roundrobin ...
package roundrobin

import (
	"errors"
	"sync/atomic"
)

// ErrorNoObjectsProvided is the error that occurs when no objects are provided.
var ErrorNoObjectsProvided = errors.New("no objects provided")

// RoundRobin is an interface for representing round-robin balancing.
type RoundRobin[O any] interface {
	Next() *O
}

type roundrobin[O any] struct {
	objects []*O
	next    uint32
}

// New returns RoundRobin implementation with roundrobin.
func New[O any](objects ...*O) (RoundRobin[O], error) {
	if len(objects) == 0 {
		return nil, ErrorNoObjectsProvided
	}

	return &roundrobin[O]{
		objects: objects,
	}, nil
}

// Next returns the next object.
func (r *roundrobin[O]) Next() *O {
	n := atomic.AddUint32(&r.next, 1)
	return r.objects[(int(n)-1)%len(r.objects)]
}
