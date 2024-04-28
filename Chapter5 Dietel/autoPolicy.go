package main

import "fmt"

type AutoPolicy struct {
    policyNumber int
    makeAndModel string
    state string
}

func NewAutoPolicy(policyNumber int, makeAndModel string, state string) *AutoPolicy {
    return &AutoPolicy{
        policyNumber: policyNumber,
        makeAndModel: makeAndModel,
        state: state,
    }
}

func (ap *AutoPolicy) SetState(state string) {
    if state == "CT" || state == "MA" || state == "ME" ||
        state == "NH" || state == "NJ" || state == "NY" ||
        state == "PA" || state == "VT" {
        ap.state = state
    } else {
        fmt.Println("Invalid state code for northeast states.")
    }
}

func (ap *AutoPolicy) GetState() string {
    return ap.state
}

func main() {
    ap := NewAutoPolicy(12345, "Toyota Camry", "CT")
    fmt.Println(ap.GetState())

    ap.SetState("FL")
    fmt.Println(ap.GetState())
}
