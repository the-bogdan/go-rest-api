#  Rest APi

Simple rest api server on golang

## Response codes

200 - successfully. Returns contest in response body
204 - successfully. No content in response body
404 - instance or route no found
500 - something wrong with server

GET /users -- list of users -- 200, 404, 500
GET /users/:id -- user by id -- 200, 404, 500
POST /users/:id -- create user -- 204, 4xx, Header Location: url
PUT /users/:id -- fully update user -- 204, 404, 400, 500
DELETE /users/:id -- delete user by id -- 204, 404, 500
