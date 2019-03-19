package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type data struct {
	data      []int
	readIndex int
}

func (d *data) ReadN(n int) []int {
	data := d.data[d.readIndex : d.readIndex+n]
	d.readIndex += n
	return data
}

type header struct {
	childNumber    int
	metadataNumber int
}

type node struct {
	header     header
	childNodes []*node
	metadata   []int
	value      int
	part2Value int
	index      int
}

func (n *node) Print() {
	fmt.Println("Node: ", n.header)
	fmt.Println("Value: ", n.value, n.part2Value)
	for _, c := range n.childNodes {
		c.Print()
	}
}

func (n *node) SumMetaDataP2() int {
	total := 0

	if len(n.childNodes) == 0 {
		for _, v := range n.metadata {
			total += v
		}
		return total
	}

	for _, i := range n.metadata {

		i--
		if i < len(n.childNodes) && i >= 0 {
			total += n.childNodes[i].part2Value
			//fmt.Println(i, n.header, total)
		}
	}
	return total
}

func (n *node) SumMetaData() int {
	total := 0
	for _, v := range n.metadata {
		total += v
	}
	for _, c := range n.childNodes {
		total += c.SumMetaData()
	}
	return total
}

func (n *node) Read(d *data) {
	header := d.ReadN(2)
	//fmt.Println("Header: ", header)
	n.header.childNumber = header[0]
	n.header.metadataNumber = header[1]
	for i := 0; i < n.header.childNumber; i++ {
		c := &node{index: n.index + i + 1}
		c.Read(d)
		//fmt.Printf("Child %v, Value:%v  Value2 %v\n", c.header, c.value, c.part2Value)
		n.childNodes = append(n.childNodes, c)

	}
	n.metadata = d.ReadN(n.header.metadataNumber)
	n.value = n.SumMetaData()
	n.part2Value = n.SumMetaDataP2()

}

func main() {
	ii := readFileToInts("data")
	d := &data{data: ii}

	n := &node{}
	n.Read(d)

	n.Print()
	fmt.Println("P1: ", n.SumMetaData())
	fmt.Println("P2: ,", n.part2Value)
}

//

// Pull all lines into a int slice
func readFileToInts(file string) []int {
	// open data
	fh, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	r := bufio.NewReader(fh)
	scanner := bufio.NewScanner(r)
	list := make([]int, 0)
	// read it all in
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		for _, f := range fields {
			v, err := strconv.ParseInt(f, 10, 64)
			if err != nil {
				panic(err)
			}
			list = append(list, int(v))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}
