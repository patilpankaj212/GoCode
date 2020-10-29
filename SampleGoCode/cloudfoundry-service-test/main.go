package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
)

//VCapServiceCredentials will hold a single service instance credential
type VCapServiceCredentials struct {
	uri              string
	user             string
	password         string
	isServicePresent bool
}

const (
	//TciServiceBrokerName is the service broker name of tci-bridge
	TciServiceBrokerName = "TCI REST Endpoint"
)

var (
	vCapServiceCreds *VCapServiceCredentials
)

func main() {
	Init()
	RegisterControllers()

	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "8000"
	}
	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, nil)
}

//Init function will initialize the vcapservicecreds
func Init() {
	vCapServiceCreds = &VCapServiceCredentials{}
	appEnv, err := cfenv.Current()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("reading VCAP_SERVICES")
	if len(appEnv.Services) == 0 {
		fmt.Println("No Services Bind to the application")
	} else {
		service, ok := appEnv.Services["TCI REST Endpoint"]
		if !ok {
			fmt.Println("No tci-bridge services bind to the application")
			return
		}

		vCapServiceCreds.isServicePresent = true
		vCapServiceCreds.user = service[0].Credentials["username"].(string)
		vCapServiceCreds.password = service[0].Credentials["password"].(string)
		vCapServiceCreds.uri = service[0].Credentials["uri"].(string)
	}
}

//RegisterControllers method registers the controllers
func RegisterControllers() {

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/callService", serviceHandler)
	http.HandleFunc("/healthCheck", healthCheckHandler)
	http.HandleFunc("/callService/", serviceHandler)
	http.HandleFunc("/getDetails", detailsHandler)
}

func rootHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("root handler called")
	appInformationMap := make(map[string]string)
	appInformationMap["App Name"] = "Dummy TCI-BRIDGE app"
	appInformationMap["App Version"] = "1.0"

	byteArr, err := json.Marshal(appInformationMap)
	if err != nil {
		response.WriteHeader(http.StatusFailedDependency)
		return
	}
	response.Write(byteArr)
}

func serviceHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("service handler called")
	if vCapServiceCreds.isServicePresent {
		queryParams := request.URL.Query()
		if len(queryParams) == 0 {
			errorMap := make(map[string]string)
			errorMap["error"] = "true"
			errorMap["errorMessage"] = "Resource not provided"
			response.WriteHeader(http.StatusPreconditionRequired)
			byteValue, _ := json.Marshal(errorMap)
			response.Write(byteValue)
			return
		}

		resource := queryParams.Get("resource")

		client := &http.Client{}
		request, err := http.NewRequest("GET", vCapServiceCreds.uri+"/"+resource, nil)

		resp, err := client.Do(request)
		if err != nil {
			errorMap := make(map[string]string)
			errorMap["error"] = "true"
			errorMap["errorMessage"] = err.Error()
			byteValue, _ := json.Marshal(errorMap)
			response.WriteHeader(resp.StatusCode)
			response.Write(byteValue)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		response.WriteHeader(resp.StatusCode)
		response.Write(body)

	} else {
		errorMap := make(map[string]string)
		errorMap["error"] = "true"
		errorMap["errorMessage"] = "No TCI-BRIDGE service instance bind with the application"
		response.WriteHeader(http.StatusPreconditionFailed)
		byteValue, _ := json.Marshal(errorMap)
		response.Write(byteValue)
	}
}

func healthCheckHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("healthcheck handler called")
	response.WriteHeader(http.StatusOK)
}

func detailsHandler(response http.ResponseWriter, request *http.Request) {
	if vCapServiceCreds.isServicePresent {
		detailsMap := make(map[string]string)
		detailsMap["uri"] = vCapServiceCreds.uri
		detailsMap["user"] = vCapServiceCreds.user
		detailsMap["password"] = vCapServiceCreds.password

		byteValue, _ := json.Marshal(detailsMap)
		response.Write(byteValue)

	} else {
		errorMap := make(map[string]string)
		errorMap["error"] = "true"
		errorMap["errorMessage"] = "No TCI-BRIDGE service instance bind with the application"
		byteValue, _ := json.Marshal(errorMap)
		response.Write(byteValue)
	}
}
