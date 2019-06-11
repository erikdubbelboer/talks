package sync

type Pool struct {
	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
	New func() interface{}
}

func (p *Pool) Get() interface{}

func (p *Pool) Put(x interface{})
