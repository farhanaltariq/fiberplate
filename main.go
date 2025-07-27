//go:build !production
// +build !production

package main

import "github.com/farhanaltariq/fiberplate/app"

func main() {
	app.RunServer()
}
