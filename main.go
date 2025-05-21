package main

import (
    "fmt"
    "net/http"

    "github.com/miekg/dns"
)

func main() {
    website := "example.com"

    // Step 1: Check DNS resolution
    client := new(dns.Client)
    msg := new(dns.Msg)
    msg.SetQuestion(dns.Fqdn(website), dns.TypeA)
    dnsResp, _, err := client.Exchange(msg, "8.8.8.8:53")
    if err != nil {
        fmt.Printf("Status for %s: DOWN (DNS lookup failed: %v)\n", website, err)
        return
    }
    if len(dnsResp.Answer) == 0 {
        fmt.Printf("Status for %s: DOWN (no DNS records found)\n", website)
        return
    }
    var ip string
    for _, answer := range dnsResp.Answer {
        if a, ok := answer.(*dns.A); ok {
            ip = a.A.String()
            break
        }
    }
    fmt.Printf("DNS resolved %s to %s\n", website, ip)

    // Step 2: Check HTTP reachability
    httpResp, err := http.Get("http://" + website)
    if err != nil {
        fmt.Printf("Status for %s: DOWN (HTTP request failed: %v)\n", website, err)
        return
    }
    defer httpResp.Body.Close()

    if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
        fmt.Printf("Status for %s: UP (HTTP %d)\n", website, httpResp.StatusCode)
    } else {
        fmt.Printf("Status for %s: DOWN (HTTP %d)\n", website, httpResp.StatusCode)
    }
}