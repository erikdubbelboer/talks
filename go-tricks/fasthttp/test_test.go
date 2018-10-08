package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

var (
	responseBody = []byte("test")
)

func BenchmarkHttp(b *testing.B) {
	ln := fasthttputil.NewInmemoryListener()

	go http.Serve(ln, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write(responseBody)
	}))

	c := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return ln.Dial()
			},
		},
	}

	for i := 0; i < b.N; i++ {
		res, err := c.Get("http://example.com")
		if err != nil {
			panic(err)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		if !bytes.Equal(body, responseBody) {
			panic("wrong body")
		}

		// Close the body for connection reuse.
		res.Body.Close()
	}
}

func BenchmarkFast(b *testing.B) {
	ln := fasthttputil.NewInmemoryListener()

	go fasthttp.Serve(ln, func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(200)
		ctx.SetBody(responseBody)
	})

	c := fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}

	// Pointer to a heap allocated slice for reuse.
	var buf []byte

	for i := 0; i < b.N; i++ {
		_, body, err := c.Get(buf, "http://example.com")
		if err != nil {
			panic(err)
		}

		if !bytes.Equal(body, responseBody) {
			panic("wrong body")
		}

		// Keep for reuse.
		buf = body
	}
}
