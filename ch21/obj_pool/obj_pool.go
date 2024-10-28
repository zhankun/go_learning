package obj_pool

import (
	"errors"
	"time"
)

type ReusableObj struct {
	i int
}

type ObjPool struct {
	bufChan chan *ReusableObj
}

// InitObjPool 初始化对象池
func InitObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{i: i}
	}
	return &objPool
}

// GetObj 从池子中获取对象
func (p *ObjPool) GetObj(duration time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(duration):
		return nil, errors.New("time out")
	}
}

// ReleaseObj 归还对象到池子中
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
