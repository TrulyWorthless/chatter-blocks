// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;

interface ISimpleChannel {
    struct Message {
        address sender;
        string text;
        uint256 timestamp;
    }

    enum Actor {
        Alice,
        Bob
    }

    event MessageSubmit(uint256 index, Actor role, string message);

    function submitMessage(string memory text) external returns (bool);

    function readResponse() external view returns(Message memory);

    function readResponseAt(uint256 index) external view returns (Message memory);

    function readThread() external view returns (Message[] memory);

    function messageCount() external view returns(uint256);
}