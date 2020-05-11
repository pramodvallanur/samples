package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type products struct {
	productList []string `json:"product"`
}
type product struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Version string `json:"version"`
}

type api struct {
	ProductUrls []string `json:"product_urls"`
}

func main() {

	//Get the server name as input
	fmt.Println("Please enter the server name.")
	var serverName string
	fmt.Scanln(&serverName)

	//Get the providerOrg Name as input
	fmt.Println("Please enter the providerOrg name.")
	var pOrgName string
	fmt.Scanln(&pOrgName)

	//Get the Catalog name as input
	fmt.Println("Please enter the catalog name.")
	var catalogName string
	fmt.Scanln(&catalogName)

	//Get the Product name as input
	fmt.Println("Please enter the API name.")
	var apiName string
	fmt.Scanln(&apiName)

	//Get the Product version as input
	fmt.Println("Please enter the API version.")
	var apiVer string
	fmt.Scanln(&apiVer)

	//Construct the api url	
	myAPIUrl := fmt.Sprintf("https://%s/api/catalogs/%s/%s/apis/%s/%s?fields=add(product_urls)", serverName, pOrgName, catalogName, apiName, apiVer)
	//fmt.Printf(myAPIUrl)
	//Set to ignore certs
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	//Make the API call
	apiRequest, err := http.NewRequest("GET", myAPIUrl, nil)	
	apiRequest.Header.Set("Content-type", "application/json")
	//TODO: Replace the bearer token
	apiRequest.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJiMDFiYTdiMi00YjQ4LTQ3YjUtYTE3OS0wNDlkMTVlZmVhYTMiLCJuYW1lc3BhY2UiOiJiNmM5NTdhZi04YzBkLTRhZGUtODc0Zi00ZmNmMjUzMjM0YjU6YmJhZDQ3ODQtMTQ4NC00MzQ0LWIyOTUtYWVmYzY5Yjg4ZDcyOjc1YzkxYjZlLTNhYmYtNDU4NS04NDgwLWEyZWVlNTM3ZmI3ZCIsImF1ZCI6Ii9hcGkvY2xvdWQvcmVnaXN0cmF0aW9ucy9lOWNlNGVjZC0wNWI4LTQyZWYtYjJkMi01ZmY3NmJhNmQzNmYiLCJzdWIiOiIvYXBpL3VzZXItcmVnaXN0cmllcy9iNmM5NTdhZi04YzBkLTRhZGUtODc0Zi00ZmNmMjUzMjM0YjUvYmJhZDQ3ODQtMTQ4NC00MzQ0LWIyOTUtYWVmYzY5Yjg4ZDcyL3VzZXJzLzc1YzkxYjZlLTNhYmYtNDU4NS04NDgwLWEyZWVlNTM3ZmI3ZCIsImlzcyI6IklCTSBBUEkgQ29ubmVjdCIsImV4cCI6MTU4ODkyODE2NiwiaWF0IjoxNTg4ODk5MzY2LCJncmFudF90eXBlIjoicGFzc3dvcmQiLCJ1c2VyX3JlZ2lzdHJ5X3VybCI6Ii9hcGkvdXNlci1yZWdpc3RyaWVzL2I2Yzk1N2FmLThjMGQtNGFkZS04NzRmLTRmY2YyNTMyMzRiNS9iYmFkNDc4NC0xNDg0LTQzNDQtYjI5NS1hZWZjNjliODhkNzIiLCJyZWFsbSI6InByb3ZpZGVyL2RlZmF1bHQtaWRwLTIiLCJ1c2VybmFtZSI6InN0ZXZlIiwiaWRfdG9rZW4iOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKbWFYSnpkRjl1WVcxbElqb2lVM1JsZG1VaUxDSnNZWE4wWDI1aGJXVWlPaUpQZDI1bGNpSXNJblZ6WlhKZmFXUWlPaUkwTldJME1XWXpOUzA1TlRRNUxUUmhNVFF0T1RVNE9TMWxZekZtWVRBMk1ESTVPV1lpTENKMWMyVnlibUZ0WlNJNkluTjBaWFpsSWl3aWFXRjBJam94TlRnNE9EazVNelkyZlEuTVEtTk8xMmlWMDh3SGdzUWtuckVWenlVV3B3QVN5aHBsMzkzNHpVUndXNCIsInNjb3BlcyI6WyJjbG91ZDp2aWV3IiwiY2xvdWQ6bWFuYWdlIiwicHJvdmlkZXItb3JnOnZpZXciLCJwcm92aWRlci1vcmc6bWFuYWdlIiwib3JnOnZpZXciLCJvcmc6bWFuYWdlIiwiZHJhZnRzOnZpZXciLCJkcmFmdHM6ZWRpdCIsImNoaWxkOnZpZXciLCJjaGlsZDpjcmVhdGUiLCJjaGlsZDptYW5hZ2UiLCJwcm9kdWN0OnZpZXciLCJwcm9kdWN0OnN0YWdlIiwicHJvZHVjdDptYW5hZ2UiLCJhcHByb3ZhbDp2aWV3IiwiYXBwcm92YWw6bWFuYWdlIiwiYXBpLWFuYWx5dGljczp2aWV3IiwiYXBpLWFuYWx5dGljczptYW5hZ2UiLCJjb25zdW1lci1vcmc6dmlldyIsImNvbnN1bWVyLW9yZzptYW5hZ2UiLCJhcHA6dmlldzphbGwiLCJhcHA6bWFuYWdlOmFsbCIsIm15OnZpZXciLCJteTptYW5hZ2UiLCJ3ZWJob29rOnZpZXciXX0.Pm_O4CAoerBXibheszJlcBo__Fict6_D2MPfQR-t678")
	if err != nil {
		log.Fatalln(err)
	}
	client := http.Client{}
	apiResp, err := client.Do(apiRequest)
	defer apiResp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	//Parse the response
	var apiObj api
	json.NewDecoder(apiResp.Body).Decode(&apiObj)
	for _, productUrl := range apiObj.ProductUrls {
		productRequest, err := http.NewRequest("GET", productUrl, nil)
		productRequest.Header.Set("Content-type", "application/json")
		//TODO: Put in your bearer Token, in the next iteration, will add Login support
		//productRequest.Header.Set("Authorization", "Bearer fyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJiMDFiYTdiMi00YjQ4LTQ3YjUtYTE3OS0wNDlkMTVlZmVhYTMiLCJuYW1lc3BhY2UiOiJiNmM5NTdhZi04YzBkLTRhZGUtODc0Zi00ZmNmMjUzMjM0YjU6YmJhZDQ3ODQtMTQ4NC00MzQ0LWIyOTUtYWVmYzY5Yjg4ZDcyOjc1YzkxYjZlLTNhYmYtNDU4NS04NDgwLWEyZWVlNTM3ZmI3ZCIsImF1ZCI6Ii9hcGkvY2xvdWQvcmVnaXN0cmF0aW9ucy9lOWNlNGVjZC0wNWI4LTQyZWYtYjJkMi01ZmY3NmJhNmQzNmYiLCJzdWIiOiIvYXBpL3VzZXItcmVnaXN0cmllcy9iNmM5NTdhZi04YzBkLTRhZGUtODc0Zi00ZmNmMjUzMjM0YjUvYmJhZDQ3ODQtMTQ4NC00MzQ0LWIyOTUtYWVmYzY5Yjg4ZDcyL3VzZXJzLzc1YzkxYjZlLTNhYmYtNDU4NS04NDgwLWEyZWVlNTM3ZmI3ZCIsImlzcyI6IklCTSBBUEkgQ29ubmVjdCIsImV4cCI6MTU4ODkyODE2NiwiaWF0IjoxNTg4ODk5MzY2LCJncmFudF90eXBlIjoicGFzc3dvcmQiLCJ1c2VyX3JlZ2lzdHJ5X3VybCI6Ii9hcGkvdXNlci1yZWdpc3RyaWVzL2I2Yzk1N2FmLThjMGQtNGFkZS04NzRmLTRmY2YyNTMyMzRiNS9iYmFkNDc4NC0xNDg0LTQzNDQtYjI5NS1hZWZjNjliODhkNzIiLCJyZWFsbSI6InByb3ZpZGVyL2RlZmF1bHQtaWRwLTIiLCJ1c2VybmFtZSI6InN0ZXZlIiwiaWRfdG9rZW4iOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKbWFYSnpkRjl1WVcxbElqb2lVM1JsZG1VaUxDSnNZWE4wWDI1aGJXVWlPaUpQZDI1bGNpSXNJblZ6WlhKZmFXUWlPaUkwTldJME1XWXpOUzA1TlRRNUxUUmhNVFF0T1RVNE9TMWxZekZtWVRBMk1ESTVPV1lpTENKMWMyVnlibUZ0WlNJNkluTjBaWFpsSWl3aWFXRjBJam94TlRnNE9EazVNelkyZlEuTVEtTk8xMmlWMDh3SGdzUWtuckVWenlVV3B3QVN5aHBsMzkzNHpVUndXNCIsInNjb3BlcyI6WyJjbG91ZDp2aWV3IiwiY2xvdWQ6bWFuYWdlIiwicHJvdmlkZXItb3JnOnZpZXciLCJwcm92aWRlci1vcmc6bWFuYWdlIiwib3JnOnZpZXciLCJvcmc6bWFuYWdlIiwiZHJhZnRzOnZpZXciLCJkcmFmdHM6ZWRpdCIsImNoaWxkOnZpZXciLCJjaGlsZDpjcmVhdGUiLCJjaGlsZDptYW5hZ2UiLCJwcm9kdWN0OnZpZXciLCJwcm9kdWN0OnN0YWdlIiwicHJvZHVjdDptYW5hZ2UiLCJhcHByb3ZhbDp2aWV3IiwiYXBwcm92YWw6bWFuYWdlIiwiYXBpLWFuYWx5dGljczp2aWV3IiwiYXBpLWFuYWx5dGljczptYW5hZ2UiLCJjb25zdW1lci1vcmc6dmlldyIsImNvbnN1bWVyLW9yZzptYW5hZ2UiLCJhcHA6dmlldzphbGwiLCJhcHA6bWFuYWdlOmFsbCIsIm15OnZpZXciLCJteTptYW5hZ2UiLCJ3ZWJob29rOnZpZXciXX0.Pm_O4CAoerBXibheszJlcBo__Fict6_D2MPfQR-t678")
		if err != nil {
			log.Fatalln(err)
		}
		productResp, err := client.Do(productRequest)
		defer productResp.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
		//Parse the response
		var productObj product
		json.NewDecoder(productResp.Body).Decode(&productObj)
		json, err := json.MarshalIndent(productObj, "", "  ")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string([]byte(json)))
	}
}
