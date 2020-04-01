package main

import (
  "net/http"  
  "log"  
  "encoding/json"
  "fmt"
  "crypto/tls"
  //"io/ioutil" 
  //"time" 
)

type appList struct {
	TotalResults int `json:"total_results"`
	Results    []apps `json:"results"`
}

type apps struct {
  //Type              string    `json:"type"`
  // APIVersion        string    `json:"api_version"`
  // ID                string    `json:"id"`
  Name              string    `json:"name"`
  Title             string    `json:"title"`
  Summary           string    `json:"summary"`
  State             string    `json:"state"`
  LifecycleState    string    `json:"lifecycle_state"`
  // AppCredentialUrls []string  `json:"app_credential_urls"`
  // ImageEndpoint     string    `json:"image_endpoint"`
  // CreatedAt         time.Time `json:"created_at"`
  // UpdatedAt         time.Time `json:"updated_at"`
  // OrgURL            string    `json:"org_url"`
  // CatalogURL        string    `json:"catalog_url"`
  // ConsumerOrgURL    string    `json:"consumer_org_url"`
  // URL               string    `json:"url"`
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

  //Get the consumerOrg name as input
  fmt.Println("Please enter the Consumer org name.")
  var corgName string
  fmt.Scanln(&corgName)    
 

  //Construct the consumerOrg url
  consumerOrgUrl := fmt.Sprintf("https://%s/api/consumer-orgs/%s/%s/%s/apps",serverName,pOrgName,catalogName,corgName)
  fmt.Println(consumerOrgUrl)
  //Set to ignore certs
  http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} 

  //Make the API call
  //api/consumer-orgs/{org}/{catalog}/{consumer-org}/apps
  appRequest, err := http.NewRequest("GET",consumerOrgUrl, nil)  
  //appRequest, err := http.NewRequest("GET","https://mystack.loki.dev.ciondemand.com/api/consumer-orgs/lob-one/sandbox/andre-owner/apps", nil)  
  appRequest.Header.Set("Content-type","application/json")
  //TODO: Replace the bearer token
  appRequest.Header.Set("Authorization", "Bearer token redacted for security reasons. To get the bearer token use token login")
  if err != nil{
    log.Fatalln(err)
  }
  client := http.Client{   
  }
  appResp, err := client.Do(appRequest)
  defer appResp.Body.Close()  
  if err != nil{
    log.Fatalln(err)
  }

  //Parse the response
  var totalResult appList
  json.NewDecoder(appResp.Body).Decode(&totalResult)
  for _, apps := range totalResult.Results{
    json, err := json.MarshalIndent(apps,"", "  ")
    if err != nil{
      log.Fatalln(err)
    }    
    fmt.Println(string([]byte(json))) 
  } 
  
}
