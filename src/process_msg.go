package main

func process_msg_server(msg message) message {
    var outMsg message
    switch msg.mtype() {
        case MSG_TYPE_CONN_INIT_REQ:
            outMsg = handle_conn_init_req(msg)
        case MSG_TYPE_CONN_SEND_DATA:
            outMsg = handle_conn_send_data(msg)
        case MSG_TYPE_CONN_TEARDOWN:
            outMsg = handle_conn_shutdown(msg)
        default:
            ;
    }

    return outMsg
}
