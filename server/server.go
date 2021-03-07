package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"sacache"
)

const (
	apiCacheServiceName = "sacache"
	apiBaseURL          = "/api/"

	// url for cache service.
	apiCacheURL = apiBaseURL + apiCacheServiceName + "/"
	version     = "0.1.0"
)

var (
	port    int
	logfile string

	cache *sacache.SaCache
)

func init() {
	flag.IntVar(&port, "port", 9999, "the port to listen.")
	flag.StringVar(&logfile, "logfile", "", "Path of logfile.")
}

func cacheGetHandler(rw http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len(apiCacheURL):]
	if key == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("key is empty."))
		log.Print("key is empty.")
		return
	}
	val, ok := cache.Get(key)
	if !ok {
		log.Print("entry not found.")
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("entry not found."))
		return
	}
	b, err := val.JSON()
	if err != nil {
		panic(err)
	}
	rw.Write(b)
}

func cachePutHandler(rw http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len(apiCacheURL):]
	if key == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("key is empty."))
		log.Print("key is empty.")
		return
	}

	entry, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	v := sacache.NewCacheItem(0, time.Duration(0))
	json.Unmarshal(entry, v)
	cache.Set(key, v)
	log.Printf("stored \"%s\" in cache.", key)
	rw.WriteHeader(http.StatusCreated)
}

func cacheRequestHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			cacheGetHandler(rw, r)
		case http.MethodPut:
			cachePutHandler(rw, r)
		}
	})
}

func main() {
	flag.Parse()

	fmt.Printf("SaCache HTTP Server v%s\n", version)

	var logger *log.Logger

	if logfile == "" {
		logger = log.New(os.Stdout, "", log.LstdFlags)

	} else {
		f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		logger = log.New(f, "", log.LstdFlags)
	}

	cache = sacache.NewSaCache("ServerCacheService")

	logger.Print("cache initialized successfully.")

	// handle requests.
	http.Handle(apiCacheURL, cacheRequestHandler())

	logger.Printf("starting server on :%d", port)
	strPort := ":" + strconv.Itoa(port)
	log.Fatal("ListenAndServe: ", http.ListenAndServe(strPort, nil))
}
