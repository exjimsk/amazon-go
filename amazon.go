package amazon

import (
	"fmt"
	"net/url"
	"strings"
	"errors"
	"time"
	"sort"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

type Request struct {
	URL	url.URL
	parameters map[string]string
	Credentials Credentials
}

type Credentials struct {
	AssociateTag 	string
	AccessKeyId 	string
	SecretKey 		string
	Marketplace		string
}


func NewRequest(c Credentials) Request {
	r := Request{}
	r.URL = url.URL{}
	r.parameters = make(map[string]string)
	
	// standard stuff
	r.URL.Scheme = "http"
	r.URL.Host = c.Marketplace
	r.URL.Path = "/onca/xml"
		
	// standard stuff
	r.SetParameter("Service", "AWSECommerceService")
	r.SetParameter("Operation", "ItemSearch")
	r.SetParameter("Version", "2013-08-01")
	
	// mandatory stuff
	r.SetParameter("SubscriptionId", c.AccessKeyId)
	r.SetParameter("AssociateTag", c.AssociateTag)
	
	return r
}

// main method to compose a request
func (r Request) SetParameter(key string, val string) (err error) {
	if _, keyExists := r.parameters[key]; keyExists {
		// return error if parameter is already set
		return errors.New(fmt.Sprintf("parameter key %v already set", key))
	} else {
		r.parameters[key] = val
		return nil
	}
}

// cf. http://webservices.amazon.com/scratchpad/index.html
// not really that useful
func (r Request) UnsignedURL() (url string) {
	return fmt.Sprintf("%v?%v", r.hostAndPath(), r.sortedParametersAsString(false))
}

// THE method
func (r Request) SignedURL() string {
	r.SetParameter("Signature", r.signature())
	return fmt.Sprintf("%v/%v", r.hostAndPath(), r.sortedParametersAsString(true))
}


/*
	helps
*/
func (r Request) timestamp() {
	r.parameters["Timestamp"] = time.Now().UTC().Format(time.RFC3339)
}

func (r Request) hostAndPath() string {
	return r.URL.String()
}

func (r Request) sortedParametersAsString(escape bool) string {
	// instantiate container slice
	parameters := make([]string, 0, len(r.parameters))
	
	// append escaped/unescaped parameters to slice
	for p, _ := range r.parameters {
		var _p string
		if escape {
			_p = fmt.Sprintf("%v=%v", p, url.QueryEscape(r.parameters[p]))
		} else { 
			_p = fmt.Sprintf("%v=%v", p, r.parameters[p])
		}
		parameters = append(parameters, _p)
	}
	
	// sort slice and join with ampersand
	sort.Strings(parameters)
	return strings.Join(parameters, "&")
}

func (r Request) signature() string {
	// add timestamp to parameters map
	r.timestamp()
	
	// get hash string
	signatureStr := fmt.Sprintf("GET\n%v\n%v\n%v", r.URL.Host, r.URL.Path, r.sortedParametersAsString(false))

	// do the sha-256 hash on hash string using secret key
	hasher := hmac.New(sha256.New, []byte(r.Credentials.SecretKey))
	_, err := hasher.Write([]byte(signatureStr))
	if err != nil {	return "" }

	// return escaped base64 signature hash for use in signed URL
	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	hash = url.QueryEscape(hash)
	return hash
}

