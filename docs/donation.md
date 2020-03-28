# Create Donation

* Endpoint: `/api/v1/donations`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer jwt`
* Request Body:
    ```
    {
        "donation": [
            {
                "item_id": 1,
                "unit_id": 1,
                "quantity: 10
            },
            {
                "item_id": 2,
                "unit_id": 2,
                "quantity: 10
            },
        ]
    }
    ```
* Response Body:
    ```
    {
        "donation": {
            "id": "",
            "date": "",
            "is_accepted": true,
            "is_donated": false,
            "donator_id": ""
        },
        "items": [
            {
                "id": ""
                "item_id": 1,
                "unit_id": 1,
                "quantity: 10
            },
        ]
    }
    ```