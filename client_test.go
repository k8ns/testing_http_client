package resttest

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	go func() {
		err := NewServerApiMock().Run(":8181")
		if err != nil {
			panic(err)
		}
	}()

	os.Exit(m.Run())
}




