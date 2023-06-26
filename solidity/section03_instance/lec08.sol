// SPDX-License-Identifier: GPL-30
pragma solidity >= 0.7.0 < 0.9.0;

// 인스턴스 사용: 어떤 컨트렉트에서 다른 컨트렉트로 이동하기 위해
// 컨트렉트 연결

contract A{
    uint256 public a = 5;

    function change(uint256 _value) public {
        a = _value;
    }
}

// B에서 A로 접근
contract B{
    A instance = new A();

    function get_A() public view returns(uint256) {
        return instance.a();    // 변수 a 반환
    }
    function change_A(uint256 _value) public {
        instance.change(_value);    // _value 변수를 파라미터로 전달한 change()함수 반환
    }
}