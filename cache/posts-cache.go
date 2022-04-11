package cache

type PostCache struct {
	mc ManagerCache
}
type ManagerCache interface {
	Set(key string, value interface{})
	Get(key string) *interface{}
	Del(key string)
}

func NewPostCache(m ManagerCache) *PostCache {
	return &PostCache{mc: m}
}

func (p *PostCache) Set(key string, value interface{}) {
	p.mc.Set(key, value)
}

func (p *PostCache) Get(key string) *interface{} {
	return p.mc.Get(key)
}
func (p *PostCache) Del(key string) {
	p.mc.Del(key)
}
