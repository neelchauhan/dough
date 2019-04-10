package main

import (
    "math/rand"
)

type Connection struct {
    ConnectionId uint32

    LastMessageTypeRecv uint8
    LastMessageRecv []byte
    RecvSeqNumber uint64
    RecvHasChecksumInvalid bool

    LastMessageTypeSent uint8
    LastMessageSent []byte
    SendSeqNumber uint64
    SendHasChecksumInvalid bool
}

var ConnectionTable map[uint32]Connection
var ConnSpecIdMap map[uint32]uint32

func init_conn_table() {
    ConnectionTable = make(map[uint32]Connection)
}

func handle_conn_init_req(msg message) message {
    connInitReq := msg.(msg_conn_init_req)

    // Reject if we already have this connection ID specifier
    if _, ok := ConnSpecIdMap[connInitReq.conn_id_spec]; ok {
        return msg_conn_rejected{2}
    }
    // TODO: Reject connections on server overload

    newConnId := connInitReq.conn_id_spec * uint32(rand.Intn(8))

    ConnectionTable[newConnId] = Connection{newConnId, MSG_TYPE_CONN_INIT_REQ, []byte{}, 0, false, MSG_TYPE_CONN_ACCEPTED, []byte{}, 0, false}
    ConnSpecIdMap[connInitReq.conn_id_spec] = newConnId

    return msg_conn_accepted{newConnId}
}
