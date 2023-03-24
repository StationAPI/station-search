package db

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Website struct {
	Name        string   `json:"name"`
	Id          string   `json:"id"`
	IconURL     string   `json:"icon_url"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func SearchWebsites(query string, db gorm.DB) []Website {
	var websites []Website

	db.Limit(20).Find(&websites, "name LIKE ?", fmt.Sprintf("%"+strings.ToLower(query)+"%"))

	return websites
}
