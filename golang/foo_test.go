package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"
)

type MyInt1 int
type MyInt2 = int

func (i MyInt1) m1() {
	fmt.Println("MyInt1.m1")
}

func TestTypeAlias(t *testing.T) {
	var i1 MyInt1
	i1.m1()

}

type Person struct {
	Name string
	age  int
}

func TestCanInterface(t *testing.T) {
	var ps Person
	vv := reflect.ValueOf(ps)
	f0 := vv.Field(0)
	f1 := vv.Field(1)
	fmt.Println(f0.Kind(), f0.CanInterface())
	fmt.Println(f1.Kind(), f1.CanInterface())
}

func TestParseInt(t *testing.T) {
	s := "-129"
	v1, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		panic(err)
	}
	v2, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%b %b\n", v1, v2)
	fmt.Println(int32(v1), int32(v2))
}

func TestBufioReadBytes(t *testing.T) {
	var rd *bufio.Reader

	f, err := os.Open("/tmp/package.json")
	if err != nil {
		panic(err)
	}
	rd = bufio.NewReader(f)
	for {
		row, err := rd.ReadBytes('\n')
		fmt.Println(row, err)
		if err == io.EOF {
			break
		}
	}

	s := `aaa
bbb
111
222`
	rd = bufio.NewReader(strings.NewReader(s))
	for {
		row, err := rd.ReadBytes('\n')
		fmt.Println(row, err)
		if err == io.EOF {
			break
		}
	}
}

type TblPartition struct {
	Id   int
	Data string
}

func TestGoroutine(*testing.T) {
	runtime.GOMAXPROCS(1)
	i := TblPartition{}
	go func(i *TblPartition) {
		fmt.Println(i)
	}(&i)
	fmt.Println(1)
}

func TestSyscallFtruncate(*testing.T) {
	filename := "/tmp/bar.txt"
	fd, err := syscall.Open(filename, syscall.O_CREAT|syscall.O_RDWR, 0644)
	if err != nil || fd < 0 {
		panic(err)
	}
	defer syscall.Close(fd)
	syscall.Ftruncate(fd, int64(1000))
	fmt.Println(os.ReadFile(filename))
}

func TestCodeBlock(*testing.T) {
	if true {
		v := 1
		{
			fmt.Println(v)
		}
		fmt.Println(v)
	}
	// “undefined: v” compilation error
	// fmt.Println(v)
}

func TestWriteTCP(*testing.T) {
	conn, err := net.Dial("tcp", ":8899")
	if err != nil {
		panic(err)
	}
	content := []byte{0x31, 0x32, 0x33, 10}
	for {
		fmt.Println(conn.Write(content))
		time.Sleep(time.Second)
	}
}

func TestReadTCP(*testing.T) {
	conn, err := net.Dial("tcp", ":8899")
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 256)
	for {
		if n, err := conn.Read(buf); err != nil {
			fmt.Println(conn, err)
			time.Sleep(time.Second)
		} else {
			fmt.Println(n, "***", string(buf[:n]))
		}
	}
}

func TestRecover(*testing.T) {
	for {
		makePanic()
	}
}

func makePanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("reovered", err)
		}
		time.Sleep(1 * time.Second)
	}()
	panic("hello")
}

func TestInsecureSkipVerify(*testing.T) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Post("https://10.1.100.106/ndr/api/ndr/sync_config/v2/notice", "application/json", strings.NewReader(`{}`))
	// resp, err := client.Post("https://cip.cc", "text/plain", bytes.NewReader([]byte("")))
	// resp, err := client.Get("https://10.1.100.106/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if bb, err := ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	} else {
		fmt.Println(string(bb))
	}
}
