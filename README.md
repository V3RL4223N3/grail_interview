# Prerequisites

- go 15.8.2
- git


# How to run

If you want to build locally you can run 
```
$ git clone https://github.com/V3RL4223N3/grail_interview
$ cd grail_interview
$ go run *.go
2021/02/13 16:46:34 Starting REST Server on :8080
```

Or if you want to run it through docker

```
$ docker run -p 8080:8080 snbsniper29/grail:latest
```

Then open a web browser on `http:localhost:8080/api/v1/participants/`

# Considerations

- Instead of using a database I used a hashmap
```
var db = make(map[string]*Participant)
```

- In order to emulate the random reference number I created the function `generateRandomReferenceNumber`, that generates a suedo-random UUID 
```
func generateRandomReferenceNumber() string {
	uuid := uuid.New()
	return uuid.String()
}

```

# Data Model

```
type Participant struct {
	Name            string `json:"name" validate:"required"`
	ReferenceNumber string `json:"referenceNumber" `
	DateOfBirth     string `json:"dateOfBirth" validate:"required"`
	PhoneNumber     string `json:"phoneNumber" validate:"required"`
	Address         string `json:"address" validate:"required"`
}
```
I just mapped the attributes mentioned in the PDF plus I added the reference number as part of the struct as well to act as a Unique Identifier


# API Design

The api is designed with a subroute on `/api/v1/participants/`

- `POST /`

Creates a Participant Entity


```
curl --location --request POST 'localhost:8080/api/v1/participant/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Daniel",
    "dateOfBirth": "08/02/1990",
    "phoneNumber": "+56981372435",
    "address": "Consistorial 2580"
}'

```

Returns

```
{
    "name": "Daniel",
    "referenceNumber": "22f13542-6c21-426c-99c5-0e0ed0d1fc64",
    "dateOfBirth": "08/02/1990",
    "phoneNumber": "+56981372435",
    "address": "Consistorial 2580"
}

```


- `GET /`

Returns a list of all Participants

```
curl --location --request GET 'localhost:8080/api/v1/participant/'

```

Returns

```
[
    {
        "name": "Daniel",
        "referenceNumber": "22f13542-6c21-426c-99c5-0e0ed0d1fc64",
        "dateOfBirth": "08/02/1990",
        "phoneNumber": "+56981372435",
        "address": "Consistorial 2580"
    }
]
```


- `GET/{id}`

Returns a specific Participant by reference ID Number

```
curl --location --request GET 'localhost:8080/api/v1/participant/9ff0178c-fe56-451e-a246-6752efc6b5ab'
```

Returns

```
{
    "name": "Daniel",
    "referenceNumber": "9ff0178c-fe56-451e-a246-6752efc6b5ab",
    "dateOfBirth": "08/02/1990",
    "phoneNumber": "+56981322435",
    "address": "Consistorial 2580"
}

```

- `PUT /{id}`

Updates a Participant by reference ID Number

```
curl --location --request PUT 'localhost:8080/api/v1/participant/9ff0178c-fe56-451e-a246-6752efc6b5ab' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "DanielTapia",
    "dateOfBirth": "08/02/1991",
    "phoneNumber": "+56981322435",
    "address": "Consistorial asdf 2580"
}'
```

Returns 

```
{
    "name": "DanielTapia",
    "referenceNumber": "9ff0178c-fe56-451e-a246-6752efc6b5ab",
    "dateOfBirth": "08/02/1991",
    "phoneNumber": "+56981322435",
    "address": "Consistorial asdf 2580"
}

```


- `DELETE /{id}`

```
curl --location --request DELETE 'localhost:8080/api/v1/participant/9ff0178c-fe56-451e-a246-6752efc6b5ab'
```

Returns
```
{
    "name": "DanielTapia",
    "referenceNumber": "9ff0178c-fe56-451e-a246-6752efc6b5ab",
    "dateOfBirth": "08/02/1991",
    "phoneNumber": "+56981322435",
    "address": "Consistorial asdf 2580"
}
```

Attempting to Delete User again results in

```
{
    "message": "No Participant Found",
    "httpStatus": 500
}

```