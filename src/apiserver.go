package main

import (
	"logger"
	"data"
	"handler"
	"net/http"
	"fmt"
	"encoding/json"
	)
func handleRequest(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")

	apiRequest := &data.ApiRequest{}
	err := json.Unmarshal([]byte(query),apiRequest)
	if(nil != err){
		errMsg := fmt.Sprintf(err.Error())
		w.Write([]byte(errMsg))
		return
	}
	resp := handler.HandleApiRequest(apiRequest)

	w.Write([]byte(resp))
}

func main() {
	logConfig := &logger.LoggerConfig{}
	logger.Init(logConfig)
	
	http.HandleFunc("/apiserver", handleRequest)
	http.ListenAndServe(":8000", nil)

	logger.Debug("Api server started successfully.")
}