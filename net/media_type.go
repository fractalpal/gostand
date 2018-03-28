/*
* The MIT License (MIT)
*
*Copyright (c) 2017 Fractal Pal
*
* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
*
* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.
*
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*
* Commonly used net media types constants names
 */
package _net

import "github.com/fractalpal/gostand/charset"

const (
	// type application
	MediaTypeApplicationJSONAPI          = "application/vnd.api+json"
	MediaTypeApplicationJSONAPI_UTF8     = MediaTypeApplicationJSONAPI + "; " + _charset.UTF8
	MediaTypeApplicationAtomXML          = "application/atom+xml"
	MediaTypeApplicationAtomXML_UTF8     = MediaTypeApplicationAtomXML + "; " + _charset.UTF8
	MediaTypeApplicationFormURLEncoded   = "application/x-www-form-urlencoded"
	MediaTypeApplicationJSON             = "application/json"
	MediaTypeApplicationJSON_UTF8        = MediaTypeApplicationJSON + "; " + _charset.UTF8
	MediaTypeApplicationOctecStream      = "application/octet-stream"
	MediaTypeApplicationOctecStream_UTF8 = MediaTypeApplicationOctecStream + "; " + _charset.UTF8
	MediaTypeApplicationPDF              = "application/pdf" // https://tools.ietf.org/html/rfc3778
	MediaTypeApplicationPDF_UTF8         = MediaTypeApplicationPDF + "; " + _charset.UTF8
	MediaTypeApplicationSVGXML           = "application/svg+xml"
	MediaTypeApplicationSVGXML_UTF8      = MediaTypeApplicationSVGXML + "; " + _charset.UTF8
	MediaTypeApplicationXHTMLXML         = "application/xhtml+xml"
	MediaTypeApplicationXHTMLXML_UTF8    = MediaTypeApplicationXHTMLXML + "; " + _charset.UTF8
	MediaTypeApplicationXML              = "application/xml"
	MediaTypeApplicationXML_UTF8         = MediaTypeApplicationXML + "; " + _charset.UTF8
	MediaTypeApplicationDNS              = "application/dns" // https://tools.ietf.org/html/rfc4027
	MediaTypeApplicationDNS_UTF8         = MediaTypeApplicationDNS + "; " + _charset.UTF8
	MediaTypeApplicationHTTP             = "application/http" // https://tools.ietf.org/html/rfc7230
	MediaTypeApplicationHTTP_UTF8        = MediaTypeApplicationHTTP + "; " + _charset.UTF8
	MediaTypeApplicationGZIP             = "application/gzip" // https://tools.ietf.org/html/rfc6713
	MediaTypeApplicationGZIP_UTF8        = MediaTypeApplicationGZIP + "; " + _charset.UTF8
	MediaTypeApplicationZLIB             = "application/zlib" // https://tools.ietf.org/html/rfc6713
	MediaTypeApplicationZLIB_UTF8        = MediaTypeApplicationZLIB + "; " + _charset.UTF8

	// type text
	MediaTypeTextCSV             = "text/csv" // https://tools.ietf.org/html/rfc4180
	MediaTypeTextCSV_UTF8        = MediaTypeTextCSV + "; " + _charset.UTF8
	MediaTypeTextCSS             = "text/css" // https://tools.ietf.org/html/rfc2318
	MediaTypeTextCSS_UTF8        = MediaTypeTextCSS + "; " + _charset.UTF8
	MediaTypeTextHTML            = "text/html"
	MediaTypeTextHTML_UTF8       = MediaTypeTextHTML + "; " + _charset.UTF8
	MediaTypeTextPlain           = "text/plain"
	MediaTypeTextPlain_UTF8      = MediaTypeTextPlain + "; " + _charset.UTF8
	MediaTypeTextXML             = "text/xml"
	MediaTypeTextXML_UTF8        = MediaTypeTextXML + "; " + _charset.UTF8
	MediaTypeTextDNS             = "text/dns" // https://tools.ietf.org/html/rfc4027
	MediaTypeTextDNS_UTF8        = MediaTypeTextDNS + "; " + _charset.UTF8
	MediaTypeTextCalendar        = "text/calendar" // https://tools.ietf.org/html/rfc5545
	MediaTypeTextCalendar_UTF8   = MediaTypeTextCalendar + "; " + _charset.UTF8
	MediaTypeTextJavaScript      = "text/javascript" // https://tools.ietf.org/html/rfc4329
	MediaTypeTextJavaScript_UTF8 = MediaTypeTextJavaScript + "; " + _charset.UTF8
	MediaTypeTextMarkdown        = "text/markdown" // https://tools.ietf.org/html/rfc4329
	MediaTypeTextMarkdownt_UTF8  = MediaTypeTextMarkdown + "; " + _charset.UTF8

	// other
	MediaTypeWildcard          = "*"
	MediaTypeMultipartFormData = "multipart/form-data"
	MediaWildcard              = "*/*"
)
