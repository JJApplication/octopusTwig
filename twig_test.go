/*
Create: 2022/8/15
Project: octopusTwig
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopusTwig
package octopusTwig

import (
	"fmt"
	"os"
	"testing"
	"time"

	udsc "github.com/JJApplication/fushin/client/uds"
	"github.com/JJApplication/fushin/server/uds"
)

// benchmark
// 后台以Hermes的ping为例

var Name = "Twig"

func init() {
	os.Setenv(SocketRoot, "/tmp")
}

func BenchmarkTwigPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := GetClientFromPool("Hermes")
		err := c.Dial()
		_, err = c.SendWithRes(uds.Req{
			Operation: "ping",
			Data:      "",
			From:      Name,
			To:        nil,
		})
		if err != nil {
			b.Error(err)
		}
		PutClientToPool(c)
	}
}

func BenchmarkTwigWithNoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := NewClient(0)
		c.Addr = GetSocket("Hermes")
		err := c.Dial()
		_, err = c.SendWithRes(uds.Req{
			Operation: "ping",
			Data:      "",
			From:      Name,
			To:        nil,
		})
		if err != nil {
			b.Error(err)
		}
		if err != nil {
			b.Error(err)
		}
	}
}

var c udsc.UDSClient

func TestPoolRequest(t *testing.T) {
	twig := TwigCore{Name: OctopusTwig}
	twig.Init()

	go func() {
		twig.Start()
	}()

	// 并发请求
	// 模拟外部client调用
	var errCount = 0

	for range [100]struct{}{} {
		c = udsc.UDSClient{Addr: "/var/run/OctopusTwig.sock"}
		err := c.Dial()
		if err != nil {
			fmt.Println("error", err)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		go func() {
			_, err := c.SendWithRes(uds.Req{
				Operation: "ping",
				Data:      "",
				From:      "Test",
				To:        []string{"Hermes"},
			})
			if err != nil {
				errCount += 1
				return
			}
		}()
		time.Sleep(10 * time.Millisecond)
	}

	t.Log("error count: ", errCount)
	defer c.Close()
	select {}
}
