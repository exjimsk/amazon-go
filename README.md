# amazon-go
Amazon product API wrapper

Create an Amazon Product API request:

    func AmazonTestRequest() amazon.Request {
    
        // use your own Amazon affiliate information
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


Call functions on the Request type to :

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


