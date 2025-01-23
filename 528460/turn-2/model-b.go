package main  
import (  
    "bytes"  
    "fmt"  
    "io"  
    "io/ioutil"  
    "log"  
    "os"  
)  
  
// Define a Huffman node struct  
type huffmanNode struct {  
    char byte  
    freq int  
    left *huffmanNode  
    right *huffmanNode  
}  
  
// A priority queue using Huffman node  
type priorityQueue []*huffmanNode  
  
func (pq priorityQueue) Len() int {  
    return len(pq)  
}  
  
func (pq priorityQueue) Less(i, j int) bool {  
    return pq[i].freq < pq[j].freq  
}  
  
func (pq *priorityQueue) Push(x interface{}) {  
    *pq = append(*pq, x.(*huffmanNode))  
}  
  
func (pq *priorityQueue) Pop() interface{} {  
    n := len(*pq)  
    x := (*pq)[n-1]  
    *pq = (*pq)[0 : n-1]  
    return x  
}  
  
// Build a Huffman tree  
func buildHuffmanTree(freq map[byte]int) *huffmanNode {  
    pq := priorityQueue(make([]*huffmanNode, 0, len(freq)))  
    for char, count := range freq {  
        pq.Push(&huffmanNode{char: char, freq: count})  
    }  
    
    for pq.Len() > 1 {  
        left := pq.Pop().(*huffmanNode)  
        right := pq.Pop().(*huffmanNode)  
        parent := &huffmanNode{left: left, right: right, freq: left.freq + right.freq}  
        pq.Push(parent)  
    }  
    
    return pq.Pop().(*huffmanNode)  
}  
  
// Create a Huffman code for each character based on the Huffman tree  
func buildHuffmanCodes(root *huffmanNode, codes map[byte][]byte, prefix []byte) {  
    if root == nil {  
        return  
    }  
    
    prefix = append(prefix, root.char)  
    if root.left == nil && root.right == nil {  
        codes[root.char] = make([]byte, len(prefix))  
        copy(codes[root.char], prefix)  
    }  
    
    buildHuffmanCodes(root.left, codes, prefix)  
    buildHuffmanCodes(root.right, codes, prefix)  
}  
  
// Compress data using Huffman codes  
func compress(data []byte, codes map[byte][]byte) []byte {  
    compressedData := []byte{}  
    for _, char := range data {  
        code := codes[char]  
        compressedData = append(compressedData, code...)  
    }  
    return compressedData  
}  
  
// Decompress data using the Huffman tree  
func decompress(compressedData []byte, root *huffmanNode) []byte {  
    current := root  
    decompressedData := []byte{}  
    for _, bit := range compressedData {  
        if bit == 0 {  
            current = current.left  
        } else {  
            current = current.right  
        }  
        if current.left == nil && current.right == nil {  
            decompressedData = append(decompressedData, current.char)  
            current = root  
        }  
    }  
    return decompressedData  
}  
  
func main() {  
    filename := "example.txt"  
    data, err := ioutil.ReadFile(filename)  
    if err != nil {  
        log.Fatal(err)  
    }
