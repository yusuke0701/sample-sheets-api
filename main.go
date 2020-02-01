package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iam/v1"
	"google.golang.org/api/sheets/v4"
)

const apiVersion = "v1"

var (
	// DatastoreClient は、データストアへの接続を担保する
	DatastoreClient *datastore.Client

	// IAMService は、IAMサービスへの接続を担保する
	IAMService *iam.Service

	// SheetsService は、スプレッドシートAPIへの接続を担保する
	SheetsService *sheets.Service
)

func main() {
	// connection
	{
		ctx := context.Background()

		projectID, err := metadata.ProjectID()
		if err != nil {
			log.Fatalf("Failed to connect datastore: %v", err)
		}

		DatastoreClient, err = datastore.NewClient(ctx, projectID)
		if err != nil {
			log.Fatalf("Failed to connect datastore: %v", err)
		}

		IAMService, err = iam.NewService(ctx)
		if err != nil {
			log.Fatalf("Failed to connect iamService: %v", err)
		}

		SheetsService, err = sheets.NewService(ctx)
		if err != nil {
			log.Fatalf("Failed to connect sheetsService: %v", err)
		}
	}
	// routing
	{
		router := gin.Default()
		{
			api := router.Group(fmt.Sprintf("/apitest/%s", apiVersion))
			api.POST("/sheets", sheetsHandler)
		}
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
	}
}
