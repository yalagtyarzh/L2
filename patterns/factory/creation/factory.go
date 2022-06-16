package creation

import (
	"errors"
	"fmt"
)

// PaymentMethod - интерфейс метода оплаты
type PaymentMethod interface {
	Pay(amount float64) string
}

// Числовые константы, которые используются для проверки метода оплаты в функции GetPaymentMethod
const (
	Cash = iota + 1
	DebitCard
	CreditCard
)

// GetPaymentMethod возвращает конкретный метод оплаты в зависимости от переданного параметра m
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return &CashPM{}, nil
	case DebitCard:
		return &DebitCardPM{}, nil
	case CreditCard:
		return &CreditCardPM{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("invalid payment method: %d\n", m))
	}
}

// CashPM представляет собой объект, отвечающий за оплата наличкой
type CashPM struct {
}

// Pay представляет собой логику оплаты для налички
func (c *CashPM) Pay(amount float64) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

type DebitCardPM struct {
}

func (d *DebitCardPM) Pay(amount float64) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}

type CreditCardPM struct {
}

func (c *CreditCardPM) Pay(amount float64) string {
	return fmt.Sprintf("%0.2f paid using new credit card implementation\n", amount)
}
