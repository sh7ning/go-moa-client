package main

import (
	"github.com/blackbeans/go-moa/core"
	"os"
	"os/signal"
)

type IGoMoaDemo interface {
	SetName(name string) error
	Ping() error
}

type GoMoaDemo struct {
}

func (self GoMoaDemo) SetName(name string) error {
	return nil
}

func (self GoMoaDemo) Ping() error {
	return nil
}

func main() {
	app := core.NewApplcation("conf/moa_server.toml", func() []core.Service {
		return []core.Service{
			core.Service{
				ServiceUri: "/service/bibi/go-moa",
				Instance:   GoMoaDemo{},
				Interface:  (*IGoMoaDemo)(nil)}}
	})

	//设置
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill)
	//kill
	<-ch
	app.DestroyApplication()

}
