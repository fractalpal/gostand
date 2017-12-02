# gostand

Commonly used constants names for Go 

The **idea** is to provide widely used constants names for common packages as nicely companions. 

This package follow the **convention** that every package name in **gostand** begins with underscore("**_**") to avoid collisions and provide smooth companion for other packages. 

For instance _net, _os, _http, _header, _media etc. This concept will provide possibility to add common constant to existing packages like net, net/http, os and others. 

### Instalation

````
go get -u github.com/fractalpal/gostand
````

### Background
This package aims lacking of commonly used Go constants from various contexts like http headers or media types. No more "application/json", "Content-Length" etc. manual writing in your code.

### Example
```
import(
    "github.com/fractalpal/gostand/net/http/header"
    "github.com/fractalpal/gostand/net/media"
)

```



```
contentType := _header.ContentType
mediaType := _media.TypeApplicationJSON
```

Example code lookup
```
...
Accept = "Accept"
AcceptCharset = "Accept-Charset"
AcceptEncoding = "Accept-Encoding"
AcceptLanguage = "Accept-Language"
AcceptRanges = "Accept-Ranges"
Age = "Age"
Allow = "Allow"
Authorization = "Authorization"
CacheControl = "Cache-Control"
Connection = "Connection"
ContentEncoding = "Content-Encoding"
ContentLanguage = "Content-Language"
ContentLength = "Content-Length"
ContentLocation = "Content-Location"
ContentMD5 = "Content-MD5"

...
TypeApplicationJSONAPI = "application/vnd.api+json"
TypeApplicationAtomXML = "application/atom+xml"
TypeApplicationFormURLEncoded = "application/x-www-form-urlencoded"
TypeApplicationJSON = "application/json"
TypeApplicationOctecStream = "application/octet-stream"
TypeApplicationSVGXML = "application/svg+xml"
TypeApplicationXHTMLXML = "application/xhtml+xml"
TypeApplicationXML = "application/xml"
TypeWildcard = "*"
TypeMultipartFormData = "multipart/form-data"
TypeTextHTML = "text/html"
TypeTextPlain = "text/plain"
TypeTextXML = "text/xml"
Wildcard = "*/*"
```

### Contribution
If you like the concept and want to contribute, please feel free to add any commonly used (RFC/KNOWN specification) constants in best appropriate package location and open PR to extend this package even further. Just remember about **underscore** naming convention.

### License
[MIT](https://github.com/fractalpal/gostand/blob/master/LICENSE "License")