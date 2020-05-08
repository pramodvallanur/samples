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
  fmt.Println("Please enter the product version.")
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
    subscriptionRequest.Header.Set("Authorization", "Bearer token redacted for security reasons. To get the bearer token use token login")
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
    //Get the consumer object
    consumerOrgRequest, err := http.NewRequest("GET",subscriptionObj.ConsumerOrgURL, nil)      
    consumerOrgRequest.Header.Set("Content-type","application/json")
    //TODO: Replace the bearer token
    consumerOrgRequest.Header.Set("Authorization", "Bearer token redacted for security reasons. To get the bearer token use token login")
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
    appRequest.Header.Set("Authorization", "Bearer Bearer token redacted for security reasons. To get the bearer token use token login")
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
