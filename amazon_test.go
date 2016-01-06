package amazon_test

import (
	"testing"
	"fmt"
	"github.com/exjimsk/affiliate-go/market/amazon"
	
	"io/ioutil"
 	"net/http"
)

func createAmazonRequest() amazon.Request {

	cred := amazon.Credentials{ 
		AssociateTag: "affilia0a1-20", 
		AccessKeyId: "AKIAJFDVMAPZJZHURAJQ", 
		SecretKey: "UTWrgiB+xYDkyvNKhJH+igrRm81CWhe57Z6/m1S",
		Marketplace: "webservices.amazon.com" }
	
	req := amazon.NewRequest(cred)

	req.SetParameter("SearchIndex", "All")
	req.SetParameter("Keywords", "headphones")
	req.SetParameter("ResponseGroup", "Images,ItemAttributes,Offers")
	
	return req
}

func TestUnsignedAmazonRequest(t *testing.T) {
	req := createAmazonRequest()
	u_URL := req.UnsignedURL()
	
	fmt.Printf("Unsigned url:\n%v\n", u_URL)
}

func TestSignedAmazonRequest(t *testing.T) {
	req := createAmazonRequest()
	s_URL := req.SignedURL()

	fmt.Printf("Signed url: \n%v\n", s_URL)
}

func TestAmazonResponse(t *testing.T) {
	req := createAmazonRequest()
	s_URL := req.SignedURL()
	
	resp, err := http.Get(s_URL)
	if err != nil { 
		fmt.Println(err) 
	} else { 
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil { fmt.Println(body) }
	}
}



