package main

import (
    "encoding/base32"
    "strings"

    "github.com/miekg/dns"
)

type handler struct{}
func (this *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
    msg := dns.Msg{}
    msg.SetReply(r)

    switch r.Question[0].Qtype {
        case dns.TypeTXT:
            msg.Authoritative = true
            domain_msg := strings.Replace(msg.Question[0].Name, "." + GlobalSettings.Domain + ".", "", 1)
            base32.StdEncoding.DecodeString(domain_msg) // TODO: Assign to variable
    }
}
