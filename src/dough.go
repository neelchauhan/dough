package main

import (
    golog "log"
    "flag"
    "os"
)

type Settings struct {
    Domain string
    Front string
    isClient bool
}

var GlobalSettings *Settings

func main() {
    GlobalSettings = new(Settings)
    flag.StringVar(&GlobalSettings.Domain, "domain", "dough.example.net", "domain to use with Dough")
    flag.StringVar(&GlobalSettings.Front, "front", "https://doh-gateway.com/dns-query", "DNS-over-HTTPS gateway to use with Dough")

    flag.Parse()

    isClient, err := ptIsClient()
    if err != nil {
        golog.Fatalf("[ERROR]: %s - must be run as a managed transport", os.Args[0])
    }

    if isClient {
        // client
    } else {
        // server
    }
}
