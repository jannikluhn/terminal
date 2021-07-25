//SPDX-License-Identifier: MIT
pragma solidity =0.8.6;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {IERC20, ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {ERC20Burnable} from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

// An endpoint is either a sink or a source contract on a particular chain.
struct Endpoint {
    uint64 chainID;
    address contractAddress;
}

error AlreadyConnectedToEndpoint(
    Endpoint endpoint
);
error ERC20TransferFailed(
    IERC20 token,
    address spender,
    address from,
    address to,
    uint256 amount
);
error MinterAlreadySet();
error SenderIsNotMinter(
    address sender,
    address minter
);

contract HasAdmin is AccessControl {

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    constructor(address admin) {
        _setupRole(ADMIN_ROLE, admin);
        _setRoleAdmin(ADMIN_ROLE, ADMIN_ROLE);
    }
}

contract HasOracle is AccessControl {

    bytes32 public constant ORACLE_ROLE = keccak256("ORACLE_ROLE");

    constructor(address oracle) {
        _setupRole(ORACLE_ROLE, oracle);
    }
}

contract HasConnectedEndpoints is HasAdmin {

    error UnconnectedEndpoint(Endpoint endpoint);

    event ConnectedToEndpoint(Endpoint endpoint);
    event DisconnectedFromEndpoint(Endpoint endpoint);

    mapping(bytes32 => bool) public connectedEndpoints;

    constructor(address admin) HasAdmin(admin) {

    }

    function connectToEndpoint(Endpoint memory endpoint) public onlyRole(ADMIN_ROLE) {
        if (isConnectedTo(endpoint)) {
            revert AlreadyConnectedToEndpoint(endpoint);
        }

        bytes32 h = endpointHash(endpoint);
        connectedEndpoints[h] = true;
        
        emit ConnectedToEndpoint(endpoint);
    }

    function disconnectFromEndpoint(Endpoint memory endpoint) public onlyRole(ADMIN_ROLE) {
        revertIfNotConnectedTo(endpoint);

        bytes32 h = endpointHash(endpoint);
        connectedEndpoints[h] = false;
        
        emit DisconnectedFromEndpoint(endpoint);
    }

    function isConnectedTo(Endpoint memory endpoint) public view returns (bool) {
        return connectedEndpoints[endpointHash(endpoint)];
    }

    function revertIfNotConnectedTo(Endpoint memory endpoint) public view {
        bytes32 h = endpointHash(endpoint);
        if (!connectedEndpoints[h]) {
            revert UnconnectedEndpoint(endpoint);
        }
    }

    function endpointHash(Endpoint memory endpoint) public pure returns (bytes32) {
        return keccak256(abi.encode(endpoint));
    }
}

abstract contract EndpointContract is HasOracle, HasConnectedEndpoints {

    event TransferRequested(
        Endpoint endpoint,
        address sender,
        address receiver,
        uint256 amount,
        uint32 feeRate
    );

    event TransferReleased(
        address receiver,
        uint256 amount
    );

    uint32 constant FEE_RATE_DIVISOR = 100000000;

    constructor(address admin, address oracle) HasConnectedEndpoints(admin) HasOracle(oracle) {
    }

    // feeRate expressed in 0,000001%, so that 1_000_000 means 1%
    function requestTransfer(Endpoint memory endpoint, address receiver, uint256 amount, uint32 feeRate) public virtual;
    function releaseTransfer(address receiver, uint256 amount) public virtual;

}

contract ERC20Sink is EndpointContract {

    IERC20 public token;

    constructor(address admin, address oracle, IERC20 sinkToken) EndpointContract(admin, oracle) {
        token = sinkToken;
    }

    function requestTransfer(Endpoint memory endpoint, address receiver, uint256 amount, uint32 feeRate) public override {
        revertIfNotConnectedTo(endpoint);
        require(feeRate <= 100 * FEE_RATE_DIVISOR, 'fee rate exceeds 100%');

        bool success = token.transferFrom(msg.sender, address(this), amount);
        if (!success) {
            revert ERC20TransferFailed({
                token: token,
                spender: address(this),
                from: msg.sender,
                to: address(this),
                amount: amount
            });
        }

        emit TransferRequested({
            endpoint: endpoint,
            sender: msg.sender,
            receiver: receiver,
            amount: amount,
            feeRate: feeRate
        });
    }

    function releaseTransfer(address receiver, uint256 amount) onlyRole(ORACLE_ROLE) public override {
        bool success = token.transfer(receiver, amount);
        if (!success) {
            revert ERC20TransferFailed({
                token: token,
                spender: address(this),
                from: address(this),
                to: receiver,
                amount: amount
            });
        }

        emit TransferReleased({
            receiver: receiver,
            amount: amount
        });
    }

}

contract SourceToken is ERC20, ERC20Burnable {

    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    address public minter;

    constructor(string memory name, string memory symbol) ERC20(name, symbol) {
    }

    function mint(address account, uint256 amount) public virtual {
        if (msg.sender != minter) {
            revert SenderIsNotMinter(msg.sender, minter);
        }
        _mint(account, amount);
    }

    function setMinter(address newMinter) public {
        if (minter != address(0)) {
            revert MinterAlreadySet();
        }
        minter = newMinter;
    }
}

contract ERC20Source is EndpointContract {

    SourceToken public token;

    constructor(address admin, address oracle, SourceToken sourceToken) EndpointContract(admin, oracle) {
        token = sourceToken;
    }

    function requestTransfer(Endpoint memory endpoint, address receiver, uint256 amount, uint32 feeRate) public override {
        revertIfNotConnectedTo(endpoint);
        require(feeRate <= 100 * FEE_RATE_DIVISOR, 'fee rate exceeds 100%');

        token.burnFrom(msg.sender, amount);

        emit TransferRequested({
            endpoint: endpoint,
            sender: msg.sender,
            receiver: receiver,
            amount: amount,
            feeRate: feeRate
        });
    }

    function releaseTransfer(address receiver, uint256 amount) public override onlyRole(ORACLE_ROLE) {
        token.mint(receiver, amount);

        emit TransferReleased({
            receiver: receiver,
            amount: amount
        });
    }

}