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

		r := amazon.NewRequest(amazon.Credentials{ 
			AssociateTag: "mytag-20", 
			AccessKeyId: "AKIAIOSFODNN7EXAMPLE", 
			SecretKey: "1234567890",
			Marketplace: "webservices.amazon.com",
		}) // use personal Amazon affiliate data

		r.Parameters["Operation"] = "ItemSearch"
		r.Parameters["ResponseGroup"] = "ItemAttributes,Offers,Reviews"
		r.Parameters["Keywords"] = "Sony laptop"
		r.Parameters["Service"] = "AWSECommerceService"
		r.Parameters["Version"] = "2013-08-01"
		r.Parameters["Timestamp"] = amazon.CurrentTimestamp() // important

		resp, err := http.Get(r.SignedURL())
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
			// read returned item
		}
	}
