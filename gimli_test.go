package gimli

import "testing"

func TestGimli(t *testing.T) {
	var got Gimli
	want := Gimli{
		0xba11c85a, 0x91bad119, 0x380ce880, 0xd24c2c68,
		0x3eceffea, 0x277a921c, 0x4f73a0bd, 0xda5a9cd8,
		0x84b673f0, 0x34e52ff7, 0x9e2bef49, 0xf41bb8d6}

	for i := uint32(0); i < 12; i++ {
		got[i] = i*i*i + i*0x9e3779b9
	}

	got.Update()

	for i := uint32(0); i < 12; i++ {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func BenchmarkGimli(b *testing.B) {
	var gimli Gimli
	for i := 0; i < b.N; i++ {
		gimli.Update()
	}
}
