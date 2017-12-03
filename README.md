# gostand

Commonly used constants names for Go 

### Motivation
This package aims lacking of commonly used Go constants from various contexts like http headers or media types. No more "application/json", "Content-Length" etc. manual writing in your code.

### Idea

The **idea** is to provide widely used constants names for common packages as nicely companions. 

This package follow the **convention** that every package name in **gostand** begins with underscore("**_**") to avoid collisions and provide smooth companion for other packages. 

For instance _net, _os, _http etc. This concept will provide possibility to add common constant to existing packages like net, net/http, os and others. 

### Instalation

**Get all packages:**

````
go get -u github.com/fractalpal/gostand/...
````

##### Packages:
_net:

````
go get -u github.com/fractalpal/gostand/net
````

_http:

````
go get -u github.com/fractalpal/gostand/net/http
````

### Example
```
import(
    "github.com/fractalpal/gostand/net/"
    "github.com/fractalpal/gostand/net/http"
)

```



```
contentType := _http.HeaderContentType
mediaType := _net.MediaTypeApplicationJSON
```

### Source code sneak peek
```
...
HeaderAccept = "Accept"
HeaderAcceptCharset = "Accept-Charset"
HeaderAcceptEncoding = "Accept-Encoding"
HeaderAcceptLanguage = "Accept-Language"
HeaderAcceptRanges = "Accept-Ranges"
HeaderAge = "Age"
HeaderAllow = "Allow"
HeaderAuthorization = "Authorization"
HeaderCacheControl = "Cache-Control"
HeaderConnection = "Connection"
HeaderContentEncoding = "Content-Encoding"
HeaderContentLanguage = "Content-Language"
HeaderContentLength = "Content-Length"
HeaderContentLocation = "Content-Location"
HeaderContentMD5 = "Content-MD5"

...
MediaTypeApplicationJSONAPI = "application/vnd.api+json"
MediaTypeApplicationAtomXML = "application/atom+xml"
MediaTypeApplicationFormURLEncoded = "application/x-www-form-urlencoded"
MediaTypeApplicationJSON = "application/json"
MediaTypeApplicationOctecStream = "application/octet-stream"
MediaTypeApplicationSVGXML = "application/svg+xml"
MediaTypeApplicationXHTMLXML = "application/xhtml+xml"
MediaTypeApplicationXML = "application/xml"
MediaTypeWildcard = "*"
MediaTypeMultipartFormData = "multipart/form-data"
MediaTypeTextHTML = "text/html"
MediaTypeTextPlain = "text/plain"
MediaTypeTextXML = "text/xml"
MediaWildcard = "*/*"
```

### Contribution
If you like the concept and want to contribute, please feel free to add any commonly used (RFC/KNOWN specification) constants in best appropriate package location and open PR to extend this package even further. Just remember about **underscore** naming convention.

### License
[MIT](https://github.com/fractalpal/gostand/blob/master/LICENSE "License")