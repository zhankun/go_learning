package error_test

import (
	"learning/ch14/series"
	"testing"
)

func TestPackage(t *testing.T) {
	ret := series.GetFiSerie(5)
	t.Log(ret)
}
