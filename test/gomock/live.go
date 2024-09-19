package gomock

import (
	"errors"
	"math/rand"
)

//go:generate mockgen -package main -source live.go -destination=live_mock.go
//go:generate mockery -name=Life
type Life interface {
	GoodGoodStudy(money int64) error
	BuyHouse(money int64) error
	Marry(money int64) error
}

type Person struct {
	Life Life
}

// Live 活着
func (p *Person) Live(money1, money2, money3 int64) error {
	if err := p.Life.GoodGoodStudy(money1); err != nil {
		return err
	}
	if err := p.Life.BuyHouse(money2); err != nil {
		return err
	}
	if err := p.Life.Marry(money3); err != nil {
		return err
	}
	return nil
}

// GoodGoodStudy 好好学习
func (p *Person) GoodGoodStudy(money int64) error {
	if rand.Intn(100) > 0 {
		return errors.New("error")
	}
	_ = money
	return nil
}

// BuyHouse 买房
func (p *Person) BuyHouse(money int64) error {
	if rand.Intn(100) > 0 {
		return errors.New("error")
	}
	_ = money
	return nil
}

// Marry 结婚
func (p *Person) Marry(money int64) error {
	if rand.Intn(100) > 0 {
		return errors.New("error")
	}
	_ = money
	return nil
}
