package main

func process_msg_server(msg message) message {
    var outMsg message
    switch msg.mtype() {
        case MSG_TYPE_CONN_INIT_REQ:
            outMsg = handle_conn_init_req(msg)
        default:
            ;
    }

    return outMsg
}
