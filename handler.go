package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/sheets/v4"
)

func sheetsHandler(g *gin.Context) {
	// token := google.AppEngineTokenSource(g.Request.Context())

	s := &sheets.Spreadsheet{Properties:&sheets.SpreadsheetProperties{Title: "dou"}}

	sh, err := SheetsService.Spreadsheets.Create(s).Do()
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}
	g.JSON(http.StatusOK, sh)
}
