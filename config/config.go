package config

import "os"

var BaseUrl = os.Getenv("BASE_URL")
var GenerateUrl = BaseUrl + "generate-data"
var SubmitUrl = BaseUrl + "submit-solution"
var TokenCodeNation = os.Getenv("TOKEN_CODENATION")
