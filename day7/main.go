package main

import (
	"bufio"
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

func main() {
	flag.Parse()
	lines := readFileToLines("data")
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
	fmt.Println(d)

	finalList := make([]*dag.Vertex, 0)
	starts := d.SourceVertices()
	sort.Sort(byID(starts))

	for true {

		if len(starts) < 1 {
			break
		}
		v := starts[0]
		starts = starts[1:]
		//fmt.Println("Node: ", v)
		kids, _ := d.Successors(v)
		sort.Sort(byID(kids))
		finalList = append(finalList, v)
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
	fmt.Println("Final")
	for _, v := range finalList {
		fmt.Print(v.ID)
	}
	fmt.Println("")
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
