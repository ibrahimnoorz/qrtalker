package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/ibrahimnoorz/qrtalker/protocol"
	"github.com/nats-io/nats"
	"github.com/nats-io/nats/encoders/protobuf"
)

func usage() {
	fmt.Println("usage:")
	fmt.Println("\tqrtalkercli <taretfile> <sourcedata>")
	fmt.Println("\tqrtalkercli file1.png \"This is a test data.\"")
}

func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}

	targetfile := os.Args[1]
	sourcedata := string(os.Args[2])
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	id := r1.Int63n(999988989)

	nc, _ := nats.Connect(nats.DefaultURL)

	//a protobuf encoded connection
	ec, _ := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	defer ec.Close()

	sendProtoCh := make(chan *Transport.QRRequest)
	ec.BindSendChan("QRCodes", sendProtoCh)

	recvProtoCh := make(chan *Transport.QRResponse)
	ec.BindRecvChan("QRCodesResp", recvProtoCh)

	//send the QRRequest
	sendProtoCh <- &Transport.QRRequest{
		Id: id, Type: string("QRCode"), Authtoken: string("somesecuritytoketocheckonserver"), Datalen: int32(len(sourcedata)),
		Data: sourcedata}

	//wait for response
	resp := <-recvProtoCh
	if len(resp.Err) == 0 {
		//fmt.Println("Received response of size", resp.Datalen)
		//_ = ioutil.WriteFile(fmt.Sprintf("png/abc-%d.png", resp.Id), resp.Data, 0644)
		_ = ioutil.WriteFile(targetfile, resp.Data, 0644)
	} else {
		fmt.Println("Error", resp.Err)
	}
}
