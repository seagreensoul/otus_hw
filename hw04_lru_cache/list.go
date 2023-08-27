package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len       int
	FirstItem *ListItem
	LastItem  *ListItem
}

func (dl list) Len() int {
	return dl.len
}

func (dl list) Front() *ListItem {
	if dl.len == 0 {
		return nil
	}

	return dl.FirstItem
}

func (dl list) Back() *ListItem {
	if dl.len == 0 {
		return nil
	}

	return dl.LastItem
}

func (dl *list) PushFront(v interface{}) *ListItem {
	if dl.len == 0 {
		li := ListItem{v, nil, nil}
		dl.LastItem = &li
		dl.FirstItem = &li
	} else {
		li := ListItem{v, dl.FirstItem, nil}
		dl.FirstItem.Prev = &li
		dl.FirstItem = &li
	}
	dl.len++

	return dl.FirstItem
}

func (dl *list) PushBack(v interface{}) *ListItem {
	if dl.len == 0 {
		li := ListItem{v, nil, nil}
		dl.LastItem = &li
		dl.FirstItem = &li
	} else {
		li := ListItem{v, nil, dl.LastItem}
		dl.LastItem.Next = &li
		dl.LastItem = &li
	}
	dl.len++

	return dl.LastItem
}

func (dl *list) Remove(i *ListItem) {
	i.Next.Prev = i.Prev
	i.Prev.Next = i.Next
	dl.len--
}

func (dl *list) MoveToFront(i *ListItem) {
	dl.Remove(i)
	dl.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
