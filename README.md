# MerchApi
RESTful Api in Go using Gin framework and mysql db

Endpoints: 
    GET
        http://localhost:8080/merchwithquantity
    
    POST
        http://localhost:8080/addmerchtodb
            sample: {"name": "Faith Over Fear", "price": 499, "type": "Rubber", "color": "Red", "size": "XL", "quantity": 1}

To Do:
- Update quantity for the given size
- Get stock by name, type, color and size

Bug:
1) addmerchtodb endpoint shows success with status 200 but data not reflected in db
    Plan:
    - Debug the API
    Fixed:
    - PostMerchtoDb function was not being called in the POST route
    - wrong error check (!= instead of ==)
    - err was not assigned
