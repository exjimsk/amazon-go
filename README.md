# amazon-go
Amazon product API wrapper

	
To create a signed Amazon Product request, use a Credentials struct to create a Request type, timestamp your request and get the signed request URL.


	import (
		"github.com/exjimsk/amazon-go"
	)
	
	fun Main() {

		// use own Amazon affiliate information here
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
		r.Parameters["Version"] = "2013-08-01"
		
		// make sure to add a timestamp
		r.Parameters["Timestamp"] = amazon.CurrentTimestamp()

		signed_url := r.SignedURL()
		
	}
