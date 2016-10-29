package main

import (
	"github.com/ziutek/telnet"
	"log"
	"time"
	"strings"
	"os"
	"bufio"
	"fmt"
)

const timeout = 10 * time.Second

func checkErr(err error) {
	if err != nil {
		log.Fatalln("Error:", err)
	}
}

func expect(t *telnet.Conn, d ...string) {
	checkErr(t.SetReadDeadline(time.Now().Add(timeout)))
	checkErr(t.SkipUntil(d...))
}

func sendln(t *telnet.Conn, s string) {
	checkErr(t.SetWriteDeadline(time.Now().Add(timeout)))
	buf := make([]byte, len(s) + 1)
	copy(buf, s)
	buf[len(s)] = telnet.LF
	_, err := t.Write(buf)
	checkErr(err)
}

// invoke the handler function reading each line of the file as its input parameter
func ReadFileLineToFunc(fileName string, handler func(int, string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	index := 1
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		index++
		handler(index, line)
	}
	return scanner.Err()
}

func DisplayFan(ip string, user string, passwd string, quit chan string) {
	dst := ip + ":23"
	if len(user) == 0 {
		user, passwd = "admin", "zjyy230061"
	}

	t, err := telnet.Dial("tcp", dst)
	checkErr(err)
	defer t.Conn.Close()
	//t.SetUnixWriteMode(true)

	var data string
	expect(t, ":")
	sendln(t, user)
	expect(t, ":")
	sendln(t, passwd)
	expect(t, ">")
	sendln(t, "display fan")
	//data, err = t.ReadUntil("---- More ----")
	data, err = t.ReadString('>')

	checkErr(err)
	rs := strings.Split(data, "\r\n")
	output := fmt.Sprintf("%v %v\r\n", ip, rs[len(rs) - 1])
	for _, l := range rs[1:len(rs) - 1] {
		if len(l) > 0 {
			output = fmt.Sprintf("%v %v\r\n", output, l)
		}
	}
	quit <- output
}

func main() {
	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	} else {
		fileName = "./devlist.txt"
	}
	quit := make(chan string)
	var effectiveLines int
	err := ReadFileLineToFunc(fileName, func(i int, line string) {
		// 忽略空行和以#开头的行
		if (len(line) > 0) && (line[0] != '#') {
			args := strings.Split(line, ",")
			switch len(args) {
			case 1:
				go DisplayFan(args[0], "", "", quit)
				effectiveLines++
			case 2:
				go DisplayFan(args[0], args[1], "", quit)
				effectiveLines++
			case 3:
				go DisplayFan(args[0], args[1], args[2], quit)
				effectiveLines++
			default:
				log.Fatalln("bad arguments from file: ")
			}
		}
	})
	checkErr(err)
	for i := 0; i < effectiveLines; i++ {
		fmt.Printf("%v\r\n", <-quit)
	}
}