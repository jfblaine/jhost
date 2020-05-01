const (
    htmlIndex    = `<html><body>Welcome!</body></html>`
    inProduction = true
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, htmlIndex)
}

func makeHTTPServer() *http.Server {
    mux := &http.ServeMux{}
    mux.HandleFunc("/", handleIndex)

    // set timeouts so that a slow or malicious client doesn't
    // hold resources forever
    return &http.Server{
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 5 * time.Second,
        IdleTimeout:  120 * time.Second,
        Handler:      handler,
    }
}

func main() {
    var httpsSrv *http.Server
    var m *autocert.Handler

        httpsSrv = makeHTTPServer()
        m := &autocert.Manager{
            Prompt:     autocert.AcceptTOS,
            HostPolicy: hostPolicy,
            Cache:      autocert.DirCache(dataDir),
        }
        httpsSrv.Addr = ":8443"
        httpsSrv.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}

        go func() {
            err := httpsSrv.ListenAndServeTLS("", "")
            if err != nil {
                log.Fatalf("httpsSrv.ListendAndServeTLS() failed with %s", err)
            }
        }()
    }

    httpSrv := makeHTTPServer()
    if m != nil {
        // allow autocert handle Let's Encrypt auth callbacks over HTTP.
        // it'll pass all other urls to our hanlder
        httpSrv.Handler = m.HTTPHandler(httpServ.Handler)
    }
    httpSrv.Addr = ":8080"
    err := httpSrv.ListenAndServe()
    if err != nil {
        log.Fatalf("httpSrv.ListenAndServe() failed with %s", err)
    }
}
