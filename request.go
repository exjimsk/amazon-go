package amazon

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Credentials struct {
	AssociateTag string
	AccessKeyId  string
	SecretKey    string
	Marketplace  string
}

type Request struct {
	URL         url.URL
	Parameters  map[string]string
	Credentials Credentials
}

func HashSignature(str string, secret string) string {
	// do the sha-256 hash on hash string using secret key
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(str))
	if err != nil {
		return ""
	}

	// return escaped base64 signature hash for use in signed URL
	hash := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	hash = url.QueryEscape(hash)
	return hash
}

func CurrentTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func NewRequest(c Credentials) Request {
	r := Request{}
	r.URL = url.URL{}
	r.Parameters = make(map[string]string)

	// copy relevant credential data into URL and Parameters fields
	r.Credentials = c
	r.URL.Scheme = "http"
	r.URL.Host = c.Marketplace
	r.URL.Path = "/onca/xml"
	r.Parameters["AWSAccessKeyId"] = c.AccessKeyId
	r.Parameters["AssociateTag"] = c.AssociateTag

	return r
}

func (r Request) UnsignedURL() string {
	r.URL.RawQuery = r.sortedParametersAsString(true)
	return r.URL.String()
}

func (r Request) CanonicalString() string {
	return fmt.Sprintf("GET\n%v\n%v\n%v", r.URL.Host, r.URL.Path, r.sortedParametersAsString(true))
}

func (r Request) SignedURL() string {
	cStr := r.CanonicalString()
	sig := HashSignature(cStr, r.Credentials.SecretKey)
	r.URL.RawQuery = fmt.Sprintf("%v&Signature=%v", r.sortedParametersAsString(true), sig)

	return r.URL.String()
}

func (r Request) sortedParametersAsString(escape bool) string {

	// instantiate container slice
	parameters := make([]string, 0, len(r.Parameters))

	// append escaped/unescaped parameters to slice
	for p, _ := range r.Parameters {
		var _p string
		if escape {
			_p = fmt.Sprintf("%v=%v", p, url.QueryEscape(r.Parameters[p]))

			// force encoding of "+" into "%20"
			_p = strings.Replace(_p, "+", "%20", -1)
		} else {
			_p = fmt.Sprintf("%v=%v", p, r.Parameters[p])
		}
		parameters = append(parameters, _p)
	}

	// sort slice and join with ampersand
	sort.Strings(parameters)
	return strings.Join(parameters, "&")
}
