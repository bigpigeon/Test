/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package tcp_test

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"io"
	"net"
	"os"
	"testing"
	"time"
	"unsafe"
)

const sockFile = "test.sock"

type UnixReq struct {
	ID string `json:"id"`
}

type UnixRes struct {
	Msg string `json:"msg"`
}

func ReadByte(usConn *net.UnixConn) ([]byte, error) {
	lenInt := uint32(unsafe.Sizeof(uint32(1)))
	lenBytes := make([]byte, lenInt)
	n, err := io.ReadFull(usConn, lenBytes)
	if err != nil {
		fmt.Printf("io.ReadFull(%v,lenBytes) Err. err: %v\n", usConn, err)
		return nil, err
	}
	if uint32(n) != lenInt {
		strErrMsg := fmt.Sprintf(" n:%d != lenInt:%d Err.", n, lenInt)
		fmt.Println(strErrMsg)
		return nil, errors.New(strErrMsg)
	}

	var reqLen uint32
	lenBuf := bytes.NewBuffer(lenBytes)
	err = binary.Read(lenBuf, binary.BigEndian, &reqLen)
	if err != nil {
		fmt.Printf("binary.Read(%v, binary.BigEndian, &reqLen) Err. err: %v\n", lenBuf, err)
		return nil, err
	}

	reqBytes := make([]byte, reqLen)
	n, err = io.ReadFull(usConn, reqBytes)
	if err != nil {
		fmt.Printf("io.ReadFull(%v,reqBytes) Err: %v", usConn, err)
		return nil, err
	}
	if uint32(n) != reqLen {
		strErrMsg := fmt.Sprintf(" n:%d != reqLen:%d Err.", n, reqLen)
		fmt.Println(strErrMsg)
		return nil, errors.New(strErrMsg)
	}
	return reqBytes, nil
}

func ReadByteTimeout(usConn *net.UnixConn, timeout time.Duration) ([]byte, error) {
	timer := time.NewTimer(timeout)
	type Result struct {
		Data  []byte
		Error error
	}
	resultChan := make(chan Result, 1)
	go func() {
		data, err := ReadByte(usConn)
		resultChan <- Result{
			Data:  data,
			Error: err,
		}
	}()
	select {
	case <-timer.C:
		return nil, errors.New("receive timeout")
	case d := <-resultChan:
		return d.Data, d.Error
	}
}

func WriteByte(usConn *net.UnixConn, data []byte) error {
	buf := new(bytes.Buffer)
	msglen := uint32(len(data))

	binary.Write(buf, binary.BigEndian, &msglen)
	data = append(buf.Bytes(), data...)

	n, err := usConn.Write(data)
	if err != nil {
		fmt.Printf("usConn.Write(%s) Err.err:%v\n", data, err)
		return err
	}
	if n != len(data) {
		strErrMsg := fmt.Sprintf("Write Num:%d != len(data) %d Err", n, len(data))
		fmt.Println(strErrMsg)
		return errors.New(strErrMsg)
	}
	return nil
}

func handlerConn(t *testing.T, us *net.UnixListener, conn *net.UnixConn) {
	//StrIp := conn.RemoteAddr().String()
	//fmt.Println(StrIp)
	// receive request
	defer conn.Close()
	var req UnixReq
	{
		data, err := ReadByte(conn)
		require.NoError(t, err)
		err = json.Unmarshal(data, &req)
		require.NoError(t, err)
	}
	// sent response
	{
		resp := UnixRes{
			Msg: req.ID,
		}
		resData, err := json.Marshal(&resp)
		require.NoError(t, err)
		err = WriteByte(conn, resData)
		require.NoError(t, err)
	}

}

func TestUnixListen(t *testing.T) {
	os.Remove(sockFile)
	usAddr, err := net.ResolveUnixAddr("unix", sockFile)
	require.NoError(t, err)
	usListener, err := net.ListenUnix("unix", usAddr)
	require.NoError(t, err)
	//go ClientSent(t, uuid.New().String())
	//go ClientSent(t, uuid.New().String())
	for {
		usConn, err := usListener.AcceptUnix()
		require.NoError(t, err)
		go handlerConn(t, usListener, usConn)
	}
}

func ClientSent(t *testing.T, id string) {
	usAddr, err := net.ResolveUnixAddr("unix", sockFile)
	require.NoError(t, err)

	for {
		usConn, err := net.DialUnix("unix", nil, usAddr)
		require.NoError(t, err)
		fmt.Println("start conn", id)
		req := UnixReq{ID: id}
		{
			data, err := json.Marshal(&req)
			require.NoError(t, err)
			err = WriteByte(usConn, data)
			require.NoError(t, err)
		}
		time.Sleep(1 * time.Second)
		var res UnixRes
		{
			data, err := ReadByteTimeout(usConn, time.Second*2)
			require.NoError(t, err)
			err = json.Unmarshal(data, &res)
			require.NoError(t, err)
		}
		if id != res.Msg {
			fmt.Printf("client %s receive msg %s\n", id, res.Msg)
		}
		//usConn.Close()
		time.Sleep(100 * time.Millisecond)
	}
}

func TestClientSent(t *testing.T) {
	ClientSent(t, uuid.New().String())
}
