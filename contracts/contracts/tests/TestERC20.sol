pragma solidity ^0.8.6;

import {SourceToken} from "../Endpoints.sol";

contract TestERC20 is SourceToken{

    constructor(string memory name_, string memory symbol_) ERC20(name_, symbol_) {

    }

    function mint(address account, uint256 amount) external overrides {
        _mint(account, amount);
    }
}


//SPDX-License-Identifier: MIT
