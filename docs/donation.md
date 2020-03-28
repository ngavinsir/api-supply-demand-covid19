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
        "donationItems": [
            {
                "item_id": "1",
                "unit_id": "1",
                "quantity: 10
            },
            {
                "item_id": "2",
                "unit_id": "2",
                "quantity": 10
            }
        ]
    }
    ```
* Response Body:
    ```
    {
        "donation": {
            "id": "1ZlfXsUn4fYgvWc7Ke6SZSNuDPn",
            "date": "2020-03-29T01:04:03.458704516+07:00",
            "is_accepted": false,
            "is_donated": false,
            "donator_id": "1ZiESUvzPPP0r9SCTtkt84eAGfP"
        },
        "items": [
            {
                "id": "1ZlfXrepQpHoq6e4YYfhpdAZ4HK",
                "donation_id": "1ZlfXsUn4fYgvWc7Ke6SZSNuDPn",
                "item_id": "1",
                "unit_id": "1",
                "quantity": "10"
            },
            {
                "id": "1ZlfXnxtLuhKsqXCK4qAM2q0WAM",
                "donation_id": "1ZlfXsUn4fYgvWc7Ke6SZSNuDPn",
                "item_id": "1",
                "unit_id": "1",
                "quantity": "10"
            }
        ]
    }
    ```