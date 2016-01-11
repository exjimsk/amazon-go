package amazon_test

import (
	"testing"
	"github.com/exjimsk/amazon-go"
	"fmt"

)


func AmazonTestRequest() amazon.Request {
	// these parameters are from http://docs.aws.amazon.com/AWSECommerceService/latest/DG/rest-signature.html to test product API requests and specifically to generate a proper signature hash
	c := amazon.Credentials{ 
		AssociateTag: "mytag-20", 
		AccessKeyId: "AKIAIOSFODNN7EXAMPLE", 
		SecretKey: "1234567890",
		Marketplace: "webservices.amazon.com" }
	
	r := amazon.NewRequest(c)

	r.Parameters["ItemId"] = "0679722769"
	r.Parameters["Operation"] = "ItemLookup"
	r.Parameters["ResponseGroup"] = "Images,ItemAttributes,Offers,Reviews"
	r.Parameters["Service"] = "AWSECommerceService"
	r.Parameters["Timestamp"] = "2014-08-18T12:00:00Z" // use amazon.CurrentTimestamp() for real request
	r.Parameters["Version"] = "2013-08-01"

	return r
}



var r = AmazonTestRequest()



func TestUnsignedAmazonRequest(t *testing.T) {
	EXPECTED := "http://webservices.amazon.com/onca/xml?AWSAccessKeyId=AKIAIOSFODNN7EXAMPLE&AssociateTag=mytag-20&ItemId=0679722769&Operation=ItemLookup&ResponseGroup=Images%2CItemAttributes%2COffers%2CReviews&Service=AWSECommerceService&Timestamp=2014-08-18T12%3A00%3A00Z&Version=2013-08-01"
	
	if RESULT := r.UnsignedURL(); RESULT != EXPECTED {
		t.Errorf("\nEXPECTED:\n%v\nRESULT:\n%v", EXPECTED, RESULT)
	} else { fmt.Println("TestUnsignedAmazonRequest: OK") }
	
	}



func TestAmazonSignature(t *testing.T) {
/*
GET
webservices.amazon.com
/onca/xml
AWSAccessKeyId=AKIAIOSFODNN7EXAMPLE&AssociateTag=mytag-20&ItemId=0679722769&Operation=ItemLookup&ResponseGroup=Images%2CItemAttributes%2COffers%2CReviews&Service=AWSECommerceService&Timestamp=2014-08-18T12%3A00%3A00Z&Version=2013-08-01
*/
	EXPECTED := "GET\nwebservices.amazon.com\n/onca/xml\nAWSAccessKeyId=AKIAIOSFODNN7EXAMPLE&AssociateTag=mytag-20&ItemId=0679722769&Operation=ItemLookup&ResponseGroup=Images%2CItemAttributes%2COffers%2CReviews&Service=AWSECommerceService&Timestamp=2014-08-18T12%3A00%3A00Z&Version=2013-08-01"

	cStr := r.CanonicalString()
	
	if RESULT := cStr; RESULT != EXPECTED {
		t.Errorf("\nEXPECTED:\n%v\nRESULT:\n%v", EXPECTED, RESULT)
	} else { fmt.Println("TestAmazonSignature: Canonical String: OK") }


	EXPECTED = "j7bZM0LXZ9eXeZruTqWm2DIvDYVUU3wxPPpp%2BiXxzQc%3D"
	
	if RESULT := amazon.HashSignature(cStr, r.Credentials.SecretKey); RESULT != EXPECTED {
		t.Errorf("\nEXPECTED:\n%v\nRESULT:\n%v", EXPECTED, RESULT)
	} else { fmt.Println("TestAmazonSignature: Signature: OK") }
	

}



func TestSignedAmazonRequest_ItemLookup(t *testing.T) {
	EXPECTED := "http://webservices.amazon.com/onca/xml?AWSAccessKeyId=AKIAIOSFODNN7EXAMPLE&AssociateTag=mytag-20&ItemId=0679722769&Operation=ItemLookup&ResponseGroup=Images%2CItemAttributes%2COffers%2CReviews&Service=AWSECommerceService&Timestamp=2014-08-18T12%3A00%3A00Z&Version=2013-08-01&Signature=j7bZM0LXZ9eXeZruTqWm2DIvDYVUU3wxPPpp%2BiXxzQc%3D"
	
	if RESULT := r.SignedURL(); RESULT != EXPECTED {
		t.Errorf("\nEXPECTED:\n%v\nRESULT:\n%v", EXPECTED, RESULT)
	} else { fmt.Println("TestSignedAmazonRequest_ItemLookup: OK") }
}

func TestSignedAmazonRequest_ItemSearch(t *testing.T) {
	EXPECTED := "Gv4kWyAAD3xgSGI86I4qZ1zIjAhZYs2H7CRTpeHLD1o%3D"
	
	cStr := "GET\nwebservices.amazon.co.uk\n/onca/xml\nAWSAccessKeyId=AKIAIOSFODNN7EXAMPLE&Actor=Johnny%20Depp&AssociateTag=mytag-20&Operation=ItemSearch&Operation=ItemSearch&ResponseGroup=ItemAttributes%2COffers%2CImages%2CReviews%2CVariations&SearchIndex=DVD&Service=AWSECommerceService&Sort=salesrank&Timestamp=2014-08-18T17%3A34%3A34.000Z&Version=2013-08-01"
	
	if RESULT := amazon.HashSignature(cStr, r.Credentials.SecretKey); RESULT != EXPECTED {
		t.Errorf("\nEXPECTED:\n%v\nRESULT:\n%v", EXPECTED, RESULT)
	} else { fmt.Println("TestSignedAmazonRequest_ItemSearch: OK") }
}

func TestSignedAmazonRequest_CartCreate(t *testing.T) {
	EXPECTED := "LpEUnc9tT4WGneeUwH0LvwxLLfbMEXgmjGX5GXQ1MEQ%3D"
	
	cStr := "GET\nwebservices.amazon.com\n/onca/xml\nAWSAccessKeyId=AKIAIOSFODNN7EXAMPLE&AssociateTag=mytag-20&Item.1.OfferListingId=j8ejq9wxDfSYWf2OCp6XQGDsVrWhl08GSQ9m5j%2Be8MS449BN1XGUC3DfU5Zw4nt%2FFBt87cspLow1QXzfvZpvzg%3D%3D&Item.1.Quantity=3&Operation=CartCreate&Operation=ItemSearch&Service=AWSECommerceService&Timestamp=2014-08-18T17%3A36%3A55.000Z&Version=2013-08-01"
	
	if RESULT := amazon.HashSignature(cStr, r.Credentials.SecretKey); RESULT != EXPECTED {
		t.Errorf("\nEXPECTED:\n%v\nRESULT:\n%v", EXPECTED, RESULT)
	} else { fmt.Println("TestSignedAmazonRequest_CartCreate: OK") }
}
