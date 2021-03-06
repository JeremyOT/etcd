package resources

import (
	"bytes"
	"compress/gzip"
	"io"
	"reflect"
	"unsafe"
)

var _views_stats_html = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x56\x4b\x8f\xdb\x36\x10\xbe\xe7\x57\x10\xec\xc1\x09\x20\xc9\x7a\xf8\xb1\x32\x2c\x03\x4d\x0b\xf4\xd2\x14\x45\x53\x14\x28\x8a\x1e\x68\x8a\xb6\x88\xa5\x29\x95\xa4\xad\x75\x5c\xff\xf7\x0e\x45\xd9\x8d\x64\x79\xbb\x41\xf7\x92\x7d\x80\xe2\x70\xe6\x9b\x8f\xc3\x6f\x28\x2d\x73\x7e\x40\x54\x10\xad\x33\xcc\x0c\xcd\x7d\x5a\x4a\x43\xb8\x64\x0a\x35\x53\x6d\x88\xd1\xe8\x74\xa2\xa5\xd8\xef\xa4\x3e\x9f\xe1\xd9\x90\xb5\x60\xbf\x71\xcd\xd7\x5c\x70\x73\x3c\x9f\xf1\xea\x0d\x82\x9f\x1b\xac\x75\x99\x1f\xdb\xb5\xc1\xf5\x4d\xa9\x76\xc4\xf8\x9a\x09\x46\x4d\xa9\x3e\x73\x1d\x74\xbf\xf8\xf9\xdc\xb0\x1d\xea\x9a\xb6\x8a\x54\x05\x46\x72\xeb\x53\xc1\xe9\x63\x86\x75\x51\xd6\x3f\x58\xe3\xdb\x77\x3d\xdc\x06\x5b\x1f\xb6\xe8\xc0\x94\xe6\xa5\xcc\x70\x14\x44\x18\x3d\xed\x84\x84\x4c\x85\x31\xd5\x62\x3c\xae\xeb\x3a\xa8\x93\xa0\x54\xdb\x71\x1c\x86\xe1\x18\xfc\x5b\x97\xc5\x93\xe0\xf2\x71\xc8\x31\x4a\xd3\x74\xdc\xac\x82\x6b\x86\xc3\xea\x09\xa3\xa3\x1b\x6f\x08\x54\x8a\x69\xa6\x0e\xec\x5b\x5d\xc1\x06\x7e\x21\x86\x97\x19\x7e\xfa\xc0\xe5\xef\xf0\x8f\xd1\x81\xb3\xfa\x7d\x69\x41\x50\x88\x66\xf6\x2f\x08\xc3\x39\x46\x4c\xda\xda\xfb\x6b\x42\x1f\xb7\xaa\xdc\xcb\x3c\xc3\x92\xd5\xa8\xe7\x05\x3c\x17\xba\x22\x94\x65\xf8\x92\x67\xa8\x06\x15\x31\x05\xda\x70\x21\x32\xfc\x4d\xf4\x3d\xfc\xbe\xc7\x08\x00\x3f\x24\xa1\x0f\x40\x69\x44\xfd\x68\x16\xcc\xe2\xc8\x0b\xfd\xc4\x1a\x26\x5e\x94\x04\x93\xf9\xe4\x32\x73\x03\x0d\xbd\xd6\xcd\xad\x7a\x9d\xd5\x76\xb8\x49\xae\x9d\xdd\xef\xc4\xb4\xc0\xdf\xcd\xae\xd9\x92\x87\xc4\x9b\x34\xe8\x8e\x92\x77\xe1\xf6\x09\x01\x4d\x6f\x32\x0f\xe2\x24\xa5\x7e\x1a\x4c\xa3\x14\x68\x46\x76\x3e\xf5\xe7\xc1\x3c\x9a\x5d\x26\x6e\xb8\x21\xf0\x31\x0e\x83\xc9\x03\x90\x8e\x83\xf9\xec\x01\x70\xdb\x27\xda\x62\x79\x2e\xce\x6b\xb0\x2e\x13\x37\x7c\x4c\x9c\x8f\xcb\xee\x5d\x79\x7c\xc2\xe3\x81\x2a\x5b\xe9\xac\xde\x74\x95\x3d\x06\x69\xff\x0f\xb1\x37\x0d\xd8\x17\xfb\xaf\xd6\xf8\xf5\x8b\x7d\x9a\x06\x93\x74\x66\x87\x69\x98\x3c\x2b\xf8\x9e\xe7\x0b\x45\xaf\x80\x41\x43\xd8\xd1\xc5\xfd\x06\xa8\x79\x6e\x8a\x0c\x3b\x70\x8c\x0a\xc6\xb7\x85\x81\xaa\x25\x41\x14\x45\x83\x07\xdc\x45\x8c\xc1\xd1\x06\xbe\x32\x2c\x34\x41\x92\xc6\xaf\x00\xeb\xe4\xf8\x8c\x1a\xfb\xd3\xbe\x30\xdd\x4d\xdb\x83\x28\xe2\xd5\xcf\x0c\xde\x19\x3f\x12\xc3\x24\x3d\x2e\xc7\x60\x78\x5e\xdf\x0d\xcc\xbf\x6f\x1b\x8c\x38\x1c\xad\x70\xe1\x7d\xf4\x2f\x23\x28\xb8\x36\x77\xf9\xc1\xda\x00\xb9\xa6\xa1\x10\x65\x42\x54\x24\xcf\xb9\xdc\x36\x65\xb7\x73\xab\xa8\x76\xde\x8f\x29\x18\xc9\x07\x0a\x6c\xf2\x0e\x1b\x49\x76\xcc\xb7\xae\xb0\x47\xc7\xe1\x27\xb0\x2c\xc7\xe6\x05\xb1\xd7\x72\x5c\xcb\xda\x8f\x02\xcb\x2d\x8d\xa5\xb1\x2f\xde\x21\x78\x65\xef\x0c\xc5\x2a\x46\x40\x24\x95\x25\xc3\x25\xb2\xa3\x1e\x68\x95\x0b\x23\x08\xd1\x35\x37\xb4\x40\xf6\xfa\x38\x19\xb5\x67\x8b\x91\x68\x76\x34\xf2\xd0\x86\x08\xcd\x16\x68\xb4\x29\x85\x28\x6b\x30\x9d\xff\x70\x6b\x76\x9b\x28\xcb\x1a\xf8\xc0\x56\xe1\xcf\x3b\x39\x9a\x3c\xf6\x10\xaf\x89\xfc\xba\x60\x90\x4a\xb4\x55\x3b\x9d\xae\x18\xe7\xf3\x12\x4e\x44\x76\xaa\x64\x17\x7d\x73\xac\xa0\xdd\xdf\xba\x90\x77\xa0\x72\xf0\x5a\x0d\xdc\xb3\xcf\x24\xcd\xd9\x86\xec\x85\xe9\xa5\xbb\x0f\x31\x7c\x86\x6d\xd5\xfe\x23\x6b\xe7\xa6\xff\x6b\x4f\x14\x73\xb7\x39\x58\x17\xe8\x34\xfa\xcc\x0e\x8d\xc2\x98\x1c\x2d\x5c\x1d\x5b\x45\x04\x74\xaf\x14\x93\x06\x2d\x51\x3c\xf5\x50\xc7\xbf\x54\x44\x6e\xd9\xfd\x80\x59\xd8\x0b\x50\x2c\xbf\xe7\xbd\xca\xc0\x1d\x3e\xef\x5e\x52\xc8\x01\xe1\xfa\x07\x22\xf6\xec\x7a\x80\x7d\xf4\xbf\x91\xdc\xef\xd6\x4c\x2d\x22\x04\xdf\x95\x3b\xfd\xc5\xb5\x06\xab\xba\xe9\x87\x5b\xed\x83\xd1\xf6\xf7\xe0\x0d\xd2\x3e\xb6\xc3\x3f\x01\x00\x00\xff\xff\x0a\x0c\x99\x1a\x0e\x0b\x00\x00"

// views_stats_html returns raw, uncompressed file data.
func views_stats_html() []byte {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&_views_stats_html))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(_views_stats_html)
	bx.Cap = bx.Len

	gz, err := gzip.NewReader(bytes.NewBuffer(b))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var buf bytes.Buffer
	io.Copy(&buf, gz)
	gz.Close()

	return buf.Bytes()
}


func init() {
	go_bindata["/views/stats.html"] = views_stats_html
}
