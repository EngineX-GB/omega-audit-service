package main

import (
	"fmt"
	"log"
	"omega-audit-service/config"
	"omega-audit-service/handlers"
)
import "net/http"

func main() {
	fmt.Println("Started Omega Audit Service " + config.GetAppVersion())
	http.HandleFunc("/api/v1/audit/add", handlers.AddAuditEntryHandler)
	http.HandleFunc("/api/v1/audit/", handlers.FindAuditEntryHandler)
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		log.Fatal("Cannot run this service : " + err.Error())
	}
}
