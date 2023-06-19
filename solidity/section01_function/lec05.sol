// SPDX-License-Identifier: GPL-30
pragma solidity >= 0.7.0 < 0.9.0;

contract lec5{
    /**
    function 이름 () public {   // (public, private, internal, external) 변경 가능.
        // 내용
    }
    */

    // 접근 제어자는 컨트랙 뿐만 아니라 변수에도 적용 가능하다

    /**
    public: 모든 곳에서 접근 가능
    external: public 처럼 모든 곳에서 접근 가능하나, external이 정의된 자기자신의 컨트랙 내에서는 접근 불가
    private: 오직 private이 정의된 자기 컨트랙에서만 가능(private이 정의된 컨트랙을 상속 받은 자식도 불가능)
    internal: private 처럼 internal이 정의된 자기 컨트랙에서만 가능하지만, internal이 정의된 컨트랙을 상속 받은 자식은 접근 가능
    */

    // 1. public
    uint256 public a = 5;

    // 2. private
    uint256 private a2 = 5; // 디플로이된 컨트랙에서, a2의 변수는 볼 수 없다
}

contract Public_example {
    uint256 public a = 3;

    function changeA(uint256 _value) external  {
        a = _value;
    }
    function get_a() view external  returns (uint256) {
        return a;
    }
}

contract Public_example2 {
    Public_example instance =  new Public_example();    // 인스턴스화

    function changeA_2(uint256 _value) public {
        instance.changeA(_value);
    }
    function use_public_example_a() view public returns (uint256) {
        return instance.get_a();
    }
}