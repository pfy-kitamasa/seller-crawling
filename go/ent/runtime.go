// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/panforyou/seller-finding/go/ent/schema"
	"github.com/panforyou/seller-finding/go/ent/seller"
	"github.com/panforyou/seller-finding/go/ent/website"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	sellerMixin := schema.Seller{}.Mixin()
	sellerMixinFields0 := sellerMixin[0].Fields()
	_ = sellerMixinFields0
	sellerFields := schema.Seller{}.Fields()
	_ = sellerFields
	// sellerDescCreatedAt is the schema descriptor for created_at field.
	sellerDescCreatedAt := sellerMixinFields0[0].Descriptor()
	// seller.DefaultCreatedAt holds the default value on creation for the created_at field.
	seller.DefaultCreatedAt = sellerDescCreatedAt.Default.(func() time.Time)
	// sellerDescUpdatedAt is the schema descriptor for updated_at field.
	sellerDescUpdatedAt := sellerMixinFields0[1].Descriptor()
	// seller.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	seller.DefaultUpdatedAt = sellerDescUpdatedAt.Default.(func() time.Time)
	// seller.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	seller.UpdateDefaultUpdatedAt = sellerDescUpdatedAt.UpdateDefault.(func() time.Time)
	websiteMixin := schema.Website{}.Mixin()
	websiteMixinFields0 := websiteMixin[0].Fields()
	_ = websiteMixinFields0
	websiteFields := schema.Website{}.Fields()
	_ = websiteFields
	// websiteDescCreatedAt is the schema descriptor for created_at field.
	websiteDescCreatedAt := websiteMixinFields0[0].Descriptor()
	// website.DefaultCreatedAt holds the default value on creation for the created_at field.
	website.DefaultCreatedAt = websiteDescCreatedAt.Default.(func() time.Time)
	// websiteDescUpdatedAt is the schema descriptor for updated_at field.
	websiteDescUpdatedAt := websiteMixinFields0[1].Descriptor()
	// website.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	website.DefaultUpdatedAt = websiteDescUpdatedAt.Default.(func() time.Time)
	// website.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	website.UpdateDefaultUpdatedAt = websiteDescUpdatedAt.UpdateDefault.(func() time.Time)
	// websiteDescMaxPage is the schema descriptor for max_page field.
	websiteDescMaxPage := websiteFields[4].Descriptor()
	// website.DefaultMaxPage holds the default value on creation for the max_page field.
	website.DefaultMaxPage = websiteDescMaxPage.Default.(int)
}
