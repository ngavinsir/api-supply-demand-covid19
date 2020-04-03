# Create Allocation

* Endpoint: `/api/v1/allocations`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body:
    ```
    {
        "requestID": "1a28qGTSDShE0chlPSemlQqRdTs",
        "items": [
            {
                "item_id": "1",
                "unit_id": "1",
                "quantity": "1.16"
            }	
        ]
    }
    ```
* Response Body:
    ```
    {
        "id": "1a29mTwDPGmztVWLwBD4WtCPvo0",
        "date": "2020-04-03T21:09:40.85304552+07:00",
        "requestID": "1a28qGTSDShE0chlPSemlQqRdTs",
        "adminID": "1a28a0aYZslZ1EBAFm3YkXvVXd0",
        "photoUrl": "",
        "items": [
            {
                "id": "1a29mXtuHtqlrmgOa9oA2IZ6hZv",
                "item": "Masker",
                "unit": "Buah",
                "quantity": "1.16"
            }
        ]
    }
    ```