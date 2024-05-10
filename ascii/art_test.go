package ascii

import
(
	"testing"
)


func TestReadAscii(t *testing.T){
	sample:=[]struct{
		input string
		want  string
	}{
		{, },
		{, },
	}
	for _, data:= range sample{
		result := ReadAscii(data.input)
		if result!= data.eant{
			t.Errorf("double(%v)=%v, want %v", data.input, result, data.want)
		}
	}
}