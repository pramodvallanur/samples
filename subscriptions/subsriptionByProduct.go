package main

import (
  "net/http"  
  "log"  
  "encoding/json"
  "fmt"
  "crypto/tls"  
)

type product struct {  
  SubscriptionUrls []string  `json:"subscription_urls"`
}
type subscription struct {	
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Title          string    `json:"title"`	
	Plan           string    `json:"plan"`
	AppURL         string    `json:"app_url,omitempty"`
	ConsumerOrgURL string    `json:"consumer_org_url,omitempty"`
	ConsumerOrgTitle string    `json:"consumerorgtitle"`
	AppTitle	       string    `json:"apptitle"`
}

type consumerOrg struct {
	Title      string    `json:"title"`	
}


type app struct {		
	Title             string    `json:"title"`	
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
  fmt.Println("Please enter the product name.")
  var productName string
  fmt.Scanln(&productName)   
  
  //Get the Product version as input
  fmt.Println("Please enter the product name.")
  var productVer string
  fmt.Scanln(&productVer)  

  //Construct the product url
  productUrl := fmt.Sprintf("https://%s/api/catalogs/%s/%s/products/%s/%s?fields=add(subscription_urls)",serverName,pOrgName,catalogName,productName,productVer)

  //Set to ignore certs
  http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} 

  //Make the API call
  productRequest, err := http.NewRequest("GET",productUrl, nil)    
  productRequest.Header.Set("Content-type","application/json")
  //TODO: Replace the bearer token
  productRequest.Header.Set("Authorization", "Bearer token")
  if err != nil{
    log.Fatalln(err)
  }
  client := http.Client{   
  }
  productResp, err := client.Do(productRequest)
  defer productResp.Body.Close()  
  if err != nil{
    log.Fatalln(err)
  }
  //Parse the response
  var productObj product
  json.NewDecoder(productResp.Body).Decode(&productObj)  
  
  subscriptionCount := len(productObj.SubscriptionUrls)
  fmt.Printf("There are total of %d subscription(s) for %s : %s. They are:",subscriptionCount,productName,productVer)
  fmt.Println("")
  //For each subscription get the subscription object
  for _, SubscriptionUrl := range productObj.SubscriptionUrls{
    subscriptionRequest, err := http.NewRequest("GET",SubscriptionUrl, nil)      
	subscriptionRequest.Header.Set("Content-type","application/json")
	//TODO: Replace the bearer token
    subscriptionRequest.Header.Set("Authorization", "Bearer token")
    if err != nil{
      log.Fatalln(err)
    }
    subscriptionResp, err := client.Do(subscriptionRequest)
    defer subscriptionResp.Body.Close()
    if err != nil{
      log.Fatalln(err)
	}
	//Parse the response
    var subscriptionObj subscription
	json.NewDecoder(subscriptionResp.Body).Decode(&subscriptionObj)  	

	//Get the application object
	consumerOrgRequest, err := http.NewRequest("GET",subscriptionObj.ConsumerOrgURL, nil)      
	consumerOrgRequest.Header.Set("Content-type","application/json")
	//TODO: Replace the bearer token
    consumerOrgRequest.Header.Set("Authorization", "Bearer token")
    if err != nil{
      log.Fatalln(err)
    }
	consumerOrgResp, err := client.Do(consumerOrgRequest)
    defer consumerOrgResp.Body.Close()
    if err != nil{
      log.Fatalln(err)
	}
	var consumerOrgObj consumerOrg
	json.NewDecoder(consumerOrgResp.Body).Decode(&consumerOrgObj)  
	subscriptionObj.ConsumerOrgTitle = consumerOrgObj.Title
	subscriptionObj.ConsumerOrgURL = ""

	//Get the application object
	appRequest, err := http.NewRequest("GET",subscriptionObj.AppURL, nil)      
    appRequest.Header.Set("Content-type","application/json")
    appRequest.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJkODVkNzM3MS01NDhiLTQ2Y2UtYjdkOC01MWQ5NjUwOTY2YWMiLCJuYW1lc3BhY2UiOiI2ODgwOTczYi0yZGY1LTRlYTItODMzMC01OTM4NjYxYTk3NmM6ZjBiMTM5MzItYTU3OS00NGZiLWFlMTEtMGJhZWY4MmY2YjQ4OmZlMGNjY2RkLTZhZWQtNGM2ZC1hNzYzLWEwNjdjNjQwOTNiOSIsImF1ZCI6Ii9hcGkvY2xvdWQvcmVnaXN0cmF0aW9ucy8xNzMzYThhNi0xYjY0LTRiYWMtOTJhNS1iNDg5OTU0MjU0NDIiLCJzdWIiOiIvYXBpL3VzZXItcmVnaXN0cmllcy82ODgwOTczYi0yZGY1LTRlYTItODMzMC01OTM4NjYxYTk3NmMvZjBiMTM5MzItYTU3OS00NGZiLWFlMTEtMGJhZWY4MmY2YjQ4L3VzZXJzL2ZlMGNjY2RkLTZhZWQtNGM2ZC1hNzYzLWEwNjdjNjQwOTNiOSIsImlzcyI6IklCTSBBUEkgQ29ubmVjdCIsImV4cCI6MTU4NTcwOTc5MCwiaWF0IjoxNTg1NjgwOTkwLCJncmFudF90eXBlIjoicGFzc3dvcmQiLCJ1c2VyX3JlZ2lzdHJ5X3VybCI6Ii9hcGkvdXNlci1yZWdpc3RyaWVzLzY4ODA5NzNiLTJkZjUtNGVhMi04MzMwLTU5Mzg2NjFhOTc2Yy9mMGIxMzkzMi1hNTc5LTQ0ZmItYWUxMS0wYmFlZjgyZjZiNDgiLCJyZWFsbSI6InByb3ZpZGVyL2RlZmF1bHQtaWRwLTIiLCJ1c2VybmFtZSI6InN0ZXZlIiwiaWRfdG9rZW4iOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKbWFYSnpkRjl1WVcxbElqb2lVM1JsZG1VaUxDSnNZWE4wWDI1aGJXVWlPaUpQZDI1bGNpSXNJblZ6WlhKZmFXUWlPaUl4TmpSalptWXdaaTAwTnpWa0xUUXpOelF0T1RRMll5MWpNbUZsWm1GbU1ERTROek1pTENKMWMyVnlibUZ0WlNJNkluTjBaWFpsSWl3aWFXRjBJam94TlRnMU5qZ3dPVGt3ZlEucFBPMU1JbVMzR2U2ZmZnOWZCREJybTloWVBVQ0xlMm1qTFp6MEhnWkxrRSIsInNjb3BlcyI6WyJjbG91ZDp2aWV3IiwiY2xvdWQ6bWFuYWdlIiwicHJvdmlkZXItb3JnOnZpZXciLCJwcm92aWRlci1vcmc6bWFuYWdlIiwib3JnOnZpZXciLCJvcmc6bWFuYWdlIiwiZHJhZnRzOnZpZXciLCJkcmFmdHM6ZWRpdCIsImNoaWxkOnZpZXciLCJjaGlsZDpjcmVhdGUiLCJjaGlsZDptYW5hZ2UiLCJwcm9kdWN0OnZpZXciLCJwcm9kdWN0OnN0YWdlIiwicHJvZHVjdDptYW5hZ2UiLCJhcHByb3ZhbDp2aWV3IiwiYXBwcm92YWw6bWFuYWdlIiwiYXBpLWFuYWx5dGljczp2aWV3IiwiYXBpLWFuYWx5dGljczptYW5hZ2UiLCJjb25zdW1lci1vcmc6dmlldyIsImNvbnN1bWVyLW9yZzptYW5hZ2UiLCJhcHA6dmlldzphbGwiLCJhcHA6bWFuYWdlOmFsbCIsIm15OnZpZXciLCJteTptYW5hZ2UiLCJ3ZWJob29rOnZpZXciXX0.KI0IxbYzY9u2ywnEB0lIRFJsQvdiPtMhguYsWNjQq-o")
    if err != nil{
      log.Fatalln(err)
    }
	appResp, err := client.Do(appRequest)
    defer appResp.Body.Close()
    if err != nil{
      log.Fatalln(err)
	}
	var appObj app
	json.NewDecoder(appResp.Body).Decode(&appObj)  
	subscriptionObj.AppTitle = appObj.Title
	subscriptionObj.AppURL = ""
	
	//Prepare the subscription Object
	
    json, err := json.MarshalIndent(subscriptionObj,"", "  ")
    if err != nil{
      log.Fatalln(err)
    }    
    fmt.Println(string([]byte(json)))    
  }  
}
