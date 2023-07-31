package main

import (
	"fmt"
	"net/url"
	"os/exec"
	"strings"
	"sync"
)

type config struct {
    port int
    env string
}

type Backend struct {
    Addr url.URL
    IsAlive bool
    mux sync.RWMutex
}

type LoadBalancer struct {
    servers []*Backend
}

func NewLoadBalancer(servers []*Backend) *LoadBalancer {
    return &LoadBalancer{
        servers: servers,
    }
}


func main() {

}

// fuction to check if each back is up
func (b Backend) CheckHealth() bool {
    out, _ := exec.Command("ping", b.Addr.Host).Output()
    if strings.Contains(string(out), "Destination Host unreachable") {
        fmt.Println("server down")
        return false
    } else {
        fmt.Println("serve alive")
    }
    return true
}





