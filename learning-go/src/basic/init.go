package basic

import (
	"flag"
	"fmt"
	"os"
)

// enum
type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func (b ByteSize) String() string {
	switch {
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

//variables
var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USERNAME")
	gopath = os.Getenv("GOPATH")
)

func init() {
	if home == "" {
		home = "/home/" + user
	}
	if gopath == "" {
		gopath = home + "/go"
	}
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}

func Init() {

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)

	fmt.Println(home)
	fmt.Println(user)
	fmt.Println(gopath)

	// init function
}
