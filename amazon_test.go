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
		AssociateTag: "associatetag", 
		AccessKeyId: "accesskeyid", 
		SecretKey: "secretkey",
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

