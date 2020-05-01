package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
	"crypto/tls"
    "time"
    "os"
)

var (
    myHostname        = os.Getenv("HOSTNAME")
    myCrt             = os.Getenv("TLS_CRT")
    myKey             = os.Getenv("TLS_KEY")
    httpPort          = os.Getenv("HTTP_PORT")
    httpsPort         = os.Getenv("HTTPS_PORT")
    hostnameResponse  = fmt.Sprintf("I am running on pod %s\n", myHostname)
)

func handleIndex(w http.ResponseWriter, r * http.Request) {
    io.WriteString(w, hostnameResponse)
}

func makeServerFromMux(mux * http.ServeMux) * http.Server {
    // set timeouts so that a slow or malicious client doesn't
    // hold resources forever
    return &http.Server {
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 5 * time.Second,
        IdleTimeout:  120 * time.Second,
        Handler:      mux,
    }
}

func makeHTTPServer() *http.Server {
    mux := &http.ServeMux{}
    mux.HandleFunc("/", handleIndex)
    return makeServerFromMux(mux)
}

func main() {

    go func() {
        mux := http.NewServeMux()
        mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
            w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
            w.Write([]byte(hostnameResponse))
        })
        cfg := &tls.Config{
            MinVersion:               tls.VersionTLS12,
            CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
            PreferServerCipherSuites: true,
            CipherSuites: []uint16{
                tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
                tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
                tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
                tls.TLS_RSA_WITH_AES_256_CBC_SHA,
            },
        }
        srv := &http.Server{
            Addr:         fmt.Sprintf("%s%s", ":", httpsPort),
            Handler:      mux,
            TLSConfig:    cfg,
            ReadTimeout:  5 * time.Second,
            WriteTimeout: 5 * time.Second,
            IdleTimeout:  120 * time.Second,
            TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
        }
        fmt.Printf("Starting HTTPS server on %s\n", httpsPort)
        log.Fatal(srv.ListenAndServeTLS(myCrt, myKey))
    }()

    var httpSrv * http.Server
    httpSrv = makeHTTPServer()
    var httpFinalPort = fmt.Sprintf("%s%s", ":", httpPort)
    httpSrv.Addr = httpFinalPort
    fmt.Printf("Starting HTTP server on %s\n", httpPort)
    err:= httpSrv.ListenAndServe()
    if err != nil {
        log.Fatalf("httpSrv.ListenAndServe() failed with %s", err)
    }
}
