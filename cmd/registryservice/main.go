package main

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//心跳注册
	registry.SetupRegistryService()
	http.Handle("/services", &registry.RegistryServices{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	// go func() {
	// 	fmt.Println("Registry service started, Press any key to stop.")
	// 	var s string
	// 	fmt.Scanln(&s)
	// 	srv.Shutdown(ctx)
	// 	cancel()
	// }()

	<-ctx.Done()
	fmt.Println("Shutting down registry service")
}
