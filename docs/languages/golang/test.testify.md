---
layout: default
title: Test - Testify
nav_exclude: true
parent: Golang
grand_parent: Programming Languages
permalink: /docs/languages/golang/test-testify
---

__[back]({% link docs/languages/golang/test.md %})__
<br/>

<br/>
<details markdown="block">
  <summary> Testify Sample </summary>

<details markdown="block">
  <summary> Testify Sample - Calculator </summary>

`calculator.go`
```golang
package calculator

func Add(a, b int) int {
  return a + b
}
```

`calculator_test.go`
```golang
package calculator

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
  // t.Fatal("not implemented")
}

func TestAdd(t *testing.T) {
  result := Add(2, 3)
  expected := 5
  assert.Equal(t, expected, result, "Addition result")
}
```

-----
<!-- Testify Sample - Calculator -->
</details>

<details markdown="block">
  <summary> Testify Sample - Ecommerce </summary>

`interfaces.go`
```golang
package ecommerce

type Inventory interface {
  IsAvailable(productID string, quantity int) bool
  DecrementStock(productID string, quantity int)
}

type PaymentGateway interface {
  ProcessPayment(amount float64) bool
}
```

`order.go`
```golang
package ecommerce

type OrderService struct {
  Inventory      Inventory
  PaymentGateway PaymentGateway
}

func (os *OrderService) PlaceOrder(productID string, quantity int, amount float64) bool {
  if os.Inventory.IsAvailable(productID, quantity) {
    os.Inventory.DecrementStock(productID, quantity)
    return os.PaymentGateway.ProcessPayment(amount)
  }
  return false
}
```

`order_test.go`
```golang
package ecommerce

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

type MockInventory struct{}

func (mi *MockInventory) IsAvailable(productID string, quantity int) bool {
  return true
}

func (mi *MockInventory) DecrementStock(productID string, quantity int) {
}

type MockPaymentGateway struct{}

func (mp *MockPaymentGateway) ProcessPayment(amount float64) bool {
  return true
}

func TestPlaceOrder(t *testing.T) {
  orderService := OrderService{
    Inventory:      &MockInventory{},
    PaymentGateway: &MockPaymentGateway{},
  }

  result := orderService.PlaceOrder("123", 2, 50.0)
  assert.True(t, result, "Order should be successful")
}
```

-----
<!-- Testify Sample - Ecommerce -->
</details>

-----
<!-- Testify Sample -->
</details>


<details markdown="block">
  <summary> Testify Mock </summary>

`main.go`
```golang
package main

import (
  "fmt"

  "github.com/stretchr/testify/mock"
)

type mockCalculateArea struct {
  mock.Mock
}

func (m *mockCalculateArea) calculateArea(width int, height int) int {
  args := m.Called(width, height)
  return args.Int(0)
}

func main() {
  fmt.Println("vim-go")
}
```

`main_test.go`
```golang
package main

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestCalculateArea(t *testing.T) {
  mockObj := new(mockCalculateArea)
  mockObj.On("calculateArea", 5, 10).Return(50)
  actualArea := mockObj.calculateArea(5, 10)
  assert.Equal(t, 50, actualArea, "The calculated area is incorrect")
  mockObj.AssertExpectations(t)
}
```

-----
<!-- Testify Mock -->
</details>
