# Serial

[![GoDoc](https://godoc.org/github.com/xluohome/serial?status.svg)](http://godoc.org/github.com/xluohome/serial)


Serial 是一个Go语言实现的串口(Uart)接口包,可使用操作系统标准的 read 和 write 文件接口
实现串口字节流的接收和发送。


## 默认配置
 
!> 8 N 1 N （数据位: 8 奇偶校验: N 停止位: 1 数据流控: N）


## Links

* [Serial](https://github.com/xluohome/serial)

## GoDoc 

* [GoDoc](http://godoc.org/github.com/xluohome/serial)


## 代码使用

```go
package main

import (
	"log"

	"github.com/xluohome/serial"
)

func main() {
	c := &serial.Config{Name: "COM9", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	txbuf := []byte{0xAA, 0x01, 0x0f, 0x00, 0x00, 0xBA}

	n, err := s.Write(txbuf)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%X\n", buf[:n])
}

```

## 非阻塞模式

By default the returned Port reads in blocking mode. Which means
`Read()` will block until at least one byte is returned. If that's not
what you want, specify a positive ReadTimeout and the Read() will
timeout returning 0 bytes if no bytes are read.  Please note that this
is the total timeout the read operation will wait and not the interval
timeout between two bytes.

```go
	c := &serial.Config{Name: "COM45", Baud: 115200, ReadTimeout: time.Second * 5}
	
	// In this mode, you will want to suppress error for read
	// as 0 bytes return EOF error on Linux / POSIX
	n, _ = s.Read(buf)
```

## License

BSD