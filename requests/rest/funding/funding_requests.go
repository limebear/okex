package funding

import (
	"github.com/limebear/okex"
	"github.com/shopspring/decimal"
)

type (
	GetBalance struct {
		Ccy []string `json:"ccy,omitempty"`
	}
	FundsTransfer struct {
		Ccy      string            `json:"ccy"`
		Amt      decimal.Decimal   `json:"amt"`
		SubAcct  string            `json:"subAcct,omitempty"`
		InstID   string            `json:"instID,omitempty"`
		ToInstID string            `json:"instId,omitempty"`
		Type     okex.TransferType `json:"type,omitempty,string"`
		From     okex.AccountType  `json:"from,string"`
		To       okex.AccountType  `json:"to,string"`
	}
	AssetBillsDetails struct {
		Type   okex.BillType `json:"type,string,omitempty"`
		After  int64         `json:"after,string,omitempty"`
		Before int64         `json:"before,string,omitempty"`
		Limit  int64         `json:"limit,string,omitempty"`
	}
	GetDepositAddress struct {
		Ccy string `json:"ccy"`
	}
	GetDepositHistory struct {
		Ccy    string            `json:"ccy,omitempty"`
		TxID   string            `json:"txId,omitempty"`
		After  int64             `json:"after,omitempty,string"`
		Before int64             `json:"before,omitempty,string"`
		Limit  int64             `json:"limit,omitempty,string"`
		State  okex.DepositState `json:"state,omitempty,string"`
	}

	ReceiverInfo struct {
		WalletType    string `json:"walletType"`
		ExchId        string `json:"exchId"`
		RcvrFirstName string `json:"rcvrFirstName"`
		RcvrLastName  string `json:"rcvrLastName"`
	}

	Withdrawal struct {
		Ccy      string                     `json:"ccy"`
		Chain    string                     `json:"chain,omitempty"`
		ToAddr   string                     `json:"toAddr"`
		Pwd      string                     `json:"pwd"`
		Amt      decimal.Decimal            `json:"amt"`
		Fee      decimal.Decimal            `json:"fee"`
		Dest     okex.WithdrawalDestination `json:"dest,string"`
		RcvrInfo ReceiverInfo               `json:"rcvrInfo"`
	}
	GetWithdrawalHistory struct {
		Ccy    string               `json:"ccy,omitempty"`
		TxID   string               `json:"txId,omitempty"`
		After  int64                `json:"after,omitempty,string"`
		Before int64                `json:"before,omitempty,string"`
		Limit  int64                `json:"limit,omitempty,string"`
		State  okex.WithdrawalState `json:"state,omitempty,string"`
	}
	PiggyBankPurchaseRedemption struct {
		Ccy    string               `json:"ccy,omitempty"`
		TxID   string               `json:"txId,omitempty"`
		After  int64                `json:"after,omitempty,string"`
		Before int64                `json:"before,omitempty,string"`
		Limit  int64                `json:"limit,omitempty,string"`
		State  okex.WithdrawalState `json:"state,omitempty,string"`
	}
	GetPiggyBankBalance struct {
		Ccy string `json:"ccy,omitempty"`
	}
)
