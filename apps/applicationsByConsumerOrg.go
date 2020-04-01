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
  appRequest.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJkODVkNzM3MS01NDhiLTQ2Y2UtYjdkOC01MWQ5NjUwOTY2YWMiLCJuYW1lc3BhY2UiOiI2ODgwOTczYi0yZGY1LTRlYTItODMzMC01OTM4NjYxYTk3NmM6ZjBiMTM5MzItYTU3OS00NGZiLWFlMTEtMGJhZWY4MmY2YjQ4OmZlMGNjY2RkLTZhZWQtNGM2ZC1hNzYzLWEwNjdjNjQwOTNiOSIsImF1ZCI6Ii9hcGkvY2xvdWQvcmVnaXN0cmF0aW9ucy8xNzMzYThhNi0xYjY0LTRiYWMtOTJhNS1iNDg5OTU0MjU0NDIiLCJzdWIiOiIvYXBpL3VzZXItcmVnaXN0cmllcy82ODgwOTczYi0yZGY1LTRlYTItODMzMC01OTM4NjYxYTk3NmMvZjBiMTM5MzItYTU3OS00NGZiLWFlMTEtMGJhZWY4MmY2YjQ4L3VzZXJzL2ZlMGNjY2RkLTZhZWQtNGM2ZC1hNzYzLWEwNjdjNjQwOTNiOSIsImlzcyI6IklCTSBBUEkgQ29ubmVjdCIsImV4cCI6MTU4NTcwOTc5MCwiaWF0IjoxNTg1NjgwOTkwLCJncmFudF90eXBlIjoicGFzc3dvcmQiLCJ1c2VyX3JlZ2lzdHJ5X3VybCI6Ii9hcGkvdXNlci1yZWdpc3RyaWVzLzY4ODA5NzNiLTJkZjUtNGVhMi04MzMwLTU5Mzg2NjFhOTc2Yy9mMGIxMzkzMi1hNTc5LTQ0ZmItYWUxMS0wYmFlZjgyZjZiNDgiLCJyZWFsbSI6InByb3ZpZGVyL2RlZmF1bHQtaWRwLTIiLCJ1c2VybmFtZSI6InN0ZXZlIiwiaWRfdG9rZW4iOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKbWFYSnpkRjl1WVcxbElqb2lVM1JsZG1VaUxDSnNZWE4wWDI1aGJXVWlPaUpQZDI1bGNpSXNJblZ6WlhKZmFXUWlPaUl4TmpSalptWXdaaTAwTnpWa0xUUXpOelF0T1RRMll5MWpNbUZsWm1GbU1ERTROek1pTENKMWMyVnlibUZ0WlNJNkluTjBaWFpsSWl3aWFXRjBJam94TlRnMU5qZ3dPVGt3ZlEucFBPMU1JbVMzR2U2ZmZnOWZCREJybTloWVBVQ0xlMm1qTFp6MEhnWkxrRSIsInNjb3BlcyI6WyJjbG91ZDp2aWV3IiwiY2xvdWQ6bWFuYWdlIiwicHJvdmlkZXItb3JnOnZpZXciLCJwcm92aWRlci1vcmc6bWFuYWdlIiwib3JnOnZpZXciLCJvcmc6bWFuYWdlIiwiZHJhZnRzOnZpZXciLCJkcmFmdHM6ZWRpdCIsImNoaWxkOnZpZXciLCJjaGlsZDpjcmVhdGUiLCJjaGlsZDptYW5hZ2UiLCJwcm9kdWN0OnZpZXciLCJwcm9kdWN0OnN0YWdlIiwicHJvZHVjdDptYW5hZ2UiLCJhcHByb3ZhbDp2aWV3IiwiYXBwcm92YWw6bWFuYWdlIiwiYXBpLWFuYWx5dGljczp2aWV3IiwiYXBpLWFuYWx5dGljczptYW5hZ2UiLCJjb25zdW1lci1vcmc6dmlldyIsImNvbnN1bWVyLW9yZzptYW5hZ2UiLCJhcHA6dmlldzphbGwiLCJhcHA6bWFuYWdlOmFsbCIsIm15OnZpZXciLCJteTptYW5hZ2UiLCJ3ZWJob29rOnZpZXciXX0.KI0IxbYzY9u2ywnEB0lIRFJsQvdiPtMhguYsWNjQq-o")
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
