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

type OrderRunnerChange struct {

	// Matched Backs - matched amounts by distinct matched price on the Back side for this runner (selection)
	Mb [][]float64 `json:"mb,omitempty"`

	// Strategy Matches - Matched Backs and Matched Lays grouped by strategy reference
	Smc map[string]StrategyMatchChange `json:"smc,omitempty"`

	// Unmatched Orders - orders on this runner (selection) that are not fully matched
	Uo []Order `json:"uo,omitempty"`

	// Selection Id - the id of the runner (selection)
	Id int64 `json:"id,omitempty"`

	// Handicap - the handicap of the runner (selection) (null if not applicable)
	Hc float64 `json:"hc,omitempty"`

	FullImage bool `json:"fullImage,omitempty"`

	// Matched Lays - matched amounts by distinct matched price on the Lay side for this runner (selection)
	Ml [][]float64 `json:"ml,omitempty"`
}
