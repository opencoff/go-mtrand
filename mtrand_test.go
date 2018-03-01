// mtrand_test.go

package mtrand

import (
	"testing"
)

func Test0(t *testing.T) {

	m := New(1)

	for i := 0; i < 1000; i++ {
		z := m.Next()
		t.Logf("%v\n", z)
	}

}
