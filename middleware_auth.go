package main

import (
	"net/http"

	"github.com/keshavsharma54126/rssagg/internal/database"
)

type authdHandler func(http.ResponseWriter,*http.Request,database.User)

