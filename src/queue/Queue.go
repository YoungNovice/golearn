package queue

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *Queue) PushInt(v interface{}) {
	*q = append(*q, v.(int))
}

// interface.(type) 转类型
func (q *Queue) PopInt() int {
	v := (*q)[0]
	*q = (*q)[1:]
	return v.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
