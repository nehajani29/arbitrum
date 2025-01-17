// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.6.11;

import "../libraries/BytesParser.sol";

/// @notice DEPRECATED - see new repo(https://github.com/OffchainLabs/token-bridge-contracts) for new updates
contract TestBytesParser {
    function bytesToString(bytes memory input)
        public
        pure
        returns (bool success, string memory res)
    {
        return BytesParser.toString(input);
    }

    function bytesToUint8(bytes memory input) public pure returns (bool, uint8) {
        return BytesParser.toUint8(input);
    }
}
