package buildin

import (
	"bytes"
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
	"io"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// ascii85 数据编码(5 个 ascii 字符表示 4 个字节)，pdf 文档的编码格式
func TestAscii85(t *testing.T) {
	Convey("test ascii85 encode/decode", t, func() {
		{
			buf := make([]byte, 32)
			n := ascii85.Encode(buf, []byte("hello world"))
			So(buf[0:n], ShouldResemble, []byte("BOu!rD]j7BEbo7"))
		}
		{
			buf := make([]byte, 32)
			n, _, _ := ascii85.Decode(buf, []byte("BOu!rD]j7BEbo7"), true)
			So(buf[0:n], ShouldResemble, []byte("hello world"))
		}
	})

	Convey("test ascii85 encoder/decoder", t, func() {
		{
			buf := make([]byte, 4)
			reader := ascii85.NewDecoder(bytes.NewReader([]byte("BOu!rD]j7BEbo7")))
			writer := &bytes.Buffer{}
			for n, err := reader.Read(buf); err != io.EOF; n, err = reader.Read(buf) {
				_, _ = writer.Write(buf[0:n])
			}
			So(writer.String(), ShouldEqual, "hello world")
		}
		{
			buffer := &bytes.Buffer{}
			writer := ascii85.NewEncoder(buffer)
			_, _ = writer.Write([]byte("hello"))
			_, _ = writer.Write([]byte(" "))
			_, _ = writer.Write([]byte("world"))
			_ = writer.Close()
			So(buffer.String(), ShouldEqual, "BOu!rD]j7BEbo7")
		}
	})

	Convey("support chinese", t, func() {
		{
			buf := make([]byte, 32)
			n := ascii85.Encode(buf, []byte("你好世界"))
			So(buf[0:n], ShouldResemble, []byte("jLq5JV7ks\"QKONl"))
		}
		{
			buf := make([]byte, 32)
			n, _, _ := ascii85.Decode(buf, []byte("jLq5JV7ks\"QKONl"), true)
			So(buf[0:n], ShouldResemble, []byte("你好世界"))
		}
	})
}

func TestBase64(t *testing.T) {
	Convey("test base64 encode/decode", t, func() {
		{
			buf := make([]byte, 32)
			base64.StdEncoding.Encode(buf, []byte("hello world"))
			So(buf[0:base64.StdEncoding.EncodedLen(len("hello world"))], ShouldResemble, []byte("aGVsbG8gd29ybGQ="))
		}
		{
			buf := make([]byte, 32)
			n, _ := base64.StdEncoding.Decode(buf, []byte("aGVsbG8gd29ybGQ="))
			So(buf[0:n], ShouldResemble, []byte("hello world"))
		}
	})

	Convey("test base64 encode/decode to string", t, func() {
		{
			So(base64.StdEncoding.EncodeToString([]byte("hello world")), ShouldEqual, "aGVsbG8gd29ybGQ=")
			So(base64.URLEncoding.EncodeToString([]byte("hello world")), ShouldEqual, "aGVsbG8gd29ybGQ=")
			So(base64.RawStdEncoding.EncodeToString([]byte("hello world")), ShouldEqual, "aGVsbG8gd29ybGQ")
			So(base64.RawURLEncoding.EncodeToString([]byte("hello world")), ShouldEqual, "aGVsbG8gd29ybGQ")
		}
		{
			buf, _ := base64.StdEncoding.DecodeString("aGVsbG8gd29ybGQ=")
			So(buf, ShouldResemble, []byte("hello world"))
		}
		{
			buf, _ := base64.URLEncoding.DecodeString("aGVsbG8gd29ybGQ=")
			So(buf, ShouldResemble, []byte("hello world"))
		}
		{
			buf, _ := base64.RawStdEncoding.DecodeString("aGVsbG8gd29ybGQ")
			So(buf, ShouldResemble, []byte("hello world"))
		}
		{
			buf, _ := base64.RawURLEncoding.DecodeString("aGVsbG8gd29ybGQ")
			So(buf, ShouldResemble, []byte("hello world"))
		}
	})

	Convey("test base64 encoder/decoder", t, func() {
		{
			buf := make([]byte, 4)
			reader := base64.NewDecoder(base64.StdEncoding, bytes.NewReader([]byte("aGVsbG8gd29ybGQ=")))
			writer := &bytes.Buffer{}
			for n, err := reader.Read(buf); err != io.EOF; n, err = reader.Read(buf) {
				_, _ = writer.Write(buf[0:n])
			}
			So(writer.String(), ShouldEqual, "hello world")
		}
		{
			buffer := &bytes.Buffer{}
			writer := base64.NewEncoder(base64.StdEncoding, buffer)
			_, _ = writer.Write([]byte("hello"))
			_, _ = writer.Write([]byte(" "))
			_, _ = writer.Write([]byte("world"))
			_ = writer.Close()
			So(buffer.String(), ShouldEqual, "aGVsbG8gd29ybGQ=")
		}
	})
}

func TestBase32(t *testing.T) {
	Convey("test base32 encode/decode", t, func() {
		{
			buf := make([]byte, 32)
			base32.StdEncoding.Encode(buf, []byte("hello world"))
			So(buf[0:base32.StdEncoding.EncodedLen(len("hello world"))], ShouldResemble, []byte("NBSWY3DPEB3W64TMMQ======"))
		}
		{
			buf := make([]byte, 32)
			n, _ := base32.StdEncoding.Decode(buf, []byte("NBSWY3DPEB3W64TMMQ======"))
			So(buf[0:n], ShouldResemble, []byte("hello world"))
		}
	})

	Convey("test base32 encode/decode to string", t, func() {
		{
			So(base32.StdEncoding.EncodeToString([]byte("hello world")), ShouldEqual, "NBSWY3DPEB3W64TMMQ======")
			So(base32.HexEncoding.EncodeToString([]byte("hello world")), ShouldEqual, "D1IMOR3F41RMUSJCCG======")
		}
		{
			buf, _ := base32.StdEncoding.DecodeString("NBSWY3DPEB3W64TMMQ======")
			So(buf, ShouldResemble, []byte("hello world"))
		}
		{
			buf, _ := base32.HexEncoding.DecodeString("D1IMOR3F41RMUSJCCG======")
			So(buf, ShouldResemble, []byte("hello world"))
		}
	})

	Convey("test base32 encoder/decoder", t, func() {
		{
			buf := make([]byte, 4)
			reader := base32.NewDecoder(base32.StdEncoding, bytes.NewReader([]byte("NBSWY3DPEB3W64TMMQ======")))
			writer := &bytes.Buffer{}
			for n, err := reader.Read(buf); err != io.EOF; n, err = reader.Read(buf) {
				_, _ = writer.Write(buf[0:n])
			}
			So(writer.String(), ShouldEqual, "hello world")
		}
		{
			buffer := &bytes.Buffer{}
			writer := base32.NewEncoder(base32.StdEncoding, buffer)
			_, _ = writer.Write([]byte("hello"))
			_, _ = writer.Write([]byte(" "))
			_, _ = writer.Write([]byte("world"))
			_ = writer.Close()
			So(buffer.String(), ShouldEqual, "NBSWY3DPEB3W64TMMQ======")
		}
	})
}
