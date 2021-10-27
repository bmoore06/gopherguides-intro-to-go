package main

import (
	"fmt"
)

var _ Entertainer = &Band{}
var _ Setuper = &Band{}
var _Teardowner = &Band{}

type Band struct {
	IsSetup     bool
	IsTorndown  bool
	MinAudience int
	PlayedFor   int
}

func (a Band) Name() string {
	return "The Thunders of Rock"
}

func (a Band) Validate(a Venue) error {
	if b.Audience < a.MinAudience {
		return fmt.Errorf("we don't play small bands")
	}
	return nil
}

func (a *Band) Perform(b Venue) error {
	if err := a.Validate(b); err != nil {
		return err
	}
	a.PlayedFor = b.Audience

	return nil
}

func (a *Band) Setup(b Venue) error {
	if a.IsSetup {
		return fmt.Errorf("we already set up your material")

	}
	if err := a.Validate(b); err != nil {
		return err

	}
	a.IsSetup = true
	return nil

}

func (a *Band) Teardown(a Venue) error {
	if a.IsTorndown {
		return fmt.Errorf("we already tore down our material")

	}
	a.IsTorndown = true
	return nil
}

var _Entertainer = Poet{}

type Poet struct{}

func (c Poet) Name() string {
	return "Maybelle Marie"

}

func (c Poet) Perform(a Venue) error {
	if b.Audience == 1 {
		return fmt.Errorf("i'm not playing for just the employees")

	}
	return nil
}
