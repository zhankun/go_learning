package remote_test

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestRemote(t *testing.T) {
	m := cm.CreateConcurrentMap(10)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
	t.Log(m.Get(cm.StrKey("jh")))

}
