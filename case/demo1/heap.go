package demo1

type Heap struct {
	val   []int
	cap   int
	count int
}

func NewHeap(cap int) *Heap {
	return &Heap{val: make([]int, cap),
		cap:   cap,
		count: 0}
}

//[ 10, 7, 2, 5, 1 ]
func (h *Heap) CreateMaxHeap(val int) {
	if h.count == h.cap {
		return
	}
	h.val[h.count] = val
	i := h.count
	parent := (i - 1) / 2
	for parent >= 0 && i-1 >= 0 {
		if h.val[parent] < h.val[i] {
			h.val[parent], h.val[i] = h.val[i], h.val[parent]
		}
		i = parent
		parent = (i - 1) / 2
	}
	h.count++
}

//调整堆,根节点与孩子结点中的最大值进行比较，进行交换
func (h *Heap) Adjust() {
	top := 0
	//j 为左孩子或者右孩子
	j := 2*top + 1
	for j < h.count{
		for j+ 1 < h.count && h.val[j+1]>h.val[j]{
			j++
		}
		if h.val[j] > h.val[top]{
			h.val[top],h.val[j] = h.val[j],h.val[top]
			top = j
			j = 2*top+1
		}else{
			break
		}
	}
}

//删跟节点
func (h *Heap) Del(){
	root :=0
	h.val[root] ,h.val[h.count-1] = h.val[h.count-1],h.val[root]
	h.val[h.count-1] = 0
	h.count--
	h.Adjust()
}