package main

import (
	"container/heap"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Circuits struct {
	Distance float64
	Boxes    [2][3]int
}
type Family struct {
	parent   [3]int
	children [][3]int
}

type CHeap []Circuits

func Day8Part1(data []string, problem string) int {
	answer := 0
	properData := inputToInt(data)
	h := &CHeap{}
	heap.Init(h)
	for i := 0; i < len(properData); i++ {
		for j := i + 1; j < len(properData); j++ {
			box1 := properData[i]
			box2 := properData[j]
			xDistance := math.Pow(float64(box1[0]-box2[0]), 2)
			yDistance := math.Pow(float64(box1[1]-box2[1]), 2)
			zDistance := math.Pow(float64(box1[2]-box2[2]), 2)
			totalDistance := math.Sqrt(xDistance + yDistance + zDistance)
			newCircuit := Circuits{
				Distance: totalDistance,
				Boxes:    [2][3]int{box1, box2},
			}
			if strings.Contains(problem, "Example") && h.Len() < 10 {
				heap.Push(h, newCircuit)
			} else if strings.Contains(problem, "Problem") && h.Len() < 1000 {
				heap.Push(h, newCircuit)
			} else if newCircuit.Distance < h.Peek() {
				heap.Pop(h)
				heap.Push(h, newCircuit)
			}

		}
	}
	familyMap := map[[3]int]Family{}
	for h.Len() > 0 {
		circut := heap.Pop(h).(Circuits)
		box1, box2 := circut.Boxes[0], circut.Boxes[1]
		_, b1Exist := familyMap[box1]
		_, b2Exist := familyMap[box2]

		switch {
		case !b1Exist && !b2Exist: //neither exist
			child := [][3]int{box1, box2}
			familyMap[box1] = Family{
				parent:   box1,
				children: child,
			}
			familyMap[box2] = Family{parent: box1}
		case b1Exist && b2Exist: //both exist
			b1Parent := findParent(familyMap, familyMap[box1].parent)
			b2Parent := findParent(familyMap, familyMap[box2].parent)
			if b1Parent == b2Parent {
				continue
			}
			b1ChildLen := len(familyMap[b1Parent].children)
			b2ChildLen := len(familyMap[b2Parent].children)
			if b1ChildLen < b2ChildLen {
				b1Parent, b2Parent = b2Parent, b1Parent
			}
			b1Family := familyMap[b1Parent]
			b2Family := familyMap[b2Parent]
			b1Family.children = append(b1Family.children, b2Family.children...)
			b2Family.parent = b1Parent
			b2Family.children = [][3]int{}
			familyMap[b1Parent] = b1Family
			familyMap[b2Parent] = b2Family
		case b1Exist || b2Exist: //one of them exist
			if b2Exist {
				box1, box2 = box2, box1
			}
			b1Parent := findParent(familyMap, familyMap[box1].parent)
			b1Family := familyMap[b1Parent]
			b1Family.children = append(b1Family.children, box2)
			familyMap[b1Parent] = b1Family
			familyMap[box2] = Family{parent: box1}
		}
	}
	circutSize := []int{}
	for _, key := range familyMap {
		size := len(key.children)
		circutSize = append(circutSize, size)
	}
	sort.Slice(circutSize, func(i, j int) bool {
		return circutSize[i] > circutSize[j]
	})
	answer = circutSize[0] * circutSize[1] * circutSize[2]
	return answer
}

func findParent(familymap map[[3]int]Family, box [3]int) [3]int {
	curParent := familymap[box].parent
	if curParent != box {
		return findParent(familymap, curParent)
	}
	return curParent
}

//	func abs(x, y int) int {
//		answer := x - y
//		if answer < 0 {
//			return answer * -1
//		}
//		return answer
//	}
func inputToInt(data []string) [][3]int {
	properData := [][3]int{}
	for _, line := range data {
		stringArray := strings.Split(line, ",")
		tempArray := [3]int{}
		for i, num := range stringArray {
			number, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Problem converting string to number %v", err)
			}
			tempArray[i] = number
		}
		properData = append(properData, tempArray)
	}
	return properData
}

func (h CHeap) Len() int           { return len(h) }
func (h CHeap) Less(i, j int) bool { return h[i].Distance > h[j].Distance }
func (h CHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *CHeap) Push(x any) {
	*h = append(*h, x.(Circuits))
}
func (h *CHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h CHeap) Peek() float64 {
	if h.Len() == 0 {
		return 0
	}
	return h[0].Distance
}
