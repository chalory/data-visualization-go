package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Count struct {
	Key   string
	Value int
}

type CountList []Count

func (c CountList) Len() int           { return len(c) }
func (c CountList) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c CountList) Less(i, j int) bool { return c[i].Value < c[j].Value }

func generateBarItems(data CountList) []opts.BarData {
	barData := []int{}
	items := make([]opts.BarData, 0)

	for i := 0; i <= 4; i++ {
		barData = append(barData, data[i].Value)

	}

	for _, v := range barData {
		items = append(items, opts.BarData{Value: v})
	}

	return items

}

func main() {

	// csv reading

	f, err := os.Open("brooklyn.csv")
	data := []string{}

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(f)
	reader.LazyQuotes = true

	for {
		col, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, col[5])

	}

	// mapping
	freq := map[string]int{}

	for _, v := range data {
		_, exists := freq[v]

		if exists {
			freq[v] += 1
		} else {
			freq[v] = 1
		}
	}

	count := make(CountList, len(freq))
	i := 0

	for k, v := range freq {
		count[i] = Count{k, v}
		i++
	}

	sort.Sort(sort.Reverse(count)) // sort in descending order
	fmt.Println(count)

	// create a new bar instance
	bar := charts.NewBar()
	bar.SetXAxis([]string{
		count[0].Key[:4],
		count[1].Key[:4],
		count[2].Key[:4],
		count[3].Key[:4],
		count[4].Key[:4],
	}).AddSeries("Values", generateBarItems)
}
