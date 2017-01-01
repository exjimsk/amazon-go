package amazon_test

import (
	"encoding/xml"
	"github.com/exjimsk/amazon-go"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestItemSearchResponse(t *testing.T) {
	isr := amazon.ItemSearchResponse{}
	b, err := ioutil.ReadFile("sample.xml")
	if err != nil {
		t.Fatal(err)
	}

	if err := xml.Unmarshal(b, &isr); err != nil {
		t.Fatal(err)
	}

	assert.True(t, len(isr.XMLNS) > 0)
	assert.True(t, len(isr.OperationRequest.HTTPHeaders) > 0)
	assert.True(t, len(isr.OperationRequest.HTTPHeaders[0].Name) > 0)
	assert.True(t, len(isr.OperationRequest.HTTPHeaders[0].Value) > 0)
	assert.True(t, len(isr.OperationRequest.RequestId) > 0)
	assert.True(t, len(isr.OperationRequest.Arguments) > 0)
	assert.True(t, len(isr.OperationRequest.Arguments[0].Name) > 0)
	assert.True(t, len(isr.OperationRequest.Arguments[0].Value) > 0)
	assert.True(t, len(isr.OperationRequest.RequestProcessingTime) > 0)

	assert.True(t, isr.Items.Request.IsValid == true || isr.Items.Request.IsValid == false)
	assert.True(t, len(isr.Items.Request.ItemSearchRequest.Keywords) > 0)
	assert.True(t, len(isr.Items.Request.ItemSearchRequest.ResponseGroups) > 0)
	assert.True(t, len(isr.Items.Request.ItemSearchRequest.SearchIndex) > 0)
	assert.True(t, len(isr.Items.Items) > 0)
	assert.True(t, len(isr.Items.Items[0].DetailPageURL) > 0)
	assert.True(t, len(isr.Items.Items[0].ItemLinks) > 0)
	assert.True(t, len(isr.Items.Items[0].ItemLinks[0].Description) > 0)
	assert.True(t, len(isr.Items.Items[0].ItemLinks[0].URL) > 0)

	assert.True(t, len(isr.Items.Items[0].ItemAttributes.EANList) > 0)
	assert.True(t, len(isr.Items.Items[0].ItemAttributes.Features) > 0)

	assert.True(t, isr.Items.Items[0].ItemAttributes.ItemDimensions.Height.Value > 0)
	assert.True(t, len(isr.Items.Items[0].ItemAttributes.ItemDimensions.Height.Units) > 0)

	assert.True(t, isr.Items.Items[0].ItemAttributes.ListPrice.Amount > 0)
	assert.True(t, len(isr.Items.Items[0].ItemAttributes.ListPrice.CurrencyCode) > 0)
	assert.True(t, len(isr.Items.Items[0].ItemAttributes.ListPrice.FormattedPrice) > 0)

	assert.True(t, len(isr.Items.Items[0].ItemAttributes.UPCList) > 0)

	assert.True(t, len(isr.Items.Items[0].OfferSummary.LowestNewPrice.FormattedPrice) > 0)
	assert.True(t, len(isr.Items.Items[0].OfferSummary.LowestUsedPrice.FormattedPrice) > 0)
	assert.True(t, isr.Items.Items[0].OfferSummary.TotalUsed > 0)
	assert.True(t, isr.Items.Items[0].OfferSummary.TotalNew > 0)
}
