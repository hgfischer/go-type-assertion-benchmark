package bench

import "testing"

func BenchmarkAppendNativeTypeAssertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		col := &FloatCol{}
		col.AppendNativeTypeAssertion(1.2)
		col.AppendNativeTypeAssertion("4.3")
		col.AppendNativeTypeAssertion(float32(123))
	}
}

func BenchmarkAppendReflectionTypeAssertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		col := &FloatCol{}
		col.AppendReflectionTypeAssertion(1.2)
		col.AppendReflectionTypeAssertion("4.3")
		col.AppendReflectionTypeAssertion(float32(123))
	}
}