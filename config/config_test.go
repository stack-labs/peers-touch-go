package config

import (
	js "encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testOptionsA struct {
	AbstractOptions
	AV1 string
	AV2 int
}

func (t *testOptionsA) Path() string {
	return "a"
}

type testOptionsB struct {
	AbstractOptions
}

func (t *testOptionsB) Path() string {
	return "a.b"
}

type testOptionsC struct {
	AbstractOptions
	TestC0 string
	TestC1 string
}

func (t *testOptionsC) Path() string {
	return "a.c"
}

type testOptionsD struct {
	AbstractOptions
	TestD0 string
	TestD1 string
}

func (t *testOptionsD) Path() string {
	return "a.b.d"
}

func TestConfig(t *testing.T) {
	a, b, c, d :=
		&testOptionsA{
			AV1: "AV1 Hello",
			AV2: 2,
		},
		&testOptionsB{},
		&testOptionsC{
			TestC0: "testC0V",
			TestC1: "testC1V",
		},
		&testOptionsD{
			TestD0: "testD0V",
			TestD1: "testD1V",
		}

	Convey("Config ", t, FailureHalts, func(cc C) {
		Convey("Register", FailureHalts, func(cc C) {
			Register(d, a, c, b)
			cc.So(configs.valuesMap, ShouldNotBeNil)

			Convey("Init", FailureHalts, func(cc C) {
				configs.init()
				j, _ := js.Marshal(configs.options)
				t.Logf("values: %s", string(j))
			})
		})
	})
}
