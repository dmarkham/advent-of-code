package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/goombaio/dag"
)

// https://adventofcode.com/2018/day/7
var part2 bool

func init() {
	flag.BoolVar(&part2, "part2", false, "Run Part2?")
}
func fileToDag(file string) *dag.DAG {
	lines := readFileToLines(file)
	d := dag.NewDAG()
	seen := make(map[string]*dag.Vertex)
	for _, l := range lines {
		ff := strings.Fields(l)
		//fmt.Println(ff[1], ff[7])
		v1, ok := seen[ff[1]]
		if !ok {
			v1 = dag.NewVertex(ff[1], nil)
			d.AddVertex(v1)
			seen[ff[1]] = v1
		}
		v2, ok := seen[ff[7]]
		if !ok {
			v2 = dag.NewVertex(ff[7], nil)
			d.AddVertex(v2)
			seen[ff[7]] = v2
		}

		d.AddEdge(v1, v2)

	}
	return d
}

type Worker struct {
	WorkerID  int
	busyUntil int
	index     int // The index of the item in the heap.
}
type PriorityQueue []*Worker

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].busyUntil < pq[j].busyUntil

}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Worker))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	//item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func main() {
	flag.Parse()
	d := fileToDag("data")
	fmt.Println(d)
	finalList := make([]*dag.Vertex, 0)
	starts := d.SourceVertices()
	sort.Sort(byID(starts))
	secondsPassed := 0
	availableWorkers := 2
	baseSeconds := 0
	_ = availableWorkers
	_ = baseSeconds
	workerQueue := make(PriorityQueue, 0)
	for i := 0; i < availableWorkers; i++ {
		workerQueue = append(workerQueue,
			&Worker{WorkerID: i},
		)
	}

	heap.Init(&workerQueue)

	//working := 0
	seen := make(map[string]bool)
	//round := 0
	//minStartthisRound := math.MaxInt64
	for true {

		if len(starts) < 1 {
			break
		}

		fmt.Println("Starts: ", starts)
		smallest := 0
		vals := make([]int, 0)
		for i := 0; i < len(starts); i++ {
			v := starts[i]
			worktime := baseSeconds + int(starts[i].ID[0]-64)
			if !seen[starts[i].ID] {
				seen[starts[i].ID] = true
				// Count work for V
				w := heap.Pop(&workerQueue).(*Worker)
				fmt.Println("V :", v.ID)
				fmt.Println("Popped:", w.WorkerID, w.busyUntil)
				if secondsPassed < w.busyUntil {
					log.Fatalf("TooSoon %v < %v\n", secondsPassed, w.busyUntil)
				}
				w.busyUntil = secondsPassed + worktime
				if secondsPassed == 0 {
					smallest = w.busyUntil

				} else if w.busyUntil < secondsPassed {
					smallest = w.busyUntil
				} else {
					//smallest = w.busyUntil
				}
				if smallest > secondsPassed {
					secondsPassed = smallest
				}
				vals = append(vals, w.WorkerID)
				fmt.Println("Pushed:", w.WorkerID, w.busyUntil, secondsPassed)
				heap.Push(&workerQueue, w)
			}
			if len(vals) > 0 {
				// of the ID used/ use the smallest busyUntill
				secondsPassed = vals[0]

			}
		}
		/*
			smallest := math.MaxInt64
			for _, w := range workerQueue {
				if w.busyUntil < smallest {
					smallest = w.busyUntil
				}
			}
			secondsPassed = smallest
		*/
		fmt.Println("SecondsPast: ", secondsPassed)
		v := starts[0]
		starts = starts[1:]
		kids, _ := d.Successors(v)
		sort.Sort(byID(kids))
		finalList = append(finalList, v)

		/*
			for workerQueue.Len() > 0 {
				w := heap.Pop(&workerQueue).(*Worker)
				fmt.Println("Popped:", w.V.ID, w.WorkerID, w.started, w.workSeconds, secondsPassed)
			}
		*/
		for _, k := range kids {
			//fmt.Println("  Kids", k)
			v.Children.Remove(k)
			k.Parents.Remove(v)
			if k.InDegree() == 0 {
				starts = append(starts, k)
				sort.Sort(byID(starts))
			}

		}
	}
	fmt.Println("Part1 : ")
	for _, v := range finalList {
		fmt.Print(v.ID)
	}
	fmt.Println("")
	fmt.Println("Part2 : Seconds ", secondsPassed)

	os.Exit(0)
}

type byID []*dag.Vertex

func (s byID) Len() int {
	return len(s)
}
func (s byID) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byID) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

/*
func foo()  {
	ped, err := d.Predecessors(v2)
	if err != nil {
		panic(err)
	}
	if len(ped) > 1 {
		SortVert(ped)

		for i := 1; i < len(ped); i++ {
			fmt.Println("Removing:", ped[i], v2)
			d.DeleteEdge(ped[i], v2)
		}
	}


}
*/

func childern(d *dag.DAG, v *dag.Vertex) {

	list, err := d.Predecessors(v)
	if err != nil {
		panic(err)
	}
	var ss []kv
	for _, v := range list {
		ss = append(ss, kv{fmt.Sprintf("%v", v.ID), v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Key > ss[j].Key
	})
	fmt.Println(" Deleteing: ", v.ID)
	for _, s := range ss {
		d.DeleteEdge(s.Value, v)
		if s.Value.InDegree() != 0 {
			fmt.Printf("Parents: %v\n", s.Value.ID)
		}
	}
	d.DeleteVertex(v)
	fmt.Println(d.String())
}

func parents2(d *dag.DAG, v *dag.Vertex) {

	list, err := d.Predecessors(v)
	if err != nil {
		panic(err)
	}
	var ss []kv
	for _, v := range list {
		ss = append(ss, kv{fmt.Sprintf("%v", v.ID), v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Key > ss[j].Key
	})
	fmt.Println(" Deleteing: ", v.ID)
	for _, s := range ss {
		d.DeleteEdge(s.Value, v)
		if s.Value.InDegree() != 0 {
			fmt.Printf("Parents: %v\n", s.Value.ID)
		}
	}
	d.DeleteVertex(v)
	fmt.Println(d.String())
	list = d.SinkVertices()
	ss = []kv{}
	for _, v := range list {
		ss = append(ss, kv{fmt.Sprintf("%v", v.ID), v})
		fmt.Println("  New Sinks: ", v.ID)
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Key > ss[j].Key
	})

	parents2(d, ss[0].Value)
}

func parents(d *dag.DAG, v *dag.Vertex) {

	list, err := d.Predecessors(v)
	if err != nil {
		panic(err)
	}
	var ss []kv
	for _, v := range list {
		ss = append(ss, kv{fmt.Sprintf("%v", v.ID), v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Key > ss[j].Key
	})

	/*
		if len(list) > 1 {

			for i := 0; i < len(list)-1; i++ {
				//fmt.Println("Removing:", list[i], v)
				err := d.DeleteEdge(list[i], v)
				list[i].Children.Remove(v)
				v.Parents.Remove(v)

				if err != nil {
					panic(err)
				}
			}

		}
	*/
	//list, err = d.Predecessors(v)
	//if err != nil {
	//	panic(err)
	//	}

	var ss2 []kv
	for _, top := range ss { // Print the x top values
		fmt.Printf("\n  ID: %v, ParentCount: %v, Checking: %v, IN: %v , out: %v\n",
			v.ID,
			len(ss),
			top.Value.ID,
			top.Value.InDegree(),
			top.Value.OutDegree())
		if top.Value.OutDegree() > 0 && top.Value.InDegree() != 0 {
			kids, err := d.Successors(top.Value)
			if err != nil {
				panic(err)
			}
			for _, v := range kids {
				ss2 = append(ss2, kv{fmt.Sprintf("%v", v.ID), v})
			}
			sort.Slice(ss2, func(i, j int) bool {
				return ss2[i].Key < ss2[j].Key
			})
			fmt.Println("   Kids Check ", kids[0].ID, v.ID)
			if kids[0].ID == v.ID {
				if top.Value.InDegree() == 0 {
					fmt.Println("Done")
					d.DeleteEdge(v, kids[0])
					break
				}
				fmt.Printf("%v", top.Value.ID)
			} else {
				d.DeleteEdge(v, kids[0])

			}
		} else {

			//fmt.Printf("%v", top.Value.ID)

		}
	}
	for _, kv := range list { // Print the x top values
		fmt.Println("\nRecurrsing on : ", kv.ID)
		parents(d, kv)
	}
}

type kv struct {
	Key   string
	Value *dag.Vertex
}

// Pull all lines into a string slice
func readFileToLines(file string) []string {
	// open data
	fh, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	r := bufio.NewReader(fh)
	scanner := bufio.NewScanner(r)
	lines := make([]string, 0)
	// read it all in
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

type Node struct {
}
