// Code generated by go-bindata. DO NOT EDIT.
// sources:
// cmd/definitions/tmpl/function.tmpl (567B)
// cmd/definitions/tmpl/info.tmpl (1.702kB)
// cmd/definitions/tmpl/object.tmpl (2.177kB)
// cmd/definitions/tmpl/operation.tmpl (1.698kB)
// cmd/definitions/tmpl/pair.tmpl (502B)
// cmd/definitions/tmpl/service.tmpl (12.312kB)

// +build tools

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _cmdDefinitionsTmplFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\x41\x4b\xc4\x30\x10\x85\xef\xfd\x15\x8f\x25\x87\xae\xec\xe6\x07\x08\x9e\x8a\x82\xb0\xc8\xa2\xde\x25\x64\xa7\x6b\xb0\x99\x94\x66\x5a\x17\x62\xfe\xbb\xa4\x55\x17\x11\x0f\xde\x3c\x25\xcc\xcc\x7b\xf3\x3e\x26\xa5\x2d\x06\xc3\x47\x82\x7a\xda\x40\x4d\xb8\xbc\x82\xd2\x37\x23\xdb\x88\x9c\xab\xd2\x76\x2d\x38\x08\xd4\xa4\x6f\x7d\xdf\x91\x27\x16\x3a\x7c\x36\x55\xcb\x2f\xb3\x66\xd2\x77\xc6\x13\xde\x20\xa1\x31\x9e\xba\x65\xa0\x88\xd5\xa4\x77\xc1\x9a\xb9\xd2\x8e\x6c\x51\x47\x5c\xa4\x04\x75\x56\xec\x4d\x5c\x06\xd6\x48\xa9\x58\xe6\x5c\xa7\xa4\x26\xbd\x37\x83\xf1\x51\x3f\x0e\xce\xef\x4c\x14\xfd\x20\x83\xe3\xe3\x35\x1f\xe2\xab\x93\xe7\x26\x78\x6f\x72\x46\xe8\x05\xbd\x71\xc3\x2f\xa6\xa5\x5c\x62\x7e\xdf\xb4\x2c\xb8\xa7\x38\x76\x12\x3f\x8c\xe7\x00\x15\x00\xf4\x86\x9d\xad\x57\x05\xdc\x9d\xa9\x57\xeb\x6a\xa6\xa2\x2e\xd2\x1f\x71\xac\x9c\x60\x03\x0b\x9d\x44\x37\xcb\xbb\xc1\x3f\x66\xdc\x82\xf8\xeb\xca\x3f\xbe\xef\x01\x00\x00\xff\xff\x41\xc8\xae\xbf\x37\x02\x00\x00")

func cmdDefinitionsTmplFunctionTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplFunctionTmpl,
		"cmd/definitions/tmpl/function.tmpl",
	)
}

func cmdDefinitionsTmplFunctionTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplFunctionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/function.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc1, 0xfc, 0xb, 0xf4, 0xfb, 0xf, 0x94, 0xca, 0x64, 0x7a, 0xb8, 0x48, 0xf9, 0x69, 0xf1, 0x88, 0x46, 0x59, 0xd3, 0xbc, 0x0, 0x4e, 0xe7, 0x8, 0x1e, 0xdf, 0x27, 0x59, 0x81, 0x75, 0x52, 0xa2}}
	return a, nil
}

var _cmdDefinitionsTmplInfoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x94\x5f\x6f\xd3\x30\x14\xc5\xdf\xfd\x29\x0e\x55\x85\x1a\xd4\x35\x9b\x84\x78\x80\xe5\x69\x1b\x68\x42\xdb\x90\x36\xf1\x82\x10\x72\x92\x9b\xca\x34\xb6\x23\xdb\x89\x56\x32\x7f\x77\xe4\x24\x74\x4d\xf7\x07\x8a\x78\xe1\xcd\x37\xf6\xbd\xe7\x77\x8f\xaf\x13\xc7\x38\xd1\x39\x61\x49\x8a\x0c\x77\x94\x23\x5d\x63\xa9\x37\x31\x32\x99\xc7\x39\x15\x42\x09\x27\xb4\xb2\xef\x70\x7a\x85\xcb\xab\x1b\x9c\x9d\x9e\xdf\x2c\x58\xc5\xb3\x15\x5f\x12\xdc\xba\x22\xcb\x98\x90\x95\x36\x0e\x33\x06\x00\x93\x42\xba\x09\x8b\x18\x6b\xdb\x03\x18\xae\x96\x84\xe9\x6a\x8e\xa9\x50\x85\xb6\x78\x9b\x60\x71\x1e\x56\x17\xbc\x82\xf7\xac\x6d\x31\xb5\x64\x1a\x91\xd1\x25\x97\x14\xf6\xa7\x2b\xdc\xc1\xe9\x13\x2e\xa9\x0c\x47\x58\x1c\xe3\xbd\xa0\x32\x87\x50\x39\xdd\x42\x28\xb4\xed\x76\x92\xf7\x48\x85\x63\x99\x56\x36\x40\xec\xe8\x36\x5d\xcd\x5e\xdd\xfb\x0e\x71\x37\xfd\x3c\xd4\x0d\x24\xcd\xa2\x83\x08\xf2\x9f\xb8\xcd\x78\xd0\x47\x82\xa3\xe3\xe3\xb0\xbb\xea\x81\x0f\x40\x2a\x0f\xcb\x88\xb1\x60\x00\x76\x7b\x18\xa7\x5b\x67\xea\xcc\xa1\x1d\x94\x37\x6c\xdf\x9e\x62\x0b\x18\x37\xeb\xaa\xaf\xe5\xfd\xd6\x97\xfb\x33\x1b\x86\x2e\x8e\xe3\x60\x00\x6a\x4b\x39\xb8\x05\x0f\x91\xe4\x15\x0a\x6d\xa0\xd3\xef\x94\x39\x34\xbc\xac\x69\x8e\x43\x48\xe2\xca\x42\x69\x07\x4b\x6e\x8e\xa3\xe1\x83\x25\xd7\x95\xea\xea\x08\xe5\xde\xbc\xee\x42\x09\xc9\xab\x2f\xd6\x19\xa1\x96\x5f\x85\x72\x64\x0a\x9e\x51\xeb\xd9\xa0\xfc\xbc\xd7\x61\x57\x14\x81\xfe\xec\xb6\x9b\x10\xef\x59\x51\xab\x0c\x33\x89\x57\xcf\xba\x16\xe1\x03\xb9\xbe\xf1\x53\x61\xab\x92\xaf\x07\x37\x66\xd1\xd8\x8f\xc1\x57\x43\xae\x36\x0a\x72\xf1\xc0\xbe\x40\xfa\xa7\x9a\xd7\x4f\x68\x36\x63\xcd\xe8\x37\x85\x06\xa6\x47\x60\x90\xa0\x19\xf1\xb2\x61\xa0\x4a\xdb\xb1\xfe\x03\x77\x66\x23\xd4\x39\x52\xad\xcb\x68\x20\x12\x05\xe4\x22\xdc\xf0\xcb\x3d\x9f\xc0\x8b\x04\x87\x43\x8d\xe7\xdd\x9e\xc3\x99\x9a\xba\x83\x7e\xbb\xd1\x2d\xa8\x3b\xfc\x20\xa3\x3f\x87\x79\xec\x12\x0a\x5e\x5a\xda\xe7\x96\x2e\x6a\xeb\xf6\x9b\x8e\xbf\xee\x3b\x19\xf7\x5d\x71\x25\xb2\x59\x21\xdd\xe2\xba\x32\x42\xb9\x62\x36\x79\x8c\xf5\x23\xa5\x3c\xbd\x7f\xb9\xbf\xee\x5e\x6c\x9e\xdd\x24\x8a\x1e\x5a\xf4\xdf\xcc\x6e\x6f\xe5\x5d\xb2\x9f\x97\x8f\x8e\x7d\xff\x0f\xdb\xf9\xa5\xdd\x2f\x7f\x06\x00\x00\xff\xff\x26\xf5\xd0\xb3\xa6\x06\x00\x00")

func cmdDefinitionsTmplInfoTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplInfoTmpl,
		"cmd/definitions/tmpl/info.tmpl",
	)
}

func cmdDefinitionsTmplInfoTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplInfoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/info.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa6, 0x9d, 0xdc, 0x29, 0x17, 0xc8, 0x8a, 0x14, 0x25, 0x13, 0xa5, 0x85, 0x13, 0x6, 0xf, 0xb3, 0x4a, 0x21, 0xd1, 0x20, 0x4e, 0x6f, 0x92, 0x60, 0xe2, 0x6e, 0x29, 0x96, 0x4a, 0xee, 0x99, 0x66}}
	return a, nil
}

var _cmdDefinitionsTmplObjectTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x95\x51\x73\xe3\x34\x10\xc7\x9f\xab\x4f\xf1\x27\xd3\x01\x9b\x49\xed\x3b\x60\x78\x28\x97\x07\xe6\x52\xe0\x66\x68\xc3\x4c\x0a\xef\x8a\xbc\x4e\x45\x6d\xc9\x23\xad\x4d\x42\xea\xef\xce\xc8\x76\x72\x4e\x9b\x83\xdc\x30\x7d\x93\xac\xdd\xbf\x76\x7f\xbb\x2b\xa7\x29\xde\xdb\x8c\xb0\x26\x43\x4e\x32\x65\x58\x6d\xb1\xb6\x87\x3d\x54\x99\xa5\x19\xe5\xda\x68\xd6\xd6\xf8\x1f\x30\x5f\xe0\x6e\x71\x8f\x9b\xf9\x87\xfb\x44\x54\x52\x3d\xca\x35\x81\xb7\x15\x79\x21\x74\x59\x59\xc7\x88\x04\x00\x4c\xf2\x92\x27\xfd\x8a\x75\x49\xc3\xd2\x6f\x8d\x9a\x88\x58\x88\x34\xc5\x4f\x9a\x8a\x0c\xda\x64\xb4\x81\x36\xb0\xab\x3f\x49\x31\x56\x9a\x85\xb2\xc6\x07\x9d\xdd\xee\x0a\x4e\x9a\x35\xe1\xf2\x71\x8a\xcb\x06\xd7\x33\x24\x8b\xce\xee\x96\x58\xa2\x6d\x3b\xd5\xde\xf3\x43\x10\xda\xed\x70\xd9\x24\x77\xb2\x24\x3c\x81\xed\x6f\xd2\x2b\x59\xa0\x6d\x51\x6b\xc3\xdf\x7f\x87\x19\xde\xbe\x7b\x17\x8c\x1e\x83\x73\xd0\x27\x93\x85\x65\x1f\x52\xaf\x0d\xed\xc1\x0f\x04\x5f\xca\xa2\x20\xcf\xa8\x8d\xe6\x10\xe2\xda\x5e\x79\xb6\x4e\xae\x29\x11\x69\x1a\x1c\xee\x16\xf7\x37\xcb\xeb\xb0\x02\xae\x06\xf7\xaf\x3c\xf2\x90\x9a\xc7\xf2\x97\xc5\xef\xbf\xce\x61\x2c\x63\x45\x50\x0f\x21\x95\x0c\xb6\x66\xaf\x33\x82\x27\xd7\x68\x45\x3e\x39\x76\xc7\xfb\x1f\xef\x02\xe2\xe0\x61\x2b\x4d\xd9\xb3\x63\xed\xa1\xac\x51\xb5\x73\x64\x18\x5e\xe6\x94\x88\x50\x80\xfd\xb9\x67\x57\x2b\xc6\xee\x5c\x7a\xc1\x4c\xe7\x01\xdb\x9c\xbc\x72\xba\x0a\x95\xfe\x78\xf8\xc9\x83\x03\xba\x8f\x76\xf7\xdb\x8a\x3a\xf6\x6d\x3b\xfa\xf2\x8c\xb4\xb8\x48\x53\xa8\x42\x87\xe8\x07\xd0\xfb\x9d\xc1\x5f\x0f\x5a\x3d\x8c\x32\x95\x85\x6e\x28\x11\x17\x83\xc5\xb2\xa7\xef\x44\x77\x69\x9a\x86\x6e\x41\xed\x29\x83\xf4\x90\x61\x57\xca\x0a\xb9\x75\xfb\x6e\x6a\x64\x51\xd3\x14\x6f\x50\x92\x34\xbe\xab\x84\x27\x9e\xe2\xed\xf0\xc1\x13\x77\x52\x9d\x4e\xd7\x22\xe2\x22\xb3\x86\xba\xcd\xb7\xdf\x88\x8b\x32\x9c\x86\xae\x4d\x6e\x6b\xa6\x8d\x68\x85\x38\x07\x6c\xc8\x5e\xfb\x39\x55\x8e\x54\x98\xa3\xeb\x59\xc8\xee\x19\xca\x49\x9a\xe2\x60\x92\x5d\x4f\xf6\xa0\xfa\x6a\xdc\x6c\xba\x69\x6a\x5b\x91\xd7\x46\x21\xb2\xf8\xba\xbf\x23\xc6\xcf\xc4\x43\x61\xb4\xaf\x0a\xb9\x1d\x98\x47\xf1\x31\x75\xec\xba\xdc\x1c\x71\xed\x0c\x6c\xf2\xa2\x48\x21\x9b\x17\xe2\xcb\x4f\x88\x37\xc7\xe2\xf1\xde\x63\xb8\xe5\x84\x3c\x66\x68\x8e\x22\x10\x43\x23\x14\xbe\xbb\x3d\x50\x0a\xb9\x8e\x41\xb5\xad\x38\xa6\x82\x5b\xdb\x50\x06\xb6\x18\x46\x0f\x25\xb1\x4c\xc6\x1d\xf5\x19\x7c\xa2\xa3\x1c\xa6\x58\x59\x5b\xc4\x87\x0c\x3c\x4b\x8e\xe2\xbe\xb9\x74\x0e\x9b\x84\xb6\xf8\xf2\xac\xf7\xe5\x8b\x19\xde\x0c\x3a\xff\xce\x7c\x0a\x76\x35\x75\x86\xad\x18\xd3\x19\x45\xf6\x84\xbf\xc9\xd9\x3f\x42\xef\x76\x1e\xb9\x2c\x3c\x89\x57\x04\x76\x5b\x7b\xfe\xbc\xa6\xfa\x5f\xb0\x66\xc7\xb0\x2a\x69\xb4\x8a\xf2\x92\x93\x65\xe5\xb4\xe1\x3c\x9a\x0c\xe3\x3b\x52\x68\xdb\xf0\x1e\x0c\x03\x3c\x89\xe3\x81\xe0\x7f\x37\xf8\x6b\x31\x7b\x85\x39\xe9\x19\x3e\xcd\xce\x81\x78\x72\xb0\xfa\x50\xc7\x8f\xed\x8b\xb0\x55\x61\x0d\x45\x9b\xd1\x97\xdd\xe1\x41\x3f\xe7\x5f\x7b\x32\xfc\xcd\x29\xf6\xcf\x7e\x13\xa3\x0c\x3b\x87\xf0\xa3\x6f\xff\x09\x00\x00\xff\xff\x74\x56\xed\x90\x81\x08\x00\x00")

func cmdDefinitionsTmplObjectTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplObjectTmpl,
		"cmd/definitions/tmpl/object.tmpl",
	)
}

func cmdDefinitionsTmplObjectTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplObjectTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/object.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfd, 0xb0, 0xe7, 0x29, 0x9a, 0x5f, 0xc1, 0x32, 0x7, 0xac, 0x52, 0x8e, 0x9c, 0x2b, 0x1f, 0xb5, 0x69, 0x93, 0x58, 0xd7, 0x7, 0x6f, 0x97, 0x44, 0xd0, 0x62, 0x7d, 0x78, 0xa3, 0x58, 0x7d, 0x5b}}
	return a, nil
}

var _cmdDefinitionsTmplOperationTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\x4d\x8f\xd3\x30\x10\xbd\xfb\x57\x8c\xac\x1e\x12\x69\x49\xef\x48\x7b\x82\x45\x5a\x09\x75\x2b\x10\xe2\x88\x5c\x67\xda\xb5\x48\x6c\x33\x9e\xec\x87\x82\xff\x3b\xb2\x93\x76\xd3\xaa\x10\x15\x10\xe2\x16\x7b\x66\xde\x9b\xf7\xc6\x13\xaf\xf4\x57\xb5\x43\xe0\x67\x8f\x41\x08\xd3\x7a\x47\x0c\x85\x00\x00\x90\xda\x59\xc6\x27\x96\xc3\xc9\x38\x29\x4a\x21\xfa\xfe\x15\x90\xb2\x3b\x84\xc5\x97\x2b\x58\x18\x78\x7d\x0d\xd5\xad\x65\xa4\xad\xd2\x18\x20\x46\xd1\xf7\xb0\x30\xd5\x5b\x0c\x9a\x8c\x67\xe3\x6c\xba\x4c\x0c\x30\x46\x4c\xf0\x8d\x7a\x5e\xa9\x16\x21\x46\x30\xfb\x62\xe8\x33\x53\x62\x30\x5b\x70\x04\x05\x7e\x4b\xf9\x39\x51\x06\xa4\x07\xa3\x91\x64\x79\x72\xcf\x8e\xd4\x2e\xdd\xc7\x98\xeb\x3f\x32\x19\xbb\x2b\x4a\x08\xf9\xe3\x80\x89\xb6\x4e\x8d\x8c\xe7\xa9\x08\xe7\x93\x8a\x85\xa9\xee\x7c\x16\x90\x32\x96\xcb\xdc\xad\xf3\x03\xcd\x77\x60\xb7\x56\x41\xab\x26\xb5\x3c\x46\x4e\x24\x8e\xc0\xe7\x6b\x8a\x31\xf2\xce\x51\xab\x78\xad\x48\xb5\x89\xab\x84\xe3\xc0\x07\x0c\x5d\xc3\xe1\xb3\xe1\xfb\xf5\x30\x9b\x23\x95\x32\x95\x4c\x5d\xb2\x8e\x73\xf5\x7b\x37\xd0\xcc\xf6\x9e\x90\xdf\x0c\x83\xfd\x1d\x19\x93\xf2\x42\xf3\x13\x8c\x6f\xa4\x1a\xef\xae\xfe\xba\xca\xfd\xd8\x8e\x4f\xf9\xd8\x76\x81\x6f\xda\x0d\xd6\x9f\xac\x69\x7d\x83\x2d\x5a\xc6\xfa\xdc\x1b\x2b\x4a\x11\x85\x58\x2e\x61\x36\x33\x83\xc2\x06\x01\x13\x70\x8d\x35\xb0\x83\x7b\xf5\x80\xb0\x75\xf4\xa8\xa8\x06\xed\x5a\xaf\xd8\x6c\x1a\x84\x03\x96\x4a\xde\x85\x6a\x78\xe4\xf3\x1c\x81\xa9\xd3\x0c\x7d\x14\x62\xdb\x59\x0d\x45\x98\x2f\x2a\x2f\x92\x7b\x29\xf6\xc9\xce\x8c\x9b\x48\xc8\x1d\x59\x90\xb3\x00\x32\xd9\xfb\xeb\x9d\xba\xa0\x99\x7f\xb2\x44\xc7\x3f\x9b\xb3\x6b\x84\x44\x70\x0d\x2b\x7c\xbc\xf3\x48\x79\xc4\x2b\xc7\xb7\x2f\xdd\xdf\x10\x39\x2a\xe4\xb4\xdb\x18\x65\x79\xfa\xc3\x79\x71\x52\x44\xf1\x33\xbe\x0b\xec\xf9\x0f\x76\x73\x34\xef\x4f\x0c\x9a\x38\xb2\x37\x6a\xfa\x75\x70\xef\x47\x00\x00\x00\xff\xff\xbd\xe8\x37\x32\xa2\x06\x00\x00")

func cmdDefinitionsTmplOperationTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplOperationTmpl,
		"cmd/definitions/tmpl/operation.tmpl",
	)
}

func cmdDefinitionsTmplOperationTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplOperationTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/operation.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc2, 0xaf, 0x9a, 0xb8, 0xf7, 0x9f, 0x1b, 0x53, 0xaa, 0x36, 0x95, 0x6d, 0xe6, 0x56, 0x7d, 0xb9, 0x8e, 0xd7, 0x4, 0x74, 0xce, 0x5b, 0x9, 0x4f, 0xb6, 0x92, 0xc, 0x1e, 0x49, 0x5, 0x4c, 0x4c}}
	return a, nil
}

var _cmdDefinitionsTmplPairTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x41\x6b\xe3\x30\x10\x85\xef\xfa\x15\x0f\xe3\x43\x02\x89\x75\xd9\x53\x96\x3d\x6d\xf6\xb0\x14\x92\x1c\x42\x7b\x2c\x8a\x3c\x51\x44\x6c\x49\xc8\x63\xb7\xc6\xf5\x7f\x2f\xb2\x93\x40\xe9\xa5\x3a\xcd\x7c\x33\xef\xf1\x46\x52\xe2\xaf\x2f\x09\x86\x1c\x45\xc5\x54\xe2\xd4\xc3\xf8\x47\x0f\x5d\x97\xb2\xa4\xb3\x75\x96\xad\x77\xcd\x6f\x6c\xf7\xd8\xed\x8f\xf8\xb7\xfd\x7f\x2c\x44\x50\xfa\xaa\x0c\x21\x28\x1b\x1b\x21\x6c\x1d\x7c\x64\x2c\x04\x00\x64\xda\x3b\xa6\x77\xce\xc4\xdc\x1a\xcb\x97\xf6\x54\x68\x5f\xcb\x13\xf5\xde\x95\x0d\xfb\xa8\x0c\x49\xe3\xd7\xf7\xb2\xfb\x25\xc3\xd5\xc8\x0b\x73\xd0\x95\x25\xc7\xd9\xa4\x2d\x7e\xac\xe6\x3e\x50\x93\x09\xb1\x14\x62\x18\xd6\x88\xca\x19\x42\xfe\xba\x42\xde\x61\xf3\x07\xc5\x21\x05\xc5\x38\x4e\xd3\x3c\x38\x55\x53\xe2\x79\x57\xec\x52\xf9\x01\xf6\x07\xd5\x68\x55\xa5\x1d\x29\xf1\x62\xf9\x32\x0c\xf7\xcd\x71\xc4\x9b\xad\x2a\xa8\x10\xaa\x1e\x89\xdf\x74\xe3\x88\x4e\x55\x2d\x81\x3d\xf6\x61\xfa\xa9\x42\x48\x29\xe6\x95\x2d\x35\x3a\xda\x09\x27\xdb\x73\xeb\xf4\x37\xe3\x45\x77\xf3\x3b\xf6\x21\xf5\x4b\xa4\xa8\x18\xa6\xfb\x23\x71\x1b\xdd\x44\x66\x90\xde\x13\xf5\x1b\x64\x5f\x42\x64\xab\xc7\xf4\x39\xc5\xd9\xa0\x9b\xc9\x28\xe6\x8b\xc9\x95\x29\xc1\x67\x00\x00\x00\xff\xff\x81\x9d\x53\x8d\xf6\x01\x00\x00")

func cmdDefinitionsTmplPairTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplPairTmpl,
		"cmd/definitions/tmpl/pair.tmpl",
	)
}

func cmdDefinitionsTmplPairTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplPairTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/pair.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1b, 0x1a, 0xbf, 0xf0, 0xe9, 0x47, 0x35, 0xe6, 0xa7, 0xc, 0xe6, 0xba, 0xc0, 0x5e, 0x39, 0xda, 0x2d, 0xb9, 0xd0, 0xee, 0xbc, 0x1, 0xc0, 0xd1, 0x5e, 0xa2, 0xbf, 0x4d, 0xcd, 0x4, 0xe9, 0xc0}}
	return a, nil
}

var _cmdDefinitionsTmplServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xdb\x8e\x1b\xb9\xd1\xbe\xd7\x53\xd4\x36\x06\x3f\x24\x43\xa6\xbc\x7f\xf6\x26\x13\xcc\xc5\x66\xec\x38\x83\xf5\x61\xe0\xf1\xee\x5e\x78\x17\x03\xaa\xbb\x5a\x62\xd4\x4d\xb6\x49\x4a\x63\x41\xe9\x77\x0f\x78\xe8\xf3\x41\x92\x1d\x8f\x11\xc0\xbe\xb1\x86\x2c\x16\x8b\x75\xf8\xf8\x91\xec\xc5\x02\xae\x45\x84\xb0\x42\x8e\x92\x6a\x8c\x60\xb9\x87\x95\x28\xff\x86\x1d\xa3\x10\xa6\xd1\x22\xc2\x98\x71\xa6\x99\xe0\xea\x6f\xf0\xfc\x2d\xbc\x79\xfb\x1e\x5e\x3c\xbf\x79\x4f\x26\x19\x0d\x37\x74\x85\x70\x38\x00\x79\x43\x53\x84\x3c\x9f\x4c\x58\x9a\x09\xa9\x61\x3a\x01\x00\x08\x42\xc1\x35\x7e\xd2\x81\xfb\x8b\x89\x60\xe2\x7e\xad\x98\x5e\x6f\x97\x24\x14\xe9\x62\x89\x7b\xc1\x23\xa5\x85\xa4\x2b\x5c\xac\xc4\xd3\xe2\xe7\xee\xa7\x45\xb6\x59\x2d\x90\x47\x99\x60\xbc\xd0\x71\xc6\xc8\x50\x62\x84\x5c\x33\x9a\x9c\x3f\x76\xad\x75\x16\x26\x0c\xcf\x9d\x57\xa1\xdc\xb1\x10\x95\x1b\x45\x4e\x1e\xa7\xf7\x99\x19\x34\x9b\x4c\x76\x54\xc2\x3d\x54\xb6\x93\x5b\x29\x76\x2c\x42\xe9\x7b\x0a\x7f\x90\xdf\x68\xb2\x45\xdf\x78\xe7\x34\x15\x32\x85\x15\xe4\xce\xfd\x78\x21\xa5\x28\xfa\xaa\x95\x91\xb7\x99\x0d\xeb\x64\xb2\x58\xc0\xfb\x7d\x86\xc0\x14\xe8\x35\x82\x31\x06\x62\x21\x1b\x91\x0d\x05\x57\xda\x89\x5d\x41\x50\xeb\x09\xec\xf8\xb7\xcb\x7f\x61\xa8\x5f\xa3\xa6\x11\xd5\x14\xcc\xd2\x50\x15\x86\x40\x5a\xb4\x1b\xad\xc2\x8a\x92\x89\x9d\xa6\x33\x4e\x6e\x43\x0d\x87\xc9\xe1\xf0\x14\x24\xe5\x2b\x84\x8b\xfb\x39\x5c\xec\xe0\xf2\x0a\xc8\x0d\x8f\x85\x32\xc6\x18\xe7\x1a\x09\x16\x03\x7e\x84\x8b\x1d\xb9\x0b\x45\x86\x10\x38\xd5\x41\x4b\x84\x0b\x6d\x64\x5e\x26\x62\x49\x93\x7a\xdf\x45\xc6\xcd\x1a\x2e\xaf\x4c\xb7\x5d\xce\xbf\x41\x8b\x5b\xaa\xc2\xa6\x1c\x8b\x8d\xc0\x73\xa6\xb2\x84\xee\x0b\x87\x80\xff\x57\x53\x74\x35\x20\x66\x44\x90\x47\xd5\x9f\x56\x0e\x55\x28\x99\x0d\x41\xbd\xc3\x69\xca\x73\x2f\x65\x1d\x3e\xa4\xa6\xfc\xb3\xf6\x33\xb7\xe1\x78\x89\xba\xe5\xd9\x07\x96\x24\xb0\x42\xdd\xf6\x78\x2c\x45\xea\xdb\xc8\x64\xb1\x30\x83\x9f\xc2\xfb\x35\x53\x10\x6f\x79\x68\xad\x53\x6b\xb1\x4d\x22\xeb\xc6\x25\x42\x48\x93\xc4\xc1\x45\x11\x5d\x96\x66\x09\xa6\xc8\x35\x4a\x52\x8c\x47\x90\xa8\xb7\x92\x33\xbe\x6a\xcf\xc8\x14\x48\xa4\x11\x08\x9e\xec\x81\xf2\xa8\xa5\x3f\x15\x11\x8b\x19\x46\x64\x62\x0c\xe8\xae\x64\x2a\xe0\x89\x6b\x99\xb5\x35\x1f\xac\x5f\x44\x3a\x07\xb1\x31\x51\x15\xe4\x25\x6a\x5f\x03\xe5\xf0\x99\x15\x62\xb1\x91\x39\x94\x51\x74\xd6\x82\x48\xc9\xb4\xa9\xd4\x89\x3b\x8f\x7b\xa1\xa6\xc0\xa1\xf0\xb9\xea\xf7\xb9\xea\xfa\x9c\x71\x2d\x4e\xf3\xb9\xf5\x51\xe5\x74\xc1\x43\x9c\x43\x96\x20\x55\x08\x29\xdd\x20\xa8\xad\x44\xa0\x49\x02\x56\xf1\x9a\x2a\x58\x22\x72\x78\x90\x4c\x6b\xe4\xb0\xc4\x58\x48\x34\x36\x78\x77\x76\x8c\xac\xdc\x39\x87\x32\x13\xca\xc5\x17\x2e\x25\x77\x5d\x47\x8a\x74\xe6\x57\xee\xe1\xe7\x6e\xaf\x34\xa6\x1d\x0c\xb0\xad\x4d\x08\xf0\xc8\x67\x1b\x3d\x10\x0c\xe9\xf8\x22\x3c\xf0\xf3\x04\xed\x7a\x1d\x02\x85\x9e\x9a\x1e\x05\x87\x96\xce\x61\x90\x38\x03\x2c\xea\xa2\xb5\x6a\xaf\x9a\x07\xc0\xa3\x2e\x30\x02\x22\xe7\x03\x49\x7f\x64\x4a\x3c\x69\xb5\x5b\x3c\xf1\x43\x5e\xdb\xf0\x56\x09\x5e\x07\x85\x7e\xad\x67\x63\x43\xaf\x9a\xa9\x82\x27\x35\x13\x66\x03\x93\xb9\xdc\x56\x25\x5c\x28\x0b\x17\x4d\x4d\xa3\x68\xa1\x52\x32\xed\x55\xdd\x03\x1a\xbd\x72\x75\xec\x18\x71\xb3\xea\xba\xd9\x42\xc8\x80\x9b\x1f\x09\x47\x4e\x71\xfd\x1c\x54\xda\xbf\xf4\x02\x5a\x94\x85\x96\x96\x0e\x87\x2c\xbd\x35\x7f\x4b\x99\x54\x45\xaa\xf6\xd5\xf1\x29\xb5\x5b\xd5\xeb\xb5\xe0\x71\xc2\x42\x6d\x5a\x17\x0b\x78\x8e\x99\xc4\xd0\xf0\xe1\x4b\xf8\x55\x21\x64\x66\x36\xf2\x3b\xd3\xeb\x46\x59\x31\xae\x34\xd2\xc8\x3a\xbd\x56\x32\x8b\x05\x74\x44\x6d\x00\x69\x96\x25\x7b\x5f\x8c\xbe\xd6\x61\x67\xf8\x1b\x98\x8d\xc0\xd1\x30\xaf\xac\xa7\xba\xad\xc3\xdb\x8a\xa7\xbb\x66\x71\xcf\xc0\x78\xa6\x9b\xa3\xa6\xb5\x6a\xfc\x05\xf7\x97\x96\xbd\x55\x86\x04\xf3\xb2\xd7\x52\xca\x4b\xd8\xcd\x7d\xfa\x36\xf0\xa0\xf6\xd3\x52\x49\xe3\x9a\xd7\x34\x83\x2b\x48\x69\xf6\x41\x69\xc9\xf8\xea\x4f\xf7\x9f\xb7\x62\x3c\x7e\xe7\xe0\x6c\xcb\xe2\x72\x09\x7e\xed\x7e\x09\x4d\xf0\xea\x9d\xdd\x68\x50\x19\x0d\xb1\x61\x82\xc6\x34\x4b\xcc\xa1\x27\x60\x86\xc4\xc4\xa6\x3f\xb0\x75\x71\x97\x18\x82\x53\x4c\x7d\xb1\x23\x37\xa5\x40\xbf\x82\x18\xa9\xde\xca\xc1\xe1\xff\xd8\xf2\x50\xd9\x1f\x5e\xae\x5f\x8b\xf1\xed\x3d\xc7\x87\xba\x96\x69\x9f\x7f\x66\xb6\x11\x1f\x86\xd5\xb4\x2c\xa9\x8c\xe8\x37\xdf\xe3\xc6\xb8\xfd\xad\x6c\x30\xbf\xed\x31\xb1\xe5\xc0\xaa\x18\x8d\xef\x19\x8f\xf0\x13\x10\x78\x56\xb6\x5b\x59\x55\xef\xfb\xd1\xf4\x99\xe4\x9a\x0e\x26\x50\x31\xca\x9b\x7f\x5f\xaf\xaa\x46\xe2\xc0\x15\xfc\x9f\x2b\x98\x66\xfb\xa1\xb3\xe9\xcd\x06\x57\x53\x45\x73\x68\x2d\x5d\x5c\xb9\x88\x79\x77\x4d\xae\xa3\xd1\xfe\xff\x76\x36\xcb\x7d\xbc\x99\x79\x5e\x66\x45\xc9\x7b\x8c\xa9\x2d\x5c\x72\x00\xcf\x30\x89\x2a\x9c\x8e\xca\x7e\x43\xcc\x5f\xde\xdd\x3e\xfd\xf1\xd9\x5f\x0d\xc2\x53\xce\x4d\x9b\xd9\x79\x53\xb1\xc3\x08\x18\x87\xdd\x4f\xe4\x2f\xe4\x19\xb1\xba\x5f\x09\xa1\xf0\x6d\x66\xce\xfc\x4c\xf0\x9f\x93\x04\x96\x42\x24\xc3\xde\x37\x8b\xf3\xae\x7f\x3c\xb3\x06\x43\xdc\xb0\xb5\x08\xe0\xd7\xb6\xed\x37\x26\xf5\x96\x26\x27\x39\x2d\xe6\x6d\xaf\x15\x5b\x4f\xcc\xc9\x1d\x4b\xb7\x89\x35\xe0\x11\x3c\xda\xb6\xda\xf8\x34\xe6\xa7\x3a\xf5\x1b\xf9\xd8\x6c\x16\x9f\xe5\xde\x04\xb9\x5d\x9e\xd7\xf3\x78\xfe\xb5\xdb\xed\x90\x6b\x3b\x4b\x69\x79\xb7\xa7\xda\x8a\x15\x3c\x62\x8a\x8c\x2e\xe1\xac\x52\xec\xcb\x9a\xfe\x75\xaa\x93\xee\x46\x4e\x9e\x78\x78\x7f\xaa\x76\xd6\x63\xbb\x53\xcc\x3b\x28\x5e\x74\x6c\x8a\xe0\xf4\xd0\x15\x43\x02\x33\xef\x41\x8b\xe9\xce\x95\x1b\x4b\x1b\xdd\xfd\x5a\x46\xa5\xc2\xc8\x63\xbc\xdb\x01\x5a\x23\xcc\x80\x3c\x6f\xee\x02\x76\x33\x87\x0f\x7f\x9a\xf0\x94\x05\xf8\x0e\x3f\x6e\x99\xc4\xc8\xf5\xf6\x39\xd8\x74\x14\xe6\x96\xd2\xde\xa7\xff\xa4\xca\x4e\x4a\x99\xec\xf3\x2c\xd4\x7d\x3b\x26\x56\x75\x8f\x5c\x57\x2d\x16\x9e\xe9\xd2\xe4\x34\x6b\x4b\xe9\xff\xb6\xb5\xc7\xed\x2d\xc2\x28\x15\xde\x0e\xc4\xd2\xf2\x7a\x2b\xe1\x88\xb7\xb2\x54\xc9\x1e\xcb\x9e\x0c\xc4\xdf\x11\xf9\x31\xad\x53\x91\xe9\x22\xc6\x33\x98\x0e\xe8\x99\x03\x4a\x29\x64\x71\x82\x92\xa8\xb6\x89\x36\x5e\x1b\x90\xaf\xe8\xbf\xf5\xfb\x25\x98\x59\x0a\x8e\x6f\xff\x8b\x85\x84\xfb\x39\xd8\x62\x74\xd1\xb0\x86\x54\x03\xd5\x03\xd3\xe1\x1a\x76\xe4\x17\xdc\xd7\x9a\xfb\x33\xf0\xcc\x2c\x34\xff\x42\x73\x10\x0d\x9a\x71\x33\x44\xbf\x71\xb3\xc1\x62\xbf\x56\x72\x24\x15\x0e\x9d\x7b\x97\x50\x70\xcd\xf8\x16\x1b\x1d\xcd\x7b\x93\xd3\x54\x5f\x81\x96\x2d\x35\x7e\xe0\xf8\xa8\x9d\xbb\xaf\x27\xd3\x76\xea\xcd\x86\x6e\x65\x86\x6b\xe6\xcc\xba\xf9\xee\xdf\x41\xff\xe6\xf5\x22\x38\x23\x65\x59\x0c\x3f\x9c\xeb\x2a\x7f\x1e\x1f\xaa\xd1\x7c\x5e\x3d\xdb\x98\xf2\x2f\xa6\xb4\x6f\x37\x07\x73\x70\x57\x97\x1f\xfc\xf9\xfa\xd0\x13\x4a\xc8\xcb\xd5\x40\xcf\xb6\xeb\x67\x77\x46\xcf\x81\xb3\xe4\xc8\x26\x79\xc2\xf9\xcd\x5e\x15\xf5\xef\x91\x99\xdf\x23\xb3\xc1\x3d\xf2\x39\xc6\x74\x9b\x68\xe7\x0a\xe3\x01\x77\x31\xc0\x94\xb1\xc2\xf4\xf8\x0d\xcf\x5e\x17\x67\x18\xb2\x98\x85\x40\xed\xe1\xd4\x6d\x99\x03\x0a\x46\x6f\x8c\x9d\xc9\xc7\x18\x85\xdf\x63\x9b\xfb\x41\x4f\x86\x74\x55\x42\xef\x11\xbc\x71\x96\xce\x38\x5c\xec\x46\x2f\x5e\x3b\xa1\xf8\x02\xae\x32\x1e\x87\xb3\xe9\xcc\xe6\x7c\x3e\xd3\x1a\x72\x84\xd2\x7c\x0b\xf2\x72\x1a\x7d\xf9\x1f\x22\x2a\x9b\xcf\x61\x2a\x9b\x1e\xaa\x32\x55\xf0\xa4\xea\x9d\x8d\x4e\x33\x48\x5d\x36\xe7\x72\x97\xcd\x23\x91\x97\xef\x14\xe5\xec\x2d\xb4\xd7\xce\x9e\x7d\xf5\x3b\x3b\xf9\x6a\xae\xf5\xbb\x63\xd3\x37\xc5\xa5\x52\xc6\xcd\x0a\xfc\x2d\xa2\xbd\x43\xb3\x15\xdf\x7a\x21\x5c\x2c\x20\x31\x7d\xf7\x36\x2e\xfe\x8e\xd3\x00\x82\x14\xd1\x36\x74\xf7\x02\xfe\xda\x80\xb4\x07\xde\xc4\xb0\x55\x28\x01\x39\x5d\x26\x08\xda\xde\x3f\x38\x0d\x25\x81\x29\x1e\x9d\xd8\x8a\x0b\x89\xf6\x85\x46\x6d\x33\xfb\x5d\x90\x9d\xd1\x22\x00\x69\x07\x57\x91\xe2\xb6\x95\x54\x96\x7f\x56\x2c\x07\x5e\x4e\x3b\x04\x6c\x33\xce\xc0\x7e\xe5\xde\xea\x82\x84\x99\xc6\xcb\xdd\x00\x7b\x5c\x2c\xe0\x7a\x8d\xe1\x06\x64\xe3\x40\x44\xbe\x09\xb5\x3c\xb2\xb2\x6f\xc9\x2d\x6b\x2f\x0c\xcd\x47\xa0\xb1\x7b\xf5\x52\xac\x9f\x72\xf6\xb8\x78\x98\x9b\xb9\x77\xbf\x57\xa2\xf5\x8c\xdf\xa4\x6d\x89\xe9\xbe\x2f\x6c\x1d\x24\x70\xa5\x37\x12\xd5\x79\x60\xef\x3e\xaa\x1c\x57\x73\xec\x4d\xbe\xd7\x93\x9f\x4b\x0f\x4f\xa1\x7e\x35\x1a\xe1\xaf\x04\x5b\x57\x73\xee\xd5\xb9\xf9\xe6\x6c\x09\x47\x28\xd1\x2c\x9e\x82\xff\x06\x10\x96\xfb\x02\xbb\x48\x9d\x5d\x5c\x64\x3c\xcf\x67\xb5\x89\xa6\xf6\xfa\x8b\xdc\x52\x49\x53\x45\xee\x6c\x52\x1a\x09\xdf\xfe\xce\x66\x59\xbd\xc3\xd5\x41\xa8\x3f\x99\xb5\xf8\xd9\xc8\xdf\x69\xb8\x59\x49\xb1\xe5\x91\x7f\xc2\x2f\x9e\xec\x49\x35\xd3\xef\x4c\xaf\xaf\x9d\xfc\x34\xd4\x9f\xe6\xd0\x98\xf9\x9a\x26\x09\x4a\x03\xd1\x6d\x57\xd4\xc6\x0d\x78\x65\x6c\x7d\xad\x59\x4b\x8b\x7d\x5b\xcb\x8a\x93\xd7\x1f\x61\x8c\xd2\xc6\x60\x3a\x6b\x72\x9c\x8b\x8c\xea\x75\x11\x67\xaf\xf6\x96\xea\xb5\x5b\x60\x4f\x79\x50\x1e\xc1\x14\x3f\xfa\x81\x41\x30\xf3\x7f\x71\x08\xfc\x37\x40\xc1\xac\xf7\xab\x13\x23\x7e\x05\xc1\xfc\x8f\xe0\x8f\xa0\xf3\xcd\x4d\x0b\x91\x51\x4a\xb8\x32\xa0\x2f\x64\x4a\xb5\x45\xa2\x69\xe0\x96\x68\x72\x31\xcf\x03\x4b\x15\x2b\xc5\x79\x0e\xfe\xb3\x8a\xe9\x6c\x52\xbf\xeb\xe7\xe8\x88\x85\xfb\x72\x49\x44\x08\x41\x13\x5b\x52\xd3\x56\xd0\x8f\x4a\xaa\x0b\x31\x06\x7b\x05\x31\x9d\xe4\xc6\x42\xaf\x1d\x99\xe7\xd3\x59\x0d\x6c\xbd\xe1\x05\x9e\x56\x0a\x6f\xf8\x8e\x26\xcc\x83\xea\x8b\x4f\x19\x86\xf6\x8a\xde\x74\xd5\x74\xcd\xe1\xe7\x50\x6f\x69\x72\x09\x6e\xa6\xbc\x85\xe2\x23\x38\xeb\x8e\x4b\x57\x40\xb3\x0c\x79\x64\x29\xb6\x9a\x83\x22\xbe\xac\xec\xf9\xb7\x96\xde\x84\x10\xe7\xb0\x1d\x95\x86\x12\x0f\x5d\x3d\x38\xe5\x22\xd3\xf3\x32\x28\xa3\x77\x94\x76\xde\xf2\xbb\x18\x33\xe4\x87\x2b\x03\xfa\x9d\x0d\xa9\xbe\x3d\xb6\xca\xaf\xc2\x9b\x6b\x9a\xa2\x09\x40\x4f\x05\xbe\x97\x2c\x7d\x45\x95\xf6\xa5\xf8\x82\x47\x86\xc9\xaf\xaf\x45\x9a\xd2\x3c\x37\x16\xcf\x46\x36\x9a\x36\x84\x8f\xed\x36\xf5\xbe\x3e\xb0\x2c\x32\xfc\x08\x60\x7a\x1e\x70\x14\x34\x9d\x5c\x21\x7e\x2e\x78\x9a\x61\x5f\x01\x40\xbf\x24\xc3\x4e\xce\x32\xbf\xe4\x1b\xc7\x0c\x2d\x11\x84\x87\x35\x4b\x10\xd6\x94\x47\x09\xe3\x2b\xb0\x71\x33\x0b\xf4\x5f\xc5\x14\xc3\x6c\x82\xde\x9f\x9c\x9e\x6d\x72\x64\xed\x1e\x48\xbd\xb3\xb3\xce\x25\x76\x3d\xf3\x6c\x40\x18\x67\xba\xc4\x8a\x13\xbf\x3f\x31\xff\x4a\x30\x79\x87\x2b\xa6\x34\xca\xa1\xab\x29\x39\x35\xe7\x86\x39\xbc\xc1\x87\x41\x91\x59\xdf\xed\x45\x67\x86\xbb\x70\x8d\x29\xf5\xea\xfc\xa7\x3c\xa6\x9a\xfe\x13\x00\x00\xff\xff\x07\x4d\x89\x28\x18\x30\x00\x00")

func cmdDefinitionsTmplServiceTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplServiceTmpl,
		"cmd/definitions/tmpl/service.tmpl",
	)
}

func cmdDefinitionsTmplServiceTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplServiceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/service.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc3, 0x6a, 0x88, 0x31, 0xf7, 0x95, 0xcf, 0x21, 0x9c, 0xed, 0x27, 0xb7, 0x6b, 0xa5, 0xfd, 0x3f, 0x89, 0xd2, 0x2a, 0xfd, 0x7f, 0x66, 0x29, 0x85, 0xa9, 0x28, 0xf4, 0xe2, 0x71, 0xc6, 0xd6, 0x17}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"cmd/definitions/tmpl/function.tmpl":  cmdDefinitionsTmplFunctionTmpl,
	"cmd/definitions/tmpl/info.tmpl":      cmdDefinitionsTmplInfoTmpl,
	"cmd/definitions/tmpl/object.tmpl":    cmdDefinitionsTmplObjectTmpl,
	"cmd/definitions/tmpl/operation.tmpl": cmdDefinitionsTmplOperationTmpl,
	"cmd/definitions/tmpl/pair.tmpl":      cmdDefinitionsTmplPairTmpl,
	"cmd/definitions/tmpl/service.tmpl":   cmdDefinitionsTmplServiceTmpl,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"cmd": {nil, map[string]*bintree{
		"definitions": {nil, map[string]*bintree{
			"tmpl": {nil, map[string]*bintree{
				"function.tmpl": {cmdDefinitionsTmplFunctionTmpl, map[string]*bintree{}},
				"info.tmpl": {cmdDefinitionsTmplInfoTmpl, map[string]*bintree{}},
				"object.tmpl": {cmdDefinitionsTmplObjectTmpl, map[string]*bintree{}},
				"operation.tmpl": {cmdDefinitionsTmplOperationTmpl, map[string]*bintree{}},
				"pair.tmpl": {cmdDefinitionsTmplPairTmpl, map[string]*bintree{}},
				"service.tmpl": {cmdDefinitionsTmplServiceTmpl, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
