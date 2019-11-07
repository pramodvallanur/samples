# Transfer Org Owner
The change can be initiated only by the owner of the Org to either to an associate or a member in the org. 

## Steps for transfering ownership to a org member

1] Login as the owner of the org to initiate the transfer

```
apic login --server apicserver
Enter your API Connect credentials
Realm? admin/default-idp-1
Username? admin
Password? *****
Logged into apicserver successfully
``` 

2] Get list of members at the org
```
apic members:list --scope org --org admin --server apicserver
admin    [state: enabled]   https://apicserver/api/orgs/5f9fba35-a5d9-46ea-ae57-6c1d7324133c/members/5e2d1d36-70fd-4334-879d-df6019d69ed3   
will    [state: enabled]   https://apicserver/api/orgs/5f9fba35-a5d9-46ea-ae57-6c1d7324133c/members/f6e80fb0-eed7-419b-a682-365a904d18eb
```  

`admin` is the current owner of the org and is wanting to transfer the ownership to `will` 

3] create a json file (in my case: transferOwner.json) with the jason's member url
```
{
    "new_owner_member_url": "https://apicserver/api/orgs/5f9fba35-a5d9-46ea-ae57-6c1d7324133c/members/f6e80fb0-eed7-419b-a682-365a904d18eb",
}
```  

4] Initiate the transfer
```
apic org:transfer-owner --server apicserver acme transferOwner.json
acme    [state: enabled]   https://apicserver/api/orgs/5f9fba35-a5d9-46ea-ae57-6c1d7324133c
```  

You have now sucessfully transfered the ownership

## Steps for transfering ownership to an associate in the rog

1] Login as the owner of the org to initiate the transfer

```
apic login --server apicserver
Enter your API Connect credentials
Realm? admin/default-idp-1
Username? admin
Password? *****
Logged into apicserver successfully
``` 

2] Get list of associates at the org
```
apic associates:list --scope org --org admin --server apicserver
admin   https://apicserver/api/orgs/5f9fba35-a5d9-46ea-ae57-6c1d7324133c/associates/69fb9b7c-7071-42a2-b76d-bf48f28cb04c   
will   https://apicserver/api/orgs/5f9fba35-a5d9-46ea-ae57-6c1d7324133c/associates/7158d40d-3983-427d-8977-294b82d6c8d8
```  

`admin` is the current owner of the org and is wanting to transfer the ownership to `will` 

3] create a json file (in my case: transferOwner.json) with the jason's associate url
```
{
    "new_owner_associate_url": "https://apicserver/api/orgs/5f9fba35-a5d9-46ea-ae57-6c1d7324133c/associates/7158d40d-3983-427d-8977-294b82d6c8d8"
}
```  

4] Initiate the transfer
```
apic org:transfer-owner --server apicserver acme transferOwner.json
acme   [state: enabled]   https://apicserver/api/orgs/5f9fba35-a5d9-46ea-ae57-6c1d7324133c
```  

You have now sucessfully transfered the ownership
