# Title-Case-API

## Example

Run this in your terminal using [httpie](https://httpie.org)

```
http POST title-case-api.herokuapp.com title="convert to title case: Go + Python interop"
```

and receive the following result

```
HTTP/1.1 200 OK
Connection: keep-alive
Content-Length: 55
Content-Type: text/plain; charset=utf-8
Date: Tue, 24 Jan 2017 11:42:55 GMT
Server: Cowboy
Via: 1.1 vegur

{
    "Title": "Convert to Title Case: Go + Python Interop"
}
```
