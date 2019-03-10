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
            // Remove the specified server's domain from the subdomain (which contains the messages)
            domainMsg := strings.Replace(msg.Question[0].Name, "." + GlobalSettings.Domain + ".", "", 1)
            decodedMsg, err := base32.StdEncoding.DecodeString(domainMsg)
            var outMsg message
            var outBytes []byte
            if err != nil {
                outMsg = msg_checksum_invalid{0}
                outBytes = msg_to_bytes(outMsg, MSG_SIZE_RECV_TXT)
            } else {
                domainMsg := bytes_to_msg(decodedMsg, uint16(len(decodedMsg)))
            }
    }
}
