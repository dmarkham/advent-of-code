package main

import (
	"container/ring"
	"flag"
	"fmt"
	"sort"
)

var players int
var maxScore int

func init() {
	flag.IntVar(&players, "Players", 9, "Players")
	flag.IntVar(&maxScore, "Marbles", 23, "Max Score")
}

func main() {
	flag.Parse()
	r := ring.New(1)
	scores := make(map[int]int)
	// Get the length of the ring
	n := r.Len()
	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = 0
		r = r.Next()
	}

	marbel := 1
DONE:
	for true {

		for player := 1; player < players+1; player++ {

			if marbel%23 == 0 {
				scores[player] += marbel
				for i := 0; i < 9; i++ {
					r = r.Prev()
				}
				v := r.Unlink(1)
				//fmt.Println("Removed: ", v.Value)
				scores[player] += v.Value.(int)
				if marbel == maxScore {
					break DONE
				}
				r = r.Next().Next()
				//fmt.Println("Active:", r.Value)
				marbel++
				continue
			}

			l := ring.New(1)
			l.Value = marbel
			r = r.Link(l)
			if marbel == maxScore {
				break DONE
			}
			marbel++
		}
	}

	// Iterate through the ring and print its contents
	r.Do(func(p interface{}) {
		//fmt.Println(p.(int))
	})

	fmt.Println(scores)

	var ss []kv
	for k, v := range scores {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool { // Then sorting the slice by value, higher first.
		return ss[i].Value > ss[j].Value
	})
	for _, kv := range ss[:1] { // Print the x top values
		fmt.Printf("Type:: %v    MinCount: %v\n", kv.Key, kv.Value)
	}

}

type kv struct {
	Key   int
	Value int
}
