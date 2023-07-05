// SPDX-License-Identifier: GPL-30
pragma solidity >= 0.7.0 < 0.9.0;

// Father 컨트렉과 Mother 컨트렉에서 getMoney()라는 동일한 이름의 함수가 존재하므로 상속받는 Son 컨트렉에서 오버라이딩을 해주어야 한다
// 다중 상속을 받기 위해서는, is 키워드 뒤에 여러 컨트렉의 이름을 명시해야 한다
// 다중 상속 시에는, override(Father, Mother)과 같이 override 키워드 내에 상속받는 컨트렉의 이름을 명시해야 한다

contract Father {
    uint256 public fatherMoney = 100;

    function getFagherName() public pure returns(string memory) {
        return "jwPark";
    }

    function getMoney() public view virtual returns(uint256) {
        return fatherMoney;
    }
}

contract Mother {
    uint256 public motherMoney = 500;
    function getMotherName() public pure returns(string memory) {
        return "Leesol";
    }
    function getMoney() public view virtual returns(uint256) {
        return motherMoney;
    }
}

contract son is Father, Mother {
     function getMoney() public view override(Father, Mother)  returns(uint256) {
        return fatherMoney + motherMoney;
    }
}