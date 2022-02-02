We want you to implement a server to begin and finish rides. The server will have 2 endpoints: Begin Ride and Finish Ride.

Request and response should be in JSON format, as this is what we are using right now.

### Endpoints

- POST `/rides`. This endpoint will receive a body request with a user ID (string) and a vehicle ID (string).
It will return a new ride with a unique ID (string) as a response. You can assume that the user and vehicle exist.
- POST `/rides/{id}/finish`. This endpoint must end a ride with the given ID. It must calculate the cost using the
  formula specified below and return the ride as a response.

### Pricing Formula

The formula to calculate the cost of a ride is the following:
> cost = unlock_price + minutes*minute_price
 
`minutes` must be treated as an integer and is always rounded up (61 seconds = 2 minutes)
 
### Pricing
 
The current pricing is the following:
 
- unlock_price: 100 cents
- minute_price: 18 cents
 
### Clarifications  

For us, a `ride` is a trip of a user and a vehicle.

Implement any consideration you think that is important to prevent starting or finishing a ride. 
  
We expect a working demo. We are more interested in how you approach the problem and the architecture rather than in the completeness of your solution. 
 
Feel free to make any change to the provided code or use your own.

Don't hesitate asking us if you have any question.
 
We like testing.

