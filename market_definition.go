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

import (
	"time"
)

type MarketDefinition struct {

	Venue string `json:"venue,omitempty"`

	RaceType string `json:"raceType,omitempty"`

	SettledTime time.Time `json:"settledTime,omitempty"`

	Timezone string `json:"timezone,omitempty"`

	EachWayDivisor float64 `json:"eachWayDivisor,omitempty"`

	// The market regulators.
	Regulators []string `json:"regulators,omitempty"`

	MarketType string `json:"marketType,omitempty"`

	MarketBaseRate float64 `json:"marketBaseRate,omitempty"`

	NumberOfWinners int32 `json:"numberOfWinners,omitempty"`

	CountryCode string `json:"countryCode,omitempty"`

	// For Handicap and Line markets, the maximum value for the outcome, in market units for this market (eg 100 runs).
	LineMaxUnit float64 `json:"lineMaxUnit,omitempty"`

	InPlay bool `json:"inPlay,omitempty"`

	BetDelay int32 `json:"betDelay,omitempty"`

	BspMarket bool `json:"bspMarket,omitempty"`

	BettingType string `json:"bettingType,omitempty"`

	NumberOfActiveRunners int32 `json:"numberOfActiveRunners,omitempty"`

	// For Handicap and Line markets, the minimum value for the outcome, in market units for this market (eg 0 runs).
	LineMinUnit float64 `json:"lineMinUnit,omitempty"`

	EventId string `json:"eventId,omitempty"`

	CrossMatching bool `json:"crossMatching,omitempty"`

	RunnersVoidable bool `json:"runnersVoidable,omitempty"`

	TurnInPlayEnabled bool `json:"turnInPlayEnabled,omitempty"`

	// Definition of the price ladder type and any related data.
	PriceLadderDefinition PriceLadderDefinition `json:"priceLadderDefinition,omitempty"`

	// Definition of a markets key line selection (for valid markets), comprising the selectionId and handicap of the team it is applied to.
	KeyLineDefinition KeyLineDefinition `json:"keyLineDefinition,omitempty"`

	SuspendTime time.Time `json:"suspendTime,omitempty"`

	DiscountAllowed bool `json:"discountAllowed,omitempty"`

	PersistenceEnabled bool `json:"persistenceEnabled,omitempty"`

	Runners []RunnerDefinition `json:"runners,omitempty"`

	Version int64 `json:"version,omitempty"`

	// The Event Type the market is contained within.
	EventTypeId string `json:"eventTypeId,omitempty"`

	Complete bool `json:"complete,omitempty"`

	OpenDate time.Time `json:"openDate,omitempty"`

	MarketTime time.Time `json:"marketTime,omitempty"`

	BspReconciled bool `json:"bspReconciled,omitempty"`

	// For Handicap and Line markets, the lines available on this market will be between the range of lineMinUnit and lineMaxUnit, in increments of the lineInterval value. e.g. If unit is runs, lineMinUnit=10, lineMaxUnit=20 and lineInterval=0.5, then valid lines include 10, 10.5, 11, 11.5 up to 20 runs.
	LineInterval float64 `json:"lineInterval,omitempty"`

	Status string `json:"status,omitempty"`
}
