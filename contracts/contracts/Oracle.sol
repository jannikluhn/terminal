pragma solidity ^0.8.6;

import "./Endpoints.sol";

contract Oracle {

    struct Request {
        bool legit;
        uint value;
        uint lastStake;
        uint32 lastChallengeTime;
        address receiver;
        address lastModifier;
        bool claimed;
        bool transferred;
    }

    struct TransferIdentifier {
        uint64 chainID;
        uint64 blockNumber;
        uint32 logIndex;
        bytes32 txHash;
    }

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
    event Claim(uint64 chaindID, uint64 blockNumber, bytes32 txHash, uint32 logIndex, uint value, address receiver);

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

    function startRequest(TransferIdentifier calldata transferIdentifier, uint value, address receiver) external {
        bytes32 hash = TransferIdentifierHash(transferIdentifier);
        require(requests[hash].lastChallengeTime == 0, 'request already exists');

        stakedToken.transferFrom(msg.sender, address(this), startingStake);

        Request storage request = requests[hash];
        request.legit = true;
        request.value = value;
        request.lastStake = startingStake;
        request.lastChallengeTime = uint32(block.timestamp);
        request.receiver = receiver;
        request.lastModifier = msg.sender;

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
        request.lastModifier = msg.sender;
        request.lastStake = newStake;
        request.lastModifier = msg.sender;

        emit UpdatedRequest(transferIdentifier.chainID, transferIdentifier.blockNumber, transferIdentifier.txHash, transferIdentifier.logIndex, newStake, request.legit);
    }

    function claimRequest(TransferIdentifier calldata transferIdentifier) external {
        bytes32 hash = TransferIdentifierHash(transferIdentifier);
        Request memory request = requests[hash];
        require(request.lastChallengeTime + challengePeriod >= block.timestamp, 'challenge period is not over');
        require(request.claimed == false, 'request already claimed');
        request.claimed = true;
        requests[hash] = request;

        require(2 * request.lastStake > request.lastStake, 'total stake overflow');
        uint totalStaked = 2 * request.lastStake - startingStake;
        stakedToken.transfer(request.lastModifier, totalStaked);

        if (! request.transferred) {
            _closeRequest(transferIdentifier);
        }

        emit Claim(transferIdentifier.chainID, transferIdentifier.blockNumber, transferIdentifier.txHash, transferIdentifier.logIndex, totalStaked, request.lastModifier);
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

    function getRequest(TransferIdentifier calldata transferIdentifier) external view returns (bool, uint, uint, uint32, address, address){
        Request memory request = requests[TransferIdentifierHash(transferIdentifier)];
        return(
            request.legit,
            request.value,
            request.lastStake,
            request.lastChallengeTime,
            request.receiver,
            request.lastModifier
        );
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
}


//SPDX-License-Identifier: MIT
