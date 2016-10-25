package test

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import "github.com/tinylib/msgp/msgp"

// DecodeMsg implements msgp.Decodable
func (z *Something) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxvk uint32
	zxvk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "A":
			z.A, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "B":
			z.B, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "C":
			z.C, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Something) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "A"
	err = en.Append(0x83, 0xa1, 0x41)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.A)
	if err != nil {
		return
	}
	// write "B"
	err = en.Append(0xa1, 0x42)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.B)
	if err != nil {
		return
	}
	// write "C"
	err = en.Append(0xa1, 0x43)
	if err != nil {
		return err
	}
	err = en.WriteString(z.C)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Something) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "A"
	o = append(o, 0x83, 0xa1, 0x41)
	o = msgp.AppendFloat64(o, z.A)
	// string "B"
	o = append(o, 0xa1, 0x42)
	o = msgp.AppendInt64(o, z.B)
	// string "C"
	o = append(o, 0xa1, 0x43)
	o = msgp.AppendString(o, z.C)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Something) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "A":
			z.A, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "B":
			z.B, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "C":
			z.C, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z Something) Msgsize() (s int) {
	s = 1 + 2 + msgp.Float64Size + 2 + msgp.Int64Size + 2 + msgp.StringPrefixSize + len(z.C)
	return
}
