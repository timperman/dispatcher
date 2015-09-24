package plugin

import (
	"encoding/json"
	"github.com/timperman/dispatcher/network"
	"log"
	"net/http"
)

type Handshake struct {
	Implements []string
}

type Capabilities struct {
	Scope string
}

func Start(addr string) {
	http.HandleFunc("/Plugin.Activate", activate)
	http.HandleFunc("/NetworkDriver.GetCapabilities", getCapabilities)

	n, _ := network.New()

	m := map[string]map[string]func(http.ResponseWriter, *http.Request){
		"POST": {
			"/NetworkDriver.CreateNetwork":    n.CreateNetwork,
			"/NetworkDriver.DeleteNetwork":    n.DeleteNetwork,
			"/NetworkDriver.CreateEndpoint":   n.CreateEndpoint,
			"/NetworkDriver.DeleteEndpoint":   n.DeleteEndpoint,
			"/NetworkDriver.EndpointOperInfo": n.EndpointOperInfo,
			"/NetworkDriver.Join":             n.Join,
			"/NetworkDriver.Leave":            n.Leave,
		},
	}

	for method, routes := range m {
		for route, f := range routes {
			http.HandleFunc(route, handleFuncByMethod(method, f))
		}
	}

	log.Fatal(http.ListenAndServe(addr, nil))
}

func activate(w http.ResponseWriter, r *http.Request) {
	log.Println("Activate call")
	if b, err := json.Marshal(&Handshake{Implements: []string{"NetworkDriver"}}); err == nil {
		w.Write(b)
	}
}

func getCapabilities(w http.ResponseWriter, r *http.Request) {
	log.Println("GetCapabilities call")
	if b, err := json.Marshal(&Capabilities{Scope: "local"}); err == nil {
		w.Write(b)
	}
}
