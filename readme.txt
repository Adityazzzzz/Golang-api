go mod init github.com/Adityazzzzz/studentAPI-Go

to install dependency:  
    go get -u <url>

Basic Skeleton:
    /*
    TODO:
        load config
        db setup
        routes
        server
    */

1   //config
        cfg := config.MustLoad()

2   //db

3   //routes
        router := http.NewServeMux()
        router.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request){
            res.Write([]byte("welcome to api"))
        })

4   //server
        server := http.Server{
            Addr: cfg.Addr,
            Handler: router,
        }
        err:=server.ListenAndServe()
        if err!=nil {
            log.Fatal("Failed to start server")
        }

        fmt.Println("server started")