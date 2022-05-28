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

type MarketChange struct {

	// Runner Changes - a list of changes to runners (or null if un-changed)
	Rc []RunnerChange `json:"rc,omitempty"`

	// Image - replace existing prices / data with the data supplied: it is not a delta (or null if delta)
	Img bool `json:"img,omitempty"`

	// The total amount matched across the market. This value is truncated at 2dp (or null if un-changed)
	Tv float64 `json:"tv,omitempty"`

	// Conflated - have more than a single change been combined (or null if not conflated)
	Con bool `json:"con,omitempty"`

	// Market Definition - the definition of the market (or null if un-changed)
	MarketDefinition MarketDefinition `json:"marketDefinition,omitempty"`

	// Market Id - the id of the market
	Id string `json:"id,omitempty"`
}
