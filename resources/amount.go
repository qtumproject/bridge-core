package resources

import (
	"encoding/json"
	"fmt"
	"math/big"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

var One = big.NewInt(1000000000000000000)

type Amount big.Int

func (a Amount) MarshalJSON() ([]byte, error) {
	i := big.Int(a)
	return json.Marshal(stringU(&i))
}

func (a *Amount) UnmarshalJSON(data []byte) error {
	var rawAmount string
	err := json.Unmarshal(data, &rawAmount)
	if err != nil {
		return errors.Wrap(err, "can't unmarshal amount")
	}

	rawA, err := parseU(rawAmount)
	*a = Amount(*rawA)

	return errors.Wrap(err, "can't parse amount")
}

func (a Amount) String() string {
	i := big.Int(a)
	return stringU(&i)
}

func parseU(v string) (*big.Int, error) {
	var f, o, r big.Rat

	_, ok := f.SetString(v)
	if !ok {
		return nil, fmt.Errorf("cannot parse amount: %s", v)
	}

	o.SetInt(One)
	r.Mul(&f, &o)

	is := r.FloatString(0)
	amount := big.NewInt(0)
	amount.SetString(is, 10)
	return amount, nil
}

func stringU(v *big.Int) string {
	var f, o, r big.Rat

	f.SetInt(v)
	o.SetInt(One)
	r.Quo(&f, &o)

	return r.FloatString(18)
}
