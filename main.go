package main

import (
	"github.com/wejectchen/ginblog/routes"
	"github.com/wejectchen/ginblog/utils"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	var g errgroup.Group
	g.Go(func() error {
		r := routes.InitRouter()
		if err := r.Run(utils.HttpPort); err != nil {
			log.Print("startup service failed, err:%v", err)
			return err
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		log.Print(err)
	}

}