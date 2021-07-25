pragma solidity ^0.8.6;

import "./Endpoints.sol";

contract Oracle {

    struct Request {
        bool legit;
        uint value;
        uint lastStake;
        uint32 startTime;
        uint32 lastChallengeTime;
        address receiver;
        bool claimed;
        bool transferred;
        address[] stakers;
        uint32 feeRate;
    }

    struct TransferIdentifier {
        uint64 chainID;
        uint64 blockNumber;
        uint32 logIndex;
        bytes32 txHash;
    }

    uint32 constant FEE_RATE_DIVISOR = 100000000;

    mapping(bytes32 => Request) internal requests;
    uint32 public challengePeriod;
    uint public startingStake;
    IERC20 public stakedToken;
    IERC20 public transferToken;
    EndpointContract public endpoint;
    bool public initialized;

    event StartRequest(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex, uint value, address receiver);
    event UpdatedRequest(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex, uint newStake, bool legit);
    event TransferSuccessful(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex, uint value, address receiver);
    event Claim(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex);
    event BuyRequest(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex, uint price);

    function init(uint32 _challengePeriod, uint _startingStake, IERC20 _stakedToken, IERC20 _transferToken, EndpointContract _endpoint) external {
        require(! initialized, "contract already initialized");
        challengePeriod = _challengePeriod;
        startingStake = _startingStake;
        stakedToken = _stakedToken;
        transferToken = _transferToken;
        endpoint = _endpoint;
        initialized = true;
    }

    function TransferIdentifierHash(TransferIdentifier memory transferIdentifier) public pure returns(bytes32) {
        return keccak256(abi.encode(transferIdentifier));
    }

    function startRequest(TransferIdentifier calldata transferIdentifier, uint value, address receiver, uint32 feeRate) external {
        require(feeRate <= 100 * FEE_RATE_DIVISOR, 'fee rate exceeds 100%');
        bytes32 hash = TransferIdentifierHash(transferIdentifier);
        require(requests[hash].lastChallengeTime == 0, 'request already exists');

        stakedToken.transferFrom(msg.sender, address(this), startingStake);

        Request storage request = requests[hash];
        request.legit = true;
        request.value = value;
        request.lastStake = startingStake;
        request.startTime = uint32(block.timestamp);
        request.lastChallengeTime = uint32(block.timestamp);
        request.receiver = receiver;
        request.stakers = [msg.sender];
        request.feeRate = feeRate;

        emit StartRequest(transferIdentifier.chainID, transferIdentifier.blockNumber, transferIdentifier.txHash, transferIdentifier.logIndex, value, receiver);
    }

     /**
     * @notice Update a request to the opposite legit status
     *
     * @param transferIdentifier the identifier of the transfer to update the request for
     * @param expectedNewStake the expected lastStake for the request, to make sure it was not updated
     *        in between the transaction signing and execution
     *
     */
    function updateRequest(TransferIdentifier calldata transferIdentifier, uint expectedNewStake) external {
        bytes32 hash = TransferIdentifierHash(transferIdentifier);
        Request storage request = requests[hash];
        require(request.lastChallengeTime != 0, 'request does not exist');
        require(request.lastChallengeTime + challengePeriod < block.timestamp, 'challenge period is over');

        uint newStake = request.lastStake * 2;
        require(newStake > request.lastStake, 'request stake overflow');
        // TODO: think about how to handle this
        require(expectedNewStake == newStake, 'expected wrong stake');

        stakedToken.transferFrom(msg.sender, address(this), newStake);

        request.legit = ! request.legit;
        request.lastChallengeTime = uint32(block.timestamp);
        request.lastStake = newStake;
        request.stakers[request.stakers.length] = msg.sender;

        emit UpdatedRequest(transferIdentifier.chainID, transferIdentifier.blockNumber, transferIdentifier.txHash, transferIdentifier.logIndex, newStake, request.legit);
    }

    function claimRequest(TransferIdentifier calldata transferIdentifier) external {
        bytes32 hash = TransferIdentifierHash(transferIdentifier);
        Request memory request = requests[hash];
        require(request.lastChallengeTime + challengePeriod >= block.timestamp, 'challenge period is not over');
        require(request.claimed == false, 'request already claimed');
        request.claimed = true;
        requests[hash] = request;

        uint stakeToWithdraw = startingStake;
        if (request.legit) {
            assert(request.stakers.length % 2 == 1);
            for (uint i = 0; i < request.stakers.length; i += 2) {
                stakedToken.transfer(request.stakers[i], stakeToWithdraw);
                stakeToWithdraw *= 6;
            }
        } else {
            stakeToWithdraw = startingStake * 2;
            assert(request.stakers.length % 2 == 0);
            for (uint i = 1; i < request.stakers.length; i += 2) {
                stakedToken.transfer(request.stakers[i], stakeToWithdraw);
                stakeToWithdraw *= 6;
            }
        }

        if (! request.transferred) {
            _closeRequest(transferIdentifier);
        }

        emit Claim(transferIdentifier.chainID, transferIdentifier.blockNumber, transferIdentifier.txHash, transferIdentifier.logIndex);
    }

    function buyRequest(TransferIdentifier calldata transferIdentifier) external {
        bytes32 hash = TransferIdentifierHash(transferIdentifier);
        Request memory request = requests[hash];
        require(request.feeRate != 0, 'request cannot be bought');

        uint price = requestBuyPrice(request.value, request.feeRate, request.startTime);
        transferToken.transferFrom(msg.sender, request.receiver, price);
        request.receiver = msg.sender;
        request.feeRate = 0;
        requests[hash] = request;

        emit BuyRequest(transferIdentifier.chainID, transferIdentifier.blockNumber, transferIdentifier.txHash, transferIdentifier.logIndex, price);
    }

    function closeRequest(TransferIdentifier calldata transferIdentifier) external {
        bytes32 hash = TransferIdentifierHash(transferIdentifier);
        Request memory request = requests[hash];
        require(request.lastChallengeTime + challengePeriod >= block.timestamp, 'request challenge period is not over');
        require(request.legit, 'request has been denied');
        _closeRequest(transferIdentifier);
    }

    function _closeRequest(TransferIdentifier memory transferIdentifier) internal {
        Request storage request = requests[TransferIdentifierHash(transferIdentifier)];
        request.transferred = true;
        endpoint.releaseTransfer(request.receiver, request.value);
        emit TransferSuccessful(transferIdentifier.chainID, transferIdentifier.blockNumber, transferIdentifier.txHash, transferIdentifier.logIndex, request.value, request.receiver);
    }

    function getRequest(TransferIdentifier calldata transferIdentifier) external view returns (Request memory){
        return requests[TransferIdentifierHash(transferIdentifier)];
    }

    function isSuccessfulRequest(TransferIdentifier calldata transferIdentifier, uint value, address receiver) external view returns (bool) {
        Request memory request = requests[TransferIdentifierHash(transferIdentifier)];
        if (
            request.value == value &&
            request.receiver == receiver &&
            request.lastChallengeTime + challengePeriod >= block.timestamp &&
            request.legit
        ) {
            return true;
        }
        return false;
    }

    function requestBuyPrice(uint value, uint32 feeRate, uint32 startTime) internal view returns (uint) {
        require(block.timestamp - startTime < challengePeriod, 'request initial challenge time over');
        uint totalFee = value / FEE_RATE_DIVISOR * feeRate;
        uint elapsedTime = challengePeriod - (block.timestamp - startTime);
        uint finalFee = totalFee * elapsedTime / challengePeriod;
        return value - finalFee;
    }
}


//SPDX-License-Identifier: MIT
