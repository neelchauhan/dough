package main

import (
    "math/rand"
)

type Connection struct {
    ConnectionId uint32
    LastMessageTypeRecv uint8
    LastMessageTypeSent uint8
}

var ConnectionTable map[uint32]Connection

func init_conn_table() {
    ConnectionTable = make(map[uint32]Connection)
}

func handle_conn_init_req(msg message) message {
    connInitReq := msg_conn_init_req(msg)

    newConnId := connInitReq.conn_id_spec * rand.Intn(8)

    ConnectionTable[newConnId] = Connection{newConnId, MSG_TYPE_CONN_INIT_REQ, MSG_TYPE_CONN_ACCEPTED}

    return msg_conn_init_req{newConnId}
}
