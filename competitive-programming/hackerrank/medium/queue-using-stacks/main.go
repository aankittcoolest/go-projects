package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type queue []string

func (q queue) dequeue() (queue, string) {
	l := len(q)
	if l == 0 {
		return q, ""
	}
	return q[1:l], q[0]
}

func (q queue) enqueue(s string) queue {
	return append(q, s)
}

func (q queue) print() string {
	return q[0]
}

func main() {
	//processHackrankInput()
	localTest()

}

func localTest() {
	f, err := os.Open("./data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	os.Stdin = f

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	checkError(err)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	var q queue

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)
		//if tItr == 0 {
		//	continue
		//}
		a := strings.Split(s, " ")
		fmt.Println(a)
		if a[0] == "1" {
			q = q.enqueue(a[1])
			fmt.Println(q)
		} else if a[0] == "2" {
			q, _ = q.dequeue()
			fmt.Println(q)
		} else {
			fmt.Println(q)
			fmt.Println(q.print())
		}
	}

}

func processHackrankInput() {

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("data.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	var q queue

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)
		if tItr == 0 {
			continue
		}

		a := strings.Split(s, " ")
		if a[0] == "1" {
			q = q.enqueue(a[1])
		} else if a[0] == "2" {
			q, _ = q.dequeue()
		} else {
			fmt.Fprintf(writer, "%s\n", q.print())
		}
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
