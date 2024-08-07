// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

// EncodeURI encodes QuerySelectParams as URI form.
func (s *QuerySelectParams) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("order", func(e uri.Encoder) error {
		if val, ok := s.Order.Get(); ok {
			return e.EncodeValue(conv.StringToString(string(val)))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"order\"")
	}
	if err := e.EncodeField("limit", func(e uri.Encoder) error {
		if val, ok := s.Limit.Get(); ok {
			return e.EncodeValue(conv.IntToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"limit\"")
	}
	if err := e.EncodeField("offset", func(e uri.Encoder) error {
		if val, ok := s.Offset.Get(); ok {
			return e.EncodeValue(conv.IntToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"offset\"")
	}
	return nil
}

var uriFieldsNameOfQuerySelectParams = [3]string{
	0: "order",
	1: "limit",
	2: "offset",
}

// DecodeURI decodes QuerySelectParams from URI form.
func (s *QuerySelectParams) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode QuerySelectParams to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "order":
			if err := func() error {
				var sDotOrderVal Order
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotOrderVal = Order(c)
					return nil
				}(); err != nil {
					return err
				}
				s.Order.SetTo(sDotOrderVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"order\"")
			}
		case "limit":
			if err := func() error {
				var sDotLimitVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					sDotLimitVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.Limit.SetTo(sDotLimitVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"limit\"")
			}
		case "offset":
			if err := func() error {
				var sDotOffsetVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					sDotOffsetVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.Offset.SetTo(sDotOffsetVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"offset\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode QuerySelectParams")
	}

	return nil
}
