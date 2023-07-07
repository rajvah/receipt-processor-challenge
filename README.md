# receipt-processor-challenge

To Run the application:

Step 1: go build -o <application name>

Step 2: ./<application name>   # for Unix users #

To Test the application
Using Postman

Process a receipt
Open a workspace in Postman
Select the request type as "POST"
Put the URL as http://localhost:8080/receipts/process
Click on "Body"
Clik on "raw"
Select "JSON" from the dropdown
Paste this in the text field { "retailer": "Target", "purchaseDate": "2022-01-02", "purchaseTime": "13:13", "total": "1.25", "items": [ {"shortDescription": "Pepsi - 12-oz", "price": "1.25"} ] }
Hit "Send"
Example Response: { "id": "926314cb-e6e2-442e-954c-eb461907210b" }
postman_post_request

Get receipt points
Add a tab for another request in the postman workspace
Select the request type as "GET"
Put the URL as http://localhost:8080/receipts/926314cb-e6e2-442e-954c-eb461907210b/points
Hit "Send"
Example Response: { "points": 31 }


Using Unit tests

CMD: go test

Note: There are more test cases to be added, due to time contraint I have not covered all the cases. 