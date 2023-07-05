// SPDX-License-Identifier: GPL-30
pragma solidity >= 0.7.0 < 0.9.0;

// super: 두 개 이상의 컨트렉트를 상속받아 두 개 이상의 컨트렉트의 함수를 오버라이딩 하는 경우, 가장 최근으로 상속받은 컨트렉트의 함수의 수행문만 복붙 된다

contract Father {
    event FatherName(string name);
    function who() public virtual {
        emit FatherName("JeongWooPark");
    }
}

contract Mother {
    event MotherName(string name);
    function who() public virtual {
        emit MotherName("JeongWooPark2");
    }
}

contract Son is Father, Mother {
    function who() public override(Father, Mother) {
        super.who();
    }
}

// 출력 결과
// "event": "MotherName",
// 		"args": {
// 			"0": "JeongWooPark2",
// 			"name": "JeongWooPark2"
// 		}