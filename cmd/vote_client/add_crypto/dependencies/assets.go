// Code generated by go-bindata.
// sources:
// ../assets/Bitcoin.png
// ../assets/ethereum.png
// ../assets/litecoin.png
// DO NOT EDIT!

package assets

import (
	"bytes"
	"compress/gzip"
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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
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

var _AssetsBitcoinPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\x8f\x09\x38\x94\xdb\x1b\xc0\x3f\x4c\x48\x14\x2d\xcc\x7c\x23\x71\x49\x76\x4f\x4a\xa5\x7f\xf7\x9a\x71\x15\xa9\x14\x91\xa5\x24\x59\x5b\x5c\x2d\x2a\x25\xa4\x94\xc2\x58\xc6\xcc\x20\x89\x4c\xd9\x9b\x5b\x1a\x09\x59\xe2\x96\x2c\x57\xb2\x65\x2b\x22\x25\x4b\xb9\x83\xf9\xf6\xef\xfb\x3f\xf4\x3c\xe7\x9c\xe7\xbc\xef\x39\xbf\xdf\xfb\xbe\x77\x0e\x38\xda\xa9\x28\x81\x4a\x00\x00\xa8\xec\xb6\xb7\x75\x06\x00\x80\xb5\xb0\x15\xe5\x00\x00\xc8\x2a\x8f\x9e\x02\x00\x39\xe0\xc0\x5e\x97\x9d\x76\x67\xf6\x52\x51\x80\x34\x15\xa4\x00\x40\xca\x05\xe7\x2e\x00\x50\x2a\x48\x51\x00\x55\xb4\x90\x94\x72\x17\xd6\xfc\xe2\x89\x0a\x7e\x5d\x18\x28\x9f\x01\xa7\x82\xf3\x5c\x75\x29\x97\x01\xa5\x82\x50\x32\xbd\x36\x6a\x81\x82\x17\x42\x06\x94\x42\x9f\x4d\xd6\x80\xb9\x20\xe5\x0a\xa0\x7c\x06\x2a\x00\xb1\x34\x26\xca\x07\x89\x74\x26\xcc\x67\x2c\xe0\x02\xb0\x35\x03\x40\x16\x25\x0b\x78\xea\x02\xb5\xa8\x62\x20\x3c\x06\xcc\x5b\x94\xf0\x40\x98\xc7\x00\x02\x81\xc5\xa2\x74\x84\xbf\x10\x22\x29\x74\x68\xe1\x75\xa1\x19\x28\x75\xf1\x27\x97\x31\x9f\x02\x16\xb9\x02\xb8\xe0\x97\x0a\x94\x26\x83\x73\x49\x20\x26\x00\x21\x2e\x03\x13\x2c\x52\xfc\x85\xfc\x7c\x0a\x13\xe3\x83\xd2\xc5\x12\x10\x17\x94\xa4\xd0\xa5\x8b\x36\xf8\xd7\x44\x69\x20\xcc\xa3\x23\x3c\x50\xc2\xfd\xd5\x09\x48\x51\x14\x45\x12\x10\x4f\x93\xa2\x48\x88\xbf\x16\xe2\xad\x85\x52\x99\x54\x96\x26\xc4\xd7\xa2\x32\x98\x14\x81\x51\x04\x4e\xdd\xd7\xa4\x48\x5c\x9a\xa0\x04\x71\x94\x29\x92\x84\xb8\x74\x88\xc7\xa4\x32\x99\x50\xa2\x2a\x94\xb2\x66\x61\x8a\xe4\xd5\x30\x0f\xc4\xf8\x20\x94\xa0\x0c\x71\x96\x53\x38\x02\xa5\xa8\x53\xd9\x9a\x54\x3a\x93\x4c\x63\x22\x3c\x10\x4a\x52\xa3\x08\x14\x4a\x5c\x21\x8d\x57\x92\x72\x94\xe1\xfb\x1b\xf1\x8e\x6c\x88\xab\x41\x66\x30\xa9\x7b\x9a\x50\xf2\x2a\x0a\x83\xf0\x0f\xc5\x14\x0e\x23\x7c\x26\x26\x60\x52\xe8\x3c\xf2\xcc\x0b\x6b\x13\xc0\x02\x1d\x28\x41\x09\x29\xb0\xc3\x3b\x1f\xa0\xe5\x01\x14\x26\xc5\xbb\xf3\xd0\xf2\x40\xbc\xf7\x31\x9c\x65\x8e\xb5\xa6\xc0\x0f\xb6\xc1\xf7\x4c\xa1\xf8\xa5\x58\x7d\x04\x95\xa9\x89\x36\x44\x11\x43\x95\xd2\x04\x65\x62\xbc\x0d\x29\xf3\xc1\xda\xd2\x89\xa9\x0f\xc4\xb7\x56\xe9\x9d\x25\x78\xff\x53\xac\x99\x03\xf1\xd7\xc1\x99\xc6\xe4\xcc\x10\x9c\xae\x87\xf2\x41\xb4\x2a\x84\x9c\xe8\x84\x12\x97\x13\x93\x5d\x70\xae\x15\xf2\xc4\x8d\xfc\xd1\x4f\xc2\x33\x68\x6d\x18\x2c\xdc\x41\x4a\xa7\xd1\x8a\x93\xd2\x38\x79\x72\xb2\x0b\xef\x7f\x82\x56\x9c\x26\x67\xbf\x4a\xe3\x15\xc9\x1f\x83\x68\xf5\x39\xe4\xe1\x1f\x52\xce\x72\xa4\xd4\x13\x7e\xc4\x22\x7f\x7e\xa2\xe0\x19\xe4\x89\x3b\x24\xd0\x26\xbe\xb7\x13\x23\xf5\x70\xf6\x66\xf2\xe7\x20\xf1\xb5\x85\xc2\x51\x28\x49\x15\x79\xee\x4b\xce\x8e\x11\xa3\x0d\xc4\x60\x19\x94\x69\x84\x35\x27\x21\x22\x67\x62\xa4\x01\x2e\xb4\x27\x86\xab\xe1\x47\x6c\x44\xec\x8d\xd5\x47\x91\xd3\xbd\xd8\xeb\x9b\x78\x97\x90\x42\x24\x48\x91\x03\xfa\xe2\x04\xde\x53\x20\xbd\x09\x10\x53\xbd\x14\x45\x60\x8d\x77\xd0\xd7\x31\xc4\x58\x23\x94\xb1\x81\x98\xe8\xc0\xfe\x15\x48\x63\x01\xf2\xbf\xcf\x48\xf1\x3e\xea\x2e\x13\x11\xb9\xa0\x75\x97\xa1\x34\x5d\x62\xec\x2d\xfe\xee\x2e\xd6\x94\x40\xce\x8d\xc3\x0f\xb6\x42\xb1\x00\x3e\x28\x2e\xb9\xb0\xcf\x1f\x00\x80\x23\xa1\xce\x8e\x07\x01\x1a\x29\x43\xca\x62\x32\x72\x04\x8e\x91\xe8\x3c\x0e\x25\xa2\xf4\xe4\x09\x75\x1a\x61\x90\x68\x68\x7a\x5c\x2e\x18\x27\x48\xf2\x52\x6f\xe7\x12\x76\x52\x4f\x37\x7c\x14\xf2\xca\x2f\x7c\x3c\x25\xf3\x2d\xf6\xd6\xf4\x1a\xa8\xa2\xb2\x8d\x53\x70\x4a\xae\xf6\x9f\x53\x3e\xc7\xea\x1b\x27\x7f\xde\xc4\x35\x27\xc2\xc2\x26\xac\xad\xdb\x2b\xdf\x7f\x33\x7a\x6c\xf0\x2e\xbb\xf0\x6e\x27\x00\xc8\xf7\xec\xb6\x65\xbb\xb8\x37\x3c\x72\x53\x70\x39\xe8\x46\xaf\xd1\xb9\xb3\x5b\x51\xfb\xed\x81\x33\xc3\xb6\x75\xd5\xf3\x39\x3d\x1e\xd0\xfd\xe0\xf3\x43\x2d\xe1\x4d\x2d\x8e\xda\x4b\x8d\x7d\x6e\x94\x5c\x17\x5e\x90\x55\x72\x87\x4a\x3c\xa1\x4f\xc5\x23\x22\x61\x4f\xcf\x0f\xaf\xb2\xaa\xb2\xc9\x1a\x8d\x1d\x5f\xab\x8f\xfc\x73\x1e\xcd\xd6\x6c\x1b\x1d\xfb\x64\x4d\x11\x3f\xa0\x28\x8a\x92\x1b\xf6\x75\x56\xc3\x36\x94\xda\xd9\xdb\x5c\x0d\x7b\xe2\xed\x30\x23\x36\xbb\xf1\xe7\x95\xd6\x35\xc2\xb4\xdf\x18\x41\x83\x7e\xe7\x80\x2f\x4f\xcf\x6e\xaa\x9a\xe4\x86\x7a\xfc\xef\xa0\x4b\x11\x2f\xcb\x63\xfc\x29\x96\xfc\xa2\xf7\xeb\x78\xe6\x2e\xaf\x9c\xe9\xd4\xed\x6f\xfc\x4e\x5e\xe8\x53\x8a\xda\xb9\xcc\x7a\x8d\x45\xad\xfa\xf8\xc6\xc3\x81\xad\x09\x19\x60\xb3\x69\x58\xe2\x6e\x7f\xed\xd0\xba\x6d\xbe\x01\xd7\x9c\x5f\x4d\x47\x04\x44\x2a\xf9\xb3\xf2\x2e\x4a\x62\x14\xdb\xdd\xd0\x97\xd9\x1a\x76\xef\x36\xe6\x19\xad\x31\x2a\x31\xe4\x15\xf6\x5b\xbd\x77\x6e\x7f\x19\x3d\x7d\xdc\x2a\xc8\xd2\xa6\x8a\x65\xe8\x64\x19\xb4\x72\x4c\x76\x3a\xfd\x10\xcb\xec\x4d\xee\x06\xb9\xc7\x37\xb8\xf7\xcb\x3f\x8e\xfe\x66\x58\x60\x59\xa5\xfe\xcc\x5b\xd3\xd9\xa2\xb6\xb0\xc8\x69\x36\xdb\xdb\xc2\x65\x6a\x92\x6d\xb4\x6a\xbe\x79\x26\xad\xe0\x68\x9d\xde\x3e\x8b\x2b\x71\x7e\xb2\x06\xfe\x1d\x9f\xab\x39\xd1\xe9\xdd\xe0\xbd\x78\xd6\xa8\xf1\xb9\x48\x65\x8b\x8a\xa9\xf5\xc2\x95\xe5\x3d\xbb\x2d\xee\xd4\x84\xf0\xdf\x1c\xb5\xd0\xca\xb7\xd5\x36\xd2\x55\x65\x9a\x67\x04\x9a\x75\x17\xd8\x3b\xab\xee\x3e\xb6\x8b\xae\x7a\x79\x88\x0a\x0a\xde\x74\x04\xdf\x59\xeb\x4d\x41\xc3\x3a\x5b\x4f\x94\x0f\xd8\x93\x6e\xff\x8a\x85\xd9\x29\xa5\x2a\xe9\xcc\xae\x2d\xc0\xf9\x8b\x1d\x91\xeb\x6c\xe0\x9a\x3e\xe9\xe4\xbb\x81\x73\xf0\x90\x82\x58\x25\x8e\xd6\x91\x4f\xe3\x26\x79\x99\xbe\xb0\xce\xe7\xe4\xda\xab\xf8\x8a\xca\x04\xab\x20\xc9\x3d\xec\xf3\xfe\xa4\xb5\xf8\x41\xf8\x6f\x2d\xc3\x60\x5d\x91\x75\x9d\xc1\x63\x80\xb5\xe1\xb5\xc9\x96\x6f\xdb\x26\xff\xfa\x14\x3d\x33\xe3\x85\x15\x8d\xd1\x4a\x36\x1f\x94\xd5\x2e\x6b\xd8\x11\x9b\xdc\x24\x3f\x7f\xa2\x31\xd0\x3a\xd8\xe3\xe2\x95\xd7\xa2\x67\x7e\x1c\x4c\x3d\x66\xdd\xea\x31\x57\x56\xda\x15\x8f\xbb\x75\x6c\x56\xd5\x1f\x6a\x83\x5e\xcd\xaf\x00\x1f\x80\x57\xa1\x9a\x95\xd5\x5b\x1b\x6a\xcc\xae\xeb\x8f\x3a\xa9\xa4\xf3\xd9\xd3\xd5\x01\x1d\xf4\x8f\xda\x42\x63\xdc\xa6\xf7\x1e\xad\x63\xe5\x7f\x0d\x86\xaf\x98\xf9\x71\xae\xf9\x63\x45\x5a\xec\x17\x29\x7a\x11\xf4\xbe\xc8\xab\x06\xf5\x4e\xb4\x3c\x45\xb9\xf6\x17\xeb\x8e\x5e\xd6\xb0\xb1\x7c\xb0\xf5\xac\xa3\x60\x95\xbc\x3b\x5c\x96\xe9\x30\xb7\xf4\x81\xee\x9e\x5b\x6d\x59\xdb\xe1\x2f\xb5\x9d\x4e\x11\xa4\x85\xb8\xb4\x6a\x5f\x85\xbc\xe2\xe3\x89\xc3\xfb\xa5\x79\x56\x34\x83\xbd\x53\xed\x85\x9b\x64\x7d\x76\x9d\x5d\xa6\x79\x73\xcb\xdf\x9e\x7f\xb9\x9e\x97\x61\xa5\x46\xc7\x0b\xcf\x84\x6f\x37\x3d\x90\x3e\xf2\x5d\xf1\x1f\x83\xd3\x2b\x93\x77\xce\xaa\xdb\xf4\xbc\xba\x3b\xa0\xe0\x6e\xd0\xf5\xb6\xa9\x29\xfc\x3b\x6f\x58\x7f\xb4\x7d\xfd\xcb\xbc\x06\x43\x7a\xb0\x47\x12\xe0\xd2\x20\x99\x10\x4f\xcc\xf3\x2b\xa6\x9e\x69\x55\xd6\x3c\x99\x3d\x7e\x60\x3c\x20\x9e\xd5\x22\xea\x8b\xa4\x2b\xec\x25\xe5\x0b\xec\x06\xb4\x0a\xed\xcb\xd7\x4f\x4f\xe9\xab\xab\xfe\x8e\xef\xca\x11\xf7\x79\xe9\xb2\xb7\x9e\x06\x8c\x38\xbe\x6b\xfb\x45\xb1\x37\x03\xf4\x53\x02\x24\xe9\x9c\x71\xeb\xb9\xbe\xc6\x58\xb6\xd3\xa0\x4a\xcc\x7a\xc4\x57\xba\x41\xff\xbc\x6a\x7a\x4e\xec\xb5\x0b\xf9\xcb\x6d\x68\xd7\xd9\x5e\x99\xef\x66\x56\xfe\x66\xa8\xdc\x5a\xbf\x8b\xe5\xa1\x75\x46\x75\xb0\x37\xb3\x75\xe5\x6d\xab\x4b\xb2\x56\xb3\x17\x72\xda\xd7\x5f\x33\xf3\xd8\xa6\xe3\xd7\x1a\x72\x46\xb5\xe4\x7f\xae\x13\x7a\x27\x3c\x9f\x21\xdb\x98\x32\x32\xeb\xdc\xc3\xec\x92\xfd\x5e\xdf\x98\x71\xb1\x9e\x52\x13\x09\xc7\x68\xf8\xc3\x11\x55\xff\xfd\x43\xc1\xc5\x96\x6a\x26\x96\x1a\x87\xac\xc2\x42\x79\x0f\x6f\xdd\x18\x5e\x3d\x18\x19\x1b\xbb\x64\x5c\x68\x1d\xe1\xdf\x92\x15\x6f\x92\xda\x65\xe6\x1e\xa4\x76\xc0\x70\x3d\xfd\x18\x3f\xf9\xb9\xc9\x71\x27\x3b\xc1\x9e\xae\x6f\xab\x4f\x4d\x7e\x32\x59\x59\x24\x5e\x57\xb7\x16\x60\xcf\xd0\xd9\x80\x0f\x62\x23\x68\x32\xfb\xbd\xca\x3b\x22\x78\x6c\x79\xec\x7e\x87\xa2\x19\xbb\x76\xbd\x10\x9b\x02\xe6\x70\xe0\x8a\xa1\x81\x33\xad\xab\xae\xb9\x7a\x1d\x3a\x19\x78\xcf\x28\x5b\x37\x26\xdf\x01\x3c\x1e\x2d\x73\xe8\x74\xcd\x89\xe7\xa6\x86\xad\x63\x3b\xea\x88\xd9\x98\xdc\x91\x58\x56\x42\xe3\x11\xdd\xd7\xec\x9c\x5b\x22\xd7\xa6\x01\xba\x6b\xdb\xc8\xb2\xf0\xe1\x72\x93\x54\xf3\x01\x1a\xdf\x81\xdb\xbd\x44\x0f\x5e\x79\xa7\xc8\x78\xe6\xc9\xb1\x24\xce\x9e\xd6\xe0\x13\x03\xfb\x8b\x96\xf6\x34\xa4\x94\xe1\xa1\x2e\xe3\x01\x88\x73\x5f\xc0\xe1\x97\xcf\x03\x56\xa5\x29\xee\xd1\x08\x0f\x73\xbb\x14\x6a\x1c\x5f\xf8\xca\x53\x3f\x39\x40\xd0\xfc\x31\xee\xb2\x5a\x8e\xa7\xc4\xf1\x73\xe9\xcf\xb9\xdb\x32\xfa\x61\xb9\x75\xcf\xee\x1b\x0e\x3b\xbf\xa2\x7d\xf8\x92\x72\x5d\xb1\x32\xda\x0d\xf9\xbe\x37\xc3\x2d\xcc\x4f\x91\xfd\x64\x46\x74\xf3\x21\x23\x2e\x20\x5d\xd8\x58\xe1\x72\x8d\xbf\x84\xa7\x9d\x75\x58\xf6\xfd\xe9\x10\x55\x55\x52\xe1\xc8\xd8\xbd\x74\xcf\xee\x58\x63\x51\x7f\x87\xe4\x83\x67\xab\xa5\xf7\xd7\x52\x34\xe8\x5a\xc8\x7f\x49\x39\xed\x2a\xec\x4a\xad\x6f\xa6\xcd\x3c\x24\x3c\x31\xec\x4a\x92\xa6\x99\xb5\xb4\x3b\xbf\x2a\xf1\xea\xcb\xbc\x3f\xf4\x3b\xed\xac\xf4\xbc\x05\xc2\x64\x7d\x47\x48\x9e\xed\x5b\xbc\xbd\x56\x12\xfb\xae\x62\xfe\x50\x3d\xad\xfa\x5f\xc0\x5a\x3c\xab\x13\xe4\xb4\x43\x6e\x8d\xba\xde\xe4\xaa\xea\xa3\xdb\x4d\xbe\x7e\x29\xda\x99\xb4\xe2\x7c\xae\xa3\xaa\x7c\x75\xd3\x17\xc3\x15\xfa\xd7\xdd\xeb\x2f\xe5\xb0\x4a\xdb\xa4\xea\x6b\x2f\x97\x96\x3c\x28\xb4\xb5\xbb\x4b\x80\x62\xb2\xac\x31\xd5\x4e\xbe\xfa\x95\xdf\xb1\xce\xcf\x09\x3b\x0d\x9a\x0f\xb6\x5f\x34\x54\x5c\x6b\x6e\x12\xbc\x57\x12\xa1\x37\xba\xdf\x24\xfc\xef\xae\x7c\x9d\x53\x4c\x0d\x0f\x7b\xf6\x09\x26\x3d\xda\x48\xb7\x28\x89\xb8\xc9\x8f\xde\x48\xc9\x87\xd4\xed\xf9\x7e\x39\x2a\x47\x73\xd0\xad\xea\xc1\xe9\x96\xb8\x0d\x69\xf6\x3e\x1e\x4e\x27\x2b\x46\xe2\x9a\xc3\x6d\xb3\x2f\xe2\x64\x8c\x63\xb1\x5e\x89\x42\xd1\x74\x54\xda\xe1\xcf\x71\xeb\x96\x49\x4c\xab\x96\x49\x3a\xb3\xea\x4b\xbb\x44\xef\xb4\xf7\x17\xeb\x2b\x34\xc0\x63\xae\xa3\x9f\xa2\x4f\x4e\x06\xdb\xa1\x57\xc6\xaf\x57\xb5\x75\x6f\xdc\xbc\x6f\x93\xf8\x80\x03\xaf\xdf\x69\xde\x8b\xfb\x82\x13\xaa\x67\xae\xb3\xf9\xc7\xb4\xff\x87\x61\x87\x96\x1f\xd6\x2a\xa2\xc1\x3f\x1f\xe9\x1d\x12\xdf\x15\x69\x19\x47\xf5\x6f\xda\xc3\xd6\xab\x6d\xea\xe8\x35\x99\xe3\x16\xff\x05\xf9\x2b\x9a\x6f\x69\x3f\x95\xab\x7f\xca\x6c\xa2\x0d\x3c\x3d\x1b\x74\x78\x79\xc7\xd0\xb2\xc0\xe7\x25\xad\x0d\x0a\xed\x25\x5f\xfc\x7f\xf6\xfd\x49\xab\xa7\x71\xf2\x6e\xd9\xa2\x59\xbe\x8e\xe6\xf5\x50\x4b\x34\xc3\xe9\x13\x47\x7f\x60\x42\x5a\xef\x72\x99\x9b\xfb\xf4\x3f\xa1\x5e\xd8\x9d\xa1\x89\xdc\x2a\xde\x68\x66\x30\xfe\x93\x19\xbe\x22\xad\xf9\x6a\x3f\x71\x9f\xb1\xef\x68\xc9\x47\xa8\x51\x8c\x0d\xd6\x6d\x14\xfb\xd5\x44\x8d\x62\x99\x6d\xe6\x4b\xd5\x46\xed\x04\xd4\x56\x49\xd8\x2d\xd1\x46\xce\x0a\x2f\x5d\x10\xac\x4c\xb6\x54\xe8\x18\x8f\xec\xff\xc6\x89\x78\x13\xf9\x3e\x82\xa9\x89\xb9\x1b\x60\xd7\x6f\x2f\x4d\x2d\x53\x51\xb9\xff\xe6\xac\x9c\x4d\xf4\xde\xfe\xde\x71\xe6\x76\x4b\xfd\xce\xb7\x9c\x1a\xa5\x9f\xef\xdd\x32\x19\x9d\x67\x9d\x7d\x85\x21\xc4\xcc\xc7\x1d\x5b\x4e\xda\x5f\xfa\x68\x5e\x59\xd2\x15\xd8\x00\xb5\xbe\x70\x47\xe2\x95\x87\x4a\x96\x09\x84\x23\x82\xdf\x6d\xe4\x2f\x65\xa8\x97\x70\x69\xb1\x8d\x1b\x62\xe6\x6e\xfb\xc9\x69\xf8\xe4\x52\x2b\x94\x47\x7a\x3d\xb4\x5d\xbe\x9f\x07\x00\x00\xd8\xbd\xd3\xd1\x56\x64\x73\xec\xfa\xff\x03\x00\x00\xff\xff\x0d\xb0\x06\xb1\x83\x0a\x00\x00")

func AssetsBitcoinPngBytes() ([]byte, error) {
	return bindataRead(
		_AssetsBitcoinPng,
		"../assets/Bitcoin.png",
	)
}

func AssetsBitcoinPng() (*asset, error) {
	bytes, err := AssetsBitcoinPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../assets/Bitcoin.png", size: 2691, mode: os.FileMode(436), modTime: time.Unix(1643293985, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _AssetsEthereumPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\x8f\x77\x34\xd5\x0f\x1f\xc7\xbf\xd7\x9e\xbf\x5f\x42\x08\x21\x64\x66\x16\xed\x28\x92\x91\x74\x23\x23\x3b\x92\x8c\xa4\x9b\x2b\xdb\xb5\xeb\x31\x1a\x2a\x65\x85\x42\xe6\x63\xff\x88\x4b\x08\xd9\x89\x92\x11\x22\xdc\xf1\x5d\x77\x99\xf7\x3e\xa7\xe7\xf7\x39\xe7\x73\xce\xeb\x9c\xf7\xfb\x8f\xd7\x3b\xcd\xce\xd6\x42\x54\x68\xbf\x10\x00\x00\xa2\x96\x17\xcd\xb0\x00\x00\x98\xfc\x79\x01\x6e\x00\x00\xf2\x9a\xe3\x29\x00\x00\xf0\xf8\x99\x5e\x32\x05\x80\xff\x66\x0a\xef\x78\xf1\x02\x00\x80\xb9\x87\xb5\x38\x07\x54\x8f\xc8\xad\x03\x00\x37\x60\x67\x63\x6f\x6e\xe3\x7f\x37\x30\x32\x90\xc3\xe1\x50\x20\x98\x0c\x42\x64\x10\xe2\xfc\xff\xfe\x05\x0a\x04\x13\x89\x44\x32\x08\x91\xa8\x20\x89\x0a\x52\x20\x98\x02\x41\x28\x83\x49\x85\x60\x32\x15\xa4\xc2\x08\x09\x04\xc9\x20\x44\x81\x11\x2a\x8c\xbc\xf1\xe1\x80\x28\xfd\x4f\x07\x84\x21\x14\xa5\x40\x30\xca\x60\x76\xbe\x7f\xbf\x4e\xa1\xc0\x74\x06\x08\x82\x64\x08\x5a\xa5\x52\x61\x94\x0e\xa2\x28\x19\x04\x29\x20\x44\x02\xa1\xa1\xa1\x21\x32\x19\x82\x68\x34\x12\x05\x04\x11\x1a\x05\x42\xc8\x20\x4c\x85\x11\x10\x46\x29\x20\xb4\xb6\x46\x5e\x59\x5d\x45\x19\x4c\x84\xc1\x44\x19\x4c\x32\x08\xc1\x34\x06\x88\xd2\xc9\x20\x04\xa1\x34\x84\xc6\x80\x68\x74\x84\xc1\x84\x69\x74\x0a\x85\x42\x86\x61\x32\x8a\x52\x61\x34\x96\xc3\x81\xe9\x0c\x84\xce\x00\x91\x3f\x62\x54\x18\x21\x43\x10\xf9\x8f\x3c\x4c\x45\xfe\xcf\x30\x8c\x32\x18\x20\x42\x27\x53\x41\xca\x9f\x14\x06\x11\x14\x44\x69\x54\xf8\x8f\xb6\xae\x9e\x1e\x21\x21\x51\x5c\x5c\xfc\xc8\xd1\xa3\x47\x8d\x8c\xfe\x9d\x9f\x98\x94\xac\x6f\x60\x00\xa1\x34\xc6\xc6\xa6\x98\x98\xd8\x61\x1d\x9d\x84\xc4\x24\xe6\xe6\x96\x81\xa1\x21\x4c\xa3\x83\x08\x1a\x15\x1d\xbd\x4e\xa1\xea\xe9\xeb\xd3\x98\x2c\x94\xc1\xdc\xe5\x70\xa8\x30\x82\x32\x58\x3b\x1c\x8e\xa4\xa4\x64\x64\x54\x74\x3c\x21\x61\x63\x67\xc7\xf8\xd8\x31\xfa\xc6\x86\x8a\xaa\xaa\x84\x84\x04\x9b\xc3\x49\x4a\x4e\x89\x8a\x8e\x51\x53\x57\xe7\x70\x38\x5b\xbb\x9c\xd8\xb8\xf8\x94\xb4\xb4\xad\x5d\xb6\x8e\xae\xae\xe1\x91\x23\x5a\xda\xda\xac\xad\xed\x75\x32\x85\xce\xda\x88\x88\x8c\xd2\xd0\xd4\x3c\xa8\xac\x1c\x13\x1b\x17\x1d\x13\xab\xa9\xa5\x65\x64\x6c\x3c\xf5\x7d\xfa\xdb\x8f\x19\x29\x29\xa9\x95\xb5\xf5\xcd\x9d\xdd\x6d\x0e\x27\xed\xe1\xa3\xb8\x78\x42\xdf\xc0\xe7\xd5\x75\xd2\xc6\xf6\x2e\x7d\x63\x53\x44\x54\x84\x87\x87\x07\xa6\xd1\xeb\x1b\x1b\x4f\x9f\x39\x33\x3e\xf1\x55\x49\x49\x49\x41\x41\xa1\xa2\xaa\x5a\x59\x45\xc5\xcd\xdd\xfd\x90\x9a\xda\x79\x33\x33\xf3\x0b\x17\x18\x9b\x5b\x4d\x2d\xff\x64\x66\x3d\xee\xea\xee\x99\x5f\x5c\xfa\x4f\x66\xe6\xdb\xd2\xb2\xab\x0e\x0e\xfc\xfc\xfc\x30\x8d\x11\x18\x14\x9c\xfd\xfc\xc5\xd0\xc8\x68\x47\x67\xd7\x9b\xe2\x12\x6b\x1b\x1b\x59\x59\x59\x7c\xf8\x83\xdc\xfc\x02\x5c\x58\xd8\xaf\xdf\xab\xe1\x0f\x22\xd6\x48\xe4\xc5\xe5\x95\x97\x39\xaf\x7a\xfa\xfa\x67\xe6\xe7\x6b\xeb\xea\xcb\x2b\xab\x26\x26\xa7\x08\x89\x49\x3e\xbe\xbe\x24\x2a\x24\x2d\x23\xf3\x3a\x37\x6f\x63\x7b\x67\x6e\x61\xd1\xc5\xd5\xf5\xc9\xb3\xec\x9b\xb7\x6e\x15\xbe\x29\xea\x1f\x1c\x6a\x6d\xef\xa8\xac\xae\x41\x18\xcc\xe1\xb1\xf1\xa3\x47\x8d\x98\x5b\xdb\xc1\x21\x21\xfe\x01\x01\x9f\x87\x47\x4c\xcf\x9d\x53\x53\x53\x9b\x9e\x99\x5d\x58\xfa\x85\x32\x59\x57\xb0\x57\x3d\xbd\xbd\x2d\x2c\x2d\xdb\x3a\x88\x67\x4d\x4d\x95\x95\x95\xbb\x7a\x7a\x29\x30\xe2\xe8\xec\x0c\xa1\xf4\xb2\xf7\x15\xb5\x75\x75\x25\xef\x4a\x6d\x2e\xd9\x9e\x38\x79\xf2\xf9\xcb\x9c\xe1\xd1\x31\x5b\x5b\x5b\x3a\x93\x85\xd0\x19\xd7\x1c\x9d\xee\xde\xbb\x97\x5f\x58\xe8\x7d\xe3\x46\x4b\x6b\x5b\x72\xc7\xac\x17\x00\x00\xd7\x71\x58\xdb\xab\x00\x86\x6b\x67\x77\x9b\x07\xc3\xe6\xde\xe0\xd9\xda\x61\xef\xb6\xd2\xa7\xb6\xb7\xa4\x33\x34\xa8\x98\x0c\xb2\x7b\x99\xf7\x2a\xb7\xd7\xce\xbe\x14\x49\x7c\x46\xd6\xc6\x77\x5e\x75\xd3\xb3\x55\xe5\x01\x50\x26\x49\x47\x9b\xf8\x71\xb7\x27\x28\xa8\x8f\x9d\x9c\x4c\x26\xc9\x85\xab\xe3\xb9\x26\xbe\xb8\xb1\xc6\xb2\xc6\x47\xaa\xca\x02\x18\x6e\xa3\x23\x53\x6c\x76\xf2\xfb\xf6\x59\x7d\x00\xe0\x17\xb7\x34\x33\xb5\x77\xea\x4e\x77\xc4\xbb\x8a\xfb\x49\xb3\x3b\x85\xf1\x98\x2b\x72\x26\xe5\x0d\xc1\x62\x8e\xd7\xcd\xdb\xc8\xf3\x15\x33\x67\x75\xf4\x07\x54\x5f\x64\x2d\x1f\x74\x4e\x2a\xf1\x57\x14\xf4\x15\x17\x50\xc6\x4a\xb8\x9e\x65\x5d\x9b\x2d\xfe\x6b\x76\xf2\x5a\x79\x53\xf5\x7b\xf9\x30\x68\x0e\x17\x83\xc7\xb4\x6f\xeb\x78\x12\xcf\x2a\x12\xdb\x22\x0b\x22\xbf\x9d\x86\x68\x5c\x8a\x27\x34\xf4\xe1\xcb\xf4\x65\x19\x96\x1a\x6e\x34\x55\x95\xf2\x5f\x99\x39\xf1\x2f\xee\x36\x9f\x16\x7b\xa5\xe5\x66\xb7\x17\xb0\x40\x49\xb7\xc5\xf0\xc3\xc9\x26\x69\xed\x86\x3c\x15\x6a\x43\xef\x58\x51\xde\x66\xe8\x0f\xf8\x58\xf8\x6a\xfe\x9e\xb1\xa8\xa5\xc2\x75\x97\xd7\x5d\x89\x5e\x4e\x8e\x84\xc4\x94\x35\xa1\xd8\x57\x22\x46\xd4\x06\xd3\xb4\x59\x6e\xf9\xdc\x5b\x72\x2a\x19\xf2\x77\x68\xf8\xc7\x9e\xb9\x87\x4f\x69\x4a\xb2\x32\x3b\x55\xdc\x1a\xa6\xc5\x3f\x31\xf9\xc3\x71\x16\x35\x99\x11\x75\xcd\xe9\xcf\x3d\x97\xf3\x8a\x8b\xfc\x1d\x66\xee\xd8\x7c\xf4\x61\xc5\x68\x0a\xf5\x34\x2a\xec\xbe\xb4\xa8\x97\x11\x34\x45\x29\xe7\xbb\x16\x1a\x4f\x0a\xa8\x60\x94\x0e\xc8\xce\xd9\x1e\x0a\x91\x3b\xf5\xbd\xca\xdc\xd3\xf6\xc9\xb3\xb0\x99\xc3\x87\x7b\xdb\x4c\x2e\x92\xda\x4e\xf7\x9f\x91\xed\x3a\xa0\xc0\x38\x21\x7f\x59\x88\x5a\xa5\x70\x40\xcb\xcd\xb1\xa0\x97\xb8\xef\xb2\xd4\xdb\xb7\xa9\x62\x45\x70\x0b\x3f\xaa\xbf\x48\xf6\x12\x7c\x11\x58\x31\x16\xe8\xe2\x8d\x45\x74\xf7\xe6\xa1\x0a\xa9\xdd\xc5\x59\xac\x52\xfa\x03\xf5\x1a\x43\x6f\x5b\x2c\xdb\x70\xfe\xf6\xb7\xfd\x99\x63\xbc\x4f\x28\xb9\xac\x2c\xfb\x35\x6e\x9d\x5e\x85\x91\x47\x52\xd3\x29\xfb\x8c\xca\xda\x1b\x7b\x73\xd2\x07\xd5\x17\xf3\x36\x3e\xa4\xdc\x35\x3d\x74\xc3\x04\xdb\xb1\xb6\x5b\x9b\x28\xf5\x21\x9f\xd7\x8a\xf6\x49\x54\x49\x32\xb4\xe2\x1e\xdb\x73\x4f\x77\x4d\xec\x4f\xf3\xd3\xe6\xcc\x67\xbe\x6d\xe7\x6b\xdc\x0b\x87\x8e\xc4\x44\xef\x11\x25\x50\xfe\x29\x83\xe6\xde\x51\xb8\x9f\x78\xac\x7d\x33\xdf\x61\xff\xee\x15\x0d\xdd\xd7\xe9\xda\x4a\x72\xe3\x43\xde\x58\x13\x83\xbb\x4d\x70\xd7\xb9\x4a\x2e\x13\xf7\xeb\x86\x2d\x4c\x34\xcb\xdd\xfe\xcb\xe4\xa6\x88\x5c\xb0\xa9\x5b\x6a\x4d\xec\xc1\xd6\xa3\xa2\x58\x38\xaa\xc9\x2c\x17\x20\x93\x34\xdb\x4f\x89\x28\x1a\x3a\x1e\xd0\x53\x7c\xa0\xa2\x3a\x9d\x99\x98\xd0\x89\xc4\xb7\x1f\x80\xac\x1e\x8e\x76\x90\x4e\x02\xfb\x7f\xfa\x85\xd1\x8d\xba\xe8\xa4\xd1\x49\xac\x10\x9f\x59\x48\x1c\xea\x51\x6e\x37\x38\x03\x78\xee\xd9\x26\x0d\x3e\xb6\x88\xf8\xce\x57\xab\xb3\xe0\xb9\x10\x2d\xa2\x34\xe2\xf9\xb4\x4f\x83\x28\xec\x39\x62\xcf\x57\x75\x71\x4f\x8f\xa1\xb2\x78\x97\xbd\xd3\xb0\x65\x93\x80\x8c\x80\x3c\xae\xec\x01\x60\xf0\x6b\xab\x7d\x49\xfb\xfb\xbb\xda\x61\x3d\x5c\xaa\x82\x78\x73\x67\x4a\x20\x57\x38\x60\x30\xfc\x29\xfb\x96\xea\x89\xcd\x14\x91\x24\x4b\xe9\x37\x92\xd3\x4f\x99\x68\x10\xaf\xba\xc9\x8a\x54\xdf\xa5\xfe\x1b\x2f\xc3\x78\x6c\xbc\xb1\xf9\x66\xb9\x5d\xd9\xae\x7c\x9b\x40\xdd\x93\x4e\xef\x06\x05\xd5\x26\x15\xa7\xc4\x17\x22\x2f\xd5\x31\x05\xfd\x33\x19\xf3\xe3\x7c\xc0\x68\x0c\x51\x4f\xf4\xfb\xf3\xe3\xfa\x8a\xc7\xa7\x12\xa7\x4c\x7a\xe5\x90\x64\xee\x90\xcf\xc2\xc2\xb7\x67\x86\xd2\xc1\xd1\xc3\x90\x00\xd9\xd3\x4a\x81\x0f\xdb\xec\x59\x24\x92\xbf\xa4\xb9\x80\x23\x09\xaf\x44\x1a\xd7\xec\x7e\xaf\xf4\x37\x77\x22\x8a\x0b\x72\x4d\xbc\x23\x74\x59\x8b\x01\xd9\xeb\xe9\xbd\xa2\x11\xb7\xaa\x75\xb4\xc3\x23\x9f\x1e\x32\x74\xc4\xcf\xbb\x55\xea\xd7\x0f\xdb\xd1\x26\x83\x5c\xe0\x2b\x5e\xd6\x7e\x09\x3c\x47\x78\x9f\xad\xbc\x3f\x4f\xc5\x45\x08\x69\xe9\x54\x3f\xf1\xbf\x9d\x7e\xfc\x44\x5e\x75\xf0\x25\x8d\x44\x75\x23\xbe\x7d\x02\xd6\x69\x32\x7b\x4d\xf0\xc2\xf1\x3e\x70\x60\xf6\x55\x10\x91\x9b\x5b\xd1\x34\xd2\xcc\xb6\xba\x94\x11\xcc\x8e\xca\xd3\xc1\x07\x89\xca\x5a\x0e\x1b\x9f\xcb\x68\x1a\xe1\x96\x29\x3a\x33\xef\xd8\xe8\xd4\xce\xba\x94\x54\x69\x56\x6e\xaf\xfe\x11\x80\xd7\x7e\x2b\xa0\x03\x56\x51\x6d\x3c\xed\x26\x63\x32\x8a\x53\x19\x22\x04\x6e\xd1\x50\xe9\x56\x1b\x4c\x95\xe4\x29\x60\x2a\x9d\x16\x5c\xbc\xfd\xf1\x60\xcc\x1e\x26\x2e\x2e\xc6\xac\x54\xd8\xa8\x73\x2c\xdb\xeb\x1e\x06\x10\xbe\x22\x65\xe7\xda\x8c\xd2\x2e\x02\x5c\xd4\xae\xbd\xcd\x1f\x27\x6f\x66\x00\x43\xa3\x03\xd3\xed\x3b\xbb\x06\xa6\xc6\xf1\x38\x2f\x81\xf7\x90\x2e\xd1\xf1\x6a\x94\x54\xfc\xdf\x59\xd7\x9b\xa3\x58\x50\x69\x35\x22\x5f\xdb\x47\x10\x71\x59\x4b\x98\xbb\xb3\x8b\xe0\x0b\xc6\x55\xf9\xca\x26\xaa\xc6\x28\x0d\xfb\x6f\x61\xb4\x7c\xe3\x1e\x5c\xc3\xc4\xad\xf3\xf4\x4a\x67\x78\x6c\xf5\x96\x56\x4a\x72\x2d\x2f\x57\x1f\x2b\xf1\x08\xbd\xe1\x3c\x75\xce\x5d\x5e\xc0\x1a\xdf\xf0\x78\xed\x9d\x02\x41\xa9\x55\x3b\xbc\xa1\xe5\xd9\xa3\x62\xc9\x38\x51\x76\x41\xbd\xf5\x79\xdf\xf4\x58\xde\x07\x22\xa3\x46\x5f\x02\xbd\xb4\xae\xb2\x3f\x91\x0c\x87\x87\x8e\x52\x99\x8f\xb2\x6f\x0a\x74\x61\x3a\x17\x1c\xa3\xe8\x77\x72\x26\xc7\x70\xc1\xbf\xfe\x36\x97\x20\x66\xf3\x27\xd8\x27\x5e\x56\xa7\xe6\x14\x9f\x6c\x7a\xda\xf0\x93\x3e\x71\xd3\x23\xed\x11\x6f\xff\x78\x15\x9c\x57\xd2\xb8\x0f\x85\x32\x43\x47\x7c\xfd\x04\x3b\x56\xed\xf5\xb4\xdd\xe8\xc7\x0b\x70\xe1\xb1\x6e\xc3\x63\xbe\x6a\x42\x67\xcb\x70\xd9\x5f\x57\xc0\x91\xd3\x51\x6d\xaa\xcb\xc4\xba\x41\xf9\x8d\xa7\x0d\x15\x1b\x05\xf7\x02\x0d\x52\x73\x12\xa2\x46\x35\x7f\x37\xb8\x78\x38\x34\xd6\x6b\xf4\x3c\x73\x23\x3a\xb8\xb4\x46\xb1\xaa\x07\x0f\x2b\x75\x4f\xbf\x3e\x53\xf0\x6e\xd9\x98\xd8\x89\xe9\x76\xed\x27\x5b\xaa\x6b\xd4\x9f\x9f\x2c\xc6\x87\x37\x1d\xf0\xab\x8c\x0d\xdc\xfd\x87\x49\x32\x18\xac\xfb\x2d\xff\x90\x71\xd3\x7c\xa1\x30\xe4\xd8\x45\x54\xf1\xad\xb9\xa2\x83\xc3\xad\xd2\xf3\x55\xf8\x6b\x81\x01\x47\xc6\x7d\x87\x87\x7d\x6b\xf4\xee\x04\xfd\xe5\xb3\xdf\x7a\xea\xa7\xd9\xcb\x3c\x65\x07\xb9\x2f\x11\x6a\xcf\xb0\xcf\x29\x71\x25\x95\x3a\xf0\x01\xc1\x64\xe9\x0a\xe6\x57\x3e\x7c\xfa\x78\x65\xf9\xec\xec\x7f\xc6\x04\x4b\xee\x5f\xbb\x6e\x23\x26\xba\xef\x73\x6b\x6b\x9c\xaa\x7d\x7b\xb3\x97\x1b\x6d\xfd\x21\xb6\x79\x7c\x08\x1d\x8f\x97\x75\x16\xd0\xb8\x4c\xf8\x7d\xd5\x74\xa9\xdf\x4f\x40\xac\xe6\xce\x4a\xb8\x9b\x2e\x5e\x22\xd8\xa5\xfb\x6e\xd1\xdd\xe1\x8c\xa2\x50\x9f\xd6\x79\xe3\x3d\xb6\x17\xc2\x5a\x7e\x44\xa4\x5e\x78\xb4\x4e\x78\xf1\x5a\x24\xff\x06\x65\xc6\xd9\x6a\xbe\x36\xc5\x5e\x89\x96\x27\x5b\x1d\xf2\x81\x05\xf5\xa6\xe6\xb4\x97\x8f\x0d\x60\xdb\x5a\x5b\x9f\xcb\x56\xbf\xd5\xc1\xb5\x84\x45\x7f\xf8\xd6\x59\x6b\xff\x94\x60\xd1\x43\xcb\xab\x3b\x12\x5a\xff\xe6\xbe\x81\xff\x2c\xd9\xe5\xf3\x8c\x98\xb4\xdd\x92\xd9\xd9\x2b\xba\x53\xe1\x53\xf7\x29\x38\x01\x6e\x69\xee\xbe\xb1\x81\xfd\x1e\x9c\x86\x86\x4b\xeb\xe6\x09\x03\x73\x3d\x07\x07\xf3\xef\x0b\xde\xce\xad\xbd\x80\xee\xba\xea\xd8\x57\x54\xf6\x5b\x25\x93\x2a\x95\x73\x2c\x92\x08\x23\xb5\xfc\xbf\x97\x96\x2e\x04\x5d\xe3\x37\xc8\x0c\x4a\xee\xda\x14\x4b\x04\x5f\x54\x9d\xd0\x0f\xd9\x5b\xae\x28\xf1\xca\x32\xea\xcd\x74\xc6\x57\x07\x7d\xe8\x64\x21\x76\xb2\xe5\x27\xad\xc2\xae\x7c\x3b\xd7\x8e\xf4\xd1\xc5\xaa\xd7\xfb\x75\xf6\x2b\xdb\x9a\x8a\xd0\xc9\xb6\xd4\xaf\xd4\x5c\x56\x60\x81\xfa\x29\xc2\x2c\x64\x56\xd8\x38\xeb\x61\x47\xfd\xd1\xe2\xe0\x84\xe1\x7a\xa4\xc8\x72\x56\xc0\x7e\x79\x68\x26\x9c\xe2\xbb\xb2\xf8\xf9\x9e\x79\x51\x92\xb2\x26\xfa\x99\xcb\x31\x3f\xc2\x67\x45\x29\xe7\x23\xd4\x21\x99\x65\xd7\xa6\xf5\x77\x33\x2e\x3d\x46\xc4\xb9\x55\x3d\x20\xa4\x31\x69\x29\x07\x1a\x88\xc1\xd9\x6c\x93\xfd\xe4\xda\x98\xa6\x02\xec\x1f\x7d\x7a\x5f\xe7\x1a\x33\x9b\xee\x37\x30\x93\xa4\x46\xd9\xbf\x0e\xf9\x04\x28\xeb\xe6\xd7\x4c\x7f\xca\xb4\xa6\x84\x20\x22\x35\x1d\x69\x09\xef\xcc\xdd\xa3\x4b\xff\x7e\x15\x5e\x33\xd9\x7c\x7f\x33\x43\xf6\x4c\x24\x68\xb3\x88\xca\x9f\x0c\xf8\x1a\xb2\xce\x71\xee\x70\x16\x68\xb2\x40\x96\x65\xfc\x05\x69\x06\x5f\x2c\x0f\x9e\xbe\xe1\x7e\xce\x85\xa4\x37\x85\xbc\x94\x40\x26\x38\x98\x5a\x0e\xf7\xf2\x4a\xde\xd2\x28\x00\x00\x80\xa5\xb9\xad\x59\xf5\x39\x4f\xc2\xff\x02\x00\x00\xff\xff\x2a\x62\xe5\xc0\xdc\x0a\x00\x00")

func AssetsEthereumPngBytes() ([]byte, error) {
	return bindataRead(
		_AssetsEthereumPng,
		"../assets/ethereum.png",
	)
}

func AssetsEthereumPng() (*asset, error) {
	bytes, err := AssetsEthereumPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../assets/ethereum.png", size: 2780, mode: os.FileMode(436), modTime: time.Unix(1643290341, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _AssetsLitecoinPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\x90\x79\x38\x94\x6b\x1f\xc7\x1f\x59\x2a\x44\x54\x3a\x49\xbd\x27\x25\x2a\x1c\x79\xe6\x99\x91\x64\x39\x11\x93\xa5\xb1\x25\x3c\x33\x8f\xc1\xb1\xef\x2d\x38\x92\x42\x93\x8e\x2c\x65\x19\xcc\x3e\x63\x0f\x69\xb5\x26\x9a\x8a\x28\x5b\x42\x84\x2c\x99\x22\x4e\xb2\xc5\x6c\xef\x35\x9d\x3f\xee\xeb\xba\xaf\xdf\x7d\x5f\x9f\xcf\xf7\xfb\xbb\x89\x73\xb0\xde\x22\xaf\x2e\x0f\x00\xc0\x16\xac\x8d\xa5\x13\x00\x00\xe6\x92\xb3\x49\x1a\x00\x00\x5a\xcd\xb5\x6f\x00\x00\xc8\xf8\x5b\xd8\x5b\x00\xc0\xfd\x0c\x05\x01\x51\x16\x00\x00\xa9\xf3\x4e\xd6\x7f\x02\x55\x9d\x1a\x5f\x01\x40\x1a\xc0\xd9\xb9\x58\xd9\x05\x46\x02\x09\x09\x20\x4c\x05\x61\x2a\x00\x88\x51\x30\xcd\xcc\x2c\xc1\xd5\xb5\x1c\xc4\x53\x51\x78\x2a\x00\x00\x28\x98\x06\xe2\x69\x20\x9e\x0a\xe2\x29\x09\x09\x09\x20\x9e\x8e\xc2\xd3\x40\xbf\x7c\x14\x9e\x06\xf9\xb0\x41\x98\x02\xb9\xe6\x80\x9e\x14\x08\x61\x81\x30\x15\xed\xc5\x86\xbc\xd8\x10\xc2\x42\x11\x18\xa0\x5f\xbe\x91\x6b\x39\x68\x97\x85\xf1\x29\x81\x10\x26\x0a\x4f\x43\xe1\xe9\xfa\x30\x1b\xe5\x49\x93\xe8\xf0\x34\xd0\x9b\x25\x41\xc1\x34\x23\xdf\x32\x09\xf6\xd7\x10\x42\x58\xd0\xaf\x30\x20\x9e\x86\x42\x98\x27\x5c\x5d\x21\x84\x09\xc2\x54\x14\xc2\x82\x10\x66\x50\x81\x18\x22\x30\x40\x98\x2a\x01\x12\x98\x28\x84\xf1\x2b\x0f\x55\xa2\x93\xa0\xa8\x92\x57\x09\x8a\xa6\x0f\xa7\x42\x5e\x2c\x94\x5b\x0e\x8a\xc0\x84\xf0\x0c\x43\xf7\x02\x34\x81\x06\x12\xa8\x28\xc9\x1f\x16\x0a\x4f\x47\x11\x7e\x15\x81\xe9\x28\x02\x13\x24\xd0\x41\xbf\x72\x90\x50\x24\x99\xe3\xe9\x47\x09\x94\xff\x02\x83\x9e\x14\x50\x52\x9f\x22\x89\x24\xe9\xc8\x3c\x7a\x2e\x4b\xe2\xf2\xcc\x3f\x06\xb3\xff\x5b\xd1\x7f\xde\x5f\xed\x24\x76\x08\xcf\x80\x24\x77\x9a\x58\x2c\xc6\x20\x74\x8c\x77\x91\x40\x28\xd2\x76\x26\xc7\x66\x37\x41\x78\x9a\x48\x24\x42\x23\x74\xa1\x50\x04\x11\x68\x22\x91\x18\xe3\xc5\x30\xf2\x2d\x33\x84\xa9\xc6\xde\x4c\x8c\x77\xb1\x40\x20\x9c\x9e\x5d\x94\xac\x8b\x40\x3f\xe8\x94\x0b\x21\x4c\x08\x61\xe9\xbb\x17\x1c\x39\x9b\x6f\xe4\x5b\x2a\x14\x89\xb4\x1c\x73\xf5\xce\x15\xa0\x08\x74\x34\x81\x6e\xe8\x49\xc5\xf8\x94\xa2\xbd\x38\x22\xb1\x58\xc7\x85\x8c\x26\x16\xa2\xbd\xd8\x87\x5c\xf3\x8e\x11\x19\x20\x4c\xeb\x1d\xfe\x8a\xf1\x29\x36\xf0\xa0\xa0\x89\x9c\x77\x1f\x67\x4e\xf8\xb2\x0e\xb9\xe4\x05\xdc\x78\xc2\x17\x08\xff\x70\xa7\x1c\xf5\xa4\xae\xad\x0b\x74\x9c\xc9\x28\x3c\xfd\x6a\xc1\x73\xc8\x8b\x35\xb7\xb0\x1a\x93\xd5\xa4\xeb\x96\x8f\x41\x18\x24\xfa\x0b\x3d\xb7\xfc\xe5\xd5\xf5\xc9\x2f\x0b\x9e\x97\xab\xb5\x9d\xc8\x2b\x3f\xf9\xdf\x17\x7f\x36\xbf\x1d\x27\x26\x3e\xbc\xdf\x32\x54\xde\xd0\xbf\xb8\xb2\x76\xe4\x6c\xde\x0d\xc6\xcb\xcf\xb3\x8b\x02\xa1\x10\xe3\x5d\xc8\xfb\xb6\xb4\xce\x17\xfc\x9d\xf3\xac\xed\xdd\xe7\x33\x51\x77\xc9\x15\x6f\xdd\x62\xab\x2a\x9b\x06\xb1\x61\x25\x76\x11\x65\x53\x33\x3f\x92\xa8\xdc\x53\xc1\x45\x39\xe5\x6f\x46\xa6\xfe\xb5\xf0\xe7\x1c\xf5\xa0\x18\x79\x31\xce\x67\x36\xb0\x1f\xf7\x3e\xe2\x0e\xb7\xf6\x4e\x51\xab\xbb\x5c\xa3\x2b\x39\x4f\x7a\xfb\x46\x66\xbe\x7c\x5b\x1c\x1a\x9f\xfb\x87\xfd\xea\x56\x61\xeb\x38\xef\xfb\xec\xbf\xcb\x96\x81\x85\xc3\x13\xf3\x2f\x7b\x26\xc3\x6f\xd5\x15\xd6\xbc\xeb\x1c\xe0\x55\x3c\x1d\x88\x48\xab\xb7\x08\xe0\xd4\xb5\x8d\x58\x05\x15\x9d\x8d\xa9\x22\x24\xdc\x4f\xa6\x71\x7f\xb7\xcf\x6a\x78\x3d\x1a\x97\xf3\xec\x76\x49\x7b\x7a\xd1\xeb\xe3\x3e\xcc\x03\x8e\x39\xc4\xab\x0f\x3a\xfa\xa7\xa3\xd2\x1b\x9e\xb6\x8f\x75\x0e\xf2\x76\x63\x6f\x5b\x87\x14\xdb\x47\x96\x75\x45\xd6\xc6\x00\x00\x10\x75\xc1\xc9\xc1\x19\xd8\xc0\x17\x4a\x09\x65\xa4\x05\x22\xa9\x75\x3e\x5f\xb0\xe1\xb4\x48\xf6\x88\x92\x60\x67\xb8\x6d\xa4\x68\x75\x4d\x20\xbf\x59\x5e\x34\x08\x7b\x2b\xc4\xce\x5b\x6f\x7d\x86\x88\xf4\x44\x93\xcf\x0c\xb8\x63\x81\xd2\x1f\x06\x7a\x87\xa5\xc2\x4a\x5f\x17\x16\x47\xd9\xcc\xa4\x9e\x46\xfc\x02\x14\x06\xf6\xb4\xb5\xef\x3d\x1c\x5d\x5c\xd2\x38\xb9\x66\x75\x33\x5a\x59\xf5\xd9\xba\xfe\xd3\x4f\xa1\x83\x88\xe7\xcb\x77\xea\x1d\xbb\xd4\xca\x8a\x4b\x78\xd3\x37\x67\x53\x53\xe7\xaf\xfd\xde\x02\x00\xb2\xb5\x58\x4b\x0b\x97\x73\xdc\x34\xb7\x68\x78\x9b\xff\x6f\x22\x1b\xa5\x07\x9b\x9c\x55\x34\x52\x14\xad\x64\x57\x4f\x59\xa9\x78\x96\xc5\xf5\xf7\x3c\x5d\xad\xb5\x65\x8e\xb7\x35\x0f\x1e\x48\x06\xf6\x85\x64\x3e\xb5\xb0\x92\x89\xec\x7b\xdf\xef\x38\x5b\xfa\xa4\x4a\xa9\x61\x75\x34\x2e\x3e\x9a\xac\x50\xd7\x11\x7d\x8d\xf7\xbf\x2b\x07\x62\xc7\xf8\xa3\xbc\xb8\xf9\xf8\xb1\x8d\xc0\x99\x36\x2d\xd5\x99\xae\xcc\xee\xb4\xa3\x2a\x77\xdd\xac\x5f\x84\xc0\x50\xf2\x3a\xa2\x79\x3b\xa9\xc5\xdc\xc1\x2a\x43\x79\xe8\x51\x98\x90\x44\x35\x86\x2c\xf1\x9d\xa4\xe5\x59\xb4\x6c\x28\x2b\x4a\x51\x53\x56\x71\x55\x01\xc9\x63\xa3\x7f\xab\xbc\x6e\x88\x1c\x4a\x0e\x8f\x60\xbd\x0f\xc0\x95\xe8\x2a\xde\xe9\x46\xe2\xae\x2b\xcb\x69\xfa\xbf\x39\x84\xf9\x9c\x92\xe3\xe4\x6e\x30\x6f\x59\x6d\x57\xbe\x9d\x97\xd7\xce\x7b\x93\xb2\x21\x2f\xa3\x26\xc6\x0c\xc8\xed\x3c\xac\xaf\x62\xfb\x57\xed\x46\x5f\xa7\x48\x6c\xe4\x92\xd3\xba\xf5\x86\x66\xbf\xc4\x32\x67\xbb\x1d\x27\xe3\x23\xa5\x7e\x2a\x55\x1a\x8c\xc4\x6e\xbc\x68\xb7\xb4\x20\x65\xa3\xce\x52\xda\xaa\x58\xe2\x6c\x93\xb6\xff\xce\x17\x06\x7b\x30\xf5\x74\xe7\x6c\xe6\x84\xb4\xb9\x47\x5e\x8d\x81\xfc\xcb\x22\x99\x64\x3d\x03\x9d\xb5\xc7\x59\x7e\x7a\x5b\x7a\xd3\x14\x8a\x88\xb1\x79\xfe\x72\x9b\xcf\xe5\x9f\xe5\x7c\x32\x73\x08\x7e\x58\x33\xe5\x64\x94\x3b\xb7\x52\xf6\x3c\xfb\xd2\x6e\xb3\x24\xe5\x4d\xa9\xe8\x7f\x7e\x2e\x14\x56\x36\xd4\xd7\x37\x35\x35\xd5\x8b\x14\x3a\x79\x9a\x3e\x4b\xfb\xe1\xbd\xd5\xcd\x6d\x3b\xa0\xbd\x0b\x43\x6c\x12\x47\x26\x69\xc4\xed\x52\xd0\x47\x30\x2f\x37\x43\xe9\x85\xd1\xaa\xd9\xd8\x27\x95\x8c\xac\x34\x7b\xcb\xf3\x6a\x2f\x5e\x2f\xed\x57\x3b\xb1\xaa\x48\x82\xe6\xff\xce\xbc\x4a\x94\xbb\xae\x3d\xbe\x9e\xd8\x95\xb2\xa7\x6f\xf1\x61\xca\xdd\xb9\xc3\xf3\xd1\xf2\xdc\x11\x21\x4a\x26\xb1\xfb\x36\x67\xeb\x10\x37\xd5\xf4\xe0\x0f\x62\x63\xba\xd1\xba\x63\x65\x6e\x49\x87\xe3\x7e\xb7\x6b\xb1\x9d\x32\x20\xcf\x1e\xd4\x9c\x1f\xeb\xb9\x88\xbe\x9b\x6d\x6c\x6e\xbc\x80\xd5\x7a\x2d\x46\x37\x1b\xd9\xeb\x13\x15\x52\x0e\xad\xb4\x3a\xee\x68\xec\x72\xa9\x4e\x8a\xc9\xb0\x96\x92\x2d\xfa\x6a\x4a\x92\x53\x4c\xa5\x9a\x6a\x14\x21\xbd\x97\xd4\x4d\xa9\xd3\x9f\x2a\xe3\x95\xce\xb8\x2a\x33\x4c\xd5\xb9\x5f\x76\xd7\x42\xe1\x3a\x9a\x84\xe1\x1f\x1e\xca\xba\x8c\xab\xa8\x46\xcf\x64\xd5\x8b\xe6\xa3\x7a\x4b\x2a\xd5\x99\x1f\xb4\x4a\xa5\x4b\x0d\x57\xcc\xd7\xbf\xca\x8f\xea\x77\x73\x08\x15\xf5\x1f\xf6\xe5\x69\x4c\x9f\xab\x50\xe3\xec\x65\xd6\x34\x36\xe3\x32\xb1\xa8\xd8\x28\x1b\x0b\x81\x93\x89\x22\xf9\x6a\x6b\x2b\x53\xa7\xcf\xd7\x6d\x24\xf9\x7e\x9f\xd6\x7c\x98\x4e\x2f\x5b\xf5\x67\x86\xfb\xc9\x40\x62\x5b\xad\x46\xb6\x5c\x5b\x3b\xf7\x82\xc9\x82\xe6\x02\x84\x58\x55\xa8\x45\xbd\x0a\x25\x69\xc8\x6c\x76\xc1\x7d\x3f\xe8\x6c\x61\xdc\x13\xd8\xc3\x89\x4b\x3a\xb7\x47\x79\x9b\x2e\x16\x37\xd1\x66\x1b\x73\x63\xba\x3a\x3b\xe0\x38\x23\xf8\x94\xf6\x56\x2d\x13\xd9\x6d\x76\x38\x4e\x7d\x26\xf7\x4a\x71\x7e\x0f\xe0\x18\x8a\x51\xb8\xa8\x3d\x29\x63\xd3\x92\x75\x76\xd7\x0f\xa5\x7a\xf7\x5d\x4d\xbd\x35\x23\xb7\xb6\x6f\x30\x0d\x1c\x98\x32\xca\x2c\xd5\x51\x3c\x15\x39\x18\x93\x72\x9b\x95\xfe\x53\x58\xb3\x12\xff\x8c\x9d\xd1\xcd\xda\x24\xbb\xd0\x26\x1d\x30\x16\x23\x64\x3a\x93\xc8\xc1\x5d\xb0\xc9\x08\xed\xfe\xce\x58\xe5\xfe\xdf\x0d\xee\x6b\x1c\xe2\x3e\x38\xb3\x6b\xd7\x4c\x8a\x42\x82\xb9\xe1\xca\x70\xa6\xd3\x8f\x4a\xca\x26\xb5\x7d\xc7\x2a\x0e\x1f\x7f\x8e\x03\xc2\xa5\x0e\x7e\x0d\xdd\xe1\xe5\x78\x42\xf1\xd6\xb2\x9d\xf9\xcc\x1e\x7e\x7a\xbc\xf9\xaa\xaa\x7a\x3e\xf9\x0f\xb6\x32\xb0\xed\x63\xfa\x69\xaf\x87\x83\x5d\x42\xea\x3d\x0e\xe3\xde\xf1\x88\x49\xf2\xc3\x5a\xfb\x54\x85\x83\xe6\xd9\xd5\xe2\xde\xb1\xd0\xce\x1b\xfe\x2a\x63\x62\x5c\xff\xf5\xac\xda\x88\xcd\x97\x55\x5d\xff\x72\xcb\xb7\x21\x5d\x7d\x72\xfe\xb1\xab\xd2\xb8\xe3\x66\xec\x14\xc2\x69\xfa\x64\x6c\xdf\xb4\x08\x30\x86\xf7\x10\xe4\x8a\xde\x14\xc8\xab\x5a\x69\xec\xb3\x90\xd2\xf8\x52\xd7\x13\x9b\x36\xa7\x45\x96\x5f\x22\xa8\xe2\xfe\x1c\x12\xc8\xf0\x36\x4d\xe7\x8a\x8d\xc5\x8f\xc5\x8b\x64\x67\x92\xb1\x94\x54\x2d\x36\x5f\x6c\x18\xe8\xc9\x0f\x4a\x85\x47\xaa\x07\xbe\x5d\xec\x73\xfe\x63\xdb\x65\x5a\xdd\x16\x2c\xa7\xbe\x57\xa4\xed\xd9\xa2\x92\x5a\x23\x9f\xbd\xd9\x36\x73\x8b\x55\x78\x6c\xce\x9e\x40\x24\xb7\xac\xf1\xc0\xd6\xbb\x3b\x9d\x8f\xb0\xac\x6d\x87\xb1\x2b\xc6\xd6\x43\x2d\xf7\x3c\xde\xaa\x2b\xc5\xe8\x85\x9a\x40\xe4\xe9\xc5\x82\xc3\x81\xfc\xf4\xba\x57\x17\x26\x70\x24\x34\x45\x41\xaf\xa4\xcc\x54\x8e\x2b\x34\x09\xea\xde\x09\x7b\x37\x8c\x2f\x54\x91\xe6\x34\x5c\x3e\x47\x0d\xad\x7b\x07\x3d\xf0\xda\xb7\x1a\xb1\x31\xf5\x03\x61\x6e\xbb\xe6\xc2\xf8\x4d\x4a\x57\xc3\xe4\xf3\x9b\x93\x69\x55\xc1\xcb\x7a\xac\x8d\xa9\xad\x01\x20\x8f\x1e\xbe\xff\xf4\x66\x29\x78\x76\x67\x6e\xf0\x23\x62\x64\x74\x9d\xf2\xc0\xd3\x5c\x07\x9e\xdd\x75\x94\x5f\x54\x26\x83\x1e\x3b\x08\x15\xc2\xdc\x1d\x3f\x22\xee\xf1\x63\xb4\x3b\xae\x24\x73\x65\x6e\xda\x4d\x7c\x2e\x0f\x49\x0b\x6e\x6c\xad\xf6\x74\x60\x76\x09\x85\x41\xc9\x94\xf7\x61\x2b\xf8\x09\x9f\x1b\xf3\x7f\x93\xda\x8b\x9e\x07\x46\x40\xa7\x2a\x8c\xca\x5f\x15\x57\x4c\x5f\x38\xf0\xc0\x41\xd5\x6c\xb9\x0c\x3e\xe2\xa0\xa3\x68\xa9\xeb\xde\x87\x53\xb3\xec\xc2\x3c\xf6\x7b\x71\x47\x76\xea\xca\x27\xc7\x4b\x6f\xdf\x3d\x42\xd3\x86\x16\x35\x54\xe3\x1c\xe2\x3b\x94\x8e\xdd\x26\x9a\xfa\xcf\xb5\x4e\x5b\x99\xbe\xd0\x2f\x28\x68\x57\xb4\xa8\x3e\xe9\xe0\xeb\xd1\x82\x47\x8d\x8f\xae\x7f\x87\x2f\xbd\xd7\x72\x4d\x8c\xa8\x3d\x3b\x62\x20\x6a\x57\x26\xcf\xb9\x43\xf2\x97\x71\xb8\xbe\xc2\x4a\xcb\xca\x59\x65\x7d\x83\x14\x5f\xe9\x1b\x1d\x96\x24\x59\x62\x00\xb7\x3a\x24\x94\x30\xee\xa5\x68\x1d\xb5\x7d\xa8\xb1\x85\xad\xb6\x25\x20\x76\xc3\xad\xfd\x9a\x32\xba\xca\x4c\x8d\x11\xa0\xf8\x08\xf3\x39\xea\xaf\x72\xe6\x57\xef\x10\xe4\xdb\xc7\x09\x79\xe2\x30\x06\xd3\x4f\x99\xe3\xb4\x67\x26\xd5\x16\x6e\xec\xeb\x8e\x59\xb7\xd5\x5e\xde\xdd\x85\x75\x4a\x08\x30\x79\x7f\xde\x4f\x77\x77\x73\x8c\x61\x3f\xe6\xc0\xe5\xb9\x77\xdd\xb3\xf3\x33\x11\x1e\x88\xfb\xb8\xfe\x97\x64\xe7\x3b\xdc\x93\x72\x4f\x9a\x2a\xd6\x46\x55\x45\x6f\xfa\x5f\x0f\x57\x7c\xce\x5e\xd3\xaf\x09\xa3\x34\x2e\x7f\xf8\xde\x52\x10\x8e\xe1\x33\xf2\xd4\x77\x1f\xbd\x14\x62\x19\x12\xc4\xde\xc9\x17\x9b\x55\x60\x12\xc3\x0c\xed\x4d\x23\x00\x00\x00\xb0\x56\x0e\x96\x55\x7f\x7a\x25\xfe\x3f\x00\x00\xff\xff\x03\x90\xdf\x03\x98\x09\x00\x00")

func AssetsLitecoinPngBytes() ([]byte, error) {
	return bindataRead(
		_AssetsLitecoinPng,
		"../assets/litecoin.png",
	)
}

func AssetsLitecoinPng() (*asset, error) {
	bytes, err := AssetsLitecoinPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../assets/litecoin.png", size: 2456, mode: os.FileMode(436), modTime: time.Unix(1643290387, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"../assets/Bitcoin.png": AssetsBitcoinPng,
	"../assets/ethereum.png": AssetsEthereumPng,
	"../assets/litecoin.png": AssetsLitecoinPng,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"..": &bintree{nil, map[string]*bintree{
		"assets": &bintree{nil, map[string]*bintree{
			"Bitcoin.png": &bintree{AssetsBitcoinPng, map[string]*bintree{}},
			"ethereum.png": &bintree{AssetsEthereumPng, map[string]*bintree{}},
			"litecoin.png": &bintree{AssetsLitecoinPng, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

