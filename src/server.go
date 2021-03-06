package main

import (
    "encoding/base32"
    "encoding/base64"
    "encoding/binary"
    "hash/crc32"
    "strings"

    "github.com/miekg/dns"
)

type dnsHandler struct{}
func (this *dnsHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
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
                msgCrc32 := binary.LittleEndian.Uint32(decodedMsg[1:5])
                msgPostHeader := decodedMsg[5:]
                if msgCrc32 == crc32.ChecksumIEEE(msgPostHeader) {
                    inputMsg := bytes_to_msg(decodedMsg, uint16(len(decodedMsg)))
                    outMsg = process_msg_server(inputMsg)
                } else {
                    outMsg = msg_checksum_invalid{0}
                    outBytes = msg_to_bytes(outMsg, MSG_SIZE_RECV_TXT)
                }
            }

            encodedArray := make([]string, 1)
            encodedArray[0] = base64.StdEncoding.EncodeToString(outBytes)

            msg.Answer = append(msg.Answer, &dns.TXT {
                Hdr: dns.RR_Header {
                    Name: msg.Question[0].Name,
                    Rrtype: dns.TypeTXT,
                    Class: dns.ClassINET,
                    Ttl: 30,
                },
                Txt: encodedArray,
            })
    }
}
