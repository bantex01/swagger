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

type StatusMessage struct {

	// The operation type
	Op string `json:"op,omitempty"`

	// Client generated unique id to link request with response (like json rpc)
	Id int32 `json:"id,omitempty"`

	// The number of connections available for this account at this moment in time. Present on responses to Authentication messages only.
	ConnectionsAvailable int32 `json:"connectionsAvailable,omitempty"`

	// Additional message in case of a failure
	ErrorMessage string `json:"errorMessage,omitempty"`

	// The type of error in case of a failure
	ErrorCode string `json:"errorCode,omitempty"`

	// The connection id
	ConnectionId string `json:"connectionId,omitempty"`

	// Is the connection now closed
	ConnectionClosed bool `json:"connectionClosed,omitempty"`

	// The status of the last request
	StatusCode string `json:"statusCode,omitempty"`
}
