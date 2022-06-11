package generic

import "testing"

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return first
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Rio",
		Second: "Raihan",
	}
	assert.Equal(t, "Reno", data.ChangeFirst("Reno"))
	assert.Equal(t, "Hello Rio", data.SayHello("Rio"))
}
