// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: test.go
// DO NOT EDIT!

package test

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Something) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Something) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"A":`)
	fflib.AppendFloat(buf, float64(mj.A), 'g', -1, 64)
	buf.WriteString(`,"B":`)
	fflib.FormatBits2(buf, uint64(mj.B), 10, mj.B < 0)
	buf.WriteString(`,"C":`)
	fflib.WriteJsonString(buf, string(mj.C))
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Somethingbase = iota
	ffj_t_Somethingno_such_key

	ffj_t_Something_A

	ffj_t_Something_B

	ffj_t_Something_C
)

var ffj_key_Something_A = []byte("A")

var ffj_key_Something_B = []byte("B")

var ffj_key_Something_C = []byte("C")

func (uj *Something) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Something) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Somethingbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_Somethingno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'A':

					if bytes.Equal(ffj_key_Something_A, kn) {
						currentKey = ffj_t_Something_A
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'B':

					if bytes.Equal(ffj_key_Something_B, kn) {
						currentKey = ffj_t_Something_B
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'C':

					if bytes.Equal(ffj_key_Something_C, kn) {
						currentKey = ffj_t_Something_C
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Something_C, kn) {
					currentKey = ffj_t_Something_C
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Something_B, kn) {
					currentKey = ffj_t_Something_B
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Something_A, kn) {
					currentKey = ffj_t_Something_A
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Somethingno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_Something_A:
					goto handle_A

				case ffj_t_Something_B:
					goto handle_B

				case ffj_t_Something_C:
					goto handle_C

				case ffj_t_Somethingno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_A:

	/* handler: uj.A type=float64 kind=float64 quoted=false*/

	{
		if tok != fflib.FFTok_double && tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for float64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseFloat(fs.Output.Bytes(), 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.A = float64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_B:

	/* handler: uj.B type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.B = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_C:

	/* handler: uj.C type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.C = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}
