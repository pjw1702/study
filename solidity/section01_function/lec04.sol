// SPDX-License-Identifier: GPL-30
pragma solidity >= 0.7.0 < 0.9.0;

contract lec4{
    /**
    function 이름 () public {   // (public, private, internal, external) 변경 가능.
        // 내용
    }
    */

    // 1. parameter와 return 값이 없는 function 정의 
    // 2. parameter는 있고, return 값이 없는 function 정의 
    // 3. parameter와 return 값이 모두 있는 function 정의 

    uint256 public a = 3;
    // 1. parameter와 return 값이 없는 function 정의
    function changeA1() public {
        a = 5;
    }

    // 2. parameter는 있고, return 값이 없는 function 정의 
    function changeA2(uint256 _value) public {
        a = _value;
    }

    // 3. parameter와 return 값이 모두 있는 function 정의
    function changeA3(uint256 _value) public returns(uint256) {
        a = _value;
        return a;
    }
}