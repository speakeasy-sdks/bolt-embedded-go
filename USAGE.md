<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
)

func main() {
    s := boltembeddedapi.New()
    operationSecurity := operations.AddAddressSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.AddAddress(ctx, operations.AddAddressRequest{
        IdempotencyKey: boltembeddedapi.String("corrupti"),
        XPublishableKey: boltembeddedapi.String("provident"),
        AddressAccount: &shared.AddressAccount{
            Company: boltembeddedapi.String("Bolt"),
            Country: boltembeddedapi.String("United States"),
            CountryCode: "US",
            Default: boltembeddedapi.Bool(false),
            DoorCode: boltembeddedapi.String("123456"),
            Email: "alan.watts@example.com",
            FirstName: "Alan",
            LastName: "Watts",
            Locality: "Brooklyn",
            Metadata: &shared.Metadata{
                AdditionalProperties: boltembeddedapi.String("distinctio"),
            },
            Name: boltembeddedapi.String("Alan Watts"),
            Phone: boltembeddedapi.String("+12125550199"),
            PostalCode: "10044",
            Region: "NY",
            RegionCode: boltembeddedapi.String("NY"),
            StreetAddress1: "888 main street",
            StreetAddress2: boltembeddedapi.String("apt 3021"),
            StreetAddress3: boltembeddedapi.String("c/o Alicia Watts"),
            StreetAddress4: boltembeddedapi.String("Bridge Street Apartment Building B"),
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AddAddress200ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->