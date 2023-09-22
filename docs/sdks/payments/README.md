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
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
)

func main() {
    s := boltembeddedgo.New()
    operationSecurity := operations.FinalizePaymentSecurity{
            OAuth: boltembeddedgo.String(""),
        }

    ctx := context.Background()
    res, err := s.Payments.FinalizePayment(ctx, operations.FinalizePaymentRequest{
        IdempotencyKey: boltembeddedgo.String("doloremque"),
        RequestBody: &operations.FinalizePaymentRequestBody{
            MerchantEventID: boltembeddedgo.String("dbe0cd5d-3261-41d9-ba61-49e5b9d07567"),
            ShopperIdentity: &operations.FinalizePaymentRequestBodyShopperIdentity{
                CreateBoltAccount: boltembeddedgo.Bool(true),
                Email: "Easter35@yahoo.com",
                FirstName: "Elena",
                LastName: "Kshlerin",
                Phone: "(738) 590-2655",
            },
        },
        ID: "d488e1e9-1e45-40ad-aabd-44269802d502",
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
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
)

func main() {
    s := boltembeddedgo.New()
    operationSecurity := operations.InitializePaymentSecurity{
            OAuth: boltembeddedgo.String(""),
        }

    ctx := context.Background()
    res, err := s.Payments.InitializePayment(ctx, operations.InitializePaymentRequest{
        IdempotencyKey: boltembeddedgo.String("dolorum"),
        RequestBody: &operations.InitializePaymentRequestBody{
            Cart: shared.CartCreate{
                AddOns: []shared.CartAddOn{
                    shared.CartAddOn{
                        Description: boltembeddedgo.String("excepturi"),
                        ImageURL: boltembeddedgo.String("tempora"),
                        Name: "Geoffrey Green",
                        Price: 2487.53,
                        ProductID: "eligendi",
                        ProductPageURL: boltembeddedgo.String("sint"),
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
                        DiscountCategory: shared.CartDiscountDiscountCategoryStoreCredit.ToPointer(),
                        Reference: boltembeddedgo.String("DISC-1234"),
                        Type: shared.CartDiscountTypePercentage.ToPointer(),
                    },
                },
                DisplayID: boltembeddedgo.String("displayid_100"),
                Fees: []shared.CartCreateFees{
                    shared.CartCreateFees{
                        Description: boltembeddedgo.String("Item Fee"),
                        Name: "Item Fee",
                        Quantity: 5920.42,
                        Reference: "ItemFee",
                        UnitPrice: 8960.39,
                        UnitTaxAmount: 5722.52,
                    },
                },
                Fulfillments: []shared.CartCreateFulfillments{
                    shared.CartCreateFulfillments{
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
                                            "officia": "dolor",
                                        },
                                        Name: boltembeddedgo.String("Randal Parisian"),
                                        Price: &shared.AmountView{
                                            Amount: boltembeddedgo.Float64(754),
                                            Currency: boltembeddedgo.String("USD"),
                                            CurrencySymbol: boltembeddedgo.String("$"),
                                        },
                                    },
                                },
                                Description: boltembeddedgo.String("Large tote with Bolt logo."),
                                DetailsURL: boltembeddedgo.String("https://boltswagstore.com/products/123456"),
                                ExternalInputs: &shared.ICartItemExternalInputs{
                                    ShopifyLineItemReference: boltembeddedgo.Float64(8464.09),
                                    ShopifyProductReference: boltembeddedgo.Float64(9785.71),
                                    ShopifyProductVariantReference: boltembeddedgo.Float64(6994.79),
                                },
                                GiftOption: &shared.CartItemGiftOption{
                                    Cost: boltembeddedgo.Int64(770),
                                    MerchantProductID: boltembeddedgo.String("881"),
                                    Message: boltembeddedgo.String("Happy Anniversary, Smoochy Poo!"),
                                    Wrap: boltembeddedgo.Bool(false),
                                },
                                ImageURL: boltembeddedgo.String("https://boltswagstore.com/products/123456/images/1.png"),
                                Isbn: boltembeddedgo.String("9780091347314"),
                                ItemGroup: boltembeddedgo.String("dicta"),
                                Manufacturer: boltembeddedgo.String("Bolt Textiles USA"),
                                MerchantProductID: boltembeddedgo.String("881"),
                                MerchantVariantID: boltembeddedgo.String("888"),
                                Msrp: boltembeddedgo.Float64(2974.37),
                                Name: "Bolt Swag Bag",
                                Options: boltembeddedgo.String("Special Edition"),
                                Properties: []shared.CartItemProperty{
                                    shared.CartItemProperty{
                                        Color: boltembeddedgo.String("cumque"),
                                        Display: boltembeddedgo.Bool(false),
                                        Name: boltembeddedgo.String("Nathaniel Hyatt"),
                                        NameID: boltembeddedgo.Float64(2497.96),
                                        Value: boltembeddedgo.String("occaecati"),
                                        ValueID: boltembeddedgo.Float64(3132.18),
                                    },
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
                                ShipmentType: shared.CartItemShipmentTypeInStorePickup.ToPointer(),
                                Size: boltembeddedgo.String("Large"),
                                Sku: boltembeddedgo.String("BOLT-SKU_100"),
                                Source: boltembeddedgo.String("delectus"),
                                Tags: boltembeddedgo.String("tote, blue, linen, eco-friendly"),
                                TaxAmount: boltembeddedgo.Float64(0),
                                TaxCode: boltembeddedgo.String("quidem"),
                                Taxable: boltembeddedgo.Bool(false),
                                TotalAmount: 1000,
                                Type: shared.CartItemTypePhysical.ToPointer(),
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
                        DigitalDelivery: &shared.CartCreateFulfillmentsDigitalDelivery{
                            Email: boltembeddedgo.String("Magdalena_Kuvalis@hotmail.com"),
                            Phone: boltembeddedgo.String("734.664.0437"),
                        },
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
                            DistanceUnit: shared.InStoreCartShipmentDistanceUnitMile.ToPointer(),
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
                            PickupWindowClose: boltembeddedgo.Int64(660174),
                            PickupWindowOpen: boltembeddedgo.Int64(287991),
                            StoreName: boltembeddedgo.String("Bolt Collective"),
                        },
                        Type: shared.CartCreateFulfillmentsTypePhysicalShipToStore.ToPointer(),
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
                        DistanceUnit: shared.InStoreCartShipmentDistanceUnitMile.ToPointer(),
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
                        PickupWindowClose: boltembeddedgo.Int64(383462),
                        PickupWindowOpen: boltembeddedgo.Int64(618016),
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
                                    "nobis": "eum",
                                },
                                Name: boltembeddedgo.String("Brandon Brakus V"),
                                Price: &shared.AmountView{
                                    Amount: boltembeddedgo.Float64(754),
                                    Currency: boltembeddedgo.String("USD"),
                                    CurrencySymbol: boltembeddedgo.String("$"),
                                },
                            },
                        },
                        Description: boltembeddedgo.String("Large tote with Bolt logo."),
                        DetailsURL: boltembeddedgo.String("https://boltswagstore.com/products/123456"),
                        ExternalInputs: &shared.ICartItemExternalInputs{
                            ShopifyLineItemReference: boltembeddedgo.Float64(3540.47),
                            ShopifyProductReference: boltembeddedgo.Float64(5908.73),
                            ShopifyProductVariantReference: boltembeddedgo.Float64(5518.16),
                        },
                        GiftOption: &shared.CartItemGiftOption{
                            Cost: boltembeddedgo.Int64(770),
                            MerchantProductID: boltembeddedgo.String("881"),
                            Message: boltembeddedgo.String("Happy Anniversary, Smoochy Poo!"),
                            Wrap: boltembeddedgo.Bool(false),
                        },
                        ImageURL: boltembeddedgo.String("https://boltswagstore.com/products/123456/images/1.png"),
                        Isbn: boltembeddedgo.String("9780091347314"),
                        ItemGroup: boltembeddedgo.String("sint"),
                        Manufacturer: boltembeddedgo.String("Bolt Textiles USA"),
                        MerchantProductID: boltembeddedgo.String("881"),
                        MerchantVariantID: boltembeddedgo.String("888"),
                        Msrp: boltembeddedgo.Float64(336.25),
                        Name: "Bolt Swag Bag",
                        Options: boltembeddedgo.String("Special Edition"),
                        Properties: []shared.CartItemProperty{
                            shared.CartItemProperty{
                                Color: boltembeddedgo.String("mollitia"),
                                Display: boltembeddedgo.Bool(false),
                                Name: boltembeddedgo.String("Shaun Hammes"),
                                NameID: boltembeddedgo.Float64(8965.47),
                                Value: boltembeddedgo.String("odit"),
                                ValueID: boltembeddedgo.Float64(3675.62),
                            },
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
                        ShipmentType: shared.CartItemShipmentTypeUnknown.ToPointer(),
                        Size: boltembeddedgo.String("Large"),
                        Sku: boltembeddedgo.String("BOLT-SKU_100"),
                        Source: boltembeddedgo.String("iure"),
                        Tags: boltembeddedgo.String("tote, blue, linen, eco-friendly"),
                        TaxAmount: boltembeddedgo.Float64(0),
                        TaxCode: boltembeddedgo.String("doloribus"),
                        Taxable: boltembeddedgo.Bool(false),
                        TotalAmount: 1000,
                        Type: shared.CartItemTypeBundled.ToPointer(),
                        UnitPrice: 1000,
                        Uom: boltembeddedgo.String("inches"),
                        Upc: boltembeddedgo.String("825764603119"),
                        Weight: boltembeddedgo.Float64(10),
                        WeightUnit: boltembeddedgo.String("pounds"),
                    },
                },
                LoyaltyRewards: []shared.CartLoyaltyRewards{
                    shared.CartLoyaltyRewards{
                        Amount: boltembeddedgo.Float64(2603.41),
                        CouponCode: boltembeddedgo.String("maxime"),
                        Description: boltembeddedgo.String("$5 off (100 Points)"),
                        Details: boltembeddedgo.String("{"id": 123456, "icon": "fa-dollar", "name": "$15.00 Off", "type": "Coupon", "amount": 100, "duration": "single_use", "cost_text": "150 Points",  "description": "Get $15 off your next purchase for 150 points", "discount_type": "fixed_amount", "unrendered_name": "$15.00 Off",  "discount_percentage": null, "discount_rate_cents": null, "discount_value_cents": null, "discount_amount_cents": 1500,  "unrendered_description": "Get $15 off your next purchase for 150 points", "applies_to_product_type": "ALL"}"),
                        Points: boltembeddedgo.Float64(5370.23),
                        Source: boltembeddedgo.String("facilis"),
                        Type: boltembeddedgo.String("in"),
                    },
                },
                Metadata: map[string]string{
                    "architecto": "architecto",
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
                TaxAmount: boltembeddedgo.Float64(9194.83),
                TotalAmount: 900,
            },
            ShopperIdentity: &operations.InitializePaymentRequestBodyShopperIdentity{
                CreateBoltAccount: boltembeddedgo.Bool(true),
                Email: "Michael_Koss14@yahoo.com",
                FirstName: "Sunny",
                LastName: "Stroman",
                Phone: "356-317-8884",
            },
        },
        XPublishableKey: boltembeddedgo.String("excepturi"),
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
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
)

func main() {
    s := boltembeddedgo.New()
    operationSecurity := operations.UpdatePaymentSecurity{
            OAuth: boltembeddedgo.String(""),
        }

    ctx := context.Background()
    res, err := s.Payments.UpdatePayment(ctx, operations.UpdatePaymentRequest{
        IdempotencyKey: boltembeddedgo.String("odit"),
        RequestBody: &operations.UpdatePaymentRequestBody{
            Cart: &shared.CartCreate{
                AddOns: []shared.CartAddOn{
                    shared.CartAddOn{
                        Description: boltembeddedgo.String("ea"),
                        ImageURL: boltembeddedgo.String("accusantium"),
                        Name: "Ebony Predovic",
                        Price: 4200.75,
                        ProductID: "nam",
                        ProductPageURL: boltembeddedgo.String("eaque"),
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
                        DiscountCategory: shared.CartDiscountDiscountCategoryRewardsDiscount.ToPointer(),
                        Reference: boltembeddedgo.String("DISC-1234"),
                        Type: shared.CartDiscountTypePercentage.ToPointer(),
                    },
                },
                DisplayID: boltembeddedgo.String("displayid_100"),
                Fees: []shared.CartCreateFees{
                    shared.CartCreateFees{
                        Description: boltembeddedgo.String("Item Fee"),
                        Name: "Item Fee",
                        Quantity: 3654.96,
                        Reference: "ItemFee",
                        UnitPrice: 9755.22,
                        UnitTaxAmount: 166.27,
                    },
                },
                Fulfillments: []shared.CartCreateFulfillments{
                    shared.CartCreateFulfillments{
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
                                            "fugiat": "amet",
                                        },
                                        Name: boltembeddedgo.String("Erma Hessel"),
                                        Price: &shared.AmountView{
                                            Amount: boltembeddedgo.Float64(754),
                                            Currency: boltembeddedgo.String("USD"),
                                            CurrencySymbol: boltembeddedgo.String("$"),
                                        },
                                    },
                                },
                                Description: boltembeddedgo.String("Large tote with Bolt logo."),
                                DetailsURL: boltembeddedgo.String("https://boltswagstore.com/products/123456"),
                                ExternalInputs: &shared.ICartItemExternalInputs{
                                    ShopifyLineItemReference: boltembeddedgo.Float64(7499.99),
                                    ShopifyProductReference: boltembeddedgo.Float64(1716.29),
                                    ShopifyProductVariantReference: boltembeddedgo.Float64(3394.04),
                                },
                                GiftOption: &shared.CartItemGiftOption{
                                    Cost: boltembeddedgo.Int64(770),
                                    MerchantProductID: boltembeddedgo.String("881"),
                                    Message: boltembeddedgo.String("Happy Anniversary, Smoochy Poo!"),
                                    Wrap: boltembeddedgo.Bool(false),
                                },
                                ImageURL: boltembeddedgo.String("https://boltswagstore.com/products/123456/images/1.png"),
                                Isbn: boltembeddedgo.String("9780091347314"),
                                ItemGroup: boltembeddedgo.String("totam"),
                                Manufacturer: boltembeddedgo.String("Bolt Textiles USA"),
                                MerchantProductID: boltembeddedgo.String("881"),
                                MerchantVariantID: boltembeddedgo.String("888"),
                                Msrp: boltembeddedgo.Float64(4895.49),
                                Name: "Bolt Swag Bag",
                                Options: boltembeddedgo.String("Special Edition"),
                                Properties: []shared.CartItemProperty{
                                    shared.CartItemProperty{
                                        Color: boltembeddedgo.String("eaque"),
                                        Display: boltembeddedgo.Bool(false),
                                        Name: boltembeddedgo.String("Mr. Robin Daugherty"),
                                        NameID: boltembeddedgo.Float64(4634.51),
                                        Value: boltembeddedgo.String("dolor"),
                                        ValueID: boltembeddedgo.Float64(8745.73),
                                    },
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
                                ShipmentType: shared.CartItemShipmentTypeDoorDelivery.ToPointer(),
                                Size: boltembeddedgo.String("Large"),
                                Sku: boltembeddedgo.String("BOLT-SKU_100"),
                                Source: boltembeddedgo.String("hic"),
                                Tags: boltembeddedgo.String("tote, blue, linen, eco-friendly"),
                                TaxAmount: boltembeddedgo.Float64(0),
                                TaxCode: boltembeddedgo.String("recusandae"),
                                Taxable: boltembeddedgo.Bool(false),
                                TotalAmount: 1000,
                                Type: shared.CartItemTypePhysical.ToPointer(),
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
                        DigitalDelivery: &shared.CartCreateFulfillmentsDigitalDelivery{
                            Email: boltembeddedgo.String("Lauryn.Bartoletti50@hotmail.com"),
                            Phone: boltembeddedgo.String("(256) 399-2665 x8577"),
                        },
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
                            DistanceUnit: shared.InStoreCartShipmentDistanceUnitMile.ToPointer(),
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
                            PickupWindowClose: boltembeddedgo.Int64(964490),
                            PickupWindowOpen: boltembeddedgo.Int64(311945),
                            StoreName: boltembeddedgo.String("Bolt Collective"),
                        },
                        Type: shared.CartCreateFulfillmentsTypePhysicalInStorePickup.ToPointer(),
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
                        DistanceUnit: shared.InStoreCartShipmentDistanceUnitMile.ToPointer(),
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
                        PickupWindowClose: boltembeddedgo.Int64(398221),
                        PickupWindowOpen: boltembeddedgo.Int64(212390),
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
                                    "dolorem": "dolor",
                                },
                                Name: boltembeddedgo.String("Tiffany Welch"),
                                Price: &shared.AmountView{
                                    Amount: boltembeddedgo.Float64(754),
                                    Currency: boltembeddedgo.String("USD"),
                                    CurrencySymbol: boltembeddedgo.String("$"),
                                },
                            },
                        },
                        Description: boltembeddedgo.String("Large tote with Bolt logo."),
                        DetailsURL: boltembeddedgo.String("https://boltswagstore.com/products/123456"),
                        ExternalInputs: &shared.ICartItemExternalInputs{
                            ShopifyLineItemReference: boltembeddedgo.Float64(4521.09),
                            ShopifyProductReference: boltembeddedgo.Float64(4904.59),
                            ShopifyProductVariantReference: boltembeddedgo.Float64(9702.37),
                        },
                        GiftOption: &shared.CartItemGiftOption{
                            Cost: boltembeddedgo.Int64(770),
                            MerchantProductID: boltembeddedgo.String("881"),
                            Message: boltembeddedgo.String("Happy Anniversary, Smoochy Poo!"),
                            Wrap: boltembeddedgo.Bool(false),
                        },
                        ImageURL: boltembeddedgo.String("https://boltswagstore.com/products/123456/images/1.png"),
                        Isbn: boltembeddedgo.String("9780091347314"),
                        ItemGroup: boltembeddedgo.String("amet"),
                        Manufacturer: boltembeddedgo.String("Bolt Textiles USA"),
                        MerchantProductID: boltembeddedgo.String("881"),
                        MerchantVariantID: boltembeddedgo.String("888"),
                        Msrp: boltembeddedgo.Float64(6805.45),
                        Name: "Bolt Swag Bag",
                        Options: boltembeddedgo.String("Special Edition"),
                        Properties: []shared.CartItemProperty{
                            shared.CartItemProperty{
                                Color: boltembeddedgo.String("numquam"),
                                Display: boltembeddedgo.Bool(false),
                                Name: boltembeddedgo.String("Melissa Bednar"),
                                NameID: boltembeddedgo.Float64(3117.96),
                                Value: boltembeddedgo.String("accusamus"),
                                ValueID: boltembeddedgo.Float64(6963.44),
                            },
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
                        ShipmentType: shared.CartItemShipmentTypeInStorePickup.ToPointer(),
                        Size: boltembeddedgo.String("Large"),
                        Sku: boltembeddedgo.String("BOLT-SKU_100"),
                        Source: boltembeddedgo.String("voluptas"),
                        Tags: boltembeddedgo.String("tote, blue, linen, eco-friendly"),
                        TaxAmount: boltembeddedgo.Float64(0),
                        TaxCode: boltembeddedgo.String("natus"),
                        Taxable: boltembeddedgo.Bool(false),
                        TotalAmount: 1000,
                        Type: shared.CartItemTypeUnknown.ToPointer(),
                        UnitPrice: 1000,
                        Uom: boltembeddedgo.String("inches"),
                        Upc: boltembeddedgo.String("825764603119"),
                        Weight: boltembeddedgo.Float64(10),
                        WeightUnit: boltembeddedgo.String("pounds"),
                    },
                },
                LoyaltyRewards: []shared.CartLoyaltyRewards{
                    shared.CartLoyaltyRewards{
                        Amount: boltembeddedgo.Float64(5424.99),
                        CouponCode: boltembeddedgo.String("sit"),
                        Description: boltembeddedgo.String("$5 off (100 Points)"),
                        Details: boltembeddedgo.String("{"id": 123456, "icon": "fa-dollar", "name": "$15.00 Off", "type": "Coupon", "amount": 100, "duration": "single_use", "cost_text": "150 Points",  "description": "Get $15 off your next purchase for 150 points", "discount_type": "fixed_amount", "unrendered_name": "$15.00 Off",  "discount_percentage": null, "discount_rate_cents": null, "discount_value_cents": null, "discount_amount_cents": 1500,  "unrendered_description": "Get $15 off your next purchase for 150 points", "applies_to_product_type": "ALL"}"),
                        Points: boltembeddedgo.Float64(8546.14),
                        Source: boltembeddedgo.String("ab"),
                        Type: boltembeddedgo.String("soluta"),
                    },
                },
                Metadata: map[string]string{
                    "dolorum": "iusto",
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
                TaxAmount: boltembeddedgo.Float64(4536.97),
                TotalAmount: 900,
            },
            ShopperIdentity: &operations.UpdatePaymentRequestBodyShopperIdentity{
                CreateBoltAccount: boltembeddedgo.Bool(true),
                Email: "Kaitlin_Mohr99@hotmail.com",
                FirstName: "Janessa",
                LastName: "Emmerich",
                Phone: "792.302.7839 x365",
            },
        },
        XPublishableKey: boltembeddedgo.String("minima"),
        ID: "d8a0d446-ce2a-4f7a-b3cf-3be453f870b3",
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

