# amazon-go
Amazon product API wrapper

	
How to use:


	import (
		"github.com/exjimsk/amazon-go"
		
		"io/ioutil"
		"net/http"
		"encoding/xml"
		
		"fmt"
	)
	
	fun Main() {

		// replace with your own Amazon affiliate information here
		c := amazon.Credentials{ 
			AssociateTag: "mytag-20", 
			AccessKeyId: "AKIAIOSFODNN7EXAMPLE", 
			SecretKey: "1234567890",
			Marketplace: "webservices.amazon.com" }
	  
		r := amazon.NewRequest(c)

		r.Parameters["Keywords"] = "Sony laptop"
		r.Parameters["Operation"] = "ItemSearch"
		r.Parameters["ResponseGroup"] = "ItemAttributes,Offers,Reviews"
		r.Parameters["Service"] = "AWSECommerceService"
		r.Parameters["Version"] = "2013-08-01"
		
		// make sure to add a timestamp
		r.Parameters["Timestamp"] = amazon.CurrentTimestamp()

		
		signed_url := r.SignedURL()
		resp, err := http.Get(signed_url)
		if err != nil {
			panic(err)
		}
		
		defer resp.Body.Close()
		b, _ := ioutil.ReadAll(resp.Body)
		
		isr := amazon.ItemSearchResponse{}
		if err := xml.Unmarshal(b, &isr); err != nil {
			panic(err)
		}
		
		for _, item := range isr.Items.Items {
			// read item
		}
	}
