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

func (s *ServerPool) AddBackend(backend *Backend) {
   s.backends = append(s.backends, backend) 
}

func NewBackend(addr string, alive bool, proxy httputil.ReverseProxy) *Backend {
    return &Backend{
        Addr: addr,
        Alive: alive,
        ReverseProxy: &proxy,
    }
}

func (s *ServerPool) NextServer(urls []string) *Backend {
    for _, str := range urls {
        /*
        if s.backends[str].IsAlive() {
            fmt.Println(s.backends[val])
            return s.backends[val]
            */
        fmt.Printf("%s\n", str)
        }

    return nil
}

var serverPool ServerPool

func main() {

    urls := []string{"127.0.0.2", "127.0.0.4", "127.0.0.3"}


    s := ServerPool{}
    fmt.Println(s.NextServer(urls))
    /*
    u1 := "127.0.0.2"
    u2 := "127.0.0.4"
    u3 := "127.0.0.3"

    urlMap := make(map[string]string)

    urlMap[u1] = "127.0.0.2"
    urlMap[u2] = "127.0.0.3"
    urlMap[u3] = "127.0.0.4"
    
    for _, val := range urlMap {
        serverUrl, err := url.Parse(val)
        if err != nil {
            log.Fatal(err)
        }

        s1 := NewBackend("127.0.0.2", true, *httputil.NewSingleHostReverseProxy(serverUrl))
        s2 := NewBackend("127.0.0.3", false, *httputil.NewSingleHostReverseProxy(serverUrl))
        s3 := NewBackend("127.0.0.4", true, *httputil.NewSingleHostReverseProxy(serverUrl))

        fmt.Println(s1, s2, s3)
    }
    */

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
func (lb LoadBalancer) IsNext(s ServerPool) *Backend {
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
*/




