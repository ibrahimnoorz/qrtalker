# qrtalker
On-demand qr-code generator server and client with middleware using nats.

# Motivation
qrtalker is a client server solution to generate QR-Codes on demand. This effort is a proof of concept project to learn and use several concepts around Google's <b>go</b> programming language, NATS middleware technology and Google's Protocol Buffers.

# Installation / Setup
Please follow the steps below to setup a working project. These steps assume that you have a working go environment i.e. valid GOPATH, Go binaries on Path...

 1. Get the qrtalker repo from github.com
 
 2. Setup the qrtalker server 
  * _cd qrtalkerserver_
  * get all the required imports using
      
      _go get ./..._
      
  * _go build qrtalkersrv.go_
  
 3. Setup the qrtalker client
  * _cd qrtalkercli_
  * get all the required imports using
  
      _go get./..._
      
  * _go build qrtalkersrv.go_
 
 4. Setup Nats server using:
  * _go get github.com/nats-io/gnatsd_
  
  
# Usage
  1. Startup nats server:
  
    _gnats -DV_
    
    -DV  = Debug and Trace
    refer to: http://nats.io/documentation/server/gnatsd-usage/

  2. Startup qrtalkersrv
  
    [Linux]
    
    _./qrtalkersrv_
    
    [Windows]
    
    _qrtalkersrv.exe_
    
  3. Request a QR-Code using the client
    [Linux]
    
    _./qrtalkercli file1.png "My test message to be converted to a QR-Code image"_

    [Windows]
    
    _./qrtalkercli file1.png "My test message to be converted to a QR-Code image"_

# Todo
The following are some items that may be added in future updates:
  - Make server configurable i.e. ports, QR code parameters, enforce limits etc...
  - Accept input as file on clients
  - Support sending multiple QR-Code requests in a single call i.e. generate QR-Codes for all files in a folder
  
# Notes
I have only tested this project on a Windows machine but it should work on Linux the same.

# License
This project is licensed under the MIT license. Please read the LICENSE file.
