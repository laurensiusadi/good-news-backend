package server

import (
	"github.com/laurensiusadi/good-news-backend/utils"
)

// Init <function>
// is used to initialize server and all the corresponding services such as DB, Utils, Workers
func Init() {
	// utils
	utils.InitEnvVars()

	r := NewRouter()
	r.Run(":6969")
}