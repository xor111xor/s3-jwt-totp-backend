/*
Copyright © 2024 xor111xor
*/
// @title           Server storage API
// @version         1.0
// @description     REST API server of storage with auth

// @contact.name   API Support
// @contact.url    https://github.com/xor111xor/s3-jwt-totp-backend
// @contact.email  xor111xor@hotmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
package main

import (
	"github.com/xor111xor/s3-jwt-totp-backend/cmd"
)

func main() {
	cmd.Execute()
}
