// Code generated by go-bindata.
// sources:
// dist/images/favicon.ico
// dist/scripts/application.js
// dist/scripts/vendor.js
// dist/styles/application.css
// dist/styles/vendor.css
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

var _imagesFaviconIco = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5b\x07\x58\x54\x57\xf6\x7f\x33\x80\xd8\x22\x28\x2a\xbd\xc9\x34\x9c\x42\x1f\x3a\x43\x13\x04\x94\x8e\x02\x4a\xb1\x46\xb0\x83\x82\x34\x19\x7a\x71\xe8\x0c\x30\x54\x51\x10\x24\x36\xec\x11\xc4\x12\x23\x6a\xb2\x1b\x4d\xa2\xee\x67\x88\xc1\x8d\xfb\x4f\xfc\xef\xc6\x35\x26\xf9\x7f\xae\x22\xe7\x7f\xde\x1b\x20\x98\x88\x61\x37\x26\x7e\x9b\x8f\xf3\x7d\x3f\xce\x7b\x77\xce\xbb\xf7\x77\xcf\xad\xef\x9e\x07\x41\xd0\x08\x05\x42\x55\x95\xd4\x06\xc4\x3a\x45\x82\x10\x12\x04\x61\x60\x20\xbf\x6f\xc7\xf4\x46\x4c\x33\x35\x1d\xba\x67\x12\x84\xc7\x1c\x82\xe0\xa0\x8d\x2a\x69\x47\xc8\xd3\x29\x51\x24\xfe\xd0\x32\x99\x4e\xd0\x76\xc4\x2d\x33\xa8\x2e\x4f\xf1\x97\x15\x46\xac\xcc\xdf\x68\xb0\x2c\x75\x2d\xc3\x69\x45\x88\xe5\x5c\x42\x25\x83\xf6\xaa\x67\xd9\x4c\xd6\xa4\x8a\xa2\xc4\x98\x0b\x67\x5b\x3f\xbb\x71\xb9\xf6\xd9\xb9\x3a\x2e\x1c\x2f\x50\x1c\x3c\x94\xa3\xfc\xfd\xde\x9d\x6f\xdd\x28\xda\xaa\x95\xb4\x7a\xa9\xf9\xdc\xb1\x9e\xcf\x8c\x35\x08\xeb\x90\x70\xbf\x3b\xdb\xec\x0c\x3d\x7b\x44\x70\xb1\x6e\x0e\x5c\xae\xa5\xc1\x7b\xd5\x34\x38\x57\x4e\x60\x5e\x0a\xcf\xeb\x93\x67\x9e\xde\x1a\x35\xdf\x68\xf4\x73\xc1\x8b\x6d\x67\x2f\xf6\x72\x9c\x2f\x4d\xd0\xe8\x3e\x59\x3c\x05\xce\x55\x29\x43\x97\x74\x36\x1c\x2f\xd7\x86\x9e\xaa\xa9\xd0\x2d\x9d\x06\xa7\xca\xa6\x43\x77\xf9\x54\xe8\xcc\x9f\x06\x05\x9b\xf4\x0e\xfb\x7a\x39\x30\x82\x16\xdb\xce\xc3\x67\x55\x10\x9a\x21\xbe\x36\x9e\x2d\x39\x7a\x37\x4e\x57\xb2\x06\x7b\x9b\x98\x40\xe2\x68\x29\x1b\xde\x6f\x60\xc1\x85\x3a\x36\x1c\x2b\xe3\xc0\xd5\xdd\xac\xa7\xbd\x8d\xcc\x27\x27\xca\x74\x9e\xe6\x6c\x66\x66\xfb\x7b\xdb\xbb\x7b\xb8\x3b\x91\x5d\x84\x08\xf1\x15\x2a\xa5\xc5\xf2\xad\x8b\x13\x59\xe7\x65\x62\x23\x68\x29\xe0\x41\x4d\x3a\x03\x64\x62\x16\x34\xe7\xf1\xa0\x36\xc3\x10\x24\x09\xec\xf6\xca\x64\x73\xd7\xce\x12\xf6\x81\x2b\x8d\x73\x1e\x9f\x2c\x9b\x79\xa5\x20\x8e\x11\x38\x52\x91\xb7\x22\xe9\x49\x31\x9c\x95\xb9\x71\xcc\xae\x9a\x0c\xc6\xa0\x74\x27\x0b\x48\x54\xa6\x31\x9f\xe6\x6c\x65\x1c\x4e\x7c\xdb\xd8\x9f\x34\x3b\x59\xc1\xab\xed\x6d\x62\xf4\xef\xce\xd4\xca\xce\xde\xc4\xf2\xf8\xa9\x1f\xd1\xce\xab\x68\x87\x7e\x5f\x45\x2a\xf7\x6e\x45\x9a\xf1\xa3\xd2\x64\xc3\xef\xe2\x56\x09\x6c\x86\x7f\x97\xa5\x9b\xac\xab\x15\x9b\x44\x8f\xd5\x0e\x51\x4b\x2c\xa7\xc6\x44\x98\xea\xae\x58\x6a\xa5\x9d\xb7\xcd\xb4\x52\x96\xc1\xff\xbf\x1d\xeb\x2c\x46\x9e\x37\x35\x17\x29\x99\x98\xb8\x2a\xbc\xaa\x2f\x0c\x4b\x72\xac\x85\x45\x6e\xbc\x59\xc4\x86\x28\xab\x39\xe3\xb1\xff\x6f\x14\x80\x17\xf1\x48\x44\x10\xfd\x2a\x04\x71\x5e\x19\x81\x5e\xea\x9f\x89\x69\x66\x04\xf1\x24\x9d\x20\x06\xd3\x7f\x6e\x4f\xce\x33\x06\x08\x67\x62\xd4\x3c\xa3\xfa\xa6\x6a\x33\x21\x13\xf2\x87\x10\x3a\x42\x89\xa0\xcd\x53\x22\xe8\xa9\xb8\x6a\xd7\xd3\x7f\xeb\x02\xe7\x1b\xcd\x9a\x1e\x16\xe8\xe8\x94\xb8\x25\x2c\xb9\x20\x33\x66\x6f\x75\x59\xd2\xd1\xe2\xf4\xc0\x13\x19\x31\xf3\x0e\x26\xac\x64\x37\x6e\x8a\xe0\x25\x47\x04\x59\xf9\x7a\xb8\xbb\x18\xea\x70\x22\x94\x5e\x63\xd1\x0a\xe1\x41\x4e\xee\xa5\x85\x71\x27\xba\x4e\xd4\x7d\x7b\xe3\xc3\x4e\xb8\x73\xf3\x34\xfc\xe9\x54\x3c\x9c\x29\xd3\x84\x93\x05\x74\x38\x96\xaf\x08\xb8\x0e\xc3\xde\x9d\x33\xfe\x55\x16\xaf\x79\x37\x71\x15\x7b\x77\xa8\xbf\xb5\x87\xc0\xd2\x7f\xca\xaf\x29\x58\x51\x71\xa6\xe2\x86\xb7\x97\xac\xef\x3c\x50\xf9\xbf\x9f\xdd\x3e\x0b\x7f\xbb\x77\x05\xbe\xbe\xff\x01\xdc\x3a\x9f\x08\x3d\xe5\xca\x70\xb6\x94\x80\xae\x12\x1a\xbc\x5b\x44\xc0\xbb\xbb\x10\x12\x39\x8e\xe5\x2b\x41\x7d\xb2\xda\x63\xe4\xd1\xe2\xeb\xe5\xc8\xff\x4f\xcb\x5f\x16\x60\x16\x96\xbd\xcd\xf5\xd1\x3e\xe9\x72\x68\xab\x89\x84\x43\x7b\xe2\xa1\xbb\x33\x17\xde\x6f\x32\x82\x2b\x0d\x93\xe0\x72\x9d\x02\x85\x4b\x32\x05\xb8\x20\xa5\x23\x1f\x1a\x55\xfe\x69\xe4\x72\x06\x75\x67\xde\x24\x90\x6c\xd1\xb9\xb3\x2c\x40\x18\x40\xd0\x13\xfe\xad\x3e\xb2\xc0\xdd\x45\xbf\x60\x93\xee\xc7\xa4\x7f\x7b\x2a\x08\xb8\x80\xfb\x94\x9e\x4a\x3a\x9c\x2a\x9d\x0c\x47\xa4\x76\x70\xa8\xc6\x17\x39\x4c\x1e\xe1\x30\x8c\xf7\x6a\xd0\xbe\xfc\x47\x1e\xa7\x77\xd1\xa1\x72\xbb\xfa\x03\xec\x1b\xa1\xff\x4e\xf9\x6b\x42\x4d\x63\xf6\x67\x4e\x1b\x7c\x4f\x4a\x60\x39\x74\xb8\x52\x2f\xc7\xd5\x06\x1a\x5c\x6c\x64\xc0\x85\x3d\xd6\x78\xad\x48\xa5\xf5\x0e\x81\xb2\x19\xb2\x7d\x5f\x46\xa7\xda\x87\xe4\x71\x46\x42\x83\x8a\x6d\x1a\x5f\x86\xf9\x5b\x8b\xc6\x2a\x0f\xf7\x49\x6a\x08\x73\x84\x23\xee\x9f\xec\xb6\x44\x72\x8f\xee\xdd\xa9\x0a\x1d\x79\xaa\x70\xa0\x40\x15\x3a\x25\x2a\x70\xa2\x44\x05\x8e\x23\x8e\xec\x52\xc1\x34\x15\x2a\xfd\xf0\x2e\x55\x2a\xed\x58\xb1\x0a\x1c\x2c\x94\xdb\x52\xf6\x45\xa4\xed\x0c\x38\x94\x37\x03\x5a\xd2\x67\x40\x43\xf2\x2c\xd8\x18\xc1\xbf\xb4\x78\xa1\x83\x08\xcb\xb0\x42\x38\x20\xec\x11\x1c\x84\x32\x62\x3a\xc2\x09\x21\xf3\xf5\xb2\x5f\x9b\xbd\x89\x79\xbd\xa7\x7a\x2e\xb4\xe2\xfe\xaa\x31\x53\x80\xfb\x33\x0e\x7c\xda\xce\x82\x1b\xfb\x58\x70\xa9\x9e\x03\x75\x62\x01\x14\x6e\x33\x85\x73\x32\x63\xf8\xa4\x8d\x05\xd7\x5b\xd9\x70\xb8\x98\x0b\x59\x9b\xcd\xe0\x80\x84\x0b\xd7\x9a\xd9\xf0\x29\xa6\x7f\xbc\x8f\x0d\xdd\xd5\x1c\x28\x49\x60\x41\x7d\xba\x0e\xbc\x1d\x6e\x5e\xeb\xef\x6d\x67\x87\xe5\x24\x22\xb2\x11\x0c\x84\xd2\x90\x0f\xe8\x21\x8b\x6d\xa7\xe9\x1b\xfb\x4e\xdf\x15\x3f\xef\xd2\x87\xcd\x6a\x70\xab\x83\x4d\xe1\x36\xa9\xf7\x0f\x5f\xb3\xe0\x26\x5e\x7f\xdc\x36\x94\x3e\x0a\xd7\xf7\x71\x7e\xb4\x1f\xd2\x5f\x1c\x61\xc1\xdf\x8e\xeb\xc2\xbd\xc3\xca\x50\x9b\xa6\x7b\x43\x68\xb3\xc0\x44\x60\xbe\x50\x4d\x68\xe3\x39\xed\x65\x6d\xe1\xe5\xe1\xa8\x14\x19\x6c\x15\xb0\x3e\xc2\x3c\x66\x7d\x84\x65\xab\x24\xc1\x08\xca\x52\x0c\xa0\x70\x3b\x07\xea\xb3\x04\xd0\x90\x8d\x75\x4f\x98\x0f\x25\x49\x86\x50\x94\x68\x04\xa5\xc9\x7c\x4c\xe3\xe3\x7e\x52\x80\xf7\x0c\xa8\x4c\x35\x80\xec\xad\xec\xef\xd6\x2d\x13\xe6\x6c\x8e\xb6\x8a\x49\x89\xb1\x48\xab\x4c\x31\x3b\x7e\xb6\xc6\xe8\xf1\xd5\x46\xd5\xc1\x7d\x39\xea\x7f\x6b\xcc\xd0\xfe\x73\x5a\xac\x71\x25\xb6\x07\xe7\x55\xfd\x70\x55\xa8\xf9\xdb\x65\x29\x7a\xb8\x4f\xd5\x81\xf2\x54\x06\xb4\x49\x78\xb0\x6f\x17\x0f\xaf\x59\x58\xae\x2e\x72\xd0\x03\x59\xc6\x7c\xd8\x9b\xcf\x87\xa6\x1c\x1e\xda\x91\xf7\x5a\x50\xb0\xdd\xf0\xef\x81\x8b\xec\x04\xc3\xf9\xf0\x4d\x44\x7a\xd2\x54\xd3\x8f\xef\x1f\xd3\x81\xbb\x07\x15\xe1\xf3\x03\x04\x5c\xdf\xab\x0c\xa5\x89\xf3\x2e\x7b\xba\x3b\x31\xc6\x2a\x3f\x7a\x89\xe5\xea\xc2\xed\x86\xdf\x97\x26\xeb\x7d\x9f\xbf\x8d\xfd\xb4\x39\x8f\x0b\x7b\xf2\xb9\x20\x49\x34\xc6\xb2\xf4\x91\x83\x01\x72\x91\xa7\xd5\x88\x79\xe8\x97\x79\x4f\x2a\x52\x75\x7f\xc8\xd8\xcc\xbc\x87\xe5\x8f\xcc\x3d\x61\x01\x36\xee\xef\x37\x18\xff\xf0\xf0\x0c\x03\x0e\x49\x18\xdf\xee\xc9\xd2\xe8\xeb\x28\x98\x7b\x37\x77\x0b\xeb\x38\xee\xf5\x2d\xc6\x2a\x3f\xc4\xd7\x56\x3b\x22\xd8\xca\x79\x59\xa0\x50\xb4\x29\x5a\x70\xb8\x1a\xfb\x50\x75\xba\x36\x14\x27\x19\x53\x7b\xff\x86\x1c\x3e\xfa\x80\x01\x64\x7a\xd1\x0e\x83\xe7\xc8\x37\x23\x2a\xc4\xd2\x31\x3c\x50\x68\x8f\xef\x20\xd3\x87\xf3\x59\xb1\xc4\x3a\xf0\xc3\xbd\x9c\x67\x5f\x1e\x63\x42\xc6\x46\x8b\x7a\x9e\xd9\x02\x3d\x67\x91\x8b\xbe\x81\xf1\xe2\xa9\xaf\xf2\xff\x68\xd9\xb6\x86\x5b\x25\x13\x6b\x61\x59\x5a\x58\x77\x36\x34\xe5\xf2\xa9\xf2\x2b\x52\xe7\x21\x74\x90\x93\xde\xf3\xd5\x61\xe6\x4b\x5f\xf6\x6c\xd0\x22\x5b\xe1\x69\x29\xf7\x9f\x07\x8b\x79\x7f\x45\x5e\x36\xe3\x2d\x73\xb4\xe0\x7b\x43\x56\xc1\xf6\x79\xfd\xd8\xb6\x7d\x19\x9b\x79\x7d\x15\x69\x26\x7d\xc5\x49\xa6\xf7\x0a\x13\x38\x03\xb2\x0c\x03\xac\x3f\xf3\x79\x54\xc8\xcb\xe7\x39\x27\x47\x87\x29\x9b\xa2\xad\x72\x56\x85\x5a\x07\xff\x27\x65\x93\x82\x7d\x75\xe6\x02\x57\x67\x1d\x12\x1e\xae\x8e\x3a\xae\xce\x0e\x5a\x5e\xee\xf6\x8e\x45\x3b\x4c\x1f\xb4\x14\xe2\x7b\x58\xa6\xe0\xf9\x9a\x70\xe1\x98\xf3\xac\x1e\xd3\x45\x69\x96\x8e\xdb\x6b\xdd\x2b\xf8\x79\xd9\xb1\x4b\x93\x4d\xbe\x6a\xc5\xf2\xeb\xb3\xf8\xcf\xd7\x86\x0b\xc3\x5e\x67\xfe\xbf\x24\xbe\x0b\xed\xb4\x93\x62\x2c\x9a\x24\x89\xa6\x47\x32\x36\x9b\x1f\x5a\x1e\x64\xe3\xf8\x7b\x96\x1f\xb9\xc4\x86\xa6\xa1\xb1\x10\xf7\x60\x7e\x8a\x2a\x2a\x3e\x8a\x9e\x6e\x0e\xbf\xf9\x5e\x6c\x42\x26\xe4\x8f\x26\xf0\x2a\xe9\x47\x10\x2a\x30\x48\x28\xc0\x00\x41\x83\x27\x68\x8e\x48\x7f\x44\x10\xa2\xeb\x04\xa1\x8d\x5a\x7f\x80\x50\x50\x18\x24\x94\x33\x48\x80\x58\xbf\x04\xc4\xa2\x26\xc4\x39\x18\x14\xc1\x08\x5e\x21\xe4\x26\xc0\x14\x11\x4d\x8c\x3a\xa7\x18\x73\x55\x9e\x90\x09\x99\x90\x09\x99\x90\x09\x99\x90\xff\x16\x99\xac\x40\x28\x98\x1a\x6b\xcc\xb1\xb7\x30\x64\xda\x99\x1b\xf2\x6d\x2d\x0c\xcd\x9c\x1c\x84\x66\x2e\xae\x6e\x5c\x57\x57\x37\xa6\x93\x68\x81\x36\x53\x10\x3c\xf9\x4d\xf3\x1c\x2d\x1c\x03\x55\x05\x1b\x53\x3d\x86\xaf\x87\xc5\xaa\x75\x2b\x7c\x1a\x93\xe2\xc2\xff\x24\xc9\x89\xed\xaf\x97\xa6\x3c\x68\xac\x11\x3f\x2c\x17\x2f\x7a\x98\xbb\x5e\xef\xab\xb4\xb5\x46\xfd\x5b\x22\xb9\x37\x63\xc2\x05\xa7\x96\x05\x5a\x15\xf9\x7a\x3b\x44\xb9\xba\xb8\x09\x84\x76\x0b\x5f\x7a\xc6\xf2\x7b\x88\x9d\x85\xa1\x51\x68\x80\x7d\x6e\xae\x78\xdd\x9d\xa3\x07\x2b\x9e\x7d\x74\xed\x10\x7c\xfe\x97\x6e\xf8\xf2\x8b\x4b\xd0\x7f\xfb\x18\x5c\x3b\x10\x0c\xdd\x25\x33\xa8\xb3\x4a\xf2\x8c\xee\x54\xa1\x02\x1c\xcd\x9b\x04\x6d\xe2\xb7\xa0\x3a\x61\xee\xb3\xf4\x75\x46\x5f\xad\x09\x35\x3b\xe1\xe7\xed\xb0\xce\xc9\xd9\x5d\x57\x87\xb3\xe4\x95\x31\xda\xd7\x25\x1c\xc3\x59\x0a\x1e\x22\xae\x6f\x72\xfc\xb2\xeb\xdd\x27\xeb\x06\xfb\xfb\xce\xc1\xfd\x2f\xde\x83\x7b\x7d\xe7\xa1\xbf\xef\x22\xdc\xb9\xde\x0e\xe7\x64\x02\x38\x99\x4f\x20\x67\xc4\xae\xe1\xb3\x4e\xf9\x39\x30\x59\x9f\xae\x22\xd4\x45\x34\x38\x9c\x3b\x19\xa4\xdb\x35\x9e\x6d\x8d\xe2\x7e\x12\xb8\xc8\x6e\x9b\xc8\xd9\x4d\xeb\xb7\xe4\xce\x65\xa8\x2b\x84\xf8\xda\xad\xad\x2c\x4e\xf8\x07\xe9\xeb\xaf\xbf\xec\x85\xfb\xfd\x97\xe0\x7f\xee\xf5\xc2\xd7\xf7\xaf\xc1\x57\xf7\x2e\xc2\xb5\x76\x07\x38\x5f\x41\xc0\x79\x29\x1d\xce\x55\xc8\xcf\x8d\xbb\x8a\x69\x70\x5a\xf2\x63\x3d\x46\xea\x33\x54\x97\x93\x85\x8a\x50\x9d\x38\x77\x20\x26\xdc\xa4\xd7\xdb\xd3\x69\x91\xa3\x68\xc1\xeb\x3c\xe3\x1f\x11\xef\x05\xc2\xa0\x75\x2b\x16\xfd\xa3\x49\xb6\x13\xf6\xd4\x91\x48\x87\x8e\x16\x09\x74\x9d\x6a\x81\x4f\x6f\x5c\x84\x5b\x17\x92\xe1\xb2\x8c\x3c\x6b\x26\xa8\xf3\xe6\xde\xfa\x17\xcf\x9e\xcf\x57\xd2\xb1\x4f\x0d\xd5\xa5\xf0\xc5\xba\x90\xf5\x38\x82\xed\x91\x11\x63\xf4\x30\x68\x91\xdd\x4e\x27\x91\xfb\x5b\xaf\x93\xbb\xbb\x9b\x2b\x7b\x4b\x14\xf7\xe6\x9e\x34\x15\xd8\x9b\x3a\x19\x3a\xb2\x27\xc3\xa1\x82\x29\xd0\x96\x3d\x0d\x1a\xd3\x67\x83\x34\x95\x0d\x95\xc9\xf3\xa0\xb7\xd9\x10\x7a\x5b\xf8\xc8\xfd\xe7\xe7\xe7\xc3\xf5\x19\x7d\x8e\x7e\xaa\xf0\xc5\xf6\x38\x23\xa1\x43\x69\x9c\xd6\xd3\x50\x7f\xeb\x1a\x9c\xbb\xd4\x5e\x07\x77\x43\x41\x30\x7d\xa9\x9f\x8d\x64\x77\xea\x4c\xe8\x46\x3f\x9d\xaf\xa0\x8d\x9c\xa7\x5f\x45\x5c\x6b\x24\x35\x01\x57\x1b\xa7\xc3\x9e\xf2\x08\xc8\x11\x6f\x84\x6b\xad\x5c\xea\x1b\x86\x9f\xd6\x61\x04\x58\x97\xf7\xb1\xad\x7a\xca\xe9\xa3\xe2\x01\xc3\x6d\x41\x83\xaa\x04\xf5\xe7\x58\x07\x99\xc8\xd9\x7d\xc6\xaf\xe5\xef\xe2\xe2\xc6\x4c\x58\xc9\xf9\xfc\xf4\x2e\x05\xaa\x2f\x5f\xae\x7d\x39\xa7\xde\xfa\x49\xd0\xbd\xdb\x15\x7a\xde\x89\x81\xab\xcd\xea\x68\x47\x1f\x9b\xff\xa8\x7a\x5c\xaa\x51\xa0\xc6\xc9\x4f\xeb\x50\x1e\xaf\xf9\x34\x68\xb1\x6d\xe6\xaf\x1d\x0f\xbe\x5e\x8e\xcb\xc9\x79\xa2\xa7\x84\xc0\xb2\x68\x43\xfe\x7e\x39\x3e\x68\x54\x44\x28\x53\xb1\x90\xb1\x6c\x7e\x8a\xe1\xfc\x2e\x55\xd3\x70\x7c\x20\xf7\xa1\x71\x7d\x16\x7d\x95\xbb\xc1\xe0\x11\xce\xb1\xe3\x3a\x97\x0c\x5e\x6c\xab\x8a\xb0\x45\xf8\x22\x16\x21\x16\xe3\x58\xf2\x59\xea\x67\xdd\x56\xb2\x55\x9b\x8c\x7b\xc1\x9e\xcc\xd9\xd0\x96\xab\x06\xad\xd9\xb3\x29\xb4\xe3\xf5\xfe\x3c\x35\x68\x47\xec\xcb\x51\x83\x96\xac\xd9\x08\x35\x2a\x7d\xf8\xbe\x2d\x67\xc8\x66\xe8\xb9\x96\xa1\xe7\xc8\xdf\xf7\x66\x91\xf9\xa8\x51\x79\x76\xe4\xab\x51\xb6\x0d\x29\x6a\x20\x4b\x54\xa3\xe2\x2a\xd2\xed\x9a\x10\x1e\x60\xf5\x89\xbf\x8f\x7d\x28\xc9\x67\x08\x23\xfc\x10\x0b\x83\xe5\xdf\xc3\xd0\x11\x93\x11\x6c\xc4\x1a\xc4\x2d\xc4\x05\x9c\x97\x93\x56\x2c\xb1\xfc\xa0\x35\x47\x03\xf6\x17\xe8\x83\x24\xc1\x1c\x56\x2e\xb5\x06\xf1\x46\x0b\x38\x52\xc2\x83\x8b\x75\xf3\xe1\x5a\xb3\x31\x5c\x6e\x34\x86\xae\x2a\x2e\xb4\xe5\x0b\xa8\xdf\x96\xfa\xdb\xc2\xda\x70\x6b\x68\xc8\x34\x81\xb3\x35\x5c\xb8\xba\xdb\x18\x2e\x35\x18\x43\x27\x3e\x93\xb5\xd9\x02\x42\x7c\x6d\x21\x66\xb9\x10\xc7\xbd\x29\x1c\x2b\xe3\xc1\x85\x5a\xcc\x07\x6d\xae\x34\x19\x43\x77\x35\x17\xea\x33\x78\xb0\x3a\xd4\x1c\xe2\x56\xf2\xa1\x30\x9e\x09\xa1\x7e\x36\xef\xa2\x2f\x83\x90\x53\x34\xa2\x06\x71\x17\xd1\x8c\x20\x63\x46\xb3\x11\xb4\x51\xed\xa0\x84\xf0\x44\xe8\x28\xa9\x87\x4f\xda\x10\x69\x72\xfc\x93\xd6\x49\xf0\xe0\xe4\x0c\xf8\xf6\x2c\x03\xee\x1e\x66\xc1\xdf\xdf\x65\xc0\x77\x3d\x46\x78\x8f\xe8\x96\xe3\x31\x5e\x93\x69\x0f\xf1\xfa\x5a\x33\x07\xfa\x0e\xb1\xa8\xfb\xc7\xc3\x36\x43\xbf\x7f\xd3\x65\x04\x97\x91\xe7\xdd\x23\x2c\x78\xdc\x63\xf4\xd2\x7c\x1e\xa1\xbe\xdd\xa1\x0d\xb7\xda\x94\xa1\xef\x00\x1d\x32\x37\xb2\x3f\x77\x70\x70\x33\x1e\xc5\x8f\x8c\x6f\xe9\xfe\x52\x9f\xd2\x64\x06\x4e\x5a\x1b\x6e\x7e\xfc\x74\xf9\x2c\xb8\xd6\xa4\x01\x37\xda\x38\x54\x6c\xe9\xa3\x56\x0e\xc2\x98\xc2\xf5\x7d\x6c\xf8\x73\xcb\x8f\x20\xe3\x5d\x1f\xa3\xdd\x8d\x7d\x1c\xca\xee\x3a\x65\xfb\xe3\xef\xe4\xf5\x27\xed\x1c\x2a\x1e\x25\xcf\x03\x6d\x31\xcf\x1b\xfb\xd8\x54\x5e\x9f\xb6\xb3\xe1\xde\x51\x26\x3c\xec\x62\xc0\x83\x53\x1a\xf0\xc5\x41\x05\x38\x51\x3a\x0b\xfc\xbc\xec\x8b\x39\x02\x1f\x1d\x9e\x99\x97\x9a\x11\xdb\x6f\x5c\x63\xda\xcf\xdb\x1e\xd7\x5c\x9b\xb7\x97\x05\x0a\x6b\x22\x83\x85\xd2\xe5\x81\x36\xed\x1b\xa3\x4d\x7f\xc8\xdf\xc6\x80\xdc\x38\x39\x12\xd6\x9a\x40\x59\x8a\x09\x94\x26\x9b\x40\x79\xaa\x09\x24\xc7\x9a\x41\xe6\x66\xd6\xc8\xef\x29\xb1\x5c\xc8\xdf\x66\x46\xd9\x90\xc8\xde\x6a\x06\xa9\xb1\xc6\x98\x66\x04\x79\xf1\x0c\xc8\xda\xc2\x82\xe8\x25\xc2\xdb\x51\x21\xd6\x55\x24\x56\x87\x0a\xeb\xb7\xad\xb6\x3c\x55\x94\x60\x7a\xef\xdd\x2a\xee\xc0\x17\x47\x34\xe1\x76\x3b\x19\x4f\x32\xfc\x21\x6f\x0b\xf3\xee\xce\x58\xce\x87\xab\x43\x2d\xda\xbc\x3d\x1c\x43\x9d\x1c\x5d\xc7\x35\xbf\x4e\x9d\x92\x4a\xf5\x2f\x91\xa3\xab\x49\x4a\x2c\xfb\xab\x86\x6c\x4d\x90\xee\xd4\x81\xda\x0c\x4d\x90\x24\x72\xe0\x98\x74\x3e\xb4\x4b\x78\x70\xa0\x84\x0b\x95\x69\x5c\xfc\x4d\x1b\x2a\x52\x75\x87\x62\x63\x06\x54\x3c\x66\x7f\x11\x97\x8a\x8f\x91\xb1\xa1\xf2\x14\x3d\x68\xc2\x31\x45\xc6\x8b\xea\xb2\x34\x91\xbf\x45\x13\x41\x24\xd1\x09\x22\x02\xcb\x09\xa3\xe9\x18\xba\x4f\x11\xda\x38\xb2\x7d\x3c\xec\xd2\xf2\xe2\xcc\x7f\xb8\xd7\xa9\x0e\x9f\x63\x3f\xfa\xec\x1d\x1a\xfc\x65\xbf\x02\x7c\xb0\x7b\x2a\x34\x88\x75\x9f\x44\x04\x09\x1b\x9d\x45\x2e\xe3\xfe\xc6\xca\xcd\xd9\xc5\x24\x29\x86\xf3\x75\x15\x72\x2f\x49\xd2\x87\xca\x54\x1d\x28\xd8\x6e\x4c\xf1\xde\x8d\x1c\x5b\x0a\xf8\xe8\x63\x1e\x94\x24\xeb\x62\xbd\x0c\xa0\x28\x51\x9f\x8a\x25\xca\x32\x04\xb0\x9b\x8c\xb7\x66\xf3\xc9\x58\x06\x14\xef\x30\x80\x1a\xb1\x36\xb6\x99\x1e\x54\x61\x1d\x56\x2c\x25\xf9\xe7\xfc\x6c\x0f\x1a\xe0\x63\xbb\xb1\xb3\x94\x3b\x40\xc6\xd6\xbe\xec\x9c\x8d\x75\xa0\xc1\xdd\x83\x04\x85\xbf\x1e\xc1\x3d\x47\xd1\x9c\xc1\xf0\x00\xeb\x5d\xd6\x36\x1e\x93\xc6\xc3\xdf\xd5\xd9\xc5\x34\x19\xf9\x93\xbe\x23\x63\x89\x0d\xd9\x1a\x50\x98\x60\x0c\x47\xa5\x5c\xe8\x28\xe6\xc1\x21\x9c\x4b\x48\xfe\x95\x69\x5a\x20\x4d\xd3\x86\xaa\x9d\x5a\x50\x9a\x44\xfa\x5f\x00\xef\xe0\xdc\xb3\xbf\x88\x47\xb5\x05\xe9\xff\xe6\x3c\x75\xa8\xcb\xd4\xc4\xfa\x90\xfe\xb7\x6c\x26\x88\xd4\x17\xf8\xeb\xb2\x5c\xa6\xc5\xad\xb2\xec\x26\xc7\xc2\x37\x38\x5f\xdc\x3f\xca\x80\xbd\xd9\x7a\x38\x16\xd4\xe0\xa2\x6c\x3a\x74\x49\x55\xa1\x21\x5d\xf7\x5f\x91\xc1\x56\x07\xbd\x3c\x9c\xc6\xfc\xbe\x71\xb4\xe0\x18\x62\x60\x59\x07\xe2\x56\xf1\xba\xb7\xae\xe2\x75\x6d\x5e\xc1\xbf\x10\x1b\x69\xf6\xb8\x22\x8d\x89\xfd\x85\x01\x95\x3b\x19\x90\x14\x23\x40\xff\x73\xd1\xaf\x1c\xa8\x16\x73\x20\x63\xb3\x00\xfb\x3a\x07\xeb\xc4\xa0\x90\xb5\x75\x3e\xce\xad\xe6\x9f\x61\x1e\x67\xc8\x3c\x36\x46\x09\x7a\x70\x7f\x92\x64\x65\xb5\xf0\x05\xfe\x7c\x33\x91\x66\xe6\x26\xf3\x9b\x0f\x4e\x33\xe0\x1b\xf4\xff\xad\x0e\x16\x44\x06\xdb\x9e\xc4\x36\xd9\x12\x1d\x62\x99\x1f\xe6\x6f\xbd\xc3\xd3\xdd\x69\x91\x8f\xa7\xe3\xec\xf1\x70\x27\x05\xc7\x31\x7d\x81\x8b\x68\x9a\xc0\xcc\x7b\x3a\xcf\xc4\x67\x1a\xb6\x9b\xe1\x96\x95\xbc\xeb\x2d\x05\x73\xa0\x96\xf2\xa5\x3a\x14\xed\x60\x60\x3f\x97\xf7\xa5\x56\xec\xef\xf5\xd9\x64\x7f\xd7\xa1\x7c\x2d\xc3\xf1\x52\x8e\xe3\x62\x89\x9f\x8d\x78\x8e\x7e\xc8\x14\x13\x73\xef\xe9\x96\x56\x9e\xd3\xfd\xbd\xed\x7f\xf6\x6e\x39\xc2\xff\x14\x83\x9a\x4f\x8f\x94\x72\x9f\xf9\x78\xda\x45\x8f\x97\xeb\x78\xc4\xd9\xc9\x55\x1d\xf9\x7f\x44\xf6\xa3\xf2\x14\xf9\x78\x2d\x4c\x60\x52\x71\xf4\x1a\xb1\x00\x64\x99\x02\x2a\x76\x2c\xc1\x71\x40\xf2\x26\xfb\x7b\xf1\x0e\x7d\x08\x0f\x14\xa6\xfc\x52\xde\x3a\x4c\xd7\x69\xf1\xd8\x7f\xee\x1f\x67\xc2\xcd\x77\x58\xb0\x31\xca\xea\x9c\xb5\xcd\xf8\xfa\xc9\x78\x05\xe7\x23\x75\xec\x07\x1f\xb5\xa2\xff\xeb\xd1\xbf\x8d\x58\x8f\x62\xf4\x7f\x1b\xfa\x9d\x8c\x7b\x92\xf3\x4d\x43\x0e\x17\xca\x70\x3c\x93\xbe\xaf\x16\x6b\xe1\xf8\xd0\x85\xb0\x00\xeb\x5f\xe4\x4f\x0a\xae\xd3\xf1\xad\xf9\x7c\x48\xdf\x60\xfe\x17\x17\x91\x83\xdd\xeb\xe4\x4e\x8a\x8f\x87\xe3\x6c\x1c\x0f\x47\xd3\xd6\xb3\xfb\x52\xd7\xb3\xef\xa4\x6d\x60\xdf\xd9\x10\x65\x7a\x47\xbc\xc9\xfc\x4e\x4e\x9c\x19\x85\xf8\x35\x16\xff\x24\xbf\x6f\xa8\xc2\xb1\x51\x95\xce\xc0\xf9\x7f\x3e\xe0\xda\x39\x2e\xfe\x7e\x5e\x76\x1a\x81\x3e\xb6\x45\xde\x1e\x76\x63\x7e\xdf\xf2\x6b\x64\x39\x8e\x07\x5c\x47\xd4\xb1\x0f\x6b\x93\xb0\xb0\x5c\xa8\xed\xec\xe4\xac\x6d\x66\xe1\xa4\xcd\x37\x15\x69\x1b\xb0\x5c\x34\x97\x07\x59\xd7\x92\x73\x6b\x07\xce\x3d\xe4\x1a\xd1\x98\x23\xc0\x3d\x99\xcd\xb8\xf8\x93\xe2\xe3\x61\xaf\xe4\xe9\xee\xf0\xbb\xbc\xdf\xbf\x4c\x70\x3d\x2d\x95\x65\xf0\xa9\x75\x8b\x9c\xff\xeb\x32\xf9\x80\xeb\xf8\xb8\xf9\xbf\x69\x59\x15\x2a\x2c\xdb\x9d\xcb\x1b\x5a\x7f\xe5\xeb\xdc\xb2\x40\x9b\xd4\x37\xcd\x6b\xbc\x12\x16\x60\x93\x93\xbf\xdd\xec\x1b\xdc\x1b\x7d\x4d\x02\xf7\x3f\x7f\xc7\x71\x19\xf7\xa6\x79\x8d\x57\x02\x7d\xec\x34\xdd\x9c\x1d\xf8\xce\x4e\x0e\x3c\x12\xee\x2e\x0e\x7c\x7c\x2f\x7c\xad\xf3\xe0\x84\x4c\xc8\x84\x4c\xc8\x84\x4c\xc8\x84\xbc\x19\x91\x7f\x2d\xf4\x1b\xe8\x41\x42\x44\xe9\x47\x84\x32\xa5\xfb\x09\x1a\xa5\xcf\x13\x44\x3a\xa9\xc5\xf8\x1e\x44\x6a\xe4\xa0\x82\x7f\x06\x51\x2b\x0f\x69\x85\x51\x7a\x00\x35\x0d\xf5\x7d\x92\x6c\xaf\xfc\x9e\x18\x4a\x27\xd2\xe5\x76\x84\x8a\x3c\xdf\x01\x85\x21\x4d\x93\xeb\x27\x72\x4d\x7b\x62\x82\x7a\x80\x50\x21\xff\xa7\x0b\xb5\xfe\xb0\x1e\x00\xb9\x7e\x4e\xf1\x1e\xd0\x87\x57\xea\xdf\xc6\x4f\xff\x1f\x00\x00\xff\xff\x3a\x44\xf1\x3f\xee\x3a\x00\x00")

func imagesFaviconIcoBytes() ([]byte, error) {
	return bindataRead(
		_imagesFaviconIco,
		"images/favicon.ico",
	)
}

func imagesFaviconIco() (*asset, error) {
	bytes, err := imagesFaviconIcoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "images/favicon.ico", size: 15086, mode: os.FileMode(420), modTime: time.Unix(1460986901, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scriptsApplicationJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func scriptsApplicationJsBytes() ([]byte, error) {
	return bindataRead(
		_scriptsApplicationJs,
		"scripts/application.js",
	)
}

func scriptsApplicationJs() (*asset, error) {
	bytes, err := scriptsApplicationJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/application.js", size: 0, mode: os.FileMode(420), modTime: time.Unix(1460986901, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scriptsVendorJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func scriptsVendorJsBytes() ([]byte, error) {
	return bindataRead(
		_scriptsVendorJs,
		"scripts/vendor.js",
	)
}

func scriptsVendorJs() (*asset, error) {
	bytes, err := scriptsVendorJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/vendor.js", size: 0, mode: os.FileMode(420), modTime: time.Unix(1460986901, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _stylesApplicationCSS = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func stylesApplicationCSSBytes() ([]byte, error) {
	return bindataRead(
		_stylesApplicationCSS,
		"styles/application.css",
	)
}

func stylesApplicationCSS() (*asset, error) {
	bytes, err := stylesApplicationCSSBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "styles/application.css", size: 0, mode: os.FileMode(420), modTime: time.Unix(1460986901, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _stylesVendorCSS = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func stylesVendorCSSBytes() ([]byte, error) {
	return bindataRead(
		_stylesVendorCSS,
		"styles/vendor.css",
	)
}

func stylesVendorCSS() (*asset, error) {
	bytes, err := stylesVendorCSSBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "styles/vendor.css", size: 0, mode: os.FileMode(420), modTime: time.Unix(1460986901, 0)}
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
	"images/favicon.ico":     imagesFaviconIco,
	"scripts/application.js": scriptsApplicationJs,
	"scripts/vendor.js":      scriptsVendorJs,
	"styles/application.css": stylesApplicationCSS,
	"styles/vendor.css":      stylesVendorCSS,
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
	"images": &bintree{nil, map[string]*bintree{
		"favicon.ico": &bintree{imagesFaviconIco, map[string]*bintree{}},
	}},
	"scripts": &bintree{nil, map[string]*bintree{
		"application.js": &bintree{scriptsApplicationJs, map[string]*bintree{}},
		"vendor.js":      &bintree{scriptsVendorJs, map[string]*bintree{}},
	}},
	"styles": &bintree{nil, map[string]*bintree{
		"application.css": &bintree{stylesApplicationCSS, map[string]*bintree{}},
		"vendor.css":      &bintree{stylesVendorCSS, map[string]*bintree{}},
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
