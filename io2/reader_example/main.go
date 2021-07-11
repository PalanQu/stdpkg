// package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func readFrom(reader io.Reader, num int) ([]byte, error) {
// 	p := make([]byte, num)
// 	n, err := reader.Read(p)
// 	if n > 0 {
// 		return p[:n], nil
// 	}
// 	return p, err
// }

// func main() {
// 	reader := strings.NewReader("0123456789")
// 	b := make([]byte, 5)
// 	n, err := reader.Read(b)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(n)
// 	fmt.Println(string(b))
// 	// bytes.Buffer
// 	// data, err := readFrom(strings.NewReader("from stringaaaaaaa"), 12)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fmt.Println(string(data))
// }
