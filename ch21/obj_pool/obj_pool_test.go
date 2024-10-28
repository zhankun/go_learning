package obj_pool

import "testing"

func TestObjPool(t *testing.T) {
	pool := InitObjPool(10)
	for i := 0; i < 12; i++ {
		if obj, err := pool.GetObj(1); err != nil {
			t.Error(err)
		} else {
			t.Logf("获取到的池子对象是,%x", obj)
			// 使用结束,释放
			//if err := pool.ReleaseObj(obj); err != nil {
			//	t.Error(err)
			//}
		}
	}
	t.Log("Done")
}
