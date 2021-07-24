//SPDX-License-Identifier: MIT
pragma solidity =0.8.6;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract TestToken is ERC20 {

    constructor() ERC20("TestToken", "TOK") {
    }

    function mint(address receiver, uint256 amount) public {
        _mint(receiver, amount);
    }

}