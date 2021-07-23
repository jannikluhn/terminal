pragma solidity ^0.8.6;

import "./interfaces/IERC20.sol";

contract Oracle {

    struct Request {
        bool legit;
        uint value;
        uint lastStake;
        uint32 lastChallengeTime;
        address receiver;
        address lastModifier;
        bool claimed;
    }

    mapping(bytes32 => Request) internal requests;
    uint32 public challengePeriod;
    uint public startingStake;
    IERC20 public stakedToken;
    IERC20 public transferToken;

    event StartRequest(bytes32 transferIdentifier, uint value, address receiver);
    event UpdatedRequest(bytes32 transferIdentifier, uint newStake, bool legit);

    constructor(uint32 _challengePeriod, uint _startingStake, IERC20 _stakedToken, IERC20 _transferToken){
        challengePeriod = _challengePeriod;
        startingStake = _startingStake;
        stakedToken = _stakedToken;
        transferToken = _transferToken;
    }

    function startRequest(bytes32 transferIdentifier, uint value, address receiver) external {
        require(requests[transferIdentifier].lastChallengeTime == 0, 'request already exists');

        stakedToken.transferFrom(msg.sender, address(this), startingStake);

        Request storage request = requests[transferIdentifier];
        request.legit = true;
        request.value = value;
        request.lastStake = startingStake;
        request.lastChallengeTime = uint32(block.timestamp);
        request.receiver = receiver;
        request.lastModifier = msg.sender;

        emit StartRequest(transferIdentifier, value, receiver);
    }

     /**
     * @notice Update a request to the opposite legit status
     *
     * @param transferIdentifier the identifier of the transfer to update the request for
     * @param expectedNewStake the expected lastStake for the request, to make sure it was not updated
     *        in between the transaction signing and execution
     *
     */
    function updateRequest(bytes32 transferIdentifier, uint expectedNewStake) external {
        Request storage request = requests[transferIdentifier];
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

        emit UpdatedRequest(transferIdentifier, newStake, request.legit);
    }

    function claimRequest(bytes32 transferIdentifier) external {
        Request memory request = requests[transferIdentifier];
        require(request.lastChallengeTime + challengePeriod >= block.timestamp, 'challenge period is not over');
        require(request.lastModifier == msg.sender, 'last request modifier is not msg.sender');
        require(request.claimed == false, 'request already claimed');
        request.claimed = true;
        requests[transferIdentifier] = request;

        require(2 * request.lastStake > request.lastStake, 'total stake overflow');
        uint totalStaked = 2 * request.lastStake - startingStake;
        stakedToken.transfer(msg.sender, totalStaked);
    }

    function getRequest(bytes32 transferIdentifier) external view returns (bool, uint, uint, uint32, address, address){
        Request memory request = requests[transferIdentifier];
        return(
            request.legit,
            request.value,
            request.lastStake,
            request.lastChallengeTime,
            request.receiver,
            request.lastModifier
        );
    }

    function isSuccessfulRequest(bytes32 transferIdentifier, uint value, address receiver) external view returns (bool) {
        Request memory request = requests[transferIdentifier];
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
