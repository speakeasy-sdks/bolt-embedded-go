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
        IdempotencyKey: boltembeddedgo.String("triumphantly generating SQL"),
        RequestBody: &operations.FinalizePaymentRequestBody{
            MerchantEventID: boltembeddedgo.String("dbe0cd5d-3261-41d9-ba61-49e5b9d07567"),
            ShopperIdentity: &operations.FinalizePaymentRequestBodyShopperIdentity{
                CreateBoltAccount: boltembeddedgo.Bool(true),
                Email: "Janie_Ondricka6@hotmail.com",
                FirstName: "Bryana",
                LastName: "Schuster",
                Phone: "(576) 993-3446 x04190",
            },
        },
        ID: "<ID>",
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
        IdempotencyKey: boltembeddedgo.String("Bicycle"),
        RequestBody: &operations.InitializePaymentRequestBody{
            Cart: shared.CartCreate{
                AddOns: []shared.CartAddOn{
                    shared.CartAddOn{
                        Description: boltembeddedgo.String("Cloned actuating knowledge user"),
                        ImageURL: boltembeddedgo.String("Auto"),
                        Name: "Plastic IP Architect",
                        Price: 1981.96,
                        ProductID: "sketch silently",
                        ProductPageURL: boltembeddedgo.String("red wireless Recumbent"),
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
                        Quantity: 4604.82,
                        Reference: "ItemFee",
                        UnitPrice: 9895.87,
                        UnitTaxAmount: 4855.31,
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
                                            "ratione": "Northeast",
                                        },
                                        Name: boltembeddedgo.String("simplistic Berkshire"),
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
                                    ShopifyLineItemReference: boltembeddedgo.Float64(2484.13),
                                    ShopifyProductReference: boltembeddedgo.Float64(9036.66),
                                    ShopifyProductVariantReference: boltembeddedgo.Float64(2155.16),
                                },
                                GiftOption: &shared.CartItemGiftOption{
                                    Cost: boltembeddedgo.Int64(770),
                                    MerchantProductID: boltembeddedgo.String("881"),
                                    Message: boltembeddedgo.String("Happy Anniversary, Smoochy Poo!"),
                                    Wrap: boltembeddedgo.Bool(false),
                                },
                                ImageURL: boltembeddedgo.String("https://boltswagstore.com/products/123456/images/1.png"),
                                Isbn: boltembeddedgo.String("9780091347314"),
                                ItemGroup: boltembeddedgo.String("Cisgender female"),
                                Manufacturer: boltembeddedgo.String("Bolt Textiles USA"),
                                MerchantProductID: boltembeddedgo.String("881"),
                                MerchantVariantID: boltembeddedgo.String("888"),
                                Msrp: boltembeddedgo.Float64(6633.02),
                                Name: "Bolt Swag Bag",
                                Options: boltembeddedgo.String("Special Edition"),
                                Properties: []shared.CartItemProperty{
                                    shared.CartItemProperty{
                                        Color: boltembeddedgo.String("yellow"),
                                        Display: boltembeddedgo.Bool(false),
                                        Name: boltembeddedgo.String("invoice programming"),
                                        NameID: boltembeddedgo.Float64(7703.73),
                                        Value: boltembeddedgo.String("open"),
                                        ValueID: boltembeddedgo.Float64(8814.48),
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
                                ShipmentType: shared.CartItemShipmentTypeShipToStore.ToPointer(),
                                Size: boltembeddedgo.String("Large"),
                                Sku: boltembeddedgo.String("BOLT-SKU_100"),
                                Source: boltembeddedgo.String("Strategist THX"),
                                Tags: boltembeddedgo.String("tote, blue, linen, eco-friendly"),
                                TaxAmount: boltembeddedgo.Float64(0),
                                TaxCode: boltembeddedgo.String("Sedan haptic"),
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
                            Email: boltembeddedgo.String("Weldon4@gmail.com"),
                            Phone: boltembeddedgo.String("459.697.0696 x62730"),
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
                            PickupWindowClose: boltembeddedgo.Int64(344315),
                            PickupWindowOpen: boltembeddedgo.Int64(707410),
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
                        PickupWindowClose: boltembeddedgo.Int64(14054),
                        PickupWindowOpen: boltembeddedgo.Int64(692493),
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
                                    "dignissimos": "male",
                                },
                                Name: boltembeddedgo.String("Future Altenwerth"),
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
                            ShopifyLineItemReference: boltembeddedgo.Float64(1650.3),
                            ShopifyProductReference: boltembeddedgo.Float64(7121.75),
                            ShopifyProductVariantReference: boltembeddedgo.Float64(5414.12),
                        },
                        GiftOption: &shared.CartItemGiftOption{
                            Cost: boltembeddedgo.Int64(770),
                            MerchantProductID: boltembeddedgo.String("881"),
                            Message: boltembeddedgo.String("Happy Anniversary, Smoochy Poo!"),
                            Wrap: boltembeddedgo.Bool(false),
                        },
                        ImageURL: boltembeddedgo.String("https://boltswagstore.com/products/123456/images/1.png"),
                        Isbn: boltembeddedgo.String("9780091347314"),
                        ItemGroup: boltembeddedgo.String("Virginia"),
                        Manufacturer: boltembeddedgo.String("Bolt Textiles USA"),
                        MerchantProductID: boltembeddedgo.String("881"),
                        MerchantVariantID: boltembeddedgo.String("888"),
                        Msrp: boltembeddedgo.Float64(4440.99),
                        Name: "Bolt Swag Bag",
                        Options: boltembeddedgo.String("Special Edition"),
                        Properties: []shared.CartItemProperty{
                            shared.CartItemProperty{
                                Color: boltembeddedgo.String("mint green"),
                                Display: boltembeddedgo.Bool(false),
                                Name: boltembeddedgo.String("Principal"),
                                NameID: boltembeddedgo.Float64(1271.55),
                                Value: boltembeddedgo.String("Facilitator female"),
                                ValueID: boltembeddedgo.Float64(8076.7),
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
                        Source: boltembeddedgo.String("Jazz"),
                        Tags: boltembeddedgo.String("tote, blue, linen, eco-friendly"),
                        TaxAmount: boltembeddedgo.Float64(0),
                        TaxCode: boltembeddedgo.String("Folk Porsche Berkshire"),
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
                LoyaltyRewards: []shared.CartLoyaltyRewards{
                    shared.CartLoyaltyRewards{
                        Amount: boltembeddedgo.Float64(1851.98),
                        CouponCode: boltembeddedgo.String("newton offering Infrastructure"),
                        Description: boltembeddedgo.String("$5 off (100 Points)"),
                        Details: boltembeddedgo.String("{\"id\": 123456, \"icon\": \"fa-dollar\", \"name\": \"$15.00 Off\", \"type\": \"Coupon\", \"amount\": 100, \"duration\": \"single_use\", \"cost_text\": \"150 Points\",  \"description\": \"Get $15 off your next purchase for 150 points\", \"discount_type\": \"fixed_amount\", \"unrendered_name\": \"$15.00 Off\",  \"discount_percentage\": null, \"discount_rate_cents\": null, \"discount_value_cents\": null, \"discount_amount_cents\": 1500,  \"unrendered_description\": \"Get $15 off your next purchase for 150 points\", \"applies_to_product_type\": \"ALL\"}"),
                        Points: boltembeddedgo.Float64(811.42),
                        Source: boltembeddedgo.String("Fiat by"),
                        Type: boltembeddedgo.String("hertz blah integrated"),
                    },
                },
                Metadata: map[string]string{
                    "aliquid": "Cheese",
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
                TaxAmount: boltembeddedgo.Float64(2899.19),
                TotalAmount: 900,
            },
            ShopperIdentity: &operations.InitializePaymentRequestBodyShopperIdentity{
                CreateBoltAccount: boltembeddedgo.Bool(true),
                Email: "Violet92@hotmail.com",
                FirstName: "Stephen",
                LastName: "Powlowski",
                Phone: "(354) 487-8136 x6300",
            },
        },
        XPublishableKey: boltembeddedgo.String("sky evolve"),
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
        IdempotencyKey: boltembeddedgo.String("Northeast Pataca OCR"),
        RequestBody: &operations.UpdatePaymentRequestBody{
            Cart: &shared.CartCreate{
                AddOns: []shared.CartAddOn{
                    shared.CartAddOn{
                        Description: boltembeddedgo.String("Open-source contextually-based access"),
                        ImageURL: boltembeddedgo.String("kelvin bus gadzooks"),
                        Name: "Tennessee Loan lumen",
                        Price: 621.91,
                        ProductID: "Southwest ah",
                        ProductPageURL: boltembeddedgo.String("Wagon calculate"),
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
                        DiscountCategory: shared.CartDiscountDiscountCategoryManagedGiftcard.ToPointer(),
                        Reference: boltembeddedgo.String("DISC-1234"),
                        Type: shared.CartDiscountTypePercentage.ToPointer(),
                    },
                },
                DisplayID: boltembeddedgo.String("displayid_100"),
                Fees: []shared.CartCreateFees{
                    shared.CartCreateFees{
                        Description: boltembeddedgo.String("Item Fee"),
                        Name: "Item Fee",
                        Quantity: 6255.34,
                        Reference: "ItemFee",
                        UnitPrice: 3605.33,
                        UnitTaxAmount: 2097.19,
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
                                            "cupiditate": "female",
                                        },
                                        Name: boltembeddedgo.String("Platinum finally"),
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
                                    ShopifyLineItemReference: boltembeddedgo.Float64(1604.67),
                                    ShopifyProductReference: boltembeddedgo.Float64(2435.65),
                                    ShopifyProductVariantReference: boltembeddedgo.Float64(3021.98),
                                },
                                GiftOption: &shared.CartItemGiftOption{
                                    Cost: boltembeddedgo.Int64(770),
                                    MerchantProductID: boltembeddedgo.String("881"),
                                    Message: boltembeddedgo.String("Happy Anniversary, Smoochy Poo!"),
                                    Wrap: boltembeddedgo.Bool(false),
                                },
                                ImageURL: boltembeddedgo.String("https://boltswagstore.com/products/123456/images/1.png"),
                                Isbn: boltembeddedgo.String("9780091347314"),
                                ItemGroup: boltembeddedgo.String("Face HTTP"),
                                Manufacturer: boltembeddedgo.String("Bolt Textiles USA"),
                                MerchantProductID: boltembeddedgo.String("881"),
                                MerchantVariantID: boltembeddedgo.String("888"),
                                Msrp: boltembeddedgo.Float64(7950.1),
                                Name: "Bolt Swag Bag",
                                Options: boltembeddedgo.String("Special Edition"),
                                Properties: []shared.CartItemProperty{
                                    shared.CartItemProperty{
                                        Color: boltembeddedgo.String("black"),
                                        Display: boltembeddedgo.Bool(false),
                                        Name: boltembeddedgo.String("zero Human"),
                                        NameID: boltembeddedgo.Float64(1634.01),
                                        Value: boltembeddedgo.String("Dynamic"),
                                        ValueID: boltembeddedgo.Float64(6209.8),
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
                                Source: boltembeddedgo.String("facilitate"),
                                Tags: boltembeddedgo.String("tote, blue, linen, eco-friendly"),
                                TaxAmount: boltembeddedgo.Float64(0),
                                TaxCode: boltembeddedgo.String("ivory"),
                                Taxable: boltembeddedgo.Bool(false),
                                TotalAmount: 1000,
                                Type: shared.CartItemTypeDigital.ToPointer(),
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
                            Email: boltembeddedgo.String("Tatyana_Kessler25@hotmail.com"),
                            Phone: boltembeddedgo.String("968-639-4285 x877"),
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
                            PickupWindowClose: boltembeddedgo.Int64(358861),
                            PickupWindowOpen: boltembeddedgo.Int64(924450),
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
                        PickupWindowClose: boltembeddedgo.Int64(428066),
                        PickupWindowOpen: boltembeddedgo.Int64(597057),
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
                                    "tenetur": "Northwest",
                                },
                                Name: boltembeddedgo.String("Beauty siemens"),
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
                            ShopifyLineItemReference: boltembeddedgo.Float64(1393.18),
                            ShopifyProductReference: boltembeddedgo.Float64(2477.99),
                            ShopifyProductVariantReference: boltembeddedgo.Float64(2493.09),
                        },
                        GiftOption: &shared.CartItemGiftOption{
                            Cost: boltembeddedgo.Int64(770),
                            MerchantProductID: boltembeddedgo.String("881"),
                            Message: boltembeddedgo.String("Happy Anniversary, Smoochy Poo!"),
                            Wrap: boltembeddedgo.Bool(false),
                        },
                        ImageURL: boltembeddedgo.String("https://boltswagstore.com/products/123456/images/1.png"),
                        Isbn: boltembeddedgo.String("9780091347314"),
                        ItemGroup: boltembeddedgo.String("withdrawal Southeast"),
                        Manufacturer: boltembeddedgo.String("Bolt Textiles USA"),
                        MerchantProductID: boltembeddedgo.String("881"),
                        MerchantVariantID: boltembeddedgo.String("888"),
                        Msrp: boltembeddedgo.Float64(8791.42),
                        Name: "Bolt Swag Bag",
                        Options: boltembeddedgo.String("Special Edition"),
                        Properties: []shared.CartItemProperty{
                            shared.CartItemProperty{
                                Color: boltembeddedgo.String("indigo"),
                                Display: boltembeddedgo.Bool(false),
                                Name: boltembeddedgo.String("overriding"),
                                NameID: boltembeddedgo.Float64(2419.8),
                                Value: boltembeddedgo.String("holistic salmon"),
                                ValueID: boltembeddedgo.Float64(4786.08),
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
                        Source: boltembeddedgo.String("Home transmitting"),
                        Tags: boltembeddedgo.String("tote, blue, linen, eco-friendly"),
                        TaxAmount: boltembeddedgo.Float64(0),
                        TaxCode: boltembeddedgo.String("nonconforming non although"),
                        Taxable: boltembeddedgo.Bool(false),
                        TotalAmount: 1000,
                        Type: shared.CartItemTypeDigital.ToPointer(),
                        UnitPrice: 1000,
                        Uom: boltembeddedgo.String("inches"),
                        Upc: boltembeddedgo.String("825764603119"),
                        Weight: boltembeddedgo.Float64(10),
                        WeightUnit: boltembeddedgo.String("pounds"),
                    },
                },
                LoyaltyRewards: []shared.CartLoyaltyRewards{
                    shared.CartLoyaltyRewards{
                        Amount: boltembeddedgo.Float64(9114.8),
                        CouponCode: boltembeddedgo.String("synergies Hop deposit"),
                        Description: boltembeddedgo.String("$5 off (100 Points)"),
                        Details: boltembeddedgo.String("{\"id\": 123456, \"icon\": \"fa-dollar\", \"name\": \"$15.00 Off\", \"type\": \"Coupon\", \"amount\": 100, \"duration\": \"single_use\", \"cost_text\": \"150 Points\",  \"description\": \"Get $15 off your next purchase for 150 points\", \"discount_type\": \"fixed_amount\", \"unrendered_name\": \"$15.00 Off\",  \"discount_percentage\": null, \"discount_rate_cents\": null, \"discount_value_cents\": null, \"discount_amount_cents\": 1500,  \"unrendered_description\": \"Get $15 off your next purchase for 150 points\", \"applies_to_product_type\": \"ALL\"}"),
                        Points: boltembeddedgo.Float64(8306.32),
                        Source: boltembeddedgo.String("Stracke"),
                        Type: boltembeddedgo.String("interactive saepe North"),
                    },
                },
                Metadata: map[string]string{
                    "beatae": "Sulfur",
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
                TaxAmount: boltembeddedgo.Float64(1438.01),
                TotalAmount: 900,
            },
            ShopperIdentity: &operations.UpdatePaymentRequestBodyShopperIdentity{
                CreateBoltAccount: boltembeddedgo.Bool(true),
                Email: "Cameron_Borer@hotmail.com",
                FirstName: "Mona",
                LastName: "Maggio",
                Phone: "580-305-9242 x56905",
            },
        },
        XPublishableKey: boltembeddedgo.String("Pakistan hack"),
        ID: "<ID>",
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

