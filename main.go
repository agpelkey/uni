package main

import (
	"log"
	"net"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

// Create Backend server struct
type Backend struct {
    Addr *url.URL 
    Alive bool
    mux sync.RWMutex
    Proxy httputil.ReverseProxy
}

// Create loadbalancer struct
type LoadBalancer struct {
    servers []*Backend
    current int64 //
    // maybe add a port here
}

// Factory method to create new backends
func NewBackend(addr string) *Backend {
    serverURL, err := url.Parse(addr)
    if err != nil {
        log.Fatal(err)
    }

    return &Backend{
        Addr: serverURL,
        //Alive: true,
        Proxy: *httputil.NewSingleHostReverseProxy(serverURL),
    }
}

// Factory method to create new backends

// Method to check if backend is alive
func (b *Backend) IsAlive() (alive bool) {
    b.mux.Lock()
    alive = b.Alive
    b.mux.Unlock()
    return
}

// Method to set backend as alive
func (b *Backend) SetAlive(alive bool) {
    b.mux.Lock()
    b.Alive = alive
    b.mux.Unlock()
}

// Create method to ping backend servers to see if they are up
func (b *Backend) isBackendAlive(u *url.URL) bool {
    timeout := 2 * time.Second
    conn, err := net.DialTimeout("tcp", u.Host, timeout)
    if err != nil {
        log.Println("server failing ping test and cannot be reached")
        return false
    }

    defer conn.Close()
    return true
}

// HealthCheck(); Function that essentially wraps isBackendAlive to ping the servers and update status
func (lb *LoadBalancer) HealthCheck() {
    for _, b := range lb.servers {
        status := "up"
        alive := b.isBackendAlive(b.Addr)
        b.SetAlive(alive)
        if !alive {
            status = "down"
        }
        log.Printf("%s [%s]\n", b.Addr, status)
    }
}

// healtCheck; yes a bit confusing. This function is to be run as a go routine on an interval to scan the servers and see if they are up
var loadBalancer LoadBalancer
func healthCheck() {
    time := time.NewTicker(2 * time.Minute)
    for {
        select {
        case <- time.C:
            log.Println("beginning health check...")
            loadBalancer.HealthCheck()
            log.Println("Health check complete")
        }
    }

}

// Method to get the next backend

// Load Balancer function

func main() {

}
