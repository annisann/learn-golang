## HTTP Basic Auth
A http request authentication technique, encoded username & password pasted on request header.
<br>
e.g. Request Header with `key`:`value` as 
<br>`Authorization: Basic [encryption of username and password data]`,
<br> Where:
```
// Username and password encryption
base64encode("username:password")
```
## About
A simple web service containing an endpoint:
1. Endpoint `/student`, that shows all students data.
2. Endpoint `/student?id=s001`, that shows student data by the ID.