package amazon

import ()

type ItemSearchResponse struct {
	XMLNS            string `xml:"xmlns,attr"`
	OperationRequest struct {
		HTTPHeaders           []itemSearchResponseHeader `xml:"HTTPHeaders>Header"`
		RequestId             string
		Arguments             []itemSearchResponseArgument `xml:"Arguments>Argument"`
		RequestProcessingTime string
	}
	Items struct {
		Request struct {
			IsValid           bool
			ItemSearchRequest struct {
				Keywords       string
				ResponseGroups []string `xml:"ResponseGroup"`
				SearchIndex    string
			}
		}
		TotalResults         int
		TotalPages           int
		MoreSearchResultsUrl string
		Items                []itemSearchItem `xml:"Item"`
	}
}

type itemSearchResponseHeader struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

type itemSearchResponseArgument struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

type itemSearchItem struct {
	ASIN           string
	DetailPageURL  string
	ItemLinks      []itemSearchResponseItemLink `xml:"ItemLinks>ItemLink"`
	ItemAttributes struct {
		Binding              string
		Brand                string
		Color                string
		EAN                  string
		EANList              []string
		Features             []string `xml:"Feature"`
		IsEligibleForTradeIn int
		ItemDimensions       itemSearchResponseDimensions
		Label                string
		ListPrice            itemSearchResponsePrice
		Manufacturer         string
		Model                string
		MPN                  string
		PackageDimensions    itemSearchResponseDimensions
		PackageQuantity      int
		PartNumber           string
		ProductGroup         string
		ProductTypeName      string
		Publisher            string
		Studio               string
		Title                string
		TradeInValue         itemSearchResponsePrice
		UPC                  string
		UPCList              []string `xml:"UPCList>UPCListElement"`
	}
	OfferSummary struct {
		LowestNewPrice   itemSearchResponsePrice
		LowestUsedPrice  itemSearchResponsePrice
		TotalNew         int
		TotalUsed        int
		TotalCollectible int
		TotalRefurbished int
	}
}

type itemSearchResponseItemLink struct {
	Description string
	URL         string
}

type itemSearchResponseDimensions struct {
	Height itemSearchResponseMeasurement `xml:"Height"`
	Length itemSearchResponseMeasurement `xml:"Length"`
	Weight itemSearchResponseMeasurement `xml:"Weight"`
	Width  itemSearchResponseMeasurement `xml:"Width"`
}

type itemSearchResponseMeasurement struct {
	Units string  `xml:"Units,attr"`
	Value float64 `xml:",chardata"`
}

type itemSearchResponsePrice struct {
	Amount         float64
	CurrencyCode   string
	FormattedPrice string
}
