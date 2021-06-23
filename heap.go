package zutil
import ("fmt")

type heap struct {
    xs []int
}
// MakeHeap will make a min heap from values in xs. TC: O(n)
func MakeHeap(xs []int) *heap {
    var h heap
    h.xs = make([]int,len(xs) +1)
    for i,v := range xs {
        h.xs[i+1] = v
    }

    // As in heap half of the node will be leaf node so bubbling down on non-leaf nodes.
    for i := len(xs) / 2; i>0 ;i--{
        h.bubbleDown(i)
    }
    return &h
}
func (h *heap) Len() int {
    return len(h.xs) - 1
}
func (h *heap) Insert(x int){
    h.xs = append(h.xs)
    h.bubbleUp(len(h.xs) - 1)
}
// Min: get the minimum value of the heap
// do not remove value from heap; TC := O(1)
func (h *heap) Min() (int, bool) {
    if h.Len() == 0{
        return 0, false
    }
    return h.xs[1], true
}
// ExtractMin: get minimum value of heap
// and remove value from heap;;TC:=O(logn)

func (h *heap) ExtractMin() (int,bool){
    if h.Len() == 0 {
        return 0, false
    }
    min := h.xs[1]
    //Take last element and make it root then bublle it down.
    h.xs[1] = h.xs[h.Len()]
    h.xs = h.xs[:h.Len()]
    h.bubbleDown(1)
    return min,true
}
func (h *heap) String() string {
    return fmt.Sprintf("Min-heap %v",h.xs)
}
/*
private helper fucntions
*/
//TC:O(log(k))
func (h *heap) bubbleUp(k int){
    if h.Len() == 1{
        return
    }
    parentIdx,ok := parentIndex(k)
    if parentIdx == 0 || !ok{
        return
    }
    if h.xs[parentIdx] > h.xs[k]{
        //Swap element
        h.xs[parentIdx], h.xs[k] = h.xs[k], h.xs[parentIdx]
        h.bubbleUp(parentIdx)
    }
}
//TC:O(log(k))
func (h *heap) bubbleDown(k int){
    // COmapre k,left and right child if k is min do nothing else bubbleDown recursively on min.
    l := leftChild(k)
    r := rightChild(k)
    if l > h.Len(){
        return
    }
    if r > h.Len()  {
        if h.xs[l] < h.xs[k] {
            h.xs[l],h.xs[k] = h.xs[k],h.xs[l]
            h.bubbleDown(l)
        }
    }
    //if l and r both present then fin min and do bubbledown
    if h.xs[k] > h.xs[l] ||  h.xs[k] > h.xs[r]{
        if  h.xs[l] < h.xs[r]{
            h.xs[l],h.xs[k] = h.xs[k],h.xs[l]
            h.bubbleDown(l)
        }else{
            h.xs[r],h.xs[k] = h.xs[k],h.xs[r]
            h.bubbleDown(r)
        }
    }
}
//Get index of parent of node at x.
//TC:O(1)
func parentIndex(k int)(int, bool){
    if k == 1{
        return 0,false
    }
    return k/2,true
}
//Get index of left child of indexk k
//TC:O(1)
func leftChild(k int)int{
    return k * 2
}
//Get index of righ child of indexk k
//TC:O(1)
func rightChild(k int)int{
    return k * 2 + 1
}
