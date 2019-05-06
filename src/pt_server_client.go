package main

import (
    "net"
    "git.torproject.org/pluggable-transports/goptlib.git"
)

var ptInfo pt.ServerInfo

var PTSRV_MSG_SEND_DATA uint8 = 1;
var PTSRV_MSG_SEND_DATA_ACK uint8 = 2;
var PTSRV_MSG_RECV_DATA_REQ uint8 = 3;
var PTSRV_MSG_RECV_DATA uint8 = 4;
var PTSRV_MSG_SHUTDOWN uint8 = 5;

type ptsrv_msg interface {
    mtype() uint8
}

type ptsrv_send_data struct {
    data []byte
}

func (m ptsrv_send_data) mtype() uint8 {
    return PTSRV_MSG_SEND_DATA
}

type ptsrv_send_data_ack struct {
}

func (m ptsrv_send_data_ack) mtype() uint8 {
    return PTSRV_MSG_SEND_DATA_ACK
}

type ptsrv_recv_data_req struct {
}

func (m ptsrv_recv_data_req) mtype() uint8 {
    return PTSRV_MSG_RECV_DATA_REQ
}

type ptsrv_recv_data struct {
    data []byte
}

func (m ptsrv_recv_data) mtype() uint8 {
    return PTSRV_MSG_RECV_DATA
}

type ptsrv_shutdown struct {
}

func (m ptsrv_shutdown) mtype() uint8 {
    return PTSRV_MSG_SHUTDOWN
}

type pt_conn struct {
    msg_in chan ptsrv_msg
    msg_out chan ptsrv_msg
    or_conn net.Conn
}

var pt_conn_map map[uint32]pt_conn

func init_conn_map() {
    pt_conn_map = make(map[uint32]pt_conn)
}

func init_conn(conn_id uint32) {
    or_conn, err := pt.DialOr(&ptInfo, "127.0.0.1", "dough")
    pt_conn_map[conn_id] = pt_conn{nil, nil, or_conn}
    if err != nil {
        delete(pt_conn_map, conn_id)
    } else { // Launch
        go handle_conn(conn_id)
    }
}

func handle_conn(conn_id uint32) {
    var running bool = true
    conn_chan := pt_conn_map[conn_id]

    for running {
        msg := <-conn_chan.msg_in
        msg_type := msg.mtype()
        var msg_out ptsrv_msg

        switch msg_type {
             case PTSRV_MSG_SEND_DATA:
                 msg_send_data := msg.(ptsrv_send_data)
                 conn_chan.or_conn.Write(msg_send_data.data)
                 msg_out = new(ptsrv_send_data_ack)
             case PTSRV_MSG_RECV_DATA_REQ:
                 ;
             case PTSRV_MSG_SHUTDOWN:
                 running = false
        }

        if msg_type != PTSRV_MSG_SHUTDOWN {
            conn_chan.msg_out <- msg_out
        }
    }

    conn_chan.or_conn.Close()
}
