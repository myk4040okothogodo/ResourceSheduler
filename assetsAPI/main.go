package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"
    gohandlers "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "google.golang.org/grpc"
    handlers "github.com/myk4040okothogodo/ResourceSheduler/assetsAPI/handlers"
    protos "github.com/myk4040okothogodo/ResourceSheduler/protos/assetsdefinition/protos"
    "github.com/myk4040okothogodo/ResourceSheduler/assetsAPI/data"

)

func main() {

    l :=log.New(os.Stdout, "assetsAPI", log.LstdFlags)
    v := data.NewValidation()
    db, db_err :=  data.Init()
    if db_err != nil{
        panic(db_err)
    }
    conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
    if err != nil {
        panic(err)
    }
     
    co := data.New(db)

    defer conn.Close()

    //create client
    cc := protos.NewAssetCheckClient(conn)

    //create the handlers
    ah := handlers.NewAssets(l, v, cc, co)
    
    //create a new serve mux and register the handlers
    sm := mux.NewRouter()

    //handlers for the API
    getA := sm.Methods(http.MethodGet).Subrouter()
    getA.HandleFunc("/assets", ah.ListAll)
    getA.HandleFunc("/assets/{id:[0-9]+}", ah.ListSingle)


    putA := sm.Methods(http.MethodPut).Subrouter()
    putA.HandleFunc("/assets", ah.Update)
    putA.Use(ah.MiddlewareValidateAsset)

    postA := sm.Methods(http.MethodPost).Subrouter()
    postA.HandleFunc("/assets", ah.Create)
    postA.Use(ah.MiddlewareValidateAsset)

    deleteA := sm.Methods(http.MethodDelete).Subrouter()
    deleteA.HandleFunc("/assets/{id:[0-9]+}", ah.Delete)

    //handler for documentation
    

    //getA.Handler("/docs", sh)
    //getA.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
    //
    ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))


    //create a nw server
    s := http.Server {
      Addr:       ":9092",
        Handler:    ch(sm),          //set default handler
        ErrorLog:   l,               // set the logger for the server
        ReadTimeout:  5* time.Second, // max time to read request from the client
        WriteTimeout:  10* time.Second, //Mas time to write response to the client
        IdleTimeout:   120 * time.Second,  //max time for connections to keep alive

    }

    // start the server
    go func() {
        l.Println("Starting server on port 9090")
        err := s.ListenAndServe()
        if err != nil {
          l.Printf("Error starting server: %s\n", err)
          os.Exit(1)
        }

    }()

    // Trap sigterm or interrupt and gracefully shutdown the server
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    signal.Notify(c, os.Kill)

    //Block until a signal is received
    sig := <- c
    log.Println("Got signal:", sig)

    //gracefull shutdown the server, waiting max 30 seconds for current operations to complete
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    s.Shutdown(ctx)
}
