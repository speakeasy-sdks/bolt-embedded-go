# Transactions
(*Transactions*)

## Overview

Authorize credit card transactions and perform operations on those transactions with Bolt's transaction API.


### Available Operations

* [AuthorizeTransaction](#authorizetransaction) - Authorize a Card
* [CaptureTransaction](#capturetransaction) - Capture a Transaction
* [GetTransactionDetails](#gettransactiondetails) - Transaction Details
* [RefundTransaction](#refundtransaction) - Refund a Transaction
* [UpdateTransaction](#updatetransaction) - Update a Transaction
* [VoidTransaction](#voidtransaction) - Void a Transaction

## AuthorizeTransaction

This endpoint authorizes card payments and has three main use cases:
* • Authorize a payment using an unsaved payment method for a guest or logged-in shopper.
* • Authorize a payment using a saved payment method for a logged-in shopper.
*  • Re-charge a previous transaction using the `credit_card_id` of the transaction.


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


    operationSecurity := operations.AuthorizeTransactionSecurity{
            OAuth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Transactions.AuthorizeTransaction(ctx, operations.AuthorizeTransactionRequest{
        RequestBody: operations.CreateAuthorizeTransactionRequestBodyMerchantCreditCardAuthorization(
                shared.MerchantCreditCardAuthorization{
                    Cart: shared.CartCreate{
                        AddOns: []shared.CartAddOn{
                            shared.CartAddOn{
                                Name: "string",
                                Price: 2208.11,
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
                                Quantity: 7162.31,
                                Reference: "ItemFee",
                                UnitPrice: 7867.96,
                                UnitTaxAmount: 7842.96,
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
                    CreateBoltAccount: false,
                    CreditCard: shared.CreditCard{
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
                        Bin: "411111",
                        Expiration: "2025-11",
                        Last4: "1234",
                        PostalCode: "10044",
                        Token: "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
                        TokenType: shared.TokenTypeBolt,
                    },
                    DivisionID: "4ab56ad7865ada4ad32",
                    MerchantEventID: boltembeddedgo.String("dbe0cd5d-3261-41d9-ba61-49e5b9d07567"),
                    PreviousTransactionID: boltembeddedgo.String("string"),
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
                    Source: shared.SourceDirectPayments,
                    UserIdentifier: shared.UserIdentifier{
                        Artifact: boltembeddedgo.String("string"),
                        Email: boltembeddedgo.String("alan.watts@example.com"),
                        EmailID: boltembeddedgo.String("string"),
                        Phone: "+12125550199",
                        PhoneID: boltembeddedgo.String("string"),
                    },
                    UserIdentity: shared.UserIdentity{
                        FirstName: boltembeddedgo.String("Charlotte"),
                        LastName: boltembeddedgo.String("Charles"),
                    },
                },
        ),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.IAuthorizeResultView != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.AuthorizeTransactionRequest](../../pkg/models/operations/authorizetransactionrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.AuthorizeTransactionSecurity](../../pkg/models/operations/authorizetransactionsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.AuthorizeTransactionResponse](../../pkg/models/operations/authorizetransactionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CaptureTransaction

This captures funds for the designated transaction. A capture can be done for any partial amount or for the total authorized amount.

Although the response returns the standard `transaction_view` object, only `captures` and either `id` or `reference` are needed.


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


    operationSecurity := operations.CaptureTransactionSecurity{
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Transactions.CaptureTransaction(ctx, operations.CaptureTransactionRequest{
        CaptureTransactionWithReference: &shared.CaptureTransactionWithReference{
            Amount: 754,
            Currency: "USD",
            MerchantEventID: boltembeddedgo.String("dbe0cd5d-3261-41d9-ba61-49e5b9d07567"),
            SkipHookNotification: boltembeddedgo.Bool(false),
            TransactionReference: "LBLJ-TWW7-R9VC",
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionView != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CaptureTransactionRequest](../../pkg/models/operations/capturetransactionrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.CaptureTransactionSecurity](../../pkg/models/operations/capturetransactionsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.CaptureTransactionResponse](../../pkg/models/operations/capturetransactionresponse.md), error**
| Error Object                             | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.ErrorsBoltAPIResponse          | 403,404                                  | application/json                         |
| sdkerrors.CaptureTransactionResponseBody | 422                                      | application/json                         |
| sdkerrors.SDKError                       | 400-600                                  | */*                                      |

## GetTransactionDetails

This allows you to pull the full transaction details for a given transaction.

 **Note**: All objects and fields marked `required` in the Transaction Details response are also **nullable**. This includes any sub-components (objects or fields) also marked `required`.


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


    operationSecurity := operations.GetTransactionDetailsSecurity{
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Transactions.GetTransactionDetails(ctx, operations.GetTransactionDetailsRequest{
        Reference: "string",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetTransactionDetailsRequest](../../pkg/models/operations/gettransactiondetailsrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.GetTransactionDetailsSecurity](../../pkg/models/operations/gettransactiondetailssecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


### Response

**[*operations.GetTransactionDetailsResponse](../../pkg/models/operations/gettransactiondetailsresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorsBoltAPIResponse | 403,422                         | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

## RefundTransaction

This refunds a captured transaction. Refunds can be done for any partial amount or for the total authorized amount. These refunds are processed synchronously and return information about the refunded transaction in the standard `transaction_view` object.

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


    operationSecurity := operations.RefundTransactionSecurity{
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Transactions.RefundTransaction(ctx, operations.RefundTransactionRequest{
        RequestBody: &operations.RefundTransactionRequestBody{
            Amount: 754,
            Currency: "USD",
            MerchantEventID: boltembeddedgo.String("dbe0cd5d-3261-41d9-ba61-49e5b9d07567"),
            SkipHookNotification: boltembeddedgo.Bool(false),
            TransactionReference: "LBLJ-TWW7-R9VC",
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionView != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RefundTransactionRequest](../../pkg/models/operations/refundtransactionrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.RefundTransactionSecurity](../../pkg/models/operations/refundtransactionsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.RefundTransactionResponse](../../pkg/models/operations/refundtransactionresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorsBoltAPIResponse | 422                             | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

## UpdateTransaction

This allows you to update certain transaction properties post-authorization.

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


    operationSecurity := operations.UpdateTransactionSecurity{
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Transactions.UpdateTransaction(ctx, operations.UpdateTransactionRequest{
        Reference: "string",
        RequestBody: &operations.UpdateTransactionRequestBody{
            DisplayID: boltembeddedgo.String("order-123"),
            Metadata: map[string]string{
                "key1": "value1",
                "key2": "value2",
            },
        },
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateTransactionRequest](../../pkg/models/operations/updatetransactionrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.UpdateTransactionSecurity](../../pkg/models/operations/updatetransactionsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.UpdateTransactionResponse](../../pkg/models/operations/updatetransactionresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorsBoltAPIResponse | 403,404                         | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

## VoidTransaction

This voids the authorization for a given transaction. Voids must be completed before the authorization is captured.
In the request, either `transaction_id` or `transaction_reference` is required.
Although the response returns the standard `transaction_view` object, only `status` and either `id` or `reference` are needed.


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


    operationSecurity := operations.VoidTransactionSecurity{
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Transactions.VoidTransaction(ctx, operations.VoidTransactionRequest{
        CreditCardVoid: &shared.CreditCardVoid{
            MerchantEventID: boltembeddedgo.String("dbe0cd5d-3261-41d9-ba61-49e5b9d07567"),
            SkipHookNotification: boltembeddedgo.Bool(false),
            TransactionReference: "LBLJ-TWW7-R9VC",
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionView != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.VoidTransactionRequest](../../pkg/models/operations/voidtransactionrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.VoidTransactionSecurity](../../pkg/models/operations/voidtransactionsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.VoidTransactionResponse](../../pkg/models/operations/voidtransactionresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorsBoltAPIResponse | 403,404                         | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |
