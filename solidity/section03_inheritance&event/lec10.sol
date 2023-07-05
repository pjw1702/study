// SPDX-License-Identifier: GPL-30
pragma solidity >= 0.7.0 < 0.9.0;

contract Father {
    string public familyName = "Kim";
    string public givenName = "Jung";
    uint256 public money = 100;

    // Son 컨트렉이 Father 컨트렉을 상속받았다 하여도, 아빠와 이름까지 같으면 안된다
    // 0.6.0 버전부터는 가시성 설정은 무시되고, 항상 public으로 간주되므로 public 키워드는 쓰지 않아도 됨
    // 컨트렉트를 비배포 가능하게 만들고 싶다면 'abstract'로 선언하면 됨
    constructor(string memory _givenName) {
        givenName = _givenName;
    }

    function getFamilyName() view public returns(string memory) {
        return familyName;
    }

    function getGivenName() view public returns(string memory) {
        return givenName;
    }

    function getMoney() view public returns(uint256) {
        return money;
    }
}

// Son 컨트렉이 Father를 상속
contract Son is Father("James") {
}