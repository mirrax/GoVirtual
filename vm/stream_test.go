//	TODO:	Improve tests for Stream-level operations
//	TODO: 	Add tests for GetBuffer, PutBuffer and Clear

package vm
import "testing"
import "os"

func sixIntegerStream() *Stream {
	s := new(Stream)
	s.Buffer = *sixIntegerBuffer()
	return s
}

func sixFloatStream() *Stream {
	s := new(Stream)
	s.Buffer = *sixFloatBuffer()
	return s
}

func checkStream(s, o *Stream, t *testing.T, value bool) {
	compareValues(s, t, s.Identical(o), value)
}

func TestCreateStream(t *testing.T) {
	os.Stdout.WriteString("Stream Creation\n")
	checkStream(sixIntegerStream(), sixIntegerStream(), t, true)
}

func TestCloneStream(t *testing.T) {
	os.Stdout.WriteString("Stream Cloning\n")
	checkStream(sixIntegerStream().Clone(), sixIntegerStream(), t, true)
}

func TestSliceStream(t *testing.T) {
	os.Stdout.WriteString("Stream Slicing\n")
	s := sixIntegerStream().Slice(1, 3)
	compareValues(s, t, s.Len(), 2)
	compareValues(s, t, s.Cap(), 2)
	compareValues(s, t, s.At(0), int(byte("e"[0])))
	compareValues(s, t, s.At(1), 3)
}

func TestStreamMaths(t *testing.T) {
	os.Stdout.WriteString("Stream Maths\n")
	s1 := sixIntegerStream()
	s1.Buffer.Increment(0)											//	b[0] == 38
	compareValues(s1, t, s1.At(0), 38)
	s1.Buffer.Decrement(0)											//	b[0] == 37
	compareValues(s1, t, s1.At(0), 37)
	s1.Buffer.Add(1, 3)												//	b[1] == 'j'
	compareValues(s1, t, s1.At(1), int(byte("j"[0])))
	s1.Buffer.Subtract(2, 3)										//	b[2] == -2
	compareValues(s1, t, s1.At(2), -2)
	s1.Buffer.Negate(4)												//	b[4] == -2
	compareValues(s1, t, s1.At(4), -2)
	s1.Buffer.Multiply(2, 4)										//	b[2] == 4
	compareValues(s1, t, s1.At(2), 4)
	s1.Buffer.Divide(2, 5)											//	b[2] == 2
	compareValues(s1, t, s1.At(2), 2)
	s1.Buffer.Multiply(5, 3)										//	b[5] == 10
	s1.Buffer.And(2, 5)												//	b[2] == 2
	compareValues(s1, t, s1.At(2), 2)
	s1.Buffer.Or(2, 5)												//	b[2] == 10
	compareValues(s1, t, s1.At(2), 10)
	s1.Buffer.Negate(4)												//	b[4] == 2
	compareValues(s1, t, s1.At(4), 2)
	s1.Buffer.Xor(2, 4)												//	b[2] == 8
	compareValues(s1, t, s1.At(2), 8)
	s1 = sixIntegerStream()
	s2 := sixIntegerStream()
	s1.Add(0, s2)
	compareValues(s1, t, s1.At(0), 74)
	compareValues(s1, t, s1.At(1), int(byte("e"[0])) * 2)
	compareValues(s1, t, s1.At(2), 6)
	compareValues(s1, t, s1.At(3), 10)
	compareValues(s1, t, s1.At(4), 4)
	compareValues(s1, t, s1.At(5), 4)
	s1 = sixIntegerStream()
	s1.Subtract(0, s2)
	compareValues(s1, t, s1.At(0), 0)
	compareValues(s1, t, s1.At(1), 0)
	compareValues(s1, t, s1.At(2), 0)
	compareValues(s1, t, s1.At(3), 0)
	compareValues(s1, t, s1.At(4), 0)
	compareValues(s1, t, s1.At(5), 0)
	s1 = sixIntegerStream()
	s1.Multiply(0, s2)
	compareValues(s1, t, s1.At(0), 37 * 37)
	compareValues(s1, t, s1.At(1), int(byte("e"[0])) * int(byte("e"[0])))
	compareValues(s1, t, s1.At(2), 9)
	compareValues(s1, t, s1.At(3), 25)
	compareValues(s1, t, s1.At(4), 4)
	compareValues(s1, t, s1.At(5), 4)
	s1 = sixIntegerStream()
	s1.Divide(0, s2)
	compareValues(s1, t, s1.At(0), 1)
	compareValues(s1, t, s1.At(1), 1)
	compareValues(s1, t, s1.At(2), 1)
	compareValues(s1, t, s1.At(3), 1)
	compareValues(s1, t, s1.At(4), 1)
	compareValues(s1, t, s1.At(5), 1)
	s1 = sixIntegerStream()
	s1.Negate(0, s1.Len())
	compareValues(s1, t, s1.At(0), -37)
	compareValues(s1, t, s1.At(1), -int(byte("e"[0])))
	compareValues(s1, t, s1.At(2), -3)
	compareValues(s1, t, s1.At(3), -5)
	compareValues(s1, t, s1.At(4), -2)
	compareValues(s1, t, s1.At(5), -2)
	s1 = sixIntegerStream()
	s1.Increment(0, s1.Len())
	compareValues(s1, t, s1.At(0), 38)
	compareValues(s1, t, s1.At(1), int(byte("e"[0])) + 1)
	compareValues(s1, t, s1.At(2), 4)
	compareValues(s1, t, s1.At(3), 6)
	compareValues(s1, t, s1.At(4), 3)
	compareValues(s1, t, s1.At(5), 3)
	s1 = sixIntegerStream()
	s1.Decrement(0, s1.Len())
	compareValues(s1, t, s1.At(0), 36)
	compareValues(s1, t, s1.At(1), int(byte("e"[0])) - 1)
	compareValues(s1, t, s1.At(2), 2)
	compareValues(s1, t, s1.At(3), 4)
	compareValues(s1, t, s1.At(4), 1)
	compareValues(s1, t, s1.At(5), 1)
}

func TestFloatingPointMaths(t *testing.T) {
	os.Stdout.WriteString("Stream Floating Point Maths\n")
	s1 := sixFloatStream()
	s2 := sixFloatStream()
	s1.FAdd(0, s2)
	compareValues(s1, t, s1.FAt(0), 74.0)
	compareValues(s1, t, s1.FAt(1), float(byte("e"[0]) * 2))
	compareValues(s1, t, s1.FAt(2), 7.4)
	compareValues(s1, t, s1.FAt(3), 10.0)
	compareValues(s1, t, s1.FAt(4), 4.0)
	compareValues(s1, t, s1.FAt(5), 4.0)
}

func TestStreamBitOperators(t *testing.T) {
	os.Stdout.WriteString("Stream Bit Manipulation\n")
	s := sixIntegerStream()									//	b[0] == 37, b[5] == 2
	s.Buffer.ShiftRight(0, 5)
	compareValues(s, t, s.At(0), 148)
	s.Buffer.ShiftLeft(0, 5)
	compareValues(s, t, s.At(0), 37)
	s.Buffer.Invert(0)
	compareValues(s, t, s.At(0), ^37)
}

//func TestLogic(t *testing.T) {
//	os.Stdout.WriteString("Buffer Logic\n")
//	b := defaultBuffer()
//	checkDefaultBuffer(b, t, true)
//	compareValues(b, t, b.LessThan(2, 3), true)				//	b[2] == 3, b[3] == 5
//	compareValues(b, t, b.Equals(2, 3), false)
//	compareValues(b, t, b.GreaterThan(2, 3), false)
//	compareValues(b, t, b.LessThanZero(2), false)
//	compareValues(b, t, b.EqualsZero(2), false)
//	compareValues(b, t, b.GreaterThanZero(2), true)
//	b.Copy(1, 2)											//	b[1] == 3
//	checkDefaultBuffer(b, t, false)
//	compareValues(b, t, b.At(1), 3)
//	compareValues(b, t, b.LessThan(1, 3), true)				//	b[1] == 3, b[3] == 5
//	compareValues(b, t, b.Equals(1, 2), true)				//	b[1] == 3, b[2] == 3
//	compareValues(b, t, b.GreaterThan(1, 3), false)
//	compareValues(b, t, b.LessThanZero(1), false)
//	compareValues(b, t, b.EqualsZero(1), false)
//	compareValues(b, t, b.GreaterThanZero(1), true)
//	b.Set(1, 0)												//	b[1] == 0, b[3] == 5
//	checkDefaultBuffer(b, t, false)
//	compareValues(b, t, b.LessThan(1, 3), true)
//	compareValues(b, t, b.Equals(1, 3), false)
//	compareValues(b, t, b.GreaterThan(1, 3), false)
//	compareValues(b, t, b.LessThanZero(1), false)
//	compareValues(b, t, b.EqualsZero(1), true)
//	compareValues(b, t, b.GreaterThanZero(1), false)
//}
