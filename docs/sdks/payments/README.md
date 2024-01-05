# Payments
(*Payments*)

## Overview

Create and manage transactions for non credit card payments such as Paypal in your Embedded Accounts experience.


### Available Operations

* [FinalizePayment](#finalizepayment) - Finalize Payment
* [InitializePayment](#initializepayment) - Initialize Payment
* [UpdatePayment](#updatepayment) - Update Payment

## FinalizePayment

Finalize a Bolt Payment. NOTE: The authorization header is NOT required for payments associated with users who do not have a Bolt account.


### Example Usage

```go
package main

import(
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.FinalizePaymentSecurity{
            OAuth: boltembeddedgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }

    ctx := context.Background()
    res, err := s.Payments.FinalizePayment(ctx, operations.FinalizePaymentRequest{
        RequestBody: &operations.FinalizePaymentRequestBody{
            MerchantEventID: boltembeddedgo.String("dbe0cd5d-3261-41d9-ba61-49e5b9d07567"),
            ShopperIdentity: &operations.ShopperIdentity{
                CreateBoltAccount: boltembeddedgo.Bool(true),
                Email: "Shanon_Sipes@hotmail.com",
                FirstName: "Jalyn",
                LastName: "Reichert",
                Phone: "567.701.8847 x69933",
            },
        },
        ID: "<ID>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.FinalizePaymentRequest](../../pkg/models/operations/finalizepaymentrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.FinalizePaymentSecurity](../../pkg/models/operations/finalizepaymentsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.FinalizePaymentResponse](../../pkg/models/operations/finalizepaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## InitializePayment

Initialize a Bolt payment token that will be used to reference this payment to Bolt when it is updated or finalized. NOTE: The authorization header is NOT required for payments associated with users who do not have a Bolt account.


### Example Usage

```go
package main

import(
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"context"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.InitializePaymentSecurity{
            OAuth: boltembeddedgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }

    ctx := context.Background()
    res, err := s.Payments.InitializePayment(ctx, operations.InitializePaymentRequest{
        RequestBody: operations.CreateInitializePaymentRequestBodySchemas(
                operations.Schemas{
                    Cart: shared.CartCreate{
                        AddOns: []shared.CartAddOn{
                            shared.CartAddOn{
                                Name: "string",
                                Price: 7673.67,
                                ProductID: "string",
                            },
                        },
                        BillingAddress: &shared.Address{
                            Company: boltembeddedgo.String("Bolt"),
                            Country: boltembeddedgo.String("United States"),
                            CountryCode: "US",
                            Default: boltembeddedgo.Bool(true),
                            DoorCode: boltembeddedgo.String("123456"),
                            Email: "alan.watts@example.com",
                            FirstName: "Alan",
                            LastName: "Watts",
                            Locality: "Brooklyn",
                            Name: boltembeddedgo.String("Alan Watts"),
                            Phone: boltembeddedgo.String("+12125550199"),
                            PostalCode: "10044",
                            Region: "NY",
                            RegionCode: boltembeddedgo.String("NY"),
                            StreetAddress1: "888 main street",
                            StreetAddress2: boltembeddedgo.String("apt 3021"),
                            StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                            StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                        },
                        CartURL: boltembeddedgo.String("https://boltswagstore.com/orders/123456765432"),
                        Currency: "USD",
                        Discounts: []shared.CartDiscount{
                            shared.CartDiscount{
                                Amount: boltembeddedgo.Float64(100),
                                Code: boltembeddedgo.String("SUMMER10DISCOUNT"),
                                Description: boltembeddedgo.String("10% off our summer collection"),
                                DetailsURL: boltembeddedgo.String("https://boltswagstore.com/SUMMERSALE"),
                                Reference: boltembeddedgo.String("DISC-1234"),
                                Type: shared.TypePercentage.ToPointer(),
                            },
                        },
                        DisplayID: boltembeddedgo.String("displayid_100"),
                        Fees: []shared.Fees{
                            shared.Fees{
                                Description: boltembeddedgo.String("Item Fee"),
                                Name: "Item Fee",
                                Quantity: 7770.83,
                                Reference: "ItemFee",
                                UnitPrice: 7895.06,
                                UnitTaxAmount: 807.72,
                            },
                        },
                        Fulfillments: []shared.Fulfillments{
                            shared.Fulfillments{
                                CartItems: []shared.CartItem{
                                    shared.CartItem{
                                        Brand: boltembeddedgo.String("Bolt"),
                                        Category: boltembeddedgo.String("bags"),
                                        Collections: []string{
                                            "summer",
                                        },
                                        Color: boltembeddedgo.String("Bolt Blue"),
                                        Customizations: []shared.CartItemCustomization{
                                            shared.CartItemCustomization{
                                                Attributes: map[string]string{
                                                    "key1": "value1",
                                                    "key2": "value2",
                                                },
                                                Price: &shared.AmountView{
                                                    Amount: boltembeddedgo.Float64(754),
                                                    Currency: boltembeddedgo.String("USD"),
                                                    CurrencySymbol: boltembeddedgo.String("$"),
                                                },
                                            },
                                        },
                                        Description: boltembeddedgo.String("Large tote with Bolt logo."),
                                        DetailsURL: boltembeddedgo.String("https://boltswagstore.com/products/123456"),
                                        ExternalInputs: &shared.ICartItemExternalInputs{},
                                        GiftOption: &shared.GiftOption{
                                            Cost: boltembeddedgo.Int64(770),
                                            MerchantProductID: boltembeddedgo.String("881"),
                                            Message: boltembeddedgo.String("Happy Anniversary, Smoochy Poo!"),
                                            Wrap: boltembeddedgo.Bool(false),
                                        },
                                        ImageURL: boltembeddedgo.String("https://boltswagstore.com/products/123456/images/1.png"),
                                        Isbn: boltembeddedgo.String("9780091347314"),
                                        Manufacturer: boltembeddedgo.String("Bolt Textiles USA"),
                                        MerchantProductID: boltembeddedgo.String("881"),
                                        MerchantVariantID: boltembeddedgo.String("888"),
                                        Name: "Bolt Swag Bag",
                                        Options: boltembeddedgo.String("Special Edition"),
                                        Properties: []shared.CartItemProperty{
                                            shared.CartItemProperty{},
                                        },
                                        Quantity: 1,
                                        Reference: "item_100",
                                        Shipment: &shared.CartShipment{
                                            Carrier: boltembeddedgo.String("FedEx"),
                                            Cost: boltembeddedgo.Int64(770),
                                            DiscountedByMembership: boltembeddedgo.Bool(false),
                                            EstimatedDeliveryDate: boltembeddedgo.String("08-30-2022"),
                                            Expedited: boltembeddedgo.Bool(false),
                                            PackageDepth: boltembeddedgo.Int64(90),
                                            PackageDimensionUnit: boltembeddedgo.String("cm"),
                                            PackageHeight: boltembeddedgo.Int64(103),
                                            PackageType: boltembeddedgo.String("A big package."),
                                            PackageWeightUnit: boltembeddedgo.String("kg"),
                                            PackageWidth: boltembeddedgo.Int64(222),
                                            Service: boltembeddedgo.String("Option 1"),
                                            ShippingAddress: &shared.Address{
                                                Company: boltembeddedgo.String("Bolt"),
                                                Country: boltembeddedgo.String("United States"),
                                                CountryCode: "US",
                                                Default: boltembeddedgo.Bool(true),
                                                DoorCode: boltembeddedgo.String("123456"),
                                                Email: "alan.watts@example.com",
                                                FirstName: "Alan",
                                                LastName: "Watts",
                                                Locality: "Brooklyn",
                                                Name: boltembeddedgo.String("Alan Watts"),
                                                Phone: boltembeddedgo.String("+12125550199"),
                                                PostalCode: "10044",
                                                Region: "NY",
                                                RegionCode: boltembeddedgo.String("NY"),
                                                StreetAddress1: "888 main street",
                                                StreetAddress2: boltembeddedgo.String("apt 3021"),
                                                StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                                StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                            },
                                            ShippingAddressID: boltembeddedgo.String("addres-1"),
                                            ShippingMethod: boltembeddedgo.String("Unknown"),
                                            Signature: boltembeddedgo.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                            TaxAmount: boltembeddedgo.Int64(230),
                                            TaxCode: boltembeddedgo.String("tax-12345"),
                                            TotalWeight: boltembeddedgo.Int64(55),
                                            TotalWeightUnit: boltembeddedgo.String("kg"),
                                            Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                                        },
                                        Size: boltembeddedgo.String("Large"),
                                        Sku: boltembeddedgo.String("BOLT-SKU_100"),
                                        Tags: boltembeddedgo.String("tote, blue, linen, eco-friendly"),
                                        TaxAmount: boltembeddedgo.Float64(0),
                                        TotalAmount: 1000,
                                        UnitPrice: 1000,
                                        Uom: boltembeddedgo.String("inches"),
                                        Upc: boltembeddedgo.String("825764603119"),
                                        Weight: boltembeddedgo.Float64(10),
                                        WeightUnit: boltembeddedgo.String("pounds"),
                                    },
                                },
                                CartShipment: &shared.CartShipment{
                                    Carrier: boltembeddedgo.String("FedEx"),
                                    Cost: boltembeddedgo.Int64(770),
                                    DiscountedByMembership: boltembeddedgo.Bool(false),
                                    EstimatedDeliveryDate: boltembeddedgo.String("08-30-2022"),
                                    Expedited: boltembeddedgo.Bool(false),
                                    PackageDepth: boltembeddedgo.Int64(90),
                                    PackageDimensionUnit: boltembeddedgo.String("cm"),
                                    PackageHeight: boltembeddedgo.Int64(103),
                                    PackageType: boltembeddedgo.String("A big package."),
                                    PackageWeightUnit: boltembeddedgo.String("kg"),
                                    PackageWidth: boltembeddedgo.Int64(222),
                                    Service: boltembeddedgo.String("Option 1"),
                                    ShippingAddress: &shared.Address{
                                        Company: boltembeddedgo.String("Bolt"),
                                        Country: boltembeddedgo.String("United States"),
                                        CountryCode: "US",
                                        Default: boltembeddedgo.Bool(true),
                                        DoorCode: boltembeddedgo.String("123456"),
                                        Email: "alan.watts@example.com",
                                        FirstName: "Alan",
                                        LastName: "Watts",
                                        Locality: "Brooklyn",
                                        Name: boltembeddedgo.String("Alan Watts"),
                                        Phone: boltembeddedgo.String("+12125550199"),
                                        PostalCode: "10044",
                                        Region: "NY",
                                        RegionCode: boltembeddedgo.String("NY"),
                                        StreetAddress1: "888 main street",
                                        StreetAddress2: boltembeddedgo.String("apt 3021"),
                                        StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                        StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                    },
                                    ShippingAddressID: boltembeddedgo.String("addres-1"),
                                    ShippingMethod: boltembeddedgo.String("Unknown"),
                                    Signature: boltembeddedgo.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                    TaxAmount: boltembeddedgo.Int64(230),
                                    TaxCode: boltembeddedgo.String("tax-12345"),
                                    TotalWeight: boltembeddedgo.Int64(55),
                                    TotalWeightUnit: boltembeddedgo.String("kg"),
                                    Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                                },
                                DigitalDelivery: &shared.DigitalDelivery{},
                                InStoreCartShipment: &shared.InStoreCartShipment{
                                    CartShipment: &shared.CartShipment{
                                        Carrier: boltembeddedgo.String("FedEx"),
                                        Cost: boltembeddedgo.Int64(770),
                                        DiscountedByMembership: boltembeddedgo.Bool(false),
                                        EstimatedDeliveryDate: boltembeddedgo.String("08-30-2022"),
                                        Expedited: boltembeddedgo.Bool(false),
                                        PackageDepth: boltembeddedgo.Int64(90),
                                        PackageDimensionUnit: boltembeddedgo.String("cm"),
                                        PackageHeight: boltembeddedgo.Int64(103),
                                        PackageType: boltembeddedgo.String("A big package."),
                                        PackageWeightUnit: boltembeddedgo.String("kg"),
                                        PackageWidth: boltembeddedgo.Int64(222),
                                        Service: boltembeddedgo.String("Option 1"),
                                        ShippingAddress: &shared.Address{
                                            Company: boltembeddedgo.String("Bolt"),
                                            Country: boltembeddedgo.String("United States"),
                                            CountryCode: "US",
                                            Default: boltembeddedgo.Bool(true),
                                            DoorCode: boltembeddedgo.String("123456"),
                                            Email: "alan.watts@example.com",
                                            FirstName: "Alan",
                                            LastName: "Watts",
                                            Locality: "Brooklyn",
                                            Name: boltembeddedgo.String("Alan Watts"),
                                            Phone: boltembeddedgo.String("+12125550199"),
                                            PostalCode: "10044",
                                            Region: "NY",
                                            RegionCode: boltembeddedgo.String("NY"),
                                            StreetAddress1: "888 main street",
                                            StreetAddress2: boltembeddedgo.String("apt 3021"),
                                            StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                            StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                        },
                                        ShippingAddressID: boltembeddedgo.String("addres-1"),
                                        ShippingMethod: boltembeddedgo.String("Unknown"),
                                        Signature: boltembeddedgo.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                        TaxAmount: boltembeddedgo.Int64(230),
                                        TaxCode: boltembeddedgo.String("tax-12345"),
                                        TotalWeight: boltembeddedgo.Int64(55),
                                        TotalWeightUnit: boltembeddedgo.String("kg"),
                                        Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                                    },
                                    Description: boltembeddedgo.String("Pick up in-store at 123 Main St."),
                                    Distance: boltembeddedgo.Float64(3),
                                    DistanceUnit: shared.DistanceUnitMile.ToPointer(),
                                    InStorePickupAddress: &shared.Address{
                                        Company: boltembeddedgo.String("Bolt"),
                                        Country: boltembeddedgo.String("United States"),
                                        CountryCode: "US",
                                        Default: boltembeddedgo.Bool(true),
                                        DoorCode: boltembeddedgo.String("123456"),
                                        Email: "alan.watts@example.com",
                                        FirstName: "Alan",
                                        LastName: "Watts",
                                        Locality: "Brooklyn",
                                        Name: boltembeddedgo.String("Alan Watts"),
                                        Phone: boltembeddedgo.String("+12125550199"),
                                        PostalCode: "10044",
                                        Region: "NY",
                                        RegionCode: boltembeddedgo.String("NY"),
                                        StreetAddress1: "888 main street",
                                        StreetAddress2: boltembeddedgo.String("apt 3021"),
                                        StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                        StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                    },
                                    StoreName: boltembeddedgo.String("Bolt Collective"),
                                },
                            },
                        },
                        InStoreCartShipments: []shared.InStoreCartShipment{
                            shared.InStoreCartShipment{
                                CartShipment: &shared.CartShipment{
                                    Carrier: boltembeddedgo.String("FedEx"),
                                    Cost: boltembeddedgo.Int64(770),
                                    DiscountedByMembership: boltembeddedgo.Bool(false),
                                    EstimatedDeliveryDate: boltembeddedgo.String("08-30-2022"),
                                    Expedited: boltembeddedgo.Bool(false),
                                    PackageDepth: boltembeddedgo.Int64(90),
                                    PackageDimensionUnit: boltembeddedgo.String("cm"),
                                    PackageHeight: boltembeddedgo.Int64(103),
                                    PackageType: boltembeddedgo.String("A big package."),
                                    PackageWeightUnit: boltembeddedgo.String("kg"),
                                    PackageWidth: boltembeddedgo.Int64(222),
                                    Service: boltembeddedgo.String("Option 1"),
                                    ShippingAddress: &shared.Address{
                                        Company: boltembeddedgo.String("Bolt"),
                                        Country: boltembeddedgo.String("United States"),
                                        CountryCode: "US",
                                        Default: boltembeddedgo.Bool(true),
                                        DoorCode: boltembeddedgo.String("123456"),
                                        Email: "alan.watts@example.com",
                                        FirstName: "Alan",
                                        LastName: "Watts",
                                        Locality: "Brooklyn",
                                        Name: boltembeddedgo.String("Alan Watts"),
                                        Phone: boltembeddedgo.String("+12125550199"),
                                        PostalCode: "10044",
                                        Region: "NY",
                                        RegionCode: boltembeddedgo.String("NY"),
                                        StreetAddress1: "888 main street",
                                        StreetAddress2: boltembeddedgo.String("apt 3021"),
                                        StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                        StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                    },
                                    ShippingAddressID: boltembeddedgo.String("addres-1"),
                                    ShippingMethod: boltembeddedgo.String("Unknown"),
                                    Signature: boltembeddedgo.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                    TaxAmount: boltembeddedgo.Int64(230),
                                    TaxCode: boltembeddedgo.String("tax-12345"),
                                    TotalWeight: boltembeddedgo.Int64(55),
                                    TotalWeightUnit: boltembeddedgo.String("kg"),
                                    Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                                },
                                Description: boltembeddedgo.String("Pick up in-store at 123 Main St."),
                                Distance: boltembeddedgo.Float64(3),
                                DistanceUnit: shared.DistanceUnitMile.ToPointer(),
                                InStorePickupAddress: &shared.Address{
                                    Company: boltembeddedgo.String("Bolt"),
                                    Country: boltembeddedgo.String("United States"),
                                    CountryCode: "US",
                                    Default: boltembeddedgo.Bool(true),
                                    DoorCode: boltembeddedgo.String("123456"),
                                    Email: "alan.watts@example.com",
                                    FirstName: "Alan",
                                    LastName: "Watts",
                                    Locality: "Brooklyn",
                                    Name: boltembeddedgo.String("Alan Watts"),
                                    Phone: boltembeddedgo.String("+12125550199"),
                                    PostalCode: "10044",
                                    Region: "NY",
                                    RegionCode: boltembeddedgo.String("NY"),
                                    StreetAddress1: "888 main street",
                                    StreetAddress2: boltembeddedgo.String("apt 3021"),
                                    StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                    StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                },
                                StoreName: boltembeddedgo.String("Bolt Collective"),
                            },
                        },
                        Items: []shared.CartItem{
                            shared.CartItem{
                                Brand: boltembeddedgo.String("Bolt"),
                                Category: boltembeddedgo.String("bags"),
                                Collections: []string{
                                    "summer",
                                },
                                Color: boltembeddedgo.String("Bolt Blue"),
                                Customizations: []shared.CartItemCustomization{
                                    shared.CartItemCustomization{
                                        Attributes: map[string]string{
                                            "key1": "value1",
                                            "key2": "value2",
                                        },
                                        Price: &shared.AmountView{
                                            Amount: boltembeddedgo.Float64(754),
                                            Currency: boltembeddedgo.String("USD"),
                                            CurrencySymbol: boltembeddedgo.String("$"),
                                        },
                                    },
                                },
                                Description: boltembeddedgo.String("Large tote with Bolt logo."),
                                DetailsURL: boltembeddedgo.String("https://boltswagstore.com/products/123456"),
                                ExternalInputs: &shared.ICartItemExternalInputs{},
                                GiftOption: &shared.GiftOption{
                                    Cost: boltembeddedgo.Int64(770),
                                    MerchantProductID: boltembeddedgo.String("881"),
                                    Message: boltembeddedgo.String("Happy Anniversary, Smoochy Poo!"),
                                    Wrap: boltembeddedgo.Bool(false),
                                },
                                ImageURL: boltembeddedgo.String("https://boltswagstore.com/products/123456/images/1.png"),
                                Isbn: boltembeddedgo.String("9780091347314"),
                                Manufacturer: boltembeddedgo.String("Bolt Textiles USA"),
                                MerchantProductID: boltembeddedgo.String("881"),
                                MerchantVariantID: boltembeddedgo.String("888"),
                                Name: "Bolt Swag Bag",
                                Options: boltembeddedgo.String("Special Edition"),
                                Properties: []shared.CartItemProperty{
                                    shared.CartItemProperty{},
                                },
                                Quantity: 1,
                                Reference: "item_100",
                                Shipment: &shared.CartShipment{
                                    Carrier: boltembeddedgo.String("FedEx"),
                                    Cost: boltembeddedgo.Int64(770),
                                    DiscountedByMembership: boltembeddedgo.Bool(false),
                                    EstimatedDeliveryDate: boltembeddedgo.String("08-30-2022"),
                                    Expedited: boltembeddedgo.Bool(false),
                                    PackageDepth: boltembeddedgo.Int64(90),
                                    PackageDimensionUnit: boltembeddedgo.String("cm"),
                                    PackageHeight: boltembeddedgo.Int64(103),
                                    PackageType: boltembeddedgo.String("A big package."),
                                    PackageWeightUnit: boltembeddedgo.String("kg"),
                                    PackageWidth: boltembeddedgo.Int64(222),
                                    Service: boltembeddedgo.String("Option 1"),
                                    ShippingAddress: &shared.Address{
                                        Company: boltembeddedgo.String("Bolt"),
                                        Country: boltembeddedgo.String("United States"),
                                        CountryCode: "US",
                                        Default: boltembeddedgo.Bool(true),
                                        DoorCode: boltembeddedgo.String("123456"),
                                        Email: "alan.watts@example.com",
                                        FirstName: "Alan",
                                        LastName: "Watts",
                                        Locality: "Brooklyn",
                                        Name: boltembeddedgo.String("Alan Watts"),
                                        Phone: boltembeddedgo.String("+12125550199"),
                                        PostalCode: "10044",
                                        Region: "NY",
                                        RegionCode: boltembeddedgo.String("NY"),
                                        StreetAddress1: "888 main street",
                                        StreetAddress2: boltembeddedgo.String("apt 3021"),
                                        StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                        StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                    },
                                    ShippingAddressID: boltembeddedgo.String("addres-1"),
                                    ShippingMethod: boltembeddedgo.String("Unknown"),
                                    Signature: boltembeddedgo.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                    TaxAmount: boltembeddedgo.Int64(230),
                                    TaxCode: boltembeddedgo.String("tax-12345"),
                                    TotalWeight: boltembeddedgo.Int64(55),
                                    TotalWeightUnit: boltembeddedgo.String("kg"),
                                    Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                                },
                                Size: boltembeddedgo.String("Large"),
                                Sku: boltembeddedgo.String("BOLT-SKU_100"),
                                Tags: boltembeddedgo.String("tote, blue, linen, eco-friendly"),
                                TaxAmount: boltembeddedgo.Float64(0),
                                TotalAmount: 1000,
                                UnitPrice: 1000,
                                Uom: boltembeddedgo.String("inches"),
                                Upc: boltembeddedgo.String("825764603119"),
                                Weight: boltembeddedgo.Float64(10),
                                WeightUnit: boltembeddedgo.String("pounds"),
                            },
                        },
                        LoyaltyRewards: []shared.CartLoyaltyRewards{
                            shared.CartLoyaltyRewards{
                                Description: boltembeddedgo.String("$5 off (100 Points)"),
                                Details: boltembeddedgo.String("{\"id\": 123456, \"icon\": \"fa-dollar\", \"name\": \"$15.00 Off\", \"type\": \"Coupon\", \"amount\": 100, \"duration\": \"single_use\", \"cost_text\": \"150 Points\",  \"description\": \"Get $15 off your next purchase for 150 points\", \"discount_type\": \"fixed_amount\", \"unrendered_name\": \"$15.00 Off\",  \"discount_percentage\": null, \"discount_rate_cents\": null, \"discount_value_cents\": null, \"discount_amount_cents\": 1500,  \"unrendered_description\": \"Get $15 off your next purchase for 150 points\", \"applies_to_product_type\": \"ALL\"}"),
                            },
                        },
                        Metadata: map[string]string{
                            "key1": "value1",
                            "key2": "value2",
                        },
                        OrderDescription: boltembeddedgo.String("Order #1234567890"),
                        OrderReference: "order_100",
                        Shipments: []shared.CartShipment{
                            shared.CartShipment{
                                Carrier: boltembeddedgo.String("FedEx"),
                                Cost: boltembeddedgo.Int64(770),
                                DiscountedByMembership: boltembeddedgo.Bool(false),
                                EstimatedDeliveryDate: boltembeddedgo.String("08-30-2022"),
                                Expedited: boltembeddedgo.Bool(false),
                                PackageDepth: boltembeddedgo.Int64(90),
                                PackageDimensionUnit: boltembeddedgo.String("cm"),
                                PackageHeight: boltembeddedgo.Int64(103),
                                PackageType: boltembeddedgo.String("A big package."),
                                PackageWeightUnit: boltembeddedgo.String("kg"),
                                PackageWidth: boltembeddedgo.Int64(222),
                                Service: boltembeddedgo.String("Option 1"),
                                ShippingAddress: &shared.Address{
                                    Company: boltembeddedgo.String("Bolt"),
                                    Country: boltembeddedgo.String("United States"),
                                    CountryCode: "US",
                                    Default: boltembeddedgo.Bool(true),
                                    DoorCode: boltembeddedgo.String("123456"),
                                    Email: "alan.watts@example.com",
                                    FirstName: "Alan",
                                    LastName: "Watts",
                                    Locality: "Brooklyn",
                                    Name: boltembeddedgo.String("Alan Watts"),
                                    Phone: boltembeddedgo.String("+12125550199"),
                                    PostalCode: "10044",
                                    Region: "NY",
                                    RegionCode: boltembeddedgo.String("NY"),
                                    StreetAddress1: "888 main street",
                                    StreetAddress2: boltembeddedgo.String("apt 3021"),
                                    StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                    StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                },
                                ShippingAddressID: boltembeddedgo.String("addres-1"),
                                ShippingMethod: boltembeddedgo.String("Unknown"),
                                Signature: boltembeddedgo.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                TaxAmount: boltembeddedgo.Int64(230),
                                TaxCode: boltembeddedgo.String("tax-12345"),
                                TotalWeight: boltembeddedgo.Int64(55),
                                TotalWeightUnit: boltembeddedgo.String("kg"),
                                Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                            },
                        },
                        TotalAmount: 900,
                    },
                    PaymentMethod: operations.PaymentMethod{
                        PaymentData: operations.PaymentData{},
                        Type: operations.TypePaypal,
                    },
                    ShopperIdentity: &operations.PaypalPaymentInputInitializeShopperIdentity{
                        CreateBoltAccount: boltembeddedgo.Bool(true),
                        Email: "Kathryn_DAmore@gmail.com",
                        FirstName: "Anita",
                        LastName: "Schroeder",
                        Phone: "544-741-5984 x087",
                    },
                },
        ),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.InitializePaymentRequest](../../pkg/models/operations/initializepaymentrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.InitializePaymentSecurity](../../pkg/models/operations/initializepaymentsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.InitializePaymentResponse](../../pkg/models/operations/initializepaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdatePayment

Update a Bolt payment using the token given after initializing a payment.  Updates will completely replace the original top-level resource (for example, if a cart is sent in with the request it will replace the existing cart).  Any included object should be sent as complete object. NOTE: The authorization header is NOT required for payments associated with users who do not have a Bolt account.


### Example Usage

```go
package main

import(
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"context"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.UpdatePaymentSecurity{
            OAuth: boltembeddedgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }

    ctx := context.Background()
    res, err := s.Payments.UpdatePayment(ctx, operations.UpdatePaymentRequest{
        RequestBody: operations.CreateUpdatePaymentRequestBodySavedPaymentInputUpdateSchemas(
                operations.SavedPaymentInputUpdateSchemas{
                    Cart: &shared.CartCreate{
                        AddOns: []shared.CartAddOn{
                            shared.CartAddOn{
                                Name: "string",
                                Price: 1095.6,
                                ProductID: "string",
                            },
                        },
                        BillingAddress: &shared.Address{
                            Company: boltembeddedgo.String("Bolt"),
                            Country: boltembeddedgo.String("United States"),
                            CountryCode: "US",
                            Default: boltembeddedgo.Bool(true),
                            DoorCode: boltembeddedgo.String("123456"),
                            Email: "alan.watts@example.com",
                            FirstName: "Alan",
                            LastName: "Watts",
                            Locality: "Brooklyn",
                            Name: boltembeddedgo.String("Alan Watts"),
                            Phone: boltembeddedgo.String("+12125550199"),
                            PostalCode: "10044",
                            Region: "NY",
                            RegionCode: boltembeddedgo.String("NY"),
                            StreetAddress1: "888 main street",
                            StreetAddress2: boltembeddedgo.String("apt 3021"),
                            StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                            StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                        },
                        CartURL: boltembeddedgo.String("https://boltswagstore.com/orders/123456765432"),
                        Currency: "USD",
                        Discounts: []shared.CartDiscount{
                            shared.CartDiscount{
                                Amount: boltembeddedgo.Float64(100),
                                Code: boltembeddedgo.String("SUMMER10DISCOUNT"),
                                Description: boltembeddedgo.String("10% off our summer collection"),
                                DetailsURL: boltembeddedgo.String("https://boltswagstore.com/SUMMERSALE"),
                                Reference: boltembeddedgo.String("DISC-1234"),
                                Type: shared.TypePercentage.ToPointer(),
                            },
                        },
                        DisplayID: boltembeddedgo.String("displayid_100"),
                        Fees: []shared.Fees{
                            shared.Fees{
                                Description: boltembeddedgo.String("Item Fee"),
                                Name: "Item Fee",
                                Quantity: 2053.88,
                                Reference: "ItemFee",
                                UnitPrice: 4201.73,
                                UnitTaxAmount: 5393.29,
                            },
                        },
                        Fulfillments: []shared.Fulfillments{
                            shared.Fulfillments{
                                CartItems: []shared.CartItem{
                                    shared.CartItem{
                                        Brand: boltembeddedgo.String("Bolt"),
                                        Category: boltembeddedgo.String("bags"),
                                        Collections: []string{
                                            "summer",
                                        },
                                        Color: boltembeddedgo.String("Bolt Blue"),
                                        Customizations: []shared.CartItemCustomization{
                                            shared.CartItemCustomization{
                                                Attributes: map[string]string{
                                                    "key1": "value1",
                                                    "key2": "value2",
                                                },
                                                Price: &shared.AmountView{
                                                    Amount: boltembeddedgo.Float64(754),
                                                    Currency: boltembeddedgo.String("USD"),
                                                    CurrencySymbol: boltembeddedgo.String("$"),
                                                },
                                            },
                                        },
                                        Description: boltembeddedgo.String("Large tote with Bolt logo."),
                                        DetailsURL: boltembeddedgo.String("https://boltswagstore.com/products/123456"),
                                        ExternalInputs: &shared.ICartItemExternalInputs{},
                                        GiftOption: &shared.GiftOption{
                                            Cost: boltembeddedgo.Int64(770),
                                            MerchantProductID: boltembeddedgo.String("881"),
                                            Message: boltembeddedgo.String("Happy Anniversary, Smoochy Poo!"),
                                            Wrap: boltembeddedgo.Bool(false),
                                        },
                                        ImageURL: boltembeddedgo.String("https://boltswagstore.com/products/123456/images/1.png"),
                                        Isbn: boltembeddedgo.String("9780091347314"),
                                        Manufacturer: boltembeddedgo.String("Bolt Textiles USA"),
                                        MerchantProductID: boltembeddedgo.String("881"),
                                        MerchantVariantID: boltembeddedgo.String("888"),
                                        Name: "Bolt Swag Bag",
                                        Options: boltembeddedgo.String("Special Edition"),
                                        Properties: []shared.CartItemProperty{
                                            shared.CartItemProperty{},
                                        },
                                        Quantity: 1,
                                        Reference: "item_100",
                                        Shipment: &shared.CartShipment{
                                            Carrier: boltembeddedgo.String("FedEx"),
                                            Cost: boltembeddedgo.Int64(770),
                                            DiscountedByMembership: boltembeddedgo.Bool(false),
                                            EstimatedDeliveryDate: boltembeddedgo.String("08-30-2022"),
                                            Expedited: boltembeddedgo.Bool(false),
                                            PackageDepth: boltembeddedgo.Int64(90),
                                            PackageDimensionUnit: boltembeddedgo.String("cm"),
                                            PackageHeight: boltembeddedgo.Int64(103),
                                            PackageType: boltembeddedgo.String("A big package."),
                                            PackageWeightUnit: boltembeddedgo.String("kg"),
                                            PackageWidth: boltembeddedgo.Int64(222),
                                            Service: boltembeddedgo.String("Option 1"),
                                            ShippingAddress: &shared.Address{
                                                Company: boltembeddedgo.String("Bolt"),
                                                Country: boltembeddedgo.String("United States"),
                                                CountryCode: "US",
                                                Default: boltembeddedgo.Bool(true),
                                                DoorCode: boltembeddedgo.String("123456"),
                                                Email: "alan.watts@example.com",
                                                FirstName: "Alan",
                                                LastName: "Watts",
                                                Locality: "Brooklyn",
                                                Name: boltembeddedgo.String("Alan Watts"),
                                                Phone: boltembeddedgo.String("+12125550199"),
                                                PostalCode: "10044",
                                                Region: "NY",
                                                RegionCode: boltembeddedgo.String("NY"),
                                                StreetAddress1: "888 main street",
                                                StreetAddress2: boltembeddedgo.String("apt 3021"),
                                                StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                                StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                            },
                                            ShippingAddressID: boltembeddedgo.String("addres-1"),
                                            ShippingMethod: boltembeddedgo.String("Unknown"),
                                            Signature: boltembeddedgo.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                            TaxAmount: boltembeddedgo.Int64(230),
                                            TaxCode: boltembeddedgo.String("tax-12345"),
                                            TotalWeight: boltembeddedgo.Int64(55),
                                            TotalWeightUnit: boltembeddedgo.String("kg"),
                                            Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                                        },
                                        Size: boltembeddedgo.String("Large"),
                                        Sku: boltembeddedgo.String("BOLT-SKU_100"),
                                        Tags: boltembeddedgo.String("tote, blue, linen, eco-friendly"),
                                        TaxAmount: boltembeddedgo.Float64(0),
                                        TotalAmount: 1000,
                                        UnitPrice: 1000,
                                        Uom: boltembeddedgo.String("inches"),
                                        Upc: boltembeddedgo.String("825764603119"),
                                        Weight: boltembeddedgo.Float64(10),
                                        WeightUnit: boltembeddedgo.String("pounds"),
                                    },
                                },
                                CartShipment: &shared.CartShipment{
                                    Carrier: boltembeddedgo.String("FedEx"),
                                    Cost: boltembeddedgo.Int64(770),
                                    DiscountedByMembership: boltembeddedgo.Bool(false),
                                    EstimatedDeliveryDate: boltembeddedgo.String("08-30-2022"),
                                    Expedited: boltembeddedgo.Bool(false),
                                    PackageDepth: boltembeddedgo.Int64(90),
                                    PackageDimensionUnit: boltembeddedgo.String("cm"),
                                    PackageHeight: boltembeddedgo.Int64(103),
                                    PackageType: boltembeddedgo.String("A big package."),
                                    PackageWeightUnit: boltembeddedgo.String("kg"),
                                    PackageWidth: boltembeddedgo.Int64(222),
                                    Service: boltembeddedgo.String("Option 1"),
                                    ShippingAddress: &shared.Address{
                                        Company: boltembeddedgo.String("Bolt"),
                                        Country: boltembeddedgo.String("United States"),
                                        CountryCode: "US",
                                        Default: boltembeddedgo.Bool(true),
                                        DoorCode: boltembeddedgo.String("123456"),
                                        Email: "alan.watts@example.com",
                                        FirstName: "Alan",
                                        LastName: "Watts",
                                        Locality: "Brooklyn",
                                        Name: boltembeddedgo.String("Alan Watts"),
                                        Phone: boltembeddedgo.String("+12125550199"),
                                        PostalCode: "10044",
                                        Region: "NY",
                                        RegionCode: boltembeddedgo.String("NY"),
                                        StreetAddress1: "888 main street",
                                        StreetAddress2: boltembeddedgo.String("apt 3021"),
                                        StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                        StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                    },
                                    ShippingAddressID: boltembeddedgo.String("addres-1"),
                                    ShippingMethod: boltembeddedgo.String("Unknown"),
                                    Signature: boltembeddedgo.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                    TaxAmount: boltembeddedgo.Int64(230),
                                    TaxCode: boltembeddedgo.String("tax-12345"),
                                    TotalWeight: boltembeddedgo.Int64(55),
                                    TotalWeightUnit: boltembeddedgo.String("kg"),
                                    Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                                },
                                DigitalDelivery: &shared.DigitalDelivery{},
                                InStoreCartShipment: &shared.InStoreCartShipment{
                                    CartShipment: &shared.CartShipment{
                                        Carrier: boltembeddedgo.String("FedEx"),
                                        Cost: boltembeddedgo.Int64(770),
                                        DiscountedByMembership: boltembeddedgo.Bool(false),
                                        EstimatedDeliveryDate: boltembeddedgo.String("08-30-2022"),
                                        Expedited: boltembeddedgo.Bool(false),
                                        PackageDepth: boltembeddedgo.Int64(90),
                                        PackageDimensionUnit: boltembeddedgo.String("cm"),
                                        PackageHeight: boltembeddedgo.Int64(103),
                                        PackageType: boltembeddedgo.String("A big package."),
                                        PackageWeightUnit: boltembeddedgo.String("kg"),
                                        PackageWidth: boltembeddedgo.Int64(222),
                                        Service: boltembeddedgo.String("Option 1"),
                                        ShippingAddress: &shared.Address{
                                            Company: boltembeddedgo.String("Bolt"),
                                            Country: boltembeddedgo.String("United States"),
                                            CountryCode: "US",
                                            Default: boltembeddedgo.Bool(true),
                                            DoorCode: boltembeddedgo.String("123456"),
                                            Email: "alan.watts@example.com",
                                            FirstName: "Alan",
                                            LastName: "Watts",
                                            Locality: "Brooklyn",
                                            Name: boltembeddedgo.String("Alan Watts"),
                                            Phone: boltembeddedgo.String("+12125550199"),
                                            PostalCode: "10044",
                                            Region: "NY",
                                            RegionCode: boltembeddedgo.String("NY"),
                                            StreetAddress1: "888 main street",
                                            StreetAddress2: boltembeddedgo.String("apt 3021"),
                                            StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                            StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                        },
                                        ShippingAddressID: boltembeddedgo.String("addres-1"),
                                        ShippingMethod: boltembeddedgo.String("Unknown"),
                                        Signature: boltembeddedgo.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                        TaxAmount: boltembeddedgo.Int64(230),
                                        TaxCode: boltembeddedgo.String("tax-12345"),
                                        TotalWeight: boltembeddedgo.Int64(55),
                                        TotalWeightUnit: boltembeddedgo.String("kg"),
                                        Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                                    },
                                    Description: boltembeddedgo.String("Pick up in-store at 123 Main St."),
                                    Distance: boltembeddedgo.Float64(3),
                                    DistanceUnit: shared.DistanceUnitMile.ToPointer(),
                                    InStorePickupAddress: &shared.Address{
                                        Company: boltembeddedgo.String("Bolt"),
                                        Country: boltembeddedgo.String("United States"),
                                        CountryCode: "US",
                                        Default: boltembeddedgo.Bool(true),
                                        DoorCode: boltembeddedgo.String("123456"),
                                        Email: "alan.watts@example.com",
                                        FirstName: "Alan",
                                        LastName: "Watts",
                                        Locality: "Brooklyn",
                                        Name: boltembeddedgo.String("Alan Watts"),
                                        Phone: boltembeddedgo.String("+12125550199"),
                                        PostalCode: "10044",
                                        Region: "NY",
                                        RegionCode: boltembeddedgo.String("NY"),
                                        StreetAddress1: "888 main street",
                                        StreetAddress2: boltembeddedgo.String("apt 3021"),
                                        StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                        StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                    },
                                    StoreName: boltembeddedgo.String("Bolt Collective"),
                                },
                            },
                        },
                        InStoreCartShipments: []shared.InStoreCartShipment{
                            shared.InStoreCartShipment{
                                CartShipment: &shared.CartShipment{
                                    Carrier: boltembeddedgo.String("FedEx"),
                                    Cost: boltembeddedgo.Int64(770),
                                    DiscountedByMembership: boltembeddedgo.Bool(false),
                                    EstimatedDeliveryDate: boltembeddedgo.String("08-30-2022"),
                                    Expedited: boltembeddedgo.Bool(false),
                                    PackageDepth: boltembeddedgo.Int64(90),
                                    PackageDimensionUnit: boltembeddedgo.String("cm"),
                                    PackageHeight: boltembeddedgo.Int64(103),
                                    PackageType: boltembeddedgo.String("A big package."),
                                    PackageWeightUnit: boltembeddedgo.String("kg"),
                                    PackageWidth: boltembeddedgo.Int64(222),
                                    Service: boltembeddedgo.String("Option 1"),
                                    ShippingAddress: &shared.Address{
                                        Company: boltembeddedgo.String("Bolt"),
                                        Country: boltembeddedgo.String("United States"),
                                        CountryCode: "US",
                                        Default: boltembeddedgo.Bool(true),
                                        DoorCode: boltembeddedgo.String("123456"),
                                        Email: "alan.watts@example.com",
                                        FirstName: "Alan",
                                        LastName: "Watts",
                                        Locality: "Brooklyn",
                                        Name: boltembeddedgo.String("Alan Watts"),
                                        Phone: boltembeddedgo.String("+12125550199"),
                                        PostalCode: "10044",
                                        Region: "NY",
                                        RegionCode: boltembeddedgo.String("NY"),
                                        StreetAddress1: "888 main street",
                                        StreetAddress2: boltembeddedgo.String("apt 3021"),
                                        StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                        StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                    },
                                    ShippingAddressID: boltembeddedgo.String("addres-1"),
                                    ShippingMethod: boltembeddedgo.String("Unknown"),
                                    Signature: boltembeddedgo.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                    TaxAmount: boltembeddedgo.Int64(230),
                                    TaxCode: boltembeddedgo.String("tax-12345"),
                                    TotalWeight: boltembeddedgo.Int64(55),
                                    TotalWeightUnit: boltembeddedgo.String("kg"),
                                    Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                                },
                                Description: boltembeddedgo.String("Pick up in-store at 123 Main St."),
                                Distance: boltembeddedgo.Float64(3),
                                DistanceUnit: shared.DistanceUnitMile.ToPointer(),
                                InStorePickupAddress: &shared.Address{
                                    Company: boltembeddedgo.String("Bolt"),
                                    Country: boltembeddedgo.String("United States"),
                                    CountryCode: "US",
                                    Default: boltembeddedgo.Bool(true),
                                    DoorCode: boltembeddedgo.String("123456"),
                                    Email: "alan.watts@example.com",
                                    FirstName: "Alan",
                                    LastName: "Watts",
                                    Locality: "Brooklyn",
                                    Name: boltembeddedgo.String("Alan Watts"),
                                    Phone: boltembeddedgo.String("+12125550199"),
                                    PostalCode: "10044",
                                    Region: "NY",
                                    RegionCode: boltembeddedgo.String("NY"),
                                    StreetAddress1: "888 main street",
                                    StreetAddress2: boltembeddedgo.String("apt 3021"),
                                    StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                    StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                },
                                StoreName: boltembeddedgo.String("Bolt Collective"),
                            },
                        },
                        Items: []shared.CartItem{
                            shared.CartItem{
                                Brand: boltembeddedgo.String("Bolt"),
                                Category: boltembeddedgo.String("bags"),
                                Collections: []string{
                                    "summer",
                                },
                                Color: boltembeddedgo.String("Bolt Blue"),
                                Customizations: []shared.CartItemCustomization{
                                    shared.CartItemCustomization{
                                        Attributes: map[string]string{
                                            "key1": "value1",
                                            "key2": "value2",
                                        },
                                        Price: &shared.AmountView{
                                            Amount: boltembeddedgo.Float64(754),
                                            Currency: boltembeddedgo.String("USD"),
                                            CurrencySymbol: boltembeddedgo.String("$"),
                                        },
                                    },
                                },
                                Description: boltembeddedgo.String("Large tote with Bolt logo."),
                                DetailsURL: boltembeddedgo.String("https://boltswagstore.com/products/123456"),
                                ExternalInputs: &shared.ICartItemExternalInputs{},
                                GiftOption: &shared.GiftOption{
                                    Cost: boltembeddedgo.Int64(770),
                                    MerchantProductID: boltembeddedgo.String("881"),
                                    Message: boltembeddedgo.String("Happy Anniversary, Smoochy Poo!"),
                                    Wrap: boltembeddedgo.Bool(false),
                                },
                                ImageURL: boltembeddedgo.String("https://boltswagstore.com/products/123456/images/1.png"),
                                Isbn: boltembeddedgo.String("9780091347314"),
                                Manufacturer: boltembeddedgo.String("Bolt Textiles USA"),
                                MerchantProductID: boltembeddedgo.String("881"),
                                MerchantVariantID: boltembeddedgo.String("888"),
                                Name: "Bolt Swag Bag",
                                Options: boltembeddedgo.String("Special Edition"),
                                Properties: []shared.CartItemProperty{
                                    shared.CartItemProperty{},
                                },
                                Quantity: 1,
                                Reference: "item_100",
                                Shipment: &shared.CartShipment{
                                    Carrier: boltembeddedgo.String("FedEx"),
                                    Cost: boltembeddedgo.Int64(770),
                                    DiscountedByMembership: boltembeddedgo.Bool(false),
                                    EstimatedDeliveryDate: boltembeddedgo.String("08-30-2022"),
                                    Expedited: boltembeddedgo.Bool(false),
                                    PackageDepth: boltembeddedgo.Int64(90),
                                    PackageDimensionUnit: boltembeddedgo.String("cm"),
                                    PackageHeight: boltembeddedgo.Int64(103),
                                    PackageType: boltembeddedgo.String("A big package."),
                                    PackageWeightUnit: boltembeddedgo.String("kg"),
                                    PackageWidth: boltembeddedgo.Int64(222),
                                    Service: boltembeddedgo.String("Option 1"),
                                    ShippingAddress: &shared.Address{
                                        Company: boltembeddedgo.String("Bolt"),
                                        Country: boltembeddedgo.String("United States"),
                                        CountryCode: "US",
                                        Default: boltembeddedgo.Bool(true),
                                        DoorCode: boltembeddedgo.String("123456"),
                                        Email: "alan.watts@example.com",
                                        FirstName: "Alan",
                                        LastName: "Watts",
                                        Locality: "Brooklyn",
                                        Name: boltembeddedgo.String("Alan Watts"),
                                        Phone: boltembeddedgo.String("+12125550199"),
                                        PostalCode: "10044",
                                        Region: "NY",
                                        RegionCode: boltembeddedgo.String("NY"),
                                        StreetAddress1: "888 main street",
                                        StreetAddress2: boltembeddedgo.String("apt 3021"),
                                        StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                        StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                    },
                                    ShippingAddressID: boltembeddedgo.String("addres-1"),
                                    ShippingMethod: boltembeddedgo.String("Unknown"),
                                    Signature: boltembeddedgo.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                    TaxAmount: boltembeddedgo.Int64(230),
                                    TaxCode: boltembeddedgo.String("tax-12345"),
                                    TotalWeight: boltembeddedgo.Int64(55),
                                    TotalWeightUnit: boltembeddedgo.String("kg"),
                                    Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                                },
                                Size: boltembeddedgo.String("Large"),
                                Sku: boltembeddedgo.String("BOLT-SKU_100"),
                                Tags: boltembeddedgo.String("tote, blue, linen, eco-friendly"),
                                TaxAmount: boltembeddedgo.Float64(0),
                                TotalAmount: 1000,
                                UnitPrice: 1000,
                                Uom: boltembeddedgo.String("inches"),
                                Upc: boltembeddedgo.String("825764603119"),
                                Weight: boltembeddedgo.Float64(10),
                                WeightUnit: boltembeddedgo.String("pounds"),
                            },
                        },
                        LoyaltyRewards: []shared.CartLoyaltyRewards{
                            shared.CartLoyaltyRewards{
                                Description: boltembeddedgo.String("$5 off (100 Points)"),
                                Details: boltembeddedgo.String("{\"id\": 123456, \"icon\": \"fa-dollar\", \"name\": \"$15.00 Off\", \"type\": \"Coupon\", \"amount\": 100, \"duration\": \"single_use\", \"cost_text\": \"150 Points\",  \"description\": \"Get $15 off your next purchase for 150 points\", \"discount_type\": \"fixed_amount\", \"unrendered_name\": \"$15.00 Off\",  \"discount_percentage\": null, \"discount_rate_cents\": null, \"discount_value_cents\": null, \"discount_amount_cents\": 1500,  \"unrendered_description\": \"Get $15 off your next purchase for 150 points\", \"applies_to_product_type\": \"ALL\"}"),
                            },
                        },
                        Metadata: map[string]string{
                            "key1": "value1",
                            "key2": "value2",
                        },
                        OrderDescription: boltembeddedgo.String("Order #1234567890"),
                        OrderReference: "order_100",
                        Shipments: []shared.CartShipment{
                            shared.CartShipment{
                                Carrier: boltembeddedgo.String("FedEx"),
                                Cost: boltembeddedgo.Int64(770),
                                DiscountedByMembership: boltembeddedgo.Bool(false),
                                EstimatedDeliveryDate: boltembeddedgo.String("08-30-2022"),
                                Expedited: boltembeddedgo.Bool(false),
                                PackageDepth: boltembeddedgo.Int64(90),
                                PackageDimensionUnit: boltembeddedgo.String("cm"),
                                PackageHeight: boltembeddedgo.Int64(103),
                                PackageType: boltembeddedgo.String("A big package."),
                                PackageWeightUnit: boltembeddedgo.String("kg"),
                                PackageWidth: boltembeddedgo.Int64(222),
                                Service: boltembeddedgo.String("Option 1"),
                                ShippingAddress: &shared.Address{
                                    Company: boltembeddedgo.String("Bolt"),
                                    Country: boltembeddedgo.String("United States"),
                                    CountryCode: "US",
                                    Default: boltembeddedgo.Bool(true),
                                    DoorCode: boltembeddedgo.String("123456"),
                                    Email: "alan.watts@example.com",
                                    FirstName: "Alan",
                                    LastName: "Watts",
                                    Locality: "Brooklyn",
                                    Name: boltembeddedgo.String("Alan Watts"),
                                    Phone: boltembeddedgo.String("+12125550199"),
                                    PostalCode: "10044",
                                    Region: "NY",
                                    RegionCode: boltembeddedgo.String("NY"),
                                    StreetAddress1: "888 main street",
                                    StreetAddress2: boltembeddedgo.String("apt 3021"),
                                    StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                                    StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                                },
                                ShippingAddressID: boltembeddedgo.String("addres-1"),
                                ShippingMethod: boltembeddedgo.String("Unknown"),
                                Signature: boltembeddedgo.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                TaxAmount: boltembeddedgo.Int64(230),
                                TaxCode: boltembeddedgo.String("tax-12345"),
                                TotalWeight: boltembeddedgo.Int64(55),
                                TotalWeightUnit: boltembeddedgo.String("kg"),
                                Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                            },
                        },
                        TotalAmount: 900,
                    },
                    PaymentMethod: &operations.SavedPaymentInputUpdatePaymentMethod{
                        PaymentData: &operations.SavedPaymentInputUpdatePaymentData{},
                        Type: operations.SavedPaymentInputUpdateTypeSavedPaymentMethod.ToPointer(),
                    },
                    ShopperIdentity: &operations.SavedPaymentInputUpdateShopperIdentity{
                        CreateBoltAccount: boltembeddedgo.Bool(true),
                        Email: "Jaycee1@yahoo.com",
                        FirstName: "Winfield",
                        LastName: "Robel",
                        Phone: "1-689-791-8496",
                    },
                },
        ),
        ID: "<ID>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdatePaymentRequest](../../pkg/models/operations/updatepaymentrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.UpdatePaymentSecurity](../../pkg/models/operations/updatepaymentsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.UpdatePaymentResponse](../../pkg/models/operations/updatepaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
