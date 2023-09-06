# Payments

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
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
)

func main() {
    s := boltembeddedapi.New()
    operationSecurity := operations.FinalizePaymentSecurity{
            OAuth: boltembeddedapi.String(""),
        }

    ctx := context.Background()
    res, err := s.Payments.FinalizePayment(ctx, operations.FinalizePaymentRequest{
        IdempotencyKey: boltembeddedapi.String("cum"),
        RequestBody: &operations.FinalizePaymentRequestBody{
            MerchantEventID: boltembeddedapi.String("dbe0cd5d-3261-41d9-ba61-49e5b9d07567"),
            ShopperIdentity: &operations.FinalizePaymentRequestBodyShopperIdentity{
                CreateBoltAccount: boltembeddedapi.Bool(true),
                Email: "Alison97@gmail.com",
                FirstName: "Blanche",
                LastName: "Hessel",
                Phone: "516-484-9026 x558",
            },
        },
        ID: "488e1e91-e450-4ad2-abd4-4269802d502a",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.FinalizePayment200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.FinalizePaymentRequest](../../models/operations/finalizepaymentrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.FinalizePaymentSecurity](../../models/operations/finalizepaymentsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.FinalizePaymentResponse](../../models/operations/finalizepaymentresponse.md), error**


## InitializePayment

Initialize a Bolt payment token that will be used to reference this payment to Bolt when it is updated or finalized. NOTE: The authorization header is NOT required for payments associated with users who do not have a Bolt account.


### Example Usage

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
    operationSecurity := operations.InitializePaymentSecurity{
            OAuth: boltembeddedapi.String(""),
        }

    ctx := context.Background()
    res, err := s.Payments.InitializePayment(ctx, operations.InitializePaymentRequest{
        IdempotencyKey: boltembeddedapi.String("excepturi"),
        RequestBody: &operations.InitializePaymentRequestBody{
            Cart: shared.CartCreate{
                AddOns: []shared.CartAddOn{
                    shared.CartAddOn{
                        Description: boltembeddedapi.String("tempora"),
                        ImageURL: boltembeddedapi.String("facilis"),
                        Name: "Francisco Windler",
                        Price: 7561.07,
                        ProductID: "sint",
                        ProductPageURL: boltembeddedapi.String("aliquid"),
                    },
                },
                BillingAddress: &shared.Address{
                    Company: boltembeddedapi.String("Bolt"),
                    Country: boltembeddedapi.String("United States"),
                    CountryCode: "US",
                    Default: boltembeddedapi.Bool(true),
                    DoorCode: boltembeddedapi.String("123456"),
                    Email: "alan.watts@example.com",
                    FirstName: "Alan",
                    LastName: "Watts",
                    Locality: "Brooklyn",
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
                CartURL: boltembeddedapi.String("https://boltswagstore.com/orders/123456765432"),
                Currency: "USD",
                Discounts: []shared.CartDiscount{
                    shared.CartDiscount{
                        Amount: boltembeddedapi.Float64(100),
                        Code: boltembeddedapi.String("SUMMER10DISCOUNT"),
                        Description: boltembeddedapi.String("10% off our summer collection"),
                        DetailsURL: boltembeddedapi.String("https://boltswagstore.com/SUMMERSALE"),
                        DiscountCategory: shared.CartDiscountDiscountCategoryMembershipDiscount.ToPointer(),
                        Reference: boltembeddedapi.String("DISC-1234"),
                        Type: shared.CartDiscountTypePercentage.ToPointer(),
                    },
                },
                DisplayID: boltembeddedapi.String("displayid_100"),
                Fees: []shared.CartCreateFees{
                    shared.CartCreateFees{
                        Description: boltembeddedapi.String("Item Fee"),
                        Name: "Item Fee",
                        Quantity: 8960.39,
                        Reference: "ItemFee",
                        UnitPrice: 5722.52,
                        UnitTaxAmount: 6389.21,
                    },
                },
                Fulfillments: []shared.CartCreateFulfillments{
                    shared.CartCreateFulfillments{
                        CartItems: []shared.CartItem{
                            shared.CartItem{
                                Brand: boltembeddedapi.String("Bolt"),
                                Category: boltembeddedapi.String("bags"),
                                Collections: []string{
                                    "summer",
                                },
                                Color: boltembeddedapi.String("Bolt Blue"),
                                Customizations: []shared.CartItemCustomization{
                                    shared.CartItemCustomization{
                                        Attributes: map[string]string{
                                            "dolor": "debitis",
                                        },
                                        Name: boltembeddedapi.String("Wilbur King"),
                                        Price: &shared.AmountView{
                                            Amount: boltembeddedapi.Float64(754),
                                            Currency: boltembeddedapi.String("USD"),
                                            CurrencySymbol: boltembeddedapi.String("$"),
                                        },
                                    },
                                },
                                Description: boltembeddedapi.String("Large tote with Bolt logo."),
                                DetailsURL: boltembeddedapi.String("https://boltswagstore.com/products/123456"),
                                ExternalInputs: &shared.ICartItemExternalInputs{
                                    ShopifyLineItemReference: boltembeddedapi.Float64(9785.71),
                                    ShopifyProductReference: boltembeddedapi.Float64(6994.79),
                                    ShopifyProductVariantReference: boltembeddedapi.Float64(1162.02),
                                },
                                GiftOption: &shared.CartItemGiftOption{
                                    Cost: boltembeddedapi.Int64(770),
                                    MerchantProductID: boltembeddedapi.String("881"),
                                    Message: boltembeddedapi.String("Happy Anniversary, Smoochy Poo!"),
                                    Wrap: boltembeddedapi.Bool(false),
                                },
                                ImageURL: boltembeddedapi.String("https://boltswagstore.com/products/123456/images/1.png"),
                                Isbn: boltembeddedapi.String("9780091347314"),
                                ItemGroup: boltembeddedapi.String("magnam"),
                                Manufacturer: boltembeddedapi.String("Bolt Textiles USA"),
                                MerchantProductID: boltembeddedapi.String("881"),
                                MerchantVariantID: boltembeddedapi.String("888"),
                                Msrp: boltembeddedapi.Float64(7670.24),
                                Name: "Bolt Swag Bag",
                                Options: boltembeddedapi.String("Special Edition"),
                                Properties: []shared.CartItemProperty{
                                    shared.CartItemProperty{
                                        Color: boltembeddedapi.String("facere"),
                                        Display: boltembeddedapi.Bool(false),
                                        Name: boltembeddedapi.String("Beth Padberg"),
                                        NameID: boltembeddedapi.Float64(5812.73),
                                        Value: boltembeddedapi.String("enim"),
                                        ValueID: boltembeddedapi.Float64(8817.36),
                                    },
                                },
                                Quantity: 1,
                                Reference: "item_100",
                                Shipment: &shared.CartShipment{
                                    Carrier: boltembeddedapi.String("FedEx"),
                                    Cost: boltembeddedapi.Int64(770),
                                    DiscountedByMembership: boltembeddedapi.Bool(false),
                                    EstimatedDeliveryDate: boltembeddedapi.String("08-30-2022"),
                                    Expedited: boltembeddedapi.Bool(false),
                                    PackageDepth: boltembeddedapi.Int64(90),
                                    PackageDimensionUnit: boltembeddedapi.String("cm"),
                                    PackageHeight: boltembeddedapi.Int64(103),
                                    PackageType: boltembeddedapi.String("A big package."),
                                    PackageWeightUnit: boltembeddedapi.String("kg"),
                                    PackageWidth: boltembeddedapi.Int64(222),
                                    Service: boltembeddedapi.String("Option 1"),
                                    ShippingAddress: &shared.Address{
                                        Company: boltembeddedapi.String("Bolt"),
                                        Country: boltembeddedapi.String("United States"),
                                        CountryCode: "US",
                                        Default: boltembeddedapi.Bool(true),
                                        DoorCode: boltembeddedapi.String("123456"),
                                        Email: "alan.watts@example.com",
                                        FirstName: "Alan",
                                        LastName: "Watts",
                                        Locality: "Brooklyn",
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
                                    ShippingAddressID: boltembeddedapi.String("addres-1"),
                                    ShippingMethod: boltembeddedapi.String("Unknown"),
                                    Signature: boltembeddedapi.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                    TaxAmount: boltembeddedapi.Int64(230),
                                    TaxCode: boltembeddedapi.String("tax-12345"),
                                    TotalWeight: boltembeddedapi.Int64(55),
                                    TotalWeightUnit: boltembeddedapi.String("kg"),
                                    Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                                },
                                ShipmentType: shared.CartItemShipmentTypeInStorePickup.ToPointer(),
                                Size: boltembeddedapi.String("Large"),
                                Sku: boltembeddedapi.String("BOLT-SKU_100"),
                                Source: boltembeddedapi.String("quidem"),
                                Tags: boltembeddedapi.String("tote, blue, linen, eco-friendly"),
                                TaxAmount: boltembeddedapi.Float64(0),
                                TaxCode: boltembeddedapi.String("provident"),
                                Taxable: boltembeddedapi.Bool(false),
                                TotalAmount: 1000,
                                Type: shared.CartItemTypePhysical.ToPointer(),
                                UnitPrice: 1000,
                                Uom: boltembeddedapi.String("inches"),
                                Upc: boltembeddedapi.String("825764603119"),
                                Weight: boltembeddedapi.Float64(10),
                                WeightUnit: boltembeddedapi.String("pounds"),
                            },
                        },
                        CartShipment: &shared.CartShipment{
                            Carrier: boltembeddedapi.String("FedEx"),
                            Cost: boltembeddedapi.Int64(770),
                            DiscountedByMembership: boltembeddedapi.Bool(false),
                            EstimatedDeliveryDate: boltembeddedapi.String("08-30-2022"),
                            Expedited: boltembeddedapi.Bool(false),
                            PackageDepth: boltembeddedapi.Int64(90),
                            PackageDimensionUnit: boltembeddedapi.String("cm"),
                            PackageHeight: boltembeddedapi.Int64(103),
                            PackageType: boltembeddedapi.String("A big package."),
                            PackageWeightUnit: boltembeddedapi.String("kg"),
                            PackageWidth: boltembeddedapi.Int64(222),
                            Service: boltembeddedapi.String("Option 1"),
                            ShippingAddress: &shared.Address{
                                Company: boltembeddedapi.String("Bolt"),
                                Country: boltembeddedapi.String("United States"),
                                CountryCode: "US",
                                Default: boltembeddedapi.Bool(true),
                                DoorCode: boltembeddedapi.String("123456"),
                                Email: "alan.watts@example.com",
                                FirstName: "Alan",
                                LastName: "Watts",
                                Locality: "Brooklyn",
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
                            ShippingAddressID: boltembeddedapi.String("addres-1"),
                            ShippingMethod: boltembeddedapi.String("Unknown"),
                            Signature: boltembeddedapi.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                            TaxAmount: boltembeddedapi.Int64(230),
                            TaxCode: boltembeddedapi.String("tax-12345"),
                            TotalWeight: boltembeddedapi.Int64(55),
                            TotalWeightUnit: boltembeddedapi.String("kg"),
                            Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                        },
                        DigitalDelivery: &shared.CartCreateFulfillmentsDigitalDelivery{
                            Email: boltembeddedapi.String("Jewell.Lesch64@yahoo.com"),
                            Phone: boltembeddedapi.String("1-566-504-3762 x236"),
                        },
                        InStoreCartShipment: &shared.InStoreCartShipment{
                            CartShipment: &shared.CartShipment{
                                Carrier: boltembeddedapi.String("FedEx"),
                                Cost: boltembeddedapi.Int64(770),
                                DiscountedByMembership: boltembeddedapi.Bool(false),
                                EstimatedDeliveryDate: boltembeddedapi.String("08-30-2022"),
                                Expedited: boltembeddedapi.Bool(false),
                                PackageDepth: boltembeddedapi.Int64(90),
                                PackageDimensionUnit: boltembeddedapi.String("cm"),
                                PackageHeight: boltembeddedapi.Int64(103),
                                PackageType: boltembeddedapi.String("A big package."),
                                PackageWeightUnit: boltembeddedapi.String("kg"),
                                PackageWidth: boltembeddedapi.Int64(222),
                                Service: boltembeddedapi.String("Option 1"),
                                ShippingAddress: &shared.Address{
                                    Company: boltembeddedapi.String("Bolt"),
                                    Country: boltembeddedapi.String("United States"),
                                    CountryCode: "US",
                                    Default: boltembeddedapi.Bool(true),
                                    DoorCode: boltembeddedapi.String("123456"),
                                    Email: "alan.watts@example.com",
                                    FirstName: "Alan",
                                    LastName: "Watts",
                                    Locality: "Brooklyn",
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
                                ShippingAddressID: boltembeddedapi.String("addres-1"),
                                ShippingMethod: boltembeddedapi.String("Unknown"),
                                Signature: boltembeddedapi.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                TaxAmount: boltembeddedapi.Int64(230),
                                TaxCode: boltembeddedapi.String("tax-12345"),
                                TotalWeight: boltembeddedapi.Int64(55),
                                TotalWeightUnit: boltembeddedapi.String("kg"),
                                Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                            },
                            Description: boltembeddedapi.String("Pick up in-store at 123 Main St."),
                            Distance: boltembeddedapi.Float64(3),
                            DistanceUnit: shared.InStoreCartShipmentDistanceUnitMile.ToPointer(),
                            InStorePickupAddress: &shared.Address{
                                Company: boltembeddedapi.String("Bolt"),
                                Country: boltembeddedapi.String("United States"),
                                CountryCode: "US",
                                Default: boltembeddedapi.Bool(true),
                                DoorCode: boltembeddedapi.String("123456"),
                                Email: "alan.watts@example.com",
                                FirstName: "Alan",
                                LastName: "Watts",
                                Locality: "Brooklyn",
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
                            PickupWindowClose: boltembeddedapi.Int64(749170),
                            PickupWindowOpen: boltembeddedapi.Int64(428769),
                            StoreName: boltembeddedapi.String("Bolt Collective"),
                        },
                        Type: shared.CartCreateFulfillmentsTypeDigitalNoDelivery.ToPointer(),
                    },
                },
                InStoreCartShipments: []shared.InStoreCartShipment{
                    shared.InStoreCartShipment{
                        CartShipment: &shared.CartShipment{
                            Carrier: boltembeddedapi.String("FedEx"),
                            Cost: boltembeddedapi.Int64(770),
                            DiscountedByMembership: boltembeddedapi.Bool(false),
                            EstimatedDeliveryDate: boltembeddedapi.String("08-30-2022"),
                            Expedited: boltembeddedapi.Bool(false),
                            PackageDepth: boltembeddedapi.Int64(90),
                            PackageDimensionUnit: boltembeddedapi.String("cm"),
                            PackageHeight: boltembeddedapi.Int64(103),
                            PackageType: boltembeddedapi.String("A big package."),
                            PackageWeightUnit: boltembeddedapi.String("kg"),
                            PackageWidth: boltembeddedapi.Int64(222),
                            Service: boltembeddedapi.String("Option 1"),
                            ShippingAddress: &shared.Address{
                                Company: boltembeddedapi.String("Bolt"),
                                Country: boltembeddedapi.String("United States"),
                                CountryCode: "US",
                                Default: boltembeddedapi.Bool(true),
                                DoorCode: boltembeddedapi.String("123456"),
                                Email: "alan.watts@example.com",
                                FirstName: "Alan",
                                LastName: "Watts",
                                Locality: "Brooklyn",
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
                            ShippingAddressID: boltembeddedapi.String("addres-1"),
                            ShippingMethod: boltembeddedapi.String("Unknown"),
                            Signature: boltembeddedapi.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                            TaxAmount: boltembeddedapi.Int64(230),
                            TaxCode: boltembeddedapi.String("tax-12345"),
                            TotalWeight: boltembeddedapi.Int64(55),
                            TotalWeightUnit: boltembeddedapi.String("kg"),
                            Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                        },
                        Description: boltembeddedapi.String("Pick up in-store at 123 Main St."),
                        Distance: boltembeddedapi.Float64(3),
                        DistanceUnit: shared.InStoreCartShipmentDistanceUnitMile.ToPointer(),
                        InStorePickupAddress: &shared.Address{
                            Company: boltembeddedapi.String("Bolt"),
                            Country: boltembeddedapi.String("United States"),
                            CountryCode: "US",
                            Default: boltembeddedapi.Bool(true),
                            DoorCode: boltembeddedapi.String("123456"),
                            Email: "alan.watts@example.com",
                            FirstName: "Alan",
                            LastName: "Watts",
                            Locality: "Brooklyn",
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
                        PickupWindowClose: boltembeddedapi.Int64(135474),
                        PickupWindowOpen: boltembeddedapi.Int64(102863),
                        StoreName: boltembeddedapi.String("Bolt Collective"),
                    },
                },
                Items: []shared.CartItem{
                    shared.CartItem{
                        Brand: boltembeddedapi.String("Bolt"),
                        Category: boltembeddedapi.String("bags"),
                        Collections: []string{
                            "summer",
                        },
                        Color: boltembeddedapi.String("Bolt Blue"),
                        Customizations: []shared.CartItemCustomization{
                            shared.CartItemCustomization{
                                Attributes: map[string]string{
                                    "magnam": "et",
                                },
                                Name: boltembeddedapi.String("Derrick McLaughlin"),
                                Price: &shared.AmountView{
                                    Amount: boltembeddedapi.Float64(754),
                                    Currency: boltembeddedapi.String("USD"),
                                    CurrencySymbol: boltembeddedapi.String("$"),
                                },
                            },
                        },
                        Description: boltembeddedapi.String("Large tote with Bolt logo."),
                        DetailsURL: boltembeddedapi.String("https://boltswagstore.com/products/123456"),
                        ExternalInputs: &shared.ICartItemExternalInputs{
                            ShopifyLineItemReference: boltembeddedapi.Float64(336.25),
                            ShopifyProductReference: boltembeddedapi.Float64(6532.01),
                            ShopifyProductVariantReference: boltembeddedapi.Float64(9689.62),
                        },
                        GiftOption: &shared.CartItemGiftOption{
                            Cost: boltembeddedapi.Int64(770),
                            MerchantProductID: boltembeddedapi.String("881"),
                            Message: boltembeddedapi.String("Happy Anniversary, Smoochy Poo!"),
                            Wrap: boltembeddedapi.Bool(false),
                        },
                        ImageURL: boltembeddedapi.String("https://boltswagstore.com/products/123456/images/1.png"),
                        Isbn: boltembeddedapi.String("9780091347314"),
                        ItemGroup: boltembeddedapi.String("mollitia"),
                        Manufacturer: boltembeddedapi.String("Bolt Textiles USA"),
                        MerchantProductID: boltembeddedapi.String("881"),
                        MerchantVariantID: boltembeddedapi.String("888"),
                        Msrp: boltembeddedapi.Float64(3209.97),
                        Name: "Bolt Swag Bag",
                        Options: boltembeddedapi.String("Special Edition"),
                        Properties: []shared.CartItemProperty{
                            shared.CartItemProperty{
                                Color: boltembeddedapi.String("eum"),
                                Display: boltembeddedapi.Bool(false),
                                Name: boltembeddedapi.String("Jana Connelly III"),
                                NameID: boltembeddedapi.Float64(9840.43),
                                Value: boltembeddedapi.String("debitis"),
                                ValueID: boltembeddedapi.Float64(2603.41),
                            },
                        },
                        Quantity: 1,
                        Reference: "item_100",
                        Shipment: &shared.CartShipment{
                            Carrier: boltembeddedapi.String("FedEx"),
                            Cost: boltembeddedapi.Int64(770),
                            DiscountedByMembership: boltembeddedapi.Bool(false),
                            EstimatedDeliveryDate: boltembeddedapi.String("08-30-2022"),
                            Expedited: boltembeddedapi.Bool(false),
                            PackageDepth: boltembeddedapi.Int64(90),
                            PackageDimensionUnit: boltembeddedapi.String("cm"),
                            PackageHeight: boltembeddedapi.Int64(103),
                            PackageType: boltembeddedapi.String("A big package."),
                            PackageWeightUnit: boltembeddedapi.String("kg"),
                            PackageWidth: boltembeddedapi.Int64(222),
                            Service: boltembeddedapi.String("Option 1"),
                            ShippingAddress: &shared.Address{
                                Company: boltembeddedapi.String("Bolt"),
                                Country: boltembeddedapi.String("United States"),
                                CountryCode: "US",
                                Default: boltembeddedapi.Bool(true),
                                DoorCode: boltembeddedapi.String("123456"),
                                Email: "alan.watts@example.com",
                                FirstName: "Alan",
                                LastName: "Watts",
                                Locality: "Brooklyn",
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
                            ShippingAddressID: boltembeddedapi.String("addres-1"),
                            ShippingMethod: boltembeddedapi.String("Unknown"),
                            Signature: boltembeddedapi.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                            TaxAmount: boltembeddedapi.Int64(230),
                            TaxCode: boltembeddedapi.String("tax-12345"),
                            TotalWeight: boltembeddedapi.Int64(55),
                            TotalWeightUnit: boltembeddedapi.String("kg"),
                            Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                        },
                        ShipmentType: shared.CartItemShipmentTypeInStorePickup.ToPointer(),
                        Size: boltembeddedapi.String("Large"),
                        Sku: boltembeddedapi.String("BOLT-SKU_100"),
                        Source: boltembeddedapi.String("deleniti"),
                        Tags: boltembeddedapi.String("tote, blue, linen, eco-friendly"),
                        TaxAmount: boltembeddedapi.Float64(0),
                        TaxCode: boltembeddedapi.String("facilis"),
                        Taxable: boltembeddedapi.Bool(false),
                        TotalAmount: 1000,
                        Type: shared.CartItemTypeDigital.ToPointer(),
                        UnitPrice: 1000,
                        Uom: boltembeddedapi.String("inches"),
                        Upc: boltembeddedapi.String("825764603119"),
                        Weight: boltembeddedapi.Float64(10),
                        WeightUnit: boltembeddedapi.String("pounds"),
                    },
                },
                LoyaltyRewards: []shared.CartLoyaltyRewards{
                    shared.CartLoyaltyRewards{
                        Amount: boltembeddedapi.Float64(1002.26),
                        CouponCode: boltembeddedapi.String("architecto"),
                        Description: boltembeddedapi.String("$5 off (100 Points)"),
                        Details: boltembeddedapi.String("{"id": 123456, "icon": "fa-dollar", "name": "$15.00 Off", "type": "Coupon", "amount": 100, "duration": "single_use", "cost_text": "150 Points",  "description": "Get $15 off your next purchase for 150 points", "discount_type": "fixed_amount", "unrendered_name": "$15.00 Off",  "discount_percentage": null, "discount_rate_cents": null, "discount_value_cents": null, "discount_amount_cents": 1500,  "unrendered_description": "Get $15 off your next purchase for 150 points", "applies_to_product_type": "ALL"}"),
                        Points: boltembeddedapi.Float64(9194.83),
                        Source: boltembeddedapi.String("ullam"),
                        Type: boltembeddedapi.String("expedita"),
                    },
                },
                Metadata: map[string]string{
                    "nihil": "repellat",
                },
                OrderDescription: boltembeddedapi.String("Order #1234567890"),
                OrderReference: "order_100",
                Shipments: []shared.CartShipment{
                    shared.CartShipment{
                        Carrier: boltembeddedapi.String("FedEx"),
                        Cost: boltembeddedapi.Int64(770),
                        DiscountedByMembership: boltembeddedapi.Bool(false),
                        EstimatedDeliveryDate: boltembeddedapi.String("08-30-2022"),
                        Expedited: boltembeddedapi.Bool(false),
                        PackageDepth: boltembeddedapi.Int64(90),
                        PackageDimensionUnit: boltembeddedapi.String("cm"),
                        PackageHeight: boltembeddedapi.Int64(103),
                        PackageType: boltembeddedapi.String("A big package."),
                        PackageWeightUnit: boltembeddedapi.String("kg"),
                        PackageWidth: boltembeddedapi.Int64(222),
                        Service: boltembeddedapi.String("Option 1"),
                        ShippingAddress: &shared.Address{
                            Company: boltembeddedapi.String("Bolt"),
                            Country: boltembeddedapi.String("United States"),
                            CountryCode: "US",
                            Default: boltembeddedapi.Bool(true),
                            DoorCode: boltembeddedapi.String("123456"),
                            Email: "alan.watts@example.com",
                            FirstName: "Alan",
                            LastName: "Watts",
                            Locality: "Brooklyn",
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
                        ShippingAddressID: boltembeddedapi.String("addres-1"),
                        ShippingMethod: boltembeddedapi.String("Unknown"),
                        Signature: boltembeddedapi.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                        TaxAmount: boltembeddedapi.Int64(230),
                        TaxCode: boltembeddedapi.String("tax-12345"),
                        TotalWeight: boltembeddedapi.Int64(55),
                        TotalWeightUnit: boltembeddedapi.String("kg"),
                        Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                    },
                },
                TaxAmount: boltembeddedapi.Float64(8411.4),
                TotalAmount: 900,
            },
            ShopperIdentity: &operations.InitializePaymentRequestBodyShopperIdentity{
                CreateBoltAccount: boltembeddedapi.Bool(true),
                Email: "Sunny16@gmail.com",
                FirstName: "Johan",
                LastName: "Morissette",
                Phone: "1-278-884-5140",
            },
        },
        XPublishableKey: boltembeddedapi.String("ab"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.InitializePayment200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.InitializePaymentRequest](../../models/operations/initializepaymentrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.InitializePaymentSecurity](../../models/operations/initializepaymentsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.InitializePaymentResponse](../../models/operations/initializepaymentresponse.md), error**


## UpdatePayment

Update a Bolt payment using the token given after initializing a payment.  Updates will completely replace the original top-level resource (for example, if a cart is sent in with the request it will replace the existing cart).  Any included object should be sent as complete object. NOTE: The authorization header is NOT required for payments associated with users who do not have a Bolt account.


### Example Usage

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
    operationSecurity := operations.UpdatePaymentSecurity{
            OAuth: boltembeddedapi.String(""),
        }

    ctx := context.Background()
    res, err := s.Payments.UpdatePayment(ctx, operations.UpdatePaymentRequest{
        IdempotencyKey: boltembeddedapi.String("maiores"),
        RequestBody: &operations.UpdatePaymentRequestBody{
            Cart: &shared.CartCreate{
                AddOns: []shared.CartAddOn{
                    shared.CartAddOn{
                        Description: boltembeddedapi.String("quidem"),
                        ImageURL: boltembeddedapi.String("ipsam"),
                        Name: "Dr. Stacey Reichert",
                        Price: 9755.22,
                        ProductID: "perferendis",
                        ProductPageURL: boltembeddedapi.String("fugiat"),
                    },
                },
                BillingAddress: &shared.Address{
                    Company: boltembeddedapi.String("Bolt"),
                    Country: boltembeddedapi.String("United States"),
                    CountryCode: "US",
                    Default: boltembeddedapi.Bool(true),
                    DoorCode: boltembeddedapi.String("123456"),
                    Email: "alan.watts@example.com",
                    FirstName: "Alan",
                    LastName: "Watts",
                    Locality: "Brooklyn",
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
                CartURL: boltembeddedapi.String("https://boltswagstore.com/orders/123456765432"),
                Currency: "USD",
                Discounts: []shared.CartDiscount{
                    shared.CartDiscount{
                        Amount: boltembeddedapi.Float64(100),
                        Code: boltembeddedapi.String("SUMMER10DISCOUNT"),
                        Description: boltembeddedapi.String("10% off our summer collection"),
                        DetailsURL: boltembeddedapi.String("https://boltswagstore.com/SUMMERSALE"),
                        DiscountCategory: shared.CartDiscountDiscountCategoryManagedGiftcard.ToPointer(),
                        Reference: boltembeddedapi.String("DISC-1234"),
                        Type: shared.CartDiscountTypePercentage.ToPointer(),
                    },
                },
                DisplayID: boltembeddedapi.String("displayid_100"),
                Fees: []shared.CartCreateFees{
                    shared.CartCreateFees{
                        Description: boltembeddedapi.String("Item Fee"),
                        Name: "Item Fee",
                        Quantity: 117.14,
                        Reference: "ItemFee",
                        UnitPrice: 7649.12,
                        UnitTaxAmount: 3599.78,
                    },
                },
                Fulfillments: []shared.CartCreateFulfillments{
                    shared.CartCreateFulfillments{
                        CartItems: []shared.CartItem{
                            shared.CartItem{
                                Brand: boltembeddedapi.String("Bolt"),
                                Category: boltembeddedapi.String("bags"),
                                Collections: []string{
                                    "summer",
                                },
                                Color: boltembeddedapi.String("Bolt Blue"),
                                Customizations: []shared.CartItemCustomization{
                                    shared.CartItemCustomization{
                                        Attributes: map[string]string{
                                            "hic": "libero",
                                        },
                                        Name: boltembeddedapi.String("Ernest Hayes"),
                                        Price: &shared.AmountView{
                                            Amount: boltembeddedapi.Float64(754),
                                            Currency: boltembeddedapi.String("USD"),
                                            CurrencySymbol: boltembeddedapi.String("$"),
                                        },
                                    },
                                },
                                Description: boltembeddedapi.String("Large tote with Bolt logo."),
                                DetailsURL: boltembeddedapi.String("https://boltswagstore.com/products/123456"),
                                ExternalInputs: &shared.ICartItemExternalInputs{
                                    ShopifyLineItemReference: boltembeddedapi.Float64(543.38),
                                    ShopifyProductReference: boltembeddedapi.Float64(3389.85),
                                    ShopifyProductVariantReference: boltembeddedapi.Float64(1999.96),
                                },
                                GiftOption: &shared.CartItemGiftOption{
                                    Cost: boltembeddedapi.Int64(770),
                                    MerchantProductID: boltembeddedapi.String("881"),
                                    Message: boltembeddedapi.String("Happy Anniversary, Smoochy Poo!"),
                                    Wrap: boltembeddedapi.Bool(false),
                                },
                                ImageURL: boltembeddedapi.String("https://boltswagstore.com/products/123456/images/1.png"),
                                Isbn: boltembeddedapi.String("9780091347314"),
                                ItemGroup: boltembeddedapi.String("eos"),
                                Manufacturer: boltembeddedapi.String("Bolt Textiles USA"),
                                MerchantProductID: boltembeddedapi.String("881"),
                                MerchantVariantID: boltembeddedapi.String("888"),
                                Msrp: boltembeddedapi.Float64(185.21),
                                Name: "Bolt Swag Bag",
                                Options: boltembeddedapi.String("Special Edition"),
                                Properties: []shared.CartItemProperty{
                                    shared.CartItemProperty{
                                        Color: boltembeddedapi.String("dolores"),
                                        Display: boltembeddedapi.Bool(false),
                                        Name: boltembeddedapi.String("Darryl Fadel"),
                                        NameID: boltembeddedapi.Float64(9441.2),
                                        Value: boltembeddedapi.String("recusandae"),
                                        ValueID: boltembeddedapi.Float64(6082.53),
                                    },
                                },
                                Quantity: 1,
                                Reference: "item_100",
                                Shipment: &shared.CartShipment{
                                    Carrier: boltembeddedapi.String("FedEx"),
                                    Cost: boltembeddedapi.Int64(770),
                                    DiscountedByMembership: boltembeddedapi.Bool(false),
                                    EstimatedDeliveryDate: boltembeddedapi.String("08-30-2022"),
                                    Expedited: boltembeddedapi.Bool(false),
                                    PackageDepth: boltembeddedapi.Int64(90),
                                    PackageDimensionUnit: boltembeddedapi.String("cm"),
                                    PackageHeight: boltembeddedapi.Int64(103),
                                    PackageType: boltembeddedapi.String("A big package."),
                                    PackageWeightUnit: boltembeddedapi.String("kg"),
                                    PackageWidth: boltembeddedapi.Int64(222),
                                    Service: boltembeddedapi.String("Option 1"),
                                    ShippingAddress: &shared.Address{
                                        Company: boltembeddedapi.String("Bolt"),
                                        Country: boltembeddedapi.String("United States"),
                                        CountryCode: "US",
                                        Default: boltembeddedapi.Bool(true),
                                        DoorCode: boltembeddedapi.String("123456"),
                                        Email: "alan.watts@example.com",
                                        FirstName: "Alan",
                                        LastName: "Watts",
                                        Locality: "Brooklyn",
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
                                    ShippingAddressID: boltembeddedapi.String("addres-1"),
                                    ShippingMethod: boltembeddedapi.String("Unknown"),
                                    Signature: boltembeddedapi.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                    TaxAmount: boltembeddedapi.Int64(230),
                                    TaxCode: boltembeddedapi.String("tax-12345"),
                                    TotalWeight: boltembeddedapi.Int64(55),
                                    TotalWeightUnit: boltembeddedapi.String("kg"),
                                    Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                                },
                                ShipmentType: shared.CartItemShipmentTypeShipToStore.ToPointer(),
                                Size: boltembeddedapi.String("Large"),
                                Sku: boltembeddedapi.String("BOLT-SKU_100"),
                                Source: boltembeddedapi.String("perspiciatis"),
                                Tags: boltembeddedapi.String("tote, blue, linen, eco-friendly"),
                                TaxAmount: boltembeddedapi.Float64(0),
                                TaxCode: boltembeddedapi.String("voluptatem"),
                                Taxable: boltembeddedapi.Bool(false),
                                TotalAmount: 1000,
                                Type: shared.CartItemTypeBundled.ToPointer(),
                                UnitPrice: 1000,
                                Uom: boltembeddedapi.String("inches"),
                                Upc: boltembeddedapi.String("825764603119"),
                                Weight: boltembeddedapi.Float64(10),
                                WeightUnit: boltembeddedapi.String("pounds"),
                            },
                        },
                        CartShipment: &shared.CartShipment{
                            Carrier: boltembeddedapi.String("FedEx"),
                            Cost: boltembeddedapi.Int64(770),
                            DiscountedByMembership: boltembeddedapi.Bool(false),
                            EstimatedDeliveryDate: boltembeddedapi.String("08-30-2022"),
                            Expedited: boltembeddedapi.Bool(false),
                            PackageDepth: boltembeddedapi.Int64(90),
                            PackageDimensionUnit: boltembeddedapi.String("cm"),
                            PackageHeight: boltembeddedapi.Int64(103),
                            PackageType: boltembeddedapi.String("A big package."),
                            PackageWeightUnit: boltembeddedapi.String("kg"),
                            PackageWidth: boltembeddedapi.Int64(222),
                            Service: boltembeddedapi.String("Option 1"),
                            ShippingAddress: &shared.Address{
                                Company: boltembeddedapi.String("Bolt"),
                                Country: boltembeddedapi.String("United States"),
                                CountryCode: "US",
                                Default: boltembeddedapi.Bool(true),
                                DoorCode: boltembeddedapi.String("123456"),
                                Email: "alan.watts@example.com",
                                FirstName: "Alan",
                                LastName: "Watts",
                                Locality: "Brooklyn",
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
                            ShippingAddressID: boltembeddedapi.String("addres-1"),
                            ShippingMethod: boltembeddedapi.String("Unknown"),
                            Signature: boltembeddedapi.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                            TaxAmount: boltembeddedapi.Int64(230),
                            TaxCode: boltembeddedapi.String("tax-12345"),
                            TotalWeight: boltembeddedapi.Int64(55),
                            TotalWeightUnit: boltembeddedapi.String("kg"),
                            Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                        },
                        DigitalDelivery: &shared.CartCreateFulfillmentsDigitalDelivery{
                            Email: boltembeddedapi.String("Jett57@gmail.com"),
                            Phone: boltembeddedapi.String("399.466.5857 x7935"),
                        },
                        InStoreCartShipment: &shared.InStoreCartShipment{
                            CartShipment: &shared.CartShipment{
                                Carrier: boltembeddedapi.String("FedEx"),
                                Cost: boltembeddedapi.Int64(770),
                                DiscountedByMembership: boltembeddedapi.Bool(false),
                                EstimatedDeliveryDate: boltembeddedapi.String("08-30-2022"),
                                Expedited: boltembeddedapi.Bool(false),
                                PackageDepth: boltembeddedapi.Int64(90),
                                PackageDimensionUnit: boltembeddedapi.String("cm"),
                                PackageHeight: boltembeddedapi.Int64(103),
                                PackageType: boltembeddedapi.String("A big package."),
                                PackageWeightUnit: boltembeddedapi.String("kg"),
                                PackageWidth: boltembeddedapi.Int64(222),
                                Service: boltembeddedapi.String("Option 1"),
                                ShippingAddress: &shared.Address{
                                    Company: boltembeddedapi.String("Bolt"),
                                    Country: boltembeddedapi.String("United States"),
                                    CountryCode: "US",
                                    Default: boltembeddedapi.Bool(true),
                                    DoorCode: boltembeddedapi.String("123456"),
                                    Email: "alan.watts@example.com",
                                    FirstName: "Alan",
                                    LastName: "Watts",
                                    Locality: "Brooklyn",
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
                                ShippingAddressID: boltembeddedapi.String("addres-1"),
                                ShippingMethod: boltembeddedapi.String("Unknown"),
                                Signature: boltembeddedapi.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                                TaxAmount: boltembeddedapi.Int64(230),
                                TaxCode: boltembeddedapi.String("tax-12345"),
                                TotalWeight: boltembeddedapi.Int64(55),
                                TotalWeightUnit: boltembeddedapi.String("kg"),
                                Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                            },
                            Description: boltembeddedapi.String("Pick up in-store at 123 Main St."),
                            Distance: boltembeddedapi.Float64(3),
                            DistanceUnit: shared.InStoreCartShipmentDistanceUnitMile.ToPointer(),
                            InStorePickupAddress: &shared.Address{
                                Company: boltembeddedapi.String("Bolt"),
                                Country: boltembeddedapi.String("United States"),
                                CountryCode: "US",
                                Default: boltembeddedapi.Bool(true),
                                DoorCode: boltembeddedapi.String("123456"),
                                Email: "alan.watts@example.com",
                                FirstName: "Alan",
                                LastName: "Watts",
                                Locality: "Brooklyn",
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
                            PickupWindowClose: boltembeddedapi.Int64(398221),
                            PickupWindowOpen: boltembeddedapi.Int64(212390),
                            StoreName: boltembeddedapi.String("Bolt Collective"),
                        },
                        Type: shared.CartCreateFulfillmentsTypePhysicalShipToStore.ToPointer(),
                    },
                },
                InStoreCartShipments: []shared.InStoreCartShipment{
                    shared.InStoreCartShipment{
                        CartShipment: &shared.CartShipment{
                            Carrier: boltembeddedapi.String("FedEx"),
                            Cost: boltembeddedapi.Int64(770),
                            DiscountedByMembership: boltembeddedapi.Bool(false),
                            EstimatedDeliveryDate: boltembeddedapi.String("08-30-2022"),
                            Expedited: boltembeddedapi.Bool(false),
                            PackageDepth: boltembeddedapi.Int64(90),
                            PackageDimensionUnit: boltembeddedapi.String("cm"),
                            PackageHeight: boltembeddedapi.Int64(103),
                            PackageType: boltembeddedapi.String("A big package."),
                            PackageWeightUnit: boltembeddedapi.String("kg"),
                            PackageWidth: boltembeddedapi.Int64(222),
                            Service: boltembeddedapi.String("Option 1"),
                            ShippingAddress: &shared.Address{
                                Company: boltembeddedapi.String("Bolt"),
                                Country: boltembeddedapi.String("United States"),
                                CountryCode: "US",
                                Default: boltembeddedapi.Bool(true),
                                DoorCode: boltembeddedapi.String("123456"),
                                Email: "alan.watts@example.com",
                                FirstName: "Alan",
                                LastName: "Watts",
                                Locality: "Brooklyn",
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
                            ShippingAddressID: boltembeddedapi.String("addres-1"),
                            ShippingMethod: boltembeddedapi.String("Unknown"),
                            Signature: boltembeddedapi.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                            TaxAmount: boltembeddedapi.Int64(230),
                            TaxCode: boltembeddedapi.String("tax-12345"),
                            TotalWeight: boltembeddedapi.Int64(55),
                            TotalWeightUnit: boltembeddedapi.String("kg"),
                            Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                        },
                        Description: boltembeddedapi.String("Pick up in-store at 123 Main St."),
                        Distance: boltembeddedapi.Float64(3),
                        DistanceUnit: shared.InStoreCartShipmentDistanceUnitMile.ToPointer(),
                        InStorePickupAddress: &shared.Address{
                            Company: boltembeddedapi.String("Bolt"),
                            Country: boltembeddedapi.String("United States"),
                            CountryCode: "US",
                            Default: boltembeddedapi.Bool(true),
                            DoorCode: boltembeddedapi.String("123456"),
                            Email: "alan.watts@example.com",
                            FirstName: "Alan",
                            LastName: "Watts",
                            Locality: "Brooklyn",
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
                        PickupWindowClose: boltembeddedapi.Int64(222443),
                        PickupWindowOpen: boltembeddedapi.Int64(186193),
                        StoreName: boltembeddedapi.String("Bolt Collective"),
                    },
                },
                Items: []shared.CartItem{
                    shared.CartItem{
                        Brand: boltembeddedapi.String("Bolt"),
                        Category: boltembeddedapi.String("bags"),
                        Collections: []string{
                            "summer",
                        },
                        Color: boltembeddedapi.String("Bolt Blue"),
                        Customizations: []shared.CartItemCustomization{
                            shared.CartItemCustomization{
                                Attributes: map[string]string{
                                    "ipsum": "hic",
                                },
                                Name: boltembeddedapi.String("Felipe Klein"),
                                Price: &shared.AmountView{
                                    Amount: boltembeddedapi.Float64(754),
                                    Currency: boltembeddedapi.String("USD"),
                                    CurrencySymbol: boltembeddedapi.String("$"),
                                },
                            },
                        },
                        Description: boltembeddedapi.String("Large tote with Bolt logo."),
                        DetailsURL: boltembeddedapi.String("https://boltswagstore.com/products/123456"),
                        ExternalInputs: &shared.ICartItemExternalInputs{
                            ShopifyLineItemReference: boltembeddedapi.Float64(2274.14),
                            ShopifyProductReference: boltembeddedapi.Float64(6805.45),
                            ShopifyProductVariantReference: boltembeddedapi.Float64(2543.56),
                        },
                        GiftOption: &shared.CartItemGiftOption{
                            Cost: boltembeddedapi.Int64(770),
                            MerchantProductID: boltembeddedapi.String("881"),
                            Message: boltembeddedapi.String("Happy Anniversary, Smoochy Poo!"),
                            Wrap: boltembeddedapi.Bool(false),
                        },
                        ImageURL: boltembeddedapi.String("https://boltswagstore.com/products/123456/images/1.png"),
                        Isbn: boltembeddedapi.String("9780091347314"),
                        ItemGroup: boltembeddedapi.String("veritatis"),
                        Manufacturer: boltembeddedapi.String("Bolt Textiles USA"),
                        MerchantProductID: boltembeddedapi.String("881"),
                        MerchantVariantID: boltembeddedapi.String("888"),
                        Msrp: boltembeddedapi.Float64(580.29),
                        Name: "Bolt Swag Bag",
                        Options: boltembeddedapi.String("Special Edition"),
                        Properties: []shared.CartItemProperty{
                            shared.CartItemProperty{
                                Color: boltembeddedapi.String("ipsa"),
                                Display: boltembeddedapi.Bool(false),
                                Name: boltembeddedapi.String("Viola Hahn"),
                                NameID: boltembeddedapi.Float64(9764.05),
                                Value: boltembeddedapi.String("voluptas"),
                                ValueID: boltembeddedapi.Float64(6176.58),
                            },
                        },
                        Quantity: 1,
                        Reference: "item_100",
                        Shipment: &shared.CartShipment{
                            Carrier: boltembeddedapi.String("FedEx"),
                            Cost: boltembeddedapi.Int64(770),
                            DiscountedByMembership: boltembeddedapi.Bool(false),
                            EstimatedDeliveryDate: boltembeddedapi.String("08-30-2022"),
                            Expedited: boltembeddedapi.Bool(false),
                            PackageDepth: boltembeddedapi.Int64(90),
                            PackageDimensionUnit: boltembeddedapi.String("cm"),
                            PackageHeight: boltembeddedapi.Int64(103),
                            PackageType: boltembeddedapi.String("A big package."),
                            PackageWeightUnit: boltembeddedapi.String("kg"),
                            PackageWidth: boltembeddedapi.Int64(222),
                            Service: boltembeddedapi.String("Option 1"),
                            ShippingAddress: &shared.Address{
                                Company: boltembeddedapi.String("Bolt"),
                                Country: boltembeddedapi.String("United States"),
                                CountryCode: "US",
                                Default: boltembeddedapi.Bool(true),
                                DoorCode: boltembeddedapi.String("123456"),
                                Email: "alan.watts@example.com",
                                FirstName: "Alan",
                                LastName: "Watts",
                                Locality: "Brooklyn",
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
                            ShippingAddressID: boltembeddedapi.String("addres-1"),
                            ShippingMethod: boltembeddedapi.String("Unknown"),
                            Signature: boltembeddedapi.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                            TaxAmount: boltembeddedapi.Int64(230),
                            TaxCode: boltembeddedapi.String("tax-12345"),
                            TotalWeight: boltembeddedapi.Int64(55),
                            TotalWeightUnit: boltembeddedapi.String("kg"),
                            Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                        },
                        ShipmentType: shared.CartItemShipmentTypeUnknown.ToPointer(),
                        Size: boltembeddedapi.String("Large"),
                        Sku: boltembeddedapi.String("BOLT-SKU_100"),
                        Source: boltembeddedapi.String("atque"),
                        Tags: boltembeddedapi.String("tote, blue, linen, eco-friendly"),
                        TaxAmount: boltembeddedapi.Float64(0),
                        TaxCode: boltembeddedapi.String("sit"),
                        Taxable: boltembeddedapi.Bool(false),
                        TotalAmount: 1000,
                        Type: shared.CartItemTypeBundled.ToPointer(),
                        UnitPrice: 1000,
                        Uom: boltembeddedapi.String("inches"),
                        Upc: boltembeddedapi.String("825764603119"),
                        Weight: boltembeddedapi.Float64(10),
                        WeightUnit: boltembeddedapi.String("pounds"),
                    },
                },
                LoyaltyRewards: []shared.CartLoyaltyRewards{
                    shared.CartLoyaltyRewards{
                        Amount: boltembeddedapi.Float64(672.49),
                        CouponCode: boltembeddedapi.String("soluta"),
                        Description: boltembeddedapi.String("$5 off (100 Points)"),
                        Details: boltembeddedapi.String("{"id": 123456, "icon": "fa-dollar", "name": "$15.00 Off", "type": "Coupon", "amount": 100, "duration": "single_use", "cost_text": "150 Points",  "description": "Get $15 off your next purchase for 150 points", "discount_type": "fixed_amount", "unrendered_name": "$15.00 Off",  "discount_percentage": null, "discount_rate_cents": null, "discount_value_cents": null, "discount_amount_cents": 1500,  "unrendered_description": "Get $15 off your next purchase for 150 points", "applies_to_product_type": "ALL"}"),
                        Points: boltembeddedapi.Float64(6793.93),
                        Source: boltembeddedapi.String("iusto"),
                        Type: boltembeddedapi.String("voluptate"),
                    },
                },
                Metadata: map[string]string{
                    "dolorum": "deleniti",
                },
                OrderDescription: boltembeddedapi.String("Order #1234567890"),
                OrderReference: "order_100",
                Shipments: []shared.CartShipment{
                    shared.CartShipment{
                        Carrier: boltembeddedapi.String("FedEx"),
                        Cost: boltembeddedapi.Int64(770),
                        DiscountedByMembership: boltembeddedapi.Bool(false),
                        EstimatedDeliveryDate: boltembeddedapi.String("08-30-2022"),
                        Expedited: boltembeddedapi.Bool(false),
                        PackageDepth: boltembeddedapi.Int64(90),
                        PackageDimensionUnit: boltembeddedapi.String("cm"),
                        PackageHeight: boltembeddedapi.Int64(103),
                        PackageType: boltembeddedapi.String("A big package."),
                        PackageWeightUnit: boltembeddedapi.String("kg"),
                        PackageWidth: boltembeddedapi.Int64(222),
                        Service: boltembeddedapi.String("Option 1"),
                        ShippingAddress: &shared.Address{
                            Company: boltembeddedapi.String("Bolt"),
                            Country: boltembeddedapi.String("United States"),
                            CountryCode: "US",
                            Default: boltembeddedapi.Bool(true),
                            DoorCode: boltembeddedapi.String("123456"),
                            Email: "alan.watts@example.com",
                            FirstName: "Alan",
                            LastName: "Watts",
                            Locality: "Brooklyn",
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
                        ShippingAddressID: boltembeddedapi.String("addres-1"),
                        ShippingMethod: boltembeddedapi.String("Unknown"),
                        Signature: boltembeddedapi.String("a1B2s3dC4f5g5D6hj6E7k8F9l0"),
                        TaxAmount: boltembeddedapi.Int64(230),
                        TaxCode: boltembeddedapi.String("tax-12345"),
                        TotalWeight: boltembeddedapi.Int64(55),
                        TotalWeightUnit: boltembeddedapi.String("kg"),
                        Type: shared.CartShipmentTypeDoorDelivery.ToPointer(),
                    },
                },
                TaxAmount: boltembeddedapi.Float64(6070.45),
                TotalAmount: 900,
            },
            ShopperIdentity: &operations.UpdatePaymentRequestBodyShopperIdentity{
                CreateBoltAccount: boltembeddedapi.Bool(true),
                Email: "Michaela.Zboncak@hotmail.com",
                FirstName: "Jacky",
                LastName: "Ondricka",
                Phone: "1-410-378-3936 x53856",
            },
        },
        XPublishableKey: boltembeddedapi.String("alias"),
        ID: "d446ce2a-f7a7-43cf-bbe4-53f870b326b5",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdatePayment200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdatePaymentRequest](../../models/operations/updatepaymentrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.UpdatePaymentSecurity](../../models/operations/updatepaymentsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.UpdatePaymentResponse](../../models/operations/updatepaymentresponse.md), error**

