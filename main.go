package main

import (
	"fmt"
	"log"
	"net"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)


type Backend struct {
    Addr string
    Alive bool
    mux sync.RWMutex
    ReverseProxy *httputil.ReverseProxy
}

type ServerPool struct {
    backends []*Backend
    current uint64
    
}

type LoadBalancer struct {
    servers []*Backend
}

func NewLoadBalancer(servers []*Backend) *LoadBalancer {
    return &LoadBalancer{
        servers: servers,
    }
}

// Returns true when the backend is up (alive)
func (b *Backend) IsAlive(alive bool) {
   b.mux.RLock()
   alive = b.Alive
   b.mux.RLock()
   return
}

// Sets a backends status at alive or not
func (b *Backend) SetALive(alive bool) {
    b.mux.Lock()
    b.Alive = alive
    b.mux.Unlock()
}


func main() {

    s := Servers{}
    s.backend = append(s.backend, &Backend{Addr: "127.0.0.2"})
    s.backend = append(s.backend, &Backend{Addr: "127.0.0.3"})
    s.backend = append(s.backend, &Backend{Addr: "127.0.0.4"})

    lb := NewLoadBalancer(s.backend)

    for _ = range s.backend {
        fmt.Println(lb.IsNext(s))
    }

}

// function to check whether a TCP connection can be established to the server
func BackendStatus(u *url.URL) bool {
    timeout := 2 * time.Second
    conn, err := net.DialTimeout("tcp", u.Host, timeout)
    if err != nil {
        log.Println("server is down")
        return false 
    }
    defer conn.Close()

    return true
}

// fuction to be run periodically to check on status of available servers
func (b Backend) CheckServerHealth() bool {
    
}

// function to iterate through available servers
func (lb LoadBalancer) IsNext(s Servers) *Backend {
    var i = 0
    server := s.backend[i]
    i++

    // We have reached the end of the servers
    // and we need to restart the counter from
    // the beginning
    if i >= len(s.backend) {
       i = 0 
    }

    return server
}




