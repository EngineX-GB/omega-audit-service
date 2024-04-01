package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"omega-audit-service/dao"
	"omega-audit-service/models"
)

func AddAuditEntryHandler(res http.ResponseWriter, req *http.Request) {
	var entry models.Entry
	err := json.NewDecoder(req.Body).Decode(&entry)
	if err != nil {
		log.Fatal("Error in processing the request")
	}
	result := dao.AddEntry(entry)
	if result == true {
		handleResponse(res, 200, "SUCCESS", "Added entry for "+entry.SourceUrl)
	} else {
		handleResponse(res, 500, "ERROR", "Error in adding entry for "+entry.SourceUrl)
	}
}

func FindAuditEntryHandler(res http.ResponseWriter, req *http.Request) {
	param := req.URL.Query().Get("sourceUrl")
	if param == "" {
		handleResponse(res, 400, "ERROR", "Query Parameter 'sourceUrl' is empty")
	}
	entries, err := dao.GetEntriesBySourceUrl(param)
	if err != nil {
		handleResponse(res, 500, "ERROR", err.Error())
	}
	handleEntriesResponse(res, entries)
}

func handleEntriesResponse(resp http.ResponseWriter, entries []models.Entry) {
	resp.WriteHeader(200)
	jsonResp, err := json.Marshal(entries)
	if err != nil {
		handleResponse(resp, 500, "ERROR", "Error in marshalling the response for listing entries")
	} else {
		resp.Write(jsonResp)
	}

}

func handleResponse(resp http.ResponseWriter, statusCode int, status string, message string) {
	resp.WriteHeader(statusCode)
	jsonContent := make(map[string]string)
	jsonContent["status"] = status
	jsonContent["message"] = message
	jsonResp, err := json.Marshal(jsonContent)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		jsonContent["status"] = "ERROR"
		jsonContent["message"] = http.StatusText(http.StatusBadRequest)
		errResponse, err := json.Marshal(jsonContent)
		if err != nil {
			log.Fatal("Error happened on marshalling JSON %s", err)
		}
		resp.Write(errResponse)
		return
	}
	resp.Write(jsonResp)
}
