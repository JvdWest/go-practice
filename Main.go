package main

import (
	"bufio"
	"fmt"
	"os"
)

func printFibonacciSeries(num int) {
	b := 1
	a := 0
	println(b)
	for {
		newValue := a + b
		if newValue > num {
			break
		}
		fmt.Println(newValue)
		a = b
		b = newValue
	}
}

var abc = []string{"Car", "Truck"}

func increaseArray(start, stop, step int) []int {
	var returnArray []int
	for i := start; i < stop; i += step {
		returnArray = append(returnArray, i)
	}
	return returnArray
}

func main() {

	fmt.Println(Solution("test 5 ?a0A pass007? ?xy1 a1a"))

	//fmt.Println(Solution7(1))
	//fmt.Println(Solution6(0, 1))
	//fmt.Println(Solution3([]int{1, 4, -1, 3, 2}))

	//fmt.Println(Solution2(26))
	//fmt.Println(Solution1(1073741727, 1073741631, 1073741679))
	//fmt.Println(binaryGap(32))
	//fmt.Println(solution("photo.jpg, Warsaw, 2013-09-05 14:08:15\nJay.png, London, 2015-06-20 15:13:22\nmyFriends.png, Warsaw, 2013-09-05 14:07:13\nEiffel.jpg, Paris, 2015-07-23 08:03:02\npisatower.jpg, Paris, 2015-07-22 23:59:59\nBOB.jpg, London, 2015-08-05 00:02:03\nnotredame.png, Paris, 2015-09-01 12:00:00\nme.jpg, Warsaw, 2013-09-06 15:40:22\na.png, Warsaw, 2016-02-13 13:33:50\nb.jpg, Warsaw, 2016-01-02 15:12:22\nc.jpg, Warsaw, 2016-01-02 14:34:30\nd.jpg, Warsaw, 2016-01-02 15:15:01\ne.png, Warsaw, 2016-01-02 09:49:09\nf.png, Warsaw, 2016-01-02 10:55:32\ng.jpg, Warsaw, 2016-02-29 22:13:11"))

	//
	//fmt.Println("Hey")
	//printFibonacciSeries(10)
	//var a []int
	//a = append(a, 5)
	//a = append(a, 5)
	//a = append(a, 5)
	//a = append(a, 5)
	//a[3] = 4
	//fmt.Println(append(a, 123))
	//fmt.Println(a)
	//
	//fmt.Printf("%v\n", 5.15)
	//fmt.Printf("%T\n", 5.15)
	//fmt.Printf("%#v\n", "String")
	//fmt.Printf("%b", 123)
	//
	//var b float32
	//println(b)
	//var arr2 = []int{1, 2}
	//fmt.Println(arr2)
	//fmt.Println(abc)
	//fmt.Println(len(abc))
	//fmt.Println(len("abc"))
	//fmt.Println(cap(abc))
	//
	//var test [4]int
	//fmt.Println(test)
	//fmt.Println(test)
	//
	////makeSlide := make([]int, 10)
	////var makeSlide = make([]int, 10)
	//var makeSlide = []int{1: 10}
	//fmt.Println(makeSlide)
	//arrayIncrease := increaseArray(10, 20, 2)
	//fmt.Println(arrayIncrease)
	//fmt.Println(arrayIncrease[:cap(arrayIncrease)])
	//fmt.Println(arrayIncrease)
	//g := 200
	//changeElement(g)
	//fmt.Println(g)
	//fileIO, err := ioutil.ReadFile("Test.txt")
	//fileIO2, err2 := os.ReadFile("Test.txt")
	//fmt.Println(err)
	//if err == nil && err2 == nil {
	//	fmt.Printf("%s", fileIO)
	//	fmt.Println(fileIO)
	//	fmt.Println(fileIO2)
	//}
	//
	//readFileLineByLine("Test.txt")
	//
	//for i, i2 := range arrayIncrease {
	//	fmt.Println(i)
	//	fmt.Println(i2)
	//}
}

func changeElement(array int) {
	array = 100
}

func readFileLineByLine(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
}
