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

# Get All Request

* Endpoint: `/api/v1/requests`
* HTTP Method: `GET`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body: `-`
* Response Body:
    ```
    [
        {
            "id": "1Ziy6gwcYRtuUpc31y6l511AGIA",
            "date": "2020-03-27T19:07:15.715103Z",
            "isFulfilled": false,
            "donationApplicant": {
                "id": "1ZiESUvzPPP0r9SCTtkt84eAGfP",
                "email": "admin",
                "name": "admin",
                "role": "ADMIN"
            },
            "requestItems": [
                {
                    "id": "1Ziy6gWHnNOyCIGb7SxyBd08st6",
                    "item_id": "1",
                    "unit_id": "1",
                    "quantity": "123.00",
                    "request_id": "1Ziy6gwcYRtuUpc31y6l511AGIA"
                }
            ]
        },
        {
            "id": "1Zj2J5sHcTYVgRhh4ySPeTAYRvK",
            "date": "2020-03-27T19:41:47.719429Z",
            "isFulfilled": false,
            "donationApplicant": {
                "id": "1ZiESUvzPPP0r9SCTtkt84eAGfP",
                "email": "admin",
                "name": "admin",
                "role": "ADMIN"
            },
            "requestItems": [
                {
                    "id": "1Zj2J5Vwoz0nYZiqtAHj6TttQj5",
                    "item_id": "1",
                    "unit_id": "1",
                    "quantity": "11.00",
                    "request_id": "1Zj2J5sHcTYVgRhh4ySPeTAYRvK"
                },
                {
                    "id": "1Zj2J9sdTX2H0s2Re05vnLuPxo7",
                    "item_id": "1",
                    "unit_id": "1",
                    "quantity": "2.13",
                    "request_id": "1Zj2J5sHcTYVgRhh4ySPeTAYRvK"
                }
            ]
        }
    ]
    ```