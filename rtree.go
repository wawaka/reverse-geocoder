package main

import (
	"bufio"
	// "encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	// "strings"
	"bytes"
	"strconv"
	// "net/http"
	// "net/url"
	// "regexp"
	// "io/ioutil"
	// "time"
	// "os/signal"
	"log"
	// "syscall"

	"github.com/dhconnelly/rtreego"
)

type Region struct {
	region_id int64
	name string
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius float64 `json:"radius"`
}

func to_rect(center rtreego.Point, size []float64) *rtreego.Rect {
	return
}

func (r *Region) Bounds() *rtreego.Rect {
	// log.Printf("Bounds(): %d", r.region_id)
	point := rtreego.Point{r.Latitude, r.Longitude}
	rect := point.ToRect(0.1)
	// rect, err := rtreego.NewRect(point, []float64{21.0752833-21.0462684, 105.8126282-105.7622972})
	// rect, err := rtreego.NewRect(point, []float64{1, 1})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return rect
}

func ParseRegionLine(b []byte) *Region {
	r := Region{}
	var err error
	parts := bytes.Split(b, []byte("\t"))
	// fmt.Println(len(parts))
	// fmt.Println(parts)
	r.region_id, err = strconv.ParseInt(string(parts[0]), 10, 16)
	if err != nil {
		log.Printf("failed to convert '%s' to int", parts[0])
		return nil
	}
	r.name = string(parts[1])

	err = json.Unmarshal(parts[2], &r)
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("region_id=%d, name=%s, lat=%f, lon=%f, radius=%f", r.region_id, r.name, r.Latitude, r.Longitude, r.Radius)

	return &r
}

func ReadFile(filename string) []*Region {
	regions := make([]*Region, 0)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)

	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		// s := string(line)
		region := ParseRegionLine(line)
		if region != nil {
			regions = append(regions, region)
		}
		line, isPrefix, err = r.ReadLine()
	}
	if isPrefix {
		log.Fatal("buffer size to small")
	}
	if err != io.EOF {
		log.Fatal(err)
	}
	return regions
}


var INPUT_PATH = flag.String("i", "", "input")

func main() {
	flag.Parse()

	// var err error = nil

	regions := ReadFile(*INPUT_PATH)
	fmt.Printf("regions size %d\n", len(regions))

	rt := rtreego.NewTree(2, 5, 10)
	for _, r := range regions {
		// fmt.Println(r.name)
		rt.Insert(r)
	}
	fmt.Printf("rtree size %d\n", rt.Size())


	bb, err := rtreego.NewRect(rtreego.Point{21.0752833, 105.8126282}, []float64{0.0001, 0.0001})
	if err != nil {
		log.Fatal(err)
	}

	results := rt.SearchIntersect(bb)
	for _, res := range results {
		fmt.Println(res)
	}
}
