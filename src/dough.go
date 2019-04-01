package main

import (
    golog "log"
    "flag"
    "os"
    "log"

    "github.com/miekg/dns"
)

type Settings struct {
    Domain string
    Front string
    isClient bool
    ListenAddress string
}

var GlobalSettings *Settings

func main() {
    GlobalSettings = new(Settings)
    flag.StringVar(&GlobalSettings.Domain, "domain", "dough.example.net", "domain to use with Dough")
    flag.StringVar(&GlobalSettings.Front, "front", "https://doh-gateway.com/dns-query", "DNS-over-HTTPS gateway to use with Dough")
    flag.StringVar(&GlobalSettings.ListenAddress, "listen", "0.0.0.0:53", "Listening address (for server mode only)")

    flag.Parse()

    isClient, err := ptIsClient()
    if err != nil {
        golog.Fatalf("[ERROR]: %s - must be run as a managed transport", os.Args[0])
    }

    if isClient {
        // client
    } else {
        // server
        srv := &dns.Server{Addr: GlobalSettings.ListenAddress, Net: "udp"}
        srv.Handler = &dnsHandler{}
        if err := srv.ListenAndServe(); err != nil {
            log.Fatalf("Failed to set udp listener %s\n", err.Error())
        }
    }
}
