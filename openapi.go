// Copyright 2019 Philip Lombardi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"log"
	"net/http"
)

var openapiDocument = `---
openapi: 3.0.0
info:
  title: Obscure Quote Service API
  description: The Obscure Quote Service API
  version: dev

servers:
  - url: http://foo/

paths:
  /:
    get:
      summary: Return a randomly selected quote.
      responses:
        '200':
          description: A JSON object with a quote and some additional metadata.
          content:
            application/json:
              schema: 
                type: object
                properties: 
                  server:
                    type: string
                  quote:
                    type: string
                  time: 
                    type: string
`

func (s *Server) GetOpenAPIDocument(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/yaml")

	if _, err := w.Write([]byte(openapiDocument)); err != nil {
		log.Panicln(err)
	}
}
