package main

import (
	"bufio"
	"fmt"
	"nokowebapi/nokocore"
	"nokowebapi/task"
	"os"
	"strings"
)

/*
{
    "code": "087C8D9A",
    "cashier": "Ahmad Asy Syafiq",
    "total": 50000,
    "taxRate": 12,
    "money": 100000,
    "change": 50000,
    "qrCode": "https://www.alodev.id/",
    "carts": [
        {
            "name": "Paracetamol",
            "type": "Tablets",
            "qty": 6,
            "price": 2000,
            "total": 12000
        },
        {
            "name": "Amoxicillin",
            "type": "Strips",
            "qty": 4,
            "price": 8000,
            "total": 32000
        },
        {
            "name": "Omeprazole",
            "type": "Tablets",
            "qty": 3,
            "price": 2000,
            "total": 6000
        }
    ]
}
*/

func main() {

	//mode := &serial.Mode{
	//	BaudRate: 9600,
	//	Parity:   serial.NoParity,
	//	DataBits: 8,
	//	StopBits: serial.OneStopBit,
	//}
	//
	//port, err := serial.Open("COM4", mode)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//scanner := bufio.NewScanner(port)
	//
	//for {
	//	scanner.Scan()
	//	data := scanner.Bytes()
	//	fmt.Printf("Scan: %s | %+x\n", data, data)
	//}

	//beta.SessionPrinter()

	args := []string{"/scan", "/serial", "COM4"}
	//args := []string{"/print", "/name", "PANDA ESCPOS"}

	pipeIn, stdin := nokocore.Unwrap2(os.Pipe())
	stdout, pipeOut := nokocore.Unwrap2(os.Pipe())

	nokocore.KeepVoid(pipeIn, stdin, pipeOut, stdout)

	process := nokocore.Unwrap(task.MakeProcess("../../bin/NokoHwd.exe", args, nil, pipeIn, pipeOut, nil))
	nokocore.NoErr(process.Start())

	//writer := bufio.NewWriter(stdin)
	//nokocore.Unwrap(writer.WriteString("JSON: {\"code\":\"087C8D9A\",\"cashier\":\"Ahmad Asy Syafiq\",\"total\":50000,\"taxRate\":12,\"money\":100000,\"change\":50000,\"qrCode\":\"https://www.alodev.id/\",\"carts\":[{\"name\":\"Paracetamol\",\"type\":\"Tablets\",\"qty\":6,\"price\":2000,\"total\":12000},{\"name\":\"Amoxicillin\",\"type\":\"Strips\",\"qty\":4,\"price\":8000,\"total\":32000},{\"name\":\"Omeprazole\",\"type\":\"Tablets\",\"qty\":3,\"price\":2000,\"total\":6000}]}\n"))
	//nokocore.NoErr(writer.Flush())
	//
	//nokocore.Unwrap(writer.WriteString("EXIT\n"))
	//nokocore.NoErr(writer.Flush())

	//reader := bufio.NewReader(stdout)
	//for {
	//	input, err := reader.ReadString('\n')
	//	nokocore.NoErr(err)
	//	fmt.Printf("PipeOut: %s (%d)\n", input, len(input))
	//}

	scanner := bufio.NewScanner(stdout)
	for {
		if scanner.Scan() {
			input := scanner.Text()
			if data, ok := strings.CutPrefix(input, "Test:"); ok {
				fmt.Printf("Test: %s\n", data)
			}
			if data, ok := strings.CutPrefix(input, "Data:"); ok {
				fmt.Printf("Data: %s\n", data)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			break
		}
	}

	nokocore.NoErr(process.Wait())
}
