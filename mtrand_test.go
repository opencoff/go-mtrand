// mtrand_test.go

package mtrand_test

import (
	"testing"

	"mtrand"
)

func Test0(t *testing.T) {

	m := mtrand.New(1)

	for i := 0; i < 1000; i++ {
		z := m.Next()
		t.Logf("%v\n", z)
	}

}
