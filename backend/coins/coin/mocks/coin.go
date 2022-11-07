// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/digitalbitbox/bitbox-wallet-app/backend/coins/coin"
	"github.com/digitalbitbox/bitbox-wallet-app/util/observable"
	"math/big"
	"sync"
)

// Ensure, that CoinMock does implement coin.Coin.
// If this is not the case, regenerate this file with moq.
var _ coin.Coin = &CoinMock{}

// CoinMock is a mock implementation of coin.Coin.
//
// 	func TestSomethingThatUsesCoin(t *testing.T) {
//
// 		// make and configure a mocked coin.Coin
// 		mockedCoin := &CoinMock{
// 			BlockExplorerTransactionURLPrefixFunc: func() string {
// 				panic("mock out the BlockExplorerTransactionURLPrefix method")
// 			},
// 			CloseFunc: func() error {
// 				panic("mock out the Close method")
// 			},
// 			CodeFunc: func() coin.Code {
// 				panic("mock out the Code method")
// 			},
// 			DecimalsFunc: func(isFee bool) uint {
// 				panic("mock out the Decimals method")
// 			},
// 			FormatAmountFunc: func(amount coin.Amount, isFee bool) string {
// 				panic("mock out the FormatAmount method")
// 			},
// 			GetFormatUnitFunc: func(isFee bool) string {
// 				panic("mock out the GetFormatUnit method")
// 			},
// 			InitializeFunc: func()  {
// 				panic("mock out the Initialize method")
// 			},
// 			NameFunc: func() string {
// 				panic("mock out the Name method")
// 			},
// 			ObserveFunc: func(fn func(observable.Event)) func() {
// 				panic("mock out the Observe method")
// 			},
// 			ParseAmountFunc: func(amount string) (coin.Amount, error) {
// 				panic("mock out the ParseAmount method")
// 			},
// 			SetAmountFunc: func(amount *big.Rat, isFee bool) coin.Amount {
// 				panic("mock out the SetAmount method")
// 			},
// 			SmallestUnitFunc: func() string {
// 				panic("mock out the SmallestUnit method")
// 			},
// 			ToUnitFunc: func(amount coin.Amount, isFee bool) float64 {
// 				panic("mock out the ToUnit method")
// 			},
// 			UnitFunc: func(isFee bool) string {
// 				panic("mock out the Unit method")
// 			},
// 		}
//
// 		// use mockedCoin in code that requires coin.Coin
// 		// and then make assertions.
//
// 	}
type CoinMock struct {
	// BlockExplorerTransactionURLPrefixFunc mocks the BlockExplorerTransactionURLPrefix method.
	BlockExplorerTransactionURLPrefixFunc func() string

	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// CodeFunc mocks the Code method.
	CodeFunc func() coin.Code

	// DecimalsFunc mocks the Decimals method.
	DecimalsFunc func(isFee bool) uint

	// FormatAmountFunc mocks the FormatAmount method.
	FormatAmountFunc func(amount coin.Amount, isFee bool) string

	// GetFormatUnitFunc mocks the GetFormatUnit method.
	GetFormatUnitFunc func(isFee bool) string

	// InitializeFunc mocks the Initialize method.
	InitializeFunc func()

	// NameFunc mocks the Name method.
	NameFunc func() string

	// ObserveFunc mocks the Observe method.
	ObserveFunc func(fn func(observable.Event)) func()

	// ParseAmountFunc mocks the ParseAmount method.
	ParseAmountFunc func(amount string) (coin.Amount, error)

	// SetAmountFunc mocks the SetAmount method.
	SetAmountFunc func(amount *big.Rat, isFee bool) coin.Amount

	// SmallestUnitFunc mocks the SmallestUnit method.
	SmallestUnitFunc func() string

	// ToUnitFunc mocks the ToUnit method.
	ToUnitFunc func(amount coin.Amount, isFee bool) float64

	// UnitFunc mocks the Unit method.
	UnitFunc func(isFee bool) string

	// calls tracks calls to the methods.
	calls struct {
		// BlockExplorerTransactionURLPrefix holds details about calls to the BlockExplorerTransactionURLPrefix method.
		BlockExplorerTransactionURLPrefix []struct {
		}
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// Code holds details about calls to the Code method.
		Code []struct {
		}
		// Decimals holds details about calls to the Decimals method.
		Decimals []struct {
			// IsFee is the isFee argument value.
			IsFee bool
		}
		// FormatAmount holds details about calls to the FormatAmount method.
		FormatAmount []struct {
			// Amount is the amount argument value.
			Amount coin.Amount
			// IsFee is the isFee argument value.
			IsFee bool
		}
		// GetFormatUnit holds details about calls to the GetFormatUnit method.
		GetFormatUnit []struct {
			// IsFee is the isFee argument value.
			IsFee bool
		}
		// Initialize holds details about calls to the Initialize method.
		Initialize []struct {
		}
		// Name holds details about calls to the Name method.
		Name []struct {
		}
		// Observe holds details about calls to the Observe method.
		Observe []struct {
			// Fn is the fn argument value.
			Fn func(observable.Event)
		}
		// ParseAmount holds details about calls to the ParseAmount method.
		ParseAmount []struct {
			// Amount is the amount argument value.
			Amount string
		}
		// SetAmount holds details about calls to the SetAmount method.
		SetAmount []struct {
			// Amount is the amount argument value.
			Amount *big.Rat
			// IsFee is the isFee argument value.
			IsFee bool
		}
		// SmallestUnit holds details about calls to the SmallestUnit method.
		SmallestUnit []struct {
		}
		// ToUnit holds details about calls to the ToUnit method.
		ToUnit []struct {
			// Amount is the amount argument value.
			Amount coin.Amount
			// IsFee is the isFee argument value.
			IsFee bool
		}
		// Unit holds details about calls to the Unit method.
		Unit []struct {
			// IsFee is the isFee argument value.
			IsFee bool
		}
	}
	lockBlockExplorerTransactionURLPrefix sync.RWMutex
	lockClose                             sync.RWMutex
	lockCode                              sync.RWMutex
	lockDecimals                          sync.RWMutex
	lockFormatAmount                      sync.RWMutex
	lockGetFormatUnit                     sync.RWMutex
	lockInitialize                        sync.RWMutex
	lockName                              sync.RWMutex
	lockObserve                           sync.RWMutex
	lockParseAmount                       sync.RWMutex
	lockSetAmount                         sync.RWMutex
	lockSmallestUnit                      sync.RWMutex
	lockToUnit                            sync.RWMutex
	lockUnit                              sync.RWMutex
}

// BlockExplorerTransactionURLPrefix calls BlockExplorerTransactionURLPrefixFunc.
func (mock *CoinMock) BlockExplorerTransactionURLPrefix() string {
	if mock.BlockExplorerTransactionURLPrefixFunc == nil {
		panic("CoinMock.BlockExplorerTransactionURLPrefixFunc: method is nil but Coin.BlockExplorerTransactionURLPrefix was just called")
	}
	callInfo := struct {
	}{}
	mock.lockBlockExplorerTransactionURLPrefix.Lock()
	mock.calls.BlockExplorerTransactionURLPrefix = append(mock.calls.BlockExplorerTransactionURLPrefix, callInfo)
	mock.lockBlockExplorerTransactionURLPrefix.Unlock()
	return mock.BlockExplorerTransactionURLPrefixFunc()
}

// BlockExplorerTransactionURLPrefixCalls gets all the calls that were made to BlockExplorerTransactionURLPrefix.
// Check the length with:
//     len(mockedCoin.BlockExplorerTransactionURLPrefixCalls())
func (mock *CoinMock) BlockExplorerTransactionURLPrefixCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockBlockExplorerTransactionURLPrefix.RLock()
	calls = mock.calls.BlockExplorerTransactionURLPrefix
	mock.lockBlockExplorerTransactionURLPrefix.RUnlock()
	return calls
}

// Close calls CloseFunc.
func (mock *CoinMock) Close() error {
	if mock.CloseFunc == nil {
		panic("CoinMock.CloseFunc: method is nil but Coin.Close was just called")
	}
	callInfo := struct {
	}{}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedCoin.CloseCalls())
func (mock *CoinMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// Code calls CodeFunc.
func (mock *CoinMock) Code() coin.Code {
	if mock.CodeFunc == nil {
		panic("CoinMock.CodeFunc: method is nil but Coin.Code was just called")
	}
	callInfo := struct {
	}{}
	mock.lockCode.Lock()
	mock.calls.Code = append(mock.calls.Code, callInfo)
	mock.lockCode.Unlock()
	return mock.CodeFunc()
}

// CodeCalls gets all the calls that were made to Code.
// Check the length with:
//     len(mockedCoin.CodeCalls())
func (mock *CoinMock) CodeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockCode.RLock()
	calls = mock.calls.Code
	mock.lockCode.RUnlock()
	return calls
}

// Decimals calls DecimalsFunc.
func (mock *CoinMock) Decimals(isFee bool) uint {
	if mock.DecimalsFunc == nil {
		panic("CoinMock.DecimalsFunc: method is nil but Coin.Decimals was just called")
	}
	callInfo := struct {
		IsFee bool
	}{
		IsFee: isFee,
	}
	mock.lockDecimals.Lock()
	mock.calls.Decimals = append(mock.calls.Decimals, callInfo)
	mock.lockDecimals.Unlock()
	return mock.DecimalsFunc(isFee)
}

// DecimalsCalls gets all the calls that were made to Decimals.
// Check the length with:
//     len(mockedCoin.DecimalsCalls())
func (mock *CoinMock) DecimalsCalls() []struct {
	IsFee bool
} {
	var calls []struct {
		IsFee bool
	}
	mock.lockDecimals.RLock()
	calls = mock.calls.Decimals
	mock.lockDecimals.RUnlock()
	return calls
}

// FormatAmount calls FormatAmountFunc.
func (mock *CoinMock) FormatAmount(amount coin.Amount, isFee bool) string {
	if mock.FormatAmountFunc == nil {
		panic("CoinMock.FormatAmountFunc: method is nil but Coin.FormatAmount was just called")
	}
	callInfo := struct {
		Amount coin.Amount
		IsFee  bool
	}{
		Amount: amount,
		IsFee:  isFee,
	}
	mock.lockFormatAmount.Lock()
	mock.calls.FormatAmount = append(mock.calls.FormatAmount, callInfo)
	mock.lockFormatAmount.Unlock()
	return mock.FormatAmountFunc(amount, isFee)
}

// FormatAmountCalls gets all the calls that were made to FormatAmount.
// Check the length with:
//     len(mockedCoin.FormatAmountCalls())
func (mock *CoinMock) FormatAmountCalls() []struct {
	Amount coin.Amount
	IsFee  bool
} {
	var calls []struct {
		Amount coin.Amount
		IsFee  bool
	}
	mock.lockFormatAmount.RLock()
	calls = mock.calls.FormatAmount
	mock.lockFormatAmount.RUnlock()
	return calls
}

// GetFormatUnit calls GetFormatUnitFunc.
func (mock *CoinMock) GetFormatUnit(isFee bool) string {
	if mock.GetFormatUnitFunc == nil {
		panic("CoinMock.GetFormatUnitFunc: method is nil but Coin.GetFormatUnit was just called")
	}
	callInfo := struct {
		IsFee bool
	}{
		IsFee: isFee,
	}
	mock.lockGetFormatUnit.Lock()
	mock.calls.GetFormatUnit = append(mock.calls.GetFormatUnit, callInfo)
	mock.lockGetFormatUnit.Unlock()
	return mock.GetFormatUnitFunc(isFee)
}

// GetFormatUnitCalls gets all the calls that were made to GetFormatUnit.
// Check the length with:
//     len(mockedCoin.GetFormatUnitCalls())
func (mock *CoinMock) GetFormatUnitCalls() []struct {
	IsFee bool
} {
	var calls []struct {
		IsFee bool
	}
	mock.lockGetFormatUnit.RLock()
	calls = mock.calls.GetFormatUnit
	mock.lockGetFormatUnit.RUnlock()
	return calls
}

// Initialize calls InitializeFunc.
func (mock *CoinMock) Initialize() {
	if mock.InitializeFunc == nil {
		panic("CoinMock.InitializeFunc: method is nil but Coin.Initialize was just called")
	}
	callInfo := struct {
	}{}
	mock.lockInitialize.Lock()
	mock.calls.Initialize = append(mock.calls.Initialize, callInfo)
	mock.lockInitialize.Unlock()
	mock.InitializeFunc()
}

// InitializeCalls gets all the calls that were made to Initialize.
// Check the length with:
//     len(mockedCoin.InitializeCalls())
func (mock *CoinMock) InitializeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockInitialize.RLock()
	calls = mock.calls.Initialize
	mock.lockInitialize.RUnlock()
	return calls
}

// Name calls NameFunc.
func (mock *CoinMock) Name() string {
	if mock.NameFunc == nil {
		panic("CoinMock.NameFunc: method is nil but Coin.Name was just called")
	}
	callInfo := struct {
	}{}
	mock.lockName.Lock()
	mock.calls.Name = append(mock.calls.Name, callInfo)
	mock.lockName.Unlock()
	return mock.NameFunc()
}

// NameCalls gets all the calls that were made to Name.
// Check the length with:
//     len(mockedCoin.NameCalls())
func (mock *CoinMock) NameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockName.RLock()
	calls = mock.calls.Name
	mock.lockName.RUnlock()
	return calls
}

// Observe calls ObserveFunc.
func (mock *CoinMock) Observe(fn func(observable.Event)) func() {
	if mock.ObserveFunc == nil {
		panic("CoinMock.ObserveFunc: method is nil but Coin.Observe was just called")
	}
	callInfo := struct {
		Fn func(observable.Event)
	}{
		Fn: fn,
	}
	mock.lockObserve.Lock()
	mock.calls.Observe = append(mock.calls.Observe, callInfo)
	mock.lockObserve.Unlock()
	return mock.ObserveFunc(fn)
}

// ObserveCalls gets all the calls that were made to Observe.
// Check the length with:
//     len(mockedCoin.ObserveCalls())
func (mock *CoinMock) ObserveCalls() []struct {
	Fn func(observable.Event)
} {
	var calls []struct {
		Fn func(observable.Event)
	}
	mock.lockObserve.RLock()
	calls = mock.calls.Observe
	mock.lockObserve.RUnlock()
	return calls
}

// ParseAmount calls ParseAmountFunc.
func (mock *CoinMock) ParseAmount(amount string) (coin.Amount, error) {
	if mock.ParseAmountFunc == nil {
		panic("CoinMock.ParseAmountFunc: method is nil but Coin.ParseAmount was just called")
	}
	callInfo := struct {
		Amount string
	}{
		Amount: amount,
	}
	mock.lockParseAmount.Lock()
	mock.calls.ParseAmount = append(mock.calls.ParseAmount, callInfo)
	mock.lockParseAmount.Unlock()
	return mock.ParseAmountFunc(amount)
}

// ParseAmountCalls gets all the calls that were made to ParseAmount.
// Check the length with:
//     len(mockedCoin.ParseAmountCalls())
func (mock *CoinMock) ParseAmountCalls() []struct {
	Amount string
} {
	var calls []struct {
		Amount string
	}
	mock.lockParseAmount.RLock()
	calls = mock.calls.ParseAmount
	mock.lockParseAmount.RUnlock()
	return calls
}

// SetAmount calls SetAmountFunc.
func (mock *CoinMock) SetAmount(amount *big.Rat, isFee bool) coin.Amount {
	if mock.SetAmountFunc == nil {
		panic("CoinMock.SetAmountFunc: method is nil but Coin.SetAmount was just called")
	}
	callInfo := struct {
		Amount *big.Rat
		IsFee  bool
	}{
		Amount: amount,
		IsFee:  isFee,
	}
	mock.lockSetAmount.Lock()
	mock.calls.SetAmount = append(mock.calls.SetAmount, callInfo)
	mock.lockSetAmount.Unlock()
	return mock.SetAmountFunc(amount, isFee)
}

// SetAmountCalls gets all the calls that were made to SetAmount.
// Check the length with:
//     len(mockedCoin.SetAmountCalls())
func (mock *CoinMock) SetAmountCalls() []struct {
	Amount *big.Rat
	IsFee  bool
} {
	var calls []struct {
		Amount *big.Rat
		IsFee  bool
	}
	mock.lockSetAmount.RLock()
	calls = mock.calls.SetAmount
	mock.lockSetAmount.RUnlock()
	return calls
}

// SmallestUnit calls SmallestUnitFunc.
func (mock *CoinMock) SmallestUnit() string {
	if mock.SmallestUnitFunc == nil {
		panic("CoinMock.SmallestUnitFunc: method is nil but Coin.SmallestUnit was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSmallestUnit.Lock()
	mock.calls.SmallestUnit = append(mock.calls.SmallestUnit, callInfo)
	mock.lockSmallestUnit.Unlock()
	return mock.SmallestUnitFunc()
}

// SmallestUnitCalls gets all the calls that were made to SmallestUnit.
// Check the length with:
//     len(mockedCoin.SmallestUnitCalls())
func (mock *CoinMock) SmallestUnitCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSmallestUnit.RLock()
	calls = mock.calls.SmallestUnit
	mock.lockSmallestUnit.RUnlock()
	return calls
}

// ToUnit calls ToUnitFunc.
func (mock *CoinMock) ToUnit(amount coin.Amount, isFee bool) float64 {
	if mock.ToUnitFunc == nil {
		panic("CoinMock.ToUnitFunc: method is nil but Coin.ToUnit was just called")
	}
	callInfo := struct {
		Amount coin.Amount
		IsFee  bool
	}{
		Amount: amount,
		IsFee:  isFee,
	}
	mock.lockToUnit.Lock()
	mock.calls.ToUnit = append(mock.calls.ToUnit, callInfo)
	mock.lockToUnit.Unlock()
	return mock.ToUnitFunc(amount, isFee)
}

// ToUnitCalls gets all the calls that were made to ToUnit.
// Check the length with:
//     len(mockedCoin.ToUnitCalls())
func (mock *CoinMock) ToUnitCalls() []struct {
	Amount coin.Amount
	IsFee  bool
} {
	var calls []struct {
		Amount coin.Amount
		IsFee  bool
	}
	mock.lockToUnit.RLock()
	calls = mock.calls.ToUnit
	mock.lockToUnit.RUnlock()
	return calls
}

// Unit calls UnitFunc.
func (mock *CoinMock) Unit(isFee bool) string {
	if mock.UnitFunc == nil {
		panic("CoinMock.UnitFunc: method is nil but Coin.Unit was just called")
	}
	callInfo := struct {
		IsFee bool
	}{
		IsFee: isFee,
	}
	mock.lockUnit.Lock()
	mock.calls.Unit = append(mock.calls.Unit, callInfo)
	mock.lockUnit.Unlock()
	return mock.UnitFunc(isFee)
}

// UnitCalls gets all the calls that were made to Unit.
// Check the length with:
//     len(mockedCoin.UnitCalls())
func (mock *CoinMock) UnitCalls() []struct {
	IsFee bool
} {
	var calls []struct {
		IsFee bool
	}
	mock.lockUnit.RLock()
	calls = mock.calls.Unit
	mock.lockUnit.RUnlock()
	return calls
}
