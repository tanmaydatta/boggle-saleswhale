package main

import (
	"github.com/tanmaydatta/boggle/config"
	"github.com/tanmaydatta/boggle/init"
)

func main() {
	config.Load()
	s := server.New()
	s.ServeHTTP()
}

