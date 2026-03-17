package comp

type FContext struct {
	C map[string]*GParam
}

func (f *FContext) getParam(key string) interface{} {
	return f.C[key]
}

func (f *FContext) getStr(key string) string {
	param := f.C[key]

	return param.Value.(string)
}

func (f *FContext) getInt(key string) int {
	param := f.C[key]

	return param.Value.(int)
}

type GParam struct {
	Key   string
	Value interface{}
}
