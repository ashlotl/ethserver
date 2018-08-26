package main
import (
	"os"
	"net"
	"fmt"
	"time"
	"bufio"
	"strconv"
	"strings"
	// "strconv"
)
var conn net.Conn
var READ_TIMEOUT =5*time.Second
func main() {//only testing, needs deletion
	fmt.Println(Listen("12345"))
	reader:=bufio.NewReader(os.Stdin)
	
	for {
		text1,_:=reader.ReadString(' ')
		text1=strings.Replace(text1," ","",1)
		num1,_:=strconv.Atoi(text1)
		text1b:=strconv.FormatInt(int64(num1),2)
		num1b,_:=strconv.Atoi(text1b)
		_=num1b
		arr1b:=[]byte(text1b)
		fmt.Println(arr1b)
		text2,_:=reader.ReadString('\n')
		text2=strings.Replace(text2,"\n","",1)
		num2,_:=strconv.Atoi(text2)
		text2b:=strconv.FormatInt(int64(num2),2)
		num2b,_:=strconv.Atoi(text2b)
		_=num2b
		arr2b:=[]byte(text2b)
		fmt.Println("Array",arr2b)
		for iterator:=0;iterator<4;iterator++ {
			fmt.Println(iterator)
			if iterator<len(arr1b) && arr1b[iterator]=='1' {
				WriteOut(iterator,1)
			} else {
				WriteOut(iterator,0)
			}
		}
		for iterator:=4;iterator<8;iterator++ {
			fmt.Println(iterator,2)
                        if (iterator-4)<len(arr2b) && arr2b[iterator-4]=='1' {
                                WriteOut(iterator,1)
                        } else {
                                WriteOut(iterator,0)
                        }
                }

	}
}
func Listen(pass string) bool{
	var ip string
	ifaces, err := net.Interfaces()
	checkError(err)
	for _, i := range ifaces {
	    addrs, err := i.Addrs()
	    checkError(err)
	    for _, addr := range addrs {
	    	fmt.Println(addr.String())
	        if strings.Contains(addr.String(),".") {
	        	ip=strings.Split(addr.String(),"/")[0]
	        }
	        // process IP address
	    }
	}
	_=ip
	listener, err:=net.Listen("tcp","wurstwurstmachine.local"+":12345")
	checkError(err)
	conn, err=listener.Accept()
	checkError(err)
	recv:=make([]byte,1024)
	conn.SetReadDeadline(time.Now().Add(READ_TIMEOUT))
	_,err=conn.Read(recv)
	checkError(err)
	fmt.Println(recv)
	if strings.Trim(string(recv[:]),"\x00")!=pass {
		return false
	}
	return true
}
func WriteOut(toWrite int,value int) {
	fmt.Println("Sending This:"+strconv.Itoa(toWrite)+";"+strconv.Itoa(value)+"@")
	_,err:=conn.Write([]byte(strconv.Itoa(toWrite)+";"+strconv.Itoa(value)+"@"))
	checkError(err)
}
func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "An error: %s", err.Error())
        os.Exit(1)
    }
}
