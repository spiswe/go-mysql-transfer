// Package byteutil /*
package byteutil

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
)

func StrToBytes(u string) []byte {
	return []byte(u)
}

func BytesToStr(u []byte) string {
	return string(u)
}

func Uint64ToBytes(u uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, u)
	return buf
}

func Int64ToBytes(u int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(u))
	return buf
}

func BytesToUint64(b []byte) uint64 {
	if b == nil {
		return 0
	}
	return binary.BigEndian.Uint64(b)
}

func BytesToInt64(b []byte) int64 {
	if b == nil {
		return 0
	}
	return int64(binary.BigEndian.Uint64(b))
}

func Uint8ToBytes(u uint8) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, &u)
	if err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

func BytesToUint8(b []byte) (uint8, error) {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp uint8
	err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	if err != nil {
		return 0, err
	}
	return tmp, nil
}

func BytesToUint32(b []byte) uint32 {
	if b == nil {
		return 0
	}
	return binary.BigEndian.Uint32(b)
}

func Uint32ToBytes(u uint32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, u)
	return buf
}

func JsonBytes(v interface{}) []byte {
	bytes, err := json.Marshal(v)
	if nil != err {
		return nil
	}
	return bytes
}
