/* 
 * Betfair: Exchange Streaming API
 *
 * API to receive streamed updates. This is an ssl socket connection of CRLF delimited json messages (see RequestMessage & ResponseMessage)
 *
 * OpenAPI spec version: 1.0.1423
 * Contact: bdp@betfair.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

type Order struct {

	// Side - the side of the order. For Line markets a 'B' bet refers to a SELL line and an 'L' bet refers to a BUY line.
	Side string `json:"side,omitempty"`

	// Size Voided - the amount of the order that has been voided
	Sv float64 `json:"sv,omitempty"`

	// Persistence Type - whether the order will persist at in play or not (L = LAPSE, P = PERSIST, MOC = Market On Close)
	Pt string `json:"pt,omitempty"`

	// Order Type - the type of the order (L = LIMIT, MOC = MARKET_ON_CLOSE, LOC = LIMIT_ON_CLOSE)
	Ot string `json:"ot,omitempty"`

	// Lapse Status Reason Code - the reason that some or all of this order has been lapsed (null if no portion of the order is lapsed
	Lsrc string `json:"lsrc,omitempty"`

	// Price - the original placed price of the order. Line markets operate at even-money odds of 2.0. However, price for these markets refers to the line positions available as defined by the markets min-max range and interval steps
	P float64 `json:"p,omitempty"`

	// Size Cancelled - the amount of the order that has been cancelled
	Sc float64 `json:"sc,omitempty"`

	// Regulator Code - the regulator of the order
	Rc string `json:"rc,omitempty"`

	// Size - the original placed size of the order
	S float64 `json:"s,omitempty"`

	// Placed Date - the date the order was placed
	Pd int64 `json:"pd,omitempty"`

	// Regulator Auth Code - the auth code returned by the regulator
	Rac string `json:"rac,omitempty"`

	// Matched Date - the date the order was matched (null if the order is not matched)
	Md int64 `json:"md,omitempty"`

	// Cancelled Date - the date the order was cancelled (null if the order is not cancelled)
	Cd int64 `json:"cd,omitempty"`

	// Lapsed Date - the date the order was lapsed (null if the order is not lapsed)
	Ld int64 `json:"ld,omitempty"`

	// Size Lapsed - the amount of the order that has been lapsed
	Sl float64 `json:"sl,omitempty"`

	// Average Price Matched - the average price the order was matched at (null if the order is not matched). This value is not meaningful for activity on Line markets and is not guaranteed to be returned or maintained for these markets.
	Avp float64 `json:"avp,omitempty"`

	// Size Matched - the amount of the order that has been matched
	Sm float64 `json:"sm,omitempty"`

	// Order Reference - the customer's order reference for this order (empty string if one was not set)
	Rfo string `json:"rfo,omitempty"`

	// Bet Id - the id of the order
	Id string `json:"id,omitempty"`

	// BSP Liability - the BSP liability of the order (null if the order is not a BSP order)
	Bsp float64 `json:"bsp,omitempty"`

	// Strategy Reference - the customer's strategy reference for this order (empty string if one was not set)
	Rfs string `json:"rfs,omitempty"`

	// Status - the status of the order (E = EXECUTABLE, EC = EXECUTION_COMPLETE)
	Status string `json:"status,omitempty"`

	// Size Remaining - the amount of the order that is remaining unmatched
	Sr float64 `json:"sr,omitempty"`
}
