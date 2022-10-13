## HTTP Basic Auth
A http request authentication technique, encoded username & password pasted on request header.
e.g. Request Header with `key`:`value` as 
<br>`Authorization: Basic [encryption of username and password data]`, where:
```
// Username and password encryption
base64encode("username:password")
```