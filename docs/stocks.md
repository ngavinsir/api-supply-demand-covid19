# Get All Stocks

* Endpoint: `/api/v1/stocks`
* HTTP Method: `GET`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
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

# Create or Update Stock (Update)

* Endpoint: `/api/v1/stocks`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer jwt`
* Request Body: 
    ```
    {
        "id": "1",
        "quantity": "5.47"
    }
    ```
* Response Body:
    ```
    {
        "id": "1",
        "item_id": "1",
        "unit_id": "1",
        "quantity": "11.94"
    }
    ```

# Create or Update Stock (Create)

* Endpoint: `/api/v1/stocks`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer jwt`
* Request Body: 
    ```
    {
        "item_id": "2",
        "unit_id": "1",
        "quantity": "5.47"
    }
    ```
* Response Body:
    ```
    {
        "id": "1Zwd1U3MKH0bsOBOVWZ1aSiXcdf",
        "item_id": "2",
        "unit_id": "1",
        "quantity": "5.47"
    }
    ```