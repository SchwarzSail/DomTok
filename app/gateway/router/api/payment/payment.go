/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by hertz generator. DO NOT EDIT.

package payment

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	payment "github.com/west2-online/DomTok/app/gateway/handler/api/payment"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_payment := _api.Group("/payment", _paymentMw()...)
			_payment.POST("/process", append(_processpaymentMw(), payment.ProcessPayment)...)
			_payment.POST("/refund", append(_processrefundMw(), payment.ProcessRefund)...)
			_payment.GET("/refund-token", append(_requestrefundtokenMw(), payment.RequestRefundToken)...)
			_payment.GET("/token", append(_requestpaymenttokenMw(), payment.RequestPaymentToken)...)
		}
	}
}
