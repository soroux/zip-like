package compressor

import (
	"container/heap"
)

type HuffmanNode struct {
	Char  byte
	Freq  int
	Left  *HuffmanNode
	Right *HuffmanNode
}

type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].Freq < pq[j].Freq }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*HuffmanNode)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func BuildHuffmanTree(data []byte) *HuffmanNode {
	freq := make(map[byte]int)
	for _, b := range data {
		freq[b]++
	}

	pq := make(PriorityQueue, 0)
	for char, f := range freq {
		heap.Push(&pq, &HuffmanNode{Char: char, Freq: f})
	}

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*HuffmanNode)
		right := heap.Pop(&pq).(*HuffmanNode)
		node := &HuffmanNode{
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}
		heap.Push(&pq, node)
	}

	return heap.Pop(&pq).(*HuffmanNode)
}

func buildHuffmanCodes(node *HuffmanNode, prefix string, codes map[byte]string) {
	if node.Left == nil && node.Right == nil {
		codes[node.Char] = prefix
		return
	}
	if node.Left != nil {
		buildHuffmanCodes(node.Left, prefix+"0", codes)
	}
	if node.Right != nil {
		buildHuffmanCodes(node.Right, prefix+"1", codes)
	}
}

func CompressHuffman(data []byte) (string, map[byte]string) {
	root := BuildHuffmanTree(data)
	codes := make(map[byte]string)
	buildHuffmanCodes(root, "", codes)

	var result string
	for _, b := range data {
		result += codes[b]
	}
	return result, codes
}

func DecompressHuffman(encoded string, codes map[byte]string) ([]byte, error) {
	reverseCodes := make(map[string]byte)
	for k, v := range codes {
		reverseCodes[v] = k
	}

	var result []byte
	var currentCode string
	for _, b := range encoded {
		currentCode += string(b)
		if char, found := reverseCodes[currentCode]; found {
			result = append(result, char)
			currentCode = ""
		}
	}
	return result, nil
}
