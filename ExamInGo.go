package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

type photoObj struct {
	city        string
	extension   string
	index       int
	createdDate time.Time
}

type finalName struct {
	name  string
	index int
}

func solution(s string) string {
	x := strings.Split(s, "\n")

	dataSplit := make(map[string][]photoObj)
	for i, s2 := range x {
		commaSplit := strings.Split(s2, ",")
		date, _ := time.Parse(" 2006-01-02 15:04:05", commaSplit[2])
		city := commaSplit[1]
		record := photoObj{
			city:        city,
			extension:   strings.Split(commaSplit[0], ".")[1],
			index:       i,
			createdDate: date,
		}
		dataSplit[city] = append(dataSplit[city], record)
	}
	var finalPhotoList []finalName
	for city, arr := range dataSplit {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].createdDate.Before(arr[j].createdDate)
		})
		numberOfZeros := int(math.Floor(math.Log10(float64(len(arr)))) + 1)
		for i, obj := range arr {
			finalPhotoList = append(finalPhotoList, finalName{
				index: obj.index,
				name:  fmt.Sprintf("%s%0"+strconv.Itoa(numberOfZeros)+"d.%s", city, i+1, obj.extension),
			})
		}
	}
	sort.Slice(finalPhotoList, func(i, j int) bool {
		return finalPhotoList[i].index < finalPhotoList[j].index
	})
	var returnString string
	for _, name := range finalPhotoList {
		returnString += name.name + "\n"
	}
	return returnString
}
