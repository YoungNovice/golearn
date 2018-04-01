package queue

type Queue []int

func (q *Queue) Push(v int)  {
	*q = append(*q, v)
}

func (q *Queue) Pop() int  {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *Queue) IsEmpty() bool  {
	return len(*q) == 0
}
