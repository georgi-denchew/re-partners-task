# RE Partners challenge

This repository contains the backend application for the RE Partners Software Engineering Challenge.

The backend application is written in Golang and exposes an HTTP API for interaction.  
The backend API is accessible through <http://3.71.13.28:3000> (example usage described bellow).  
The frontend (written in React/TypeScript) is available [here](http://gi-re-partners-fe.s3-website.eu-central-1.amazonaws.com/).

## Endpoints

### `POST /orders`

Creates a new order

#### Body Parameters

> | name        | type     | data type | description                                 |
> |-------------|----------|-----------|---------------------------------------------|
> | items_count | required | Number    | Number of items the API user wants to order |

#### Responses

> | http code | content-type       | response                                                   |
> |-----------|--------------------|------------------------------------------------------------|
> | `201`     | `application/json` | `{"item_packs": [{"items": 250,"packs": 1}]}`              |
> | `400`     | `application/json` | `{"message": "field 'items_count' cannot be less than 1"}` |

#### Example cURL

```sh
curl -X POST -H "Content-Type: application/json" --data '{"items_count": 4200}' http://3.71.13.28:3000/orders
```

#### Example Response

```json
{
  "item_packs": [
    {
      "items":250,
      "packs":1
    },
    {
      "items":2000,
      "packs":2
    }
  ]
}
```

### `GET /orders`

Retrieves stored orders

#### Query Parameters

> | name        | type     | data type | description                                     |
> |-------------|----------|-----------|-------------------------------------------------|
> | count       | optional | Number    | Number of orders the API user wants to retrieve |

#### Responses

> | http code | content-type       | response                                                                                 |
> |-----------|--------------------|------------------------------------------------------------------------------------------|
> | `200`     | `application/json` | `[{"created_at":"2025-03-10T23:09:59.42748779Z","item_packs":[{"items":250,"packs":1}]}]`|
> | `400`     | `application/json` | ``                               |

#### Example cURL

```sh
curl http://3.71.13.28:3000/orders
```

#### Example Response

```json
[
  {"created_at":"2025-03-10T23:25:48.858290938Z","item_packs":[{"items":250,"packs":1},{"items":500,"packs":1}]}
  {"created_at":"2025-03-10T23:09:59.42748779Z","item_packs":[{"items":250,"packs":1}]}
]
```

### `PUT /admin/packs`

Replaces the available item packs in the application. Requires Basic Auth credentials - admin:password

#### Body Parameters

> | name        | type     | data type | description                                 |
> |-------------|----------|-----------|---------------------------------------------|
> | packs       | required | Number[]  | Array with the available pack sizes         |

#### Responses

> | http code | content-type       | response                                                   |
> |-----------|--------------------|------------------------------------------------------------|
> | `201`     | `application/json` | `{"item_packs": [{"items": 250,"packs": 1}]}`              |
> | `400`     | `application/json` | `{"message": "invalid request body"}`                      |

#### Example cURL

**Disclaimer**: Basic auth credentials are only left in the example for convenience. They shouldn't be here in real-world application.

```sh
curl --request PUT 'http://3.71.13.28:3000/admin/packs' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--data '{
    "packs":[250,500,1000,2000,5000]
}'
```

#### Example Response

```json
{
  "item_packs": [
    {
      "items":250,
      "packs":1
    },
    {
      "items":2000,
      "packs":2
    }
  ]
}
```

## Local setup

### Prerequisites

- clone the repository
- set `PACK_SIZES` environment variable which specifies the sizes for available item packs, e.g.

  ```sh
  PACK_SIZES=5000,250,500,1000,2000
  ```

### Run application

## Implementation decisions

The following is a list of purposefully made decisions:

- the application is containerized using Docker and hosted on AWS ECS hosted
- the frontend is kept in the same repository for simplicity
- the frontend is hosted in an S3 bucket
- `PACK_SIZES` are configurable through an environment variable
- `PACK_SIZES` can also be replaced using `PUT /admin/packs`
- unit tests run as part of the image build process
- different unit testing practices have been utilized to showcase variety
  - some are written as whitebox tests, where the package name doesn't include the `_test` suffix, allowing the test to access private fields and functions
  - some are written as blackbox tests
  - some utilize testify's `suite.Suite`
- apart from the use of [gomock](https://github.com/uber-go/mock), no other code generation tools were used so the reader could evaluate the author's coding style
  - some useful tools would have been [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) and [cleanenv](https://github.com/ilyakaznacheev/cleanenv)

## List of potential improvements

- replace in-memory persistence with a Database and implement a repository for DB interactions
- CI/CD pipeline
- add an OpenAPI spec for API definition
  - could also be used with `oapi-codegen` to generate boilerplate types and http handlers
- there's a `createdAt` field set against each stored order which isn't displayed on the front-end
