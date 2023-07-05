// SPDX-License-Identifier: GPL-30
pragma solidity >= 0.7.0 < 0.9.0;

// super 상속받은 컨트렉트 내의 모든 함수의 수행문을 복붙

contract Father {
    event FatherName(string name);
    function who() public virtual {
        emit FatherName("ParkJeongWoo");
    }
}

contract Son is Father {
    event sonName(string name);
    function who() public override {
        // emit FatherName("ParkJeongWoo");
        super.who();
        emit sonName("JeongWoo");
    }
}