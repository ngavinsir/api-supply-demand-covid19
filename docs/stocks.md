# Get All Stocks

* Endpoint: `/api/v1/stocks`
* HTTP Method: `GET`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer jwt`
* Request Body: `-`
* Response Body:
    ```
    {
        "success": true,
        "data": [
            {
                "id" : 1,
                "name" : "string",
                "unit" : "Box",
                "quantity" : 100
            },
        ],
        "pages": {
            "current" : 1,
            "total" :  5,
            "first" : true,
            "last" : false
        }
    }
    ```