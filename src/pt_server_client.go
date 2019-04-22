package main

var PTSRV_MSG_SEND_DATA uint8 = 1;
var PTSRV_MSG_SEND_DATA_ACK uint8 = 2;
var PTSRV_MSG_RECV_DATA uint8 = 3;
var PTSRV_MSG_RECV_DATA_ACK uint8 = 4;
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

type ptsrv_recv_data struct {
    data []byte
}

func (m ptsrv_recv_data) mtype() uint8 {
    return PTSRV_MSG_RECV_DATA
}

type ptsrv_recv_data_ack struct {
}

func (m ptsrv_recv_data_ack) mtype() uint8 {
    return PTSRV_MSG_RECV_DATA_ACK
}

type ptsrv_shutdown struct {
}

func (m ptsrv_shutdown) mtype() uint8 {
    return PTSRV_MSG_SHUTDOWN
}

type pt_conn struct {
    msg_in chan ptsrv_msg
    msg_out chan ptsrv_msg
}
