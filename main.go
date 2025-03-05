package main

import (
	"fmt"
	"slot_tree/tree"
	"time"
)

var (
	// lineSetting = [][]int{{1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {2, 2, 2, 2, 2}, {0, 1, 2, 1, 0}, {2, 1, 0, 1, 2}}
	lineSetting = map[int][]int{0: {1, 1, 1, 1, 1}, 1: {0, 0, 0, 0, 0}, 2: {2, 2, 2, 2, 2}, 3: {0, 1, 2, 1, 0}, 4: {2, 1, 0, 1, 2}}
	reel        = [][]string{{"A", "A", "A"}, {"A", "A", "A"}, {"A", "B", "B"}, {"B", "B", "B"}, {"C", "C", "C"}}
	reel1       = [][]string{{"B", "B", "B"}, {"B", "B", "B"}, {"B", "B", "B"}, {"B", "B", "B"}, {"B", "B", "B"}}
	reel2       = [][]string{{"A", "B", "C"}, {"A", "B", "C"}, {"A", "B", "C"}, {"A", "B", "C"}, {"A", "B", "C"}}
	reel3       = [][]string{{"A", "A", "A"}, {"B", "B", "B"}, {"B", "B", "B"}, {"C", "C", "C"}, {"C", "C", "C"}}

	symbol = []string{}
	count  = []int{}

	slot = &tree.SlotData{
		PayLineMap: lineSetting,
		Cache:      make(map[string]tree.ResultTree),
	}
)

func main() {
	slot.TreeInit()
	slot.Node.PrintTreeHorizontal("", true)

	time1 := time.Now()
	slot.Node.ReplaceReel(reel)
	slot.PayData(slot.Node.TraverseLengthTree())
	fmt.Println("-----------reel end")
	slot.Node.ReplaceReel(reel1)
	slot.PayData(slot.Node.TraverseLengthTree())
	fmt.Println("-----------reel1 end")
	slot.Node.ReplaceReel(reel2)
	slot.PayData(slot.Node.TraverseLengthTree())
	fmt.Println("-----------reel2 end")
	slot.Node.ReplaceReel(reel3)
	slot.PayData(slot.Node.TraverseLengthTree())
	fmt.Println("-----------reel3 end")
	fmt.Println(time.Since(time1))
	time1 = time.Now()
	slot.Node.ReplaceReel(reel)
	slot.PayData(slot.Node.TraverseLengthTree())
	fmt.Println("-----------reel end")
	slot.Node.ReplaceReel(reel1)
	slot.PayData(slot.Node.TraverseLengthTree())
	fmt.Println("-----------reel1 end")
	slot.Node.ReplaceReel(reel2)
	slot.PayData(slot.Node.TraverseLengthTree())
	fmt.Println("-----------reel2 end")
	slot.Node.ReplaceReel(reel3)
	slot.PayData(slot.Node.TraverseLengthTree())
	fmt.Println("-----------reel3 end")
	fmt.Println(time.Since(time1))
}
