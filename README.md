# MerchApi
RESTful Api in Go using Gin framework

Data to be added on Post
{"name": "Jubilee Cross", "price": 20, "type": "Pendant"}

Data to be added on Post addmerchtodb
{"name": "Eagle", "price": 499, "type": "Rubber", "size": "S", "quantity": 2}

Bug:
1) addmerchtodb endpoint shows success with status 200 but data not reflected in db
    Plan:
    - Debug the API
    Fixed:
    - PostMerchtoDb function was not being called in the POST route
    - wrong error check (!= instead of ==)
    - err was not assigned
