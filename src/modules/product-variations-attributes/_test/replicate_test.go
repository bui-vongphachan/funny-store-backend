package product_variations_attributes_test

import (
	"reflect"
	"testing"

	product_variations_attributes "github.com/vongphachan/funny-store-backend/src/modules/product-variations-attributes"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.

func TestReplicate(t *testing.T) {
	type args struct {
		props *product_variations_attributes.Props_Replicate
	}
	tests := []struct {
		name    string
		args    args
		want    *[]product_variations_attributes.ProductVariationAttribute
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Replicate(tt.args.props)
			if (err != nil) != tt.wantErr {
				t.Errorf("Replicate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Replicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
