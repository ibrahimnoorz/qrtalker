package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image/png"
	"io"
	"os"
	"sync"
	"sync/atomic"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/ibrahimnoorz/qrtalker/protocol"
	"github.com/nats-io/nats"
	"github.com/nats-io/nats/encoders/protobuf"
)

func createQRCode(data string, file string) error {
	f, _ := os.Create(file)
	defer f.Close()

	qrcode, err := qr.Encode(data, 0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		qrcode, err = barcode.Scale(qrcode, 100, 100)
		if err != nil {
			fmt.Println(err)
		} else {
			png.Encode(f, qrcode)
		}
	}

	return err
}

func createQRCodeInMem(data string, f io.Writer) error {
	qrcode, err := qr.Encode(data, 0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		qrcode, err = barcode.Scale(qrcode, 150, 150)
		if err != nil {
			fmt.Println(err)
		} else {
			png.Encode(f, qrcode)
		}
	}

	return err
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var msgcnt uint64
	nc, _ := nats.Connect(nats.DefaultURL)

	//a protobuf encoded connections
	ec, _ := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	defer ec.Close()

	recvProtoCh := make(chan *Transport.QRRequest)
	ec.BindRecvChan("QRCodes", recvProtoCh)

	sendProtoCh := make(chan *Transport.QRResponse)
	ec.BindSendChan("QRCodesResp", sendProtoCh)

	quit := make(chan bool, 2)
	// Receive via Go channels
	go func() {
		for {
			select {
			case req := <-recvProtoCh:
				atomic.AddUint64(&msgcnt, 1)

				//fmt.Println(time.Now())
				err := error(nil)
				if err == nil {
					var b bytes.Buffer
					f := bufio.NewWriter(&b)
					err = createQRCodeInMem(req.Data, f)
					f.Flush()
					if err == nil {
						qrResp := &Transport.QRResponse{Id: req.Id, Type: req.Type, Err: "", Datalen: int32(len(b.Bytes())), Data: b.Bytes()}
						sendProtoCh <- qrResp
					} else {
						fmt.Println("Error was", err)
						qrResp := &Transport.QRResponse{Id: req.Id, Type: req.Type, Err: err.Error(), Datalen: 0, Data: []byte{}}
						sendProtoCh <- qrResp
					}
				}
			case <-quit:
				wg.Done()
			}
		}
	}()

	fmt.Println("Waiting to hear something...")
	wg.Wait()
	fmt.Println("Exiting ...")
}
