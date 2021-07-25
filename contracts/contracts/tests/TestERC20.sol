pragma solidity ^0.8.6;

import {SourceToken} from "../Endpoints.sol";

contract TestERC20 is SourceToken{

    constructor(string memory name_, string memory symbol_) SourceToken(name_, symbol_) {

    }

    function mint(address account, uint256 amount) public override {
        _mint(account, amount);
    }
}


//SPDX-License-Identifier: MIT
