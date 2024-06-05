// Package roundrobin
package roundrobin

import (
	"fmt"
	"reflect"
	"testing"
)

type resource struct {
	id   int
	name string
}

func TestRoundRobin(t *testing.T) {
	tests := []struct {
		resources []*resource
		iserr     bool
		expected  []string
		want      []*resource
	}{
		{
			resources: []*resource{
				{1, "resource-1"},
				{2, "resource-2"},
				{3, "resource-3"},
				{4, "resource-4"},
				{5, "resource-5"},
				{6, "resource-6"},
				{7, "resource-7"},
			},
			iserr: false,
			want: []*resource{
				{1, "resource-1"},
				{2, "resource-2"},
				{3, "resource-3"},
				{4, "resource-4"},
				{5, "resource-5"},
				{6, "resource-6"},
				{7, "resource-7"},
				{1, "resource-1"},
				{2, "resource-2"},
				{3, "resource-3"},
				{4, "resource-4"},
				{5, "resource-5"},
				{6, "resource-6"},
				{7, "resource-7"},
				{1, "resource-1"},
				{2, "resource-2"},
				{3, "resource-3"},
				{4, "resource-4"},
				{5, "resource-5"},
				{6, "resource-6"},
				{7, "resource-7"},
				{1, "resource-1"},
				{2, "resource-2"},
				{3, "resource-3"},
			},
		},
		{
			resources: []*resource{},
			iserr:     true,
			want:      []*resource{},
		},
	}

	for i, test := range tests {
		rr, err := New(test.resources...)
		if err != nil && !test.iserr {
			t.Errorf("tests[%d] - RoundRobin iserr is wrong. want: %v, but got: %v", i, test.want, true)
		}

		for j := 0; j < len(test.want); j++ {
			got := rr.Next()
			if !reflect.DeepEqual(got, test.want[j]) {
				t.Errorf("tests[%d] index[%d] - RoundRobin is wrong. want: %v, got: %v", i, j, *test.want[j], *got)
			}
		}
	}
}

func BenchmarkRoundRobinSync(b *testing.B) {
	resources := []*struct {
		id   int
		name string
	}{
		{1, "resource-1"},
		{2, "resource-2"},
		{3, "resource-3"},
		{4, "resource-4"},
		{5, "resource-5"},
		{6, "resource-6"},
		{7, "resource-7"},
	}

	for i := 1; i < len(resources)+1; i++ {
		b.Run(fmt.Sprintf("RoundRobinSize(%d)", i), func(b *testing.B) {
			rr, err := New(resources[:i]...)
			if err != nil {
				b.Fatal(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				rr.Next()
			}
		})
	}
}

func BenchmarkRoundRobinASync(b *testing.B) {
	resources := []*struct {
		id   int
		name string
	}{
		{1, "resource-1"},
		{2, "resource-2"},
		{3, "resource-3"},
		{4, "resource-4"},
		{5, "resource-5"},
		{6, "resource-6"},
		{7, "resource-7"},
	}

	for i := 1; i < len(resources)+1; i++ {
		b.Run(fmt.Sprintf("RoundRobinSize(%d)", i), func(b *testing.B) {
			rr, err := New(resources[:i]...)
			if err != nil {
				b.Fatal(err)
			}
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					rr.Next()
				}
			})
		})
	}
}
