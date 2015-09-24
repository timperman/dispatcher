package network

import (
	"log"
	"net/http"

	"github.com/timperman/dispatcher/util"
)

type Driver struct{}

func New() (*Driver, error) {
	return &Driver{}, nil
}

func (d *Driver) CreateNetwork(w http.ResponseWriter, r *http.Request) {
	req, err := util.JSONDecode(r)
	log.Printf("CreateNetwork request: %v\n", req)
	util.JSONResponse(w, map[string]interface{}{"Err": err})
}

func (d *Driver) DeleteNetwork(w http.ResponseWriter, r *http.Request) {
	req, err := util.JSONDecode(r)
	log.Printf("DeleteNetwork request: %v\n", req)
	util.JSONResponse(w, map[string]interface{}{"Err": err})
}

func (d *Driver) CreateEndpoint(w http.ResponseWriter, r *http.Request) {
	req, err := util.JSONDecode(r)
	log.Printf("CreateEndpoint request: %v\n", req)
	util.JSONResponse(w, map[string]interface{}{"Err": err})
}

func (d *Driver) DeleteEndpoint(w http.ResponseWriter, r *http.Request) {
	req, err := util.JSONDecode(r)
	log.Printf("DeleteEndpoint request: %v\n", req)
	util.JSONResponse(w, map[string]interface{}{"Err": err})
}

func (d *Driver) EndpointOperInfo(w http.ResponseWriter, r *http.Request) {
	req, _ := util.JSONDecode(r)
	log.Printf("EndpointOperInfo request: %v\n", req)
	util.JSONResponse(w, map[string]interface{}{"value": map[string]string{"driver": "dispatcher"}})
}

func (d *Driver) Join(w http.ResponseWriter, r *http.Request) {
	req, err := util.JSONDecode(r)
	log.Printf("Join request: %v\n", req)
	util.JSONResponse(w, map[string]interface{}{"Err": err})
}

func (d *Driver) Leave(w http.ResponseWriter, r *http.Request) {
	req, err := util.JSONDecode(r)
	log.Printf("Leave request: %v\n", req)
	util.JSONResponse(w, map[string]interface{}{"Err": err})
}
