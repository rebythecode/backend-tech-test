# Senior Backend Technical Challenge
# Doing a ride

We want you to implement a server to begin and finish rides. The server will have 3 endpoints: Begin Ride and Finish Ride and Metrics

Request and response should be in JSON format, as this is what we are using right now.

### Endpoints

- POST `/rides`. This endpoint will receive a body request with a user ID (string) and a vehicle ID (string).
It will return a new ride with a unique ID (string) as a response. You can assume that the user and vehicle exist.
- POST `/rides/{id}/finish`. This endpoint must end a ride with the given ID. It must calculate the cost using the formula specified below and return the ride entity as a response.
- GET `/metrics` returns a counter with the number of times each endpoint has been called, successful and failed calls and the the minimum, average and maximum delivery time of each endpoint

### Pricing Formula

The formula to calculate the cost of a ride is the following:
> cost = unlock_price + minutes*minute_price

`minutes` must be treated as an integer and is always rounded up (61 seconds = 2 minutes)
 
### Pricing
 
The current pricing is the following:
 
- unlock_price: 100 cents of €
- minute_price: 18 cents of €

### Metrics
We expect a json with enough information to see how many times each endpoint has been called. We also want to see the number of times we are returning and error response, and we want to be able of understand how the system is performing in terms of timings.
 
### Considerations  

- For us, a `ride` is a trip of a user and a vehicle.
- Implement any consideration you think that is important to prevent starting or finishing a ride.
- We are interested in how you approach the problem and the architecture rather than in the completeness of your solution. We expect a working demo written in Go, but you can mock infrastructure like databases, etc... For instance, using in memory storages or files. It is up to you.  
- Feel free to make any change to the provided code or use your own.
- Don't hesitate asking us if you have any question.
- We love testing.

### Additional notes
- The provided solutions needs to be uploaded into a public repository (Github, Gitlab, bitbucket, etc...) with a README.MD providing the following information.
  - Instructions on how to run your solution
  - Requirements
- Please make sure the name Reby are not referenced in any place in your code.
- Commit from the very beginning and commit often. We value the possibility to review your git log.

