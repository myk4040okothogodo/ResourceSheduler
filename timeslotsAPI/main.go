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
    handlers "github.com/myk4040okothogodo/ResourceSheduler/timeslotsAPI/handlers"
    "github.com/myk4040okothogodo/ReesourceScheduler/timeslotsAPI/data"
)


func main() {
    l := log.New(os.Stdout, "timeslotsAPI", log.LstdFlags)
    v := data.NewValidation()

    db, db_err := data.Init()
    if db_err != nil {
        panic(db_err)
    }
    
    co := data.New(db)
    th := handlers.NewTimeSlots(l, v, co)

    //create a new serve mux and register the handdlers
    sm := mux.NewRouter()

    //handlers for the API
    getT := sm.Methods(http.MethodGet).Subrouter()
    getT.HandleFunc("/timeslots", th.ListAll )
    getT.HandleFunc("/timeslots/{id:[0-9]+}", th.ListSingle)


    putT := sm.Methods(http.MethodPut).Subrouter()
    putT.HandleFunc("/timeslots", th.Update)
    putT.Use(th.MiddlewareValidateTimeSlot) 

    postT := sm.Methods(http.MethodPost).Subrouter()
    postT.HandleFunc("/timeslots", th.Create)
    postT.Use(th.MiddlewareValidateTimeSlot)

    deleteT := sm.Methods(http.MethodDelete).Subrouter()
    deleteT.HandleFunc("/assets/{id:[0-9]+}", th.Delete)

    ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
  

    //create a new server
    s := http.Server{
        Addr:   ":8000",         
        Handler:    ch(sm)
        ErrorLog: l,
        ReadTimeout:   5* time.Second, 
        WriteTimeout: 10* time.Second,
        IdleTimeout:  120* time.Second
    }

    //start the server
    go func() {
        l.Println("Starting server on port 8000")
        err := s.ListenAndServe()
        if err != nil {
            l.printf("Error starting server: %s\n", err)
            os.Exit(1)
        }
    }()

    //Trap sigterm or interrupt and gracefully shutdown the server
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    siganl.Notify(c, os.kill)

    //Block untill a signal is received
    sig := <- c
    log.Println("Got signal:", sig)

    //gracefully shutdown the server, waiting max 30 seconds for current operations to complete
    ctx, _ := context.WithTimeout(context.Background(),  30*time.second)
    s.Shutdown(ctx)


  }
