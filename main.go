/*******************************************************************************
*
* Copyright 2023 SAP SE
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You should have received a copy of the License along with this
* program. If not, you may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
*******************************************************************************/

package main

import (
	"net/http"
	"log"
	"os"

	"github.com/sapcc/alertflow/pkg/server"
	"github.com/sapcc/alertflow/pkg/handlers"
	"github.com/sapcc/alertflow/pkg/clients"
)

func main() {
	/////////////////
	// set configs //
	/////////////////
	// TODO: use automatic config parsing
	endpoint := os.Getenv("BILLING_ENDPOINT")
	if endpoint == "" {
		log.Fatal("Missing environment variable BILLING_ENDPOINT")
	}
	password := os.Getenv("BILLING_AUTH")
	if password == "" {
		log.Fatal("Missing environment variable BILLING_AUTH")
	}

	////////////////////
	//  setup clients //
	////////////////////
	billingClient, err := clients.NewBillingClient(os.Getenv("BILLING_ENDPOINT"), os.Getenv("BILLING_AUTH"))
	if err != nil {
		log.Fatal("Failed to create billing client: %s", err)
	}

	projectClient := clients.NewProjectClient(billingClient)

	////////////////////////////////////////
	// setup handlers and starting server //
	////////////////////////////////////////
	alertHandler := handlers.AlertWebHookHandler(projectClient)

	router := http.NewServeMux()
	router.Handle("/alerts", alertHandler)

	server := &http.Server{
      Addr:    ":8080",
      Handler: server.WrapHandler(router),
  }
	log.Fatal(server.ListenAndServe())
}
