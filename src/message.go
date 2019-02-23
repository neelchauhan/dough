package main

import (
    "bytes"
    "encoding/binary"
    "hash/crc32"
)

// Message types
var MSG_TYPE_CHECKSUM_INVALID uint8 = 0
var MSG_TYPE_CONN_INIT_REQ uint8 = 1
var MSG_TYPE_CONN_ACCEPTED uint8 = 2
var MSG_TYPE_CONN_REJECTED uint8 = 3
var MSG_TYPE_CONN_SEND_DATA uint8 = 4
var MSG_TYPE_CONN_SEND_DATA_ACK uint8 = 5
var MSG_TYPE_CONN_RET_DATA_REQ uint8 = 6
var MSG_TYPE_CONN_RET_DATA uint8 = 7
var MSG_TYPE_CONN_REQ_FAILED uint8 = 8
var MSG_TYPE_CONN_TEARDOWN uint8 = 9
var MSG_TYPE_CONN_TEARDOWN_ACK uint8 = 10
var MSG_TYPE_INVALID_CONN uint8 = 11

// Message Sizes
var MSG_SIZE_SEND uint16 = 120
var MSG_SIZE_RECV_TXT uint16 = 180
var MSG_SIZE_RECV_OPT uint16 = 3072

type message interface {
    mtype() uint8
}

// Checksum Invalid functions

type msg_checksum_invalid struct {
    conn_id uint32
}

func (m msg_checksum_invalid) mtype() uint8 {
    return MSG_TYPE_CHECKSUM_INVALID
}

// Connection Initialization Request functions

type msg_conn_init_req struct {
    conn_id_spec uint32
}

func (m msg_conn_init_req) mtype() uint8 {
    return MSG_TYPE_CONN_INIT_REQ
}

// Connection Accepted functions

type msg_conn_accepted struct {
    conn_id uint32
}

func (m msg_conn_accepted) mtype() uint8 {
    return MSG_TYPE_CONN_ACCEPTED
}

// Connection Rejected functions

type msg_conn_rejected struct {
    error_id uint16
}

func (m msg_conn_rejected) mtype() uint8 {
    return MSG_TYPE_CONN_REJECTED
}

// Connection Send Data functions

type msg_conn_send_data struct {
    conn_id uint32
    seq_no uint32
    size uint16
    data []byte
}

func (m msg_conn_send_data) mtype() uint8 {
    return MSG_TYPE_CONN_SEND_DATA
}

// Connection Send Data Acknowledgement functions

type msg_conn_send_data_ack struct {
}

func (m msg_conn_send_data_ack) mtype() uint8 {
    return MSG_TYPE_CONN_SEND_DATA_ACK
}

// Connection Retrieve Data Request functions

type msg_conn_ret_data_req struct {
    conn_id uint32
    seq_no uint32
    size uint16
}

func (m msg_conn_ret_data_req) mtype() uint8 {
    return MSG_TYPE_CONN_RET_DATA_REQ
}

// Connection Retrieve Data functions

type msg_conn_ret_data struct {
    size uint16
    data []byte
}

func (m msg_conn_ret_data) mtype() uint8 {
    return MSG_TYPE_CONN_RET_DATA
}

// Connection Request Failed functions

type msg_conn_req_failed struct {
    fail_type uint8
}

func (m msg_conn_req_failed) mtype() uint8 {
    return MSG_TYPE_CONN_REQ_FAILED
}

// Connection Teardown functions

type msg_conn_teardown struct {
    connection_type uint32
}

func (m msg_conn_teardown) mtype() uint8 {
    return MSG_TYPE_CONN_TEARDOWN
}

// Connection Teardown Acknowledgement functions

type msg_conn_teardown_ack struct {
}

func (m msg_conn_teardown_ack) mtype() uint8 {
    return MSG_TYPE_CONN_TEARDOWN_ACK
}

// Connection Teardown Acknowledgement functions

type msg_invalid_conn struct {
}

func (m msg_invalid_conn) mtype() uint8 {
    return MSG_TYPE_INVALID_CONN
}

// Message Parsing function

func msg_to_bytes(msg message, msg_size uint16) {
    msg_type := msg.mtype()
    msg_buf_contents := new(bytes.Buffer)
    var msg_crc32 uint32 = 0

    switch msg_type {
        case MSG_TYPE_CHECKSUM_INVALID:
            ;

        case MSG_TYPE_CONN_INIT_REQ:
            binary.Write(msg_buf_contents, binary.BigEndian, msg.(msg_conn_init_req).conn_id_spec)

        case MSG_TYPE_CONN_ACCEPTED:
            binary.Write(msg_buf_contents, binary.BigEndian, msg.(msg_conn_accepted).conn_id)

        case MSG_TYPE_CONN_REJECTED:
            binary.Write(msg_buf_contents, binary.BigEndian, msg.(msg_conn_rejected).error_id)

        case MSG_TYPE_CONN_SEND_DATA:
            binary.Write(msg_buf_contents, binary.BigEndian, msg.(msg_conn_send_data).conn_id)
            binary.Write(msg_buf_contents, binary.BigEndian, msg.(msg_conn_send_data).seq_no)
            binary.Write(msg_buf_contents, binary.BigEndian, msg.(msg_conn_send_data).size)
            binary.Write(msg_buf_contents, binary.BigEndian, msg.(msg_conn_send_data).data)

        case MSG_TYPE_CONN_SEND_DATA_ACK:
            ;

        case MSG_TYPE_CONN_RET_DATA_REQ:
            binary.Write(msg_buf_contents, binary.BigEndian, msg.(msg_conn_ret_data_req).conn_id)
            binary.Write(msg_buf_contents, binary.BigEndian, msg.(msg_conn_ret_data_req).seq_no)
            binary.Write(msg_buf_contents, binary.BigEndian, msg.(msg_conn_ret_data_req).size)

        case MSG_TYPE_CONN_RET_DATA:
            binary.Write(msg_buf_contents, binary.BigEndian, msg.(msg_conn_ret_data).size)
            binary.Write(msg_buf_contents, binary.BigEndian, msg.(msg_conn_ret_data).data)

        case MSG_TYPE_CONN_REQ_FAILED:
            binary.Write(msg_buf_contents, binary.BigEndian, msg.(msg_conn_req_failed).fail_type)

        case MSG_TYPE_CONN_TEARDOWN:
            binary.Write(msg_buf_contents, binary.BigEndian, msg.(msg_conn_teardown).connection_type)

        case MSG_TYPE_CONN_TEARDOWN_ACK:
            ;

        case MSG_TYPE_INVALID_CONN:
            ;
    }
    if msg_buf_contents.Len() != 0 {
        msg_crc32 = crc32.ChecksumIEEE(msg_buf_contents.Bytes())
    }

    msg_buf_header := new(bytes.Buffer)
    binary.Write(msg_buf_header, binary.BigEndian, msg_type)
    binary.Write(msg_buf_header, binary.BigEndian, msg_crc32)
}
