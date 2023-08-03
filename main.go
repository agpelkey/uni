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
    Proxy *httputil.ReverseProxy
}

/*
type ServerPool struct {
    backends []*Backend
    current uint64
    
}
*/

type LoadBalancer struct {
    servers []*Backend
    current int
    port string
}

func NewBackend(addr string) *Backend {
    serverUrl, err := url.Parse(addr)
    if err != nil {
        log.Fatal(err)
    }

    return &Backend{
        Addr: addr,
        Proxy: httputil.NewSingleHostReverseProxy(serverUrl),
    }
}

func NewLoadBalancer(port string, servers []*Backend) *LoadBalancer {
    return &LoadBalancer{
        servers: servers,
        current: 0,
        port: port,
    }
}

// Returns true when the backend is up (alive)
func (b *Backend) IsAlive() (alive bool) {
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


/*
// function to increase the counter of the serverPool 
func (s *ServerPool) NextIndex() int {
    //return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
    return int(s.current) % len(s.backends)
}
*/

func (lb *LoadBalancer) GetNextServer(b []Backend) *Backend {
    // retrieve server 
    server := lb.servers[lb.current % len(lb.servers)]
    for !server.IsAlive() {
        lb.current++
        server = lb.servers[lb.current % len(lb.servers)]
    }
    lb.current++
    return server
}

func main() {


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
    
    return false
}

// function to iterate through available servers
/*
func (s ServerPool) IsNext(urls []string) *Backend {
    var i = 0
    server := s.backends[i]
    i++

    // We have reached the end of the servers
    // and we need to restart the counter from
    // the beginning
    if i >= len(s.backends) {
       i = 0 
    }

    return server

}
*/




