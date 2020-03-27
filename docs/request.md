# Create Request

* Endpoint: `/api/v1/requests`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body:
    ```
    {
        "requestItems": [
            {
                "item_id": "1",
                "unit_id": "1",
                "quantity": "75.75"
            }
        ]
    }
    ```
* Response Body:
    ```
    {
        "id": "1ZiG6pthZHO1wHRIue7e9AK3He7",
        "date": "2020-03-27T20:05:28.121736048+07:00",
        "is_fulfilled": false,
        "donation_applicant_id": "1ZiESUvzPPP0r9SCTtkt84eAGfP"
    }
    ```