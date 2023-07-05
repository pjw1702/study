// SPDX-License-Identifier: GPL-30
pragma solidity >= 0.7.0 < 0.9.0;

// 함수 덮어쓰기라고 생각할 수 있다
// 아버지가 아들에게서 100만원을 상속 받을 때, 100만원으로 개인의 취향에 맞게 소비 활동을 할 수 있다
// virtual과 overrride 키워드를 사용

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

    function getMoney() view public virtual returns(uint256) {
        return money;
    }
}

// Son 컨트렉이 Father를 상속
contract Son is Father("James") {
    // 상속 시, 아래와 같은 방법으로도 초기화할 수 있다
    //constructor() Father("James") {}

    uint256 public earning = 0;
    function work() public {
        earning += 100;
    }

    function getMoney() view override public returns(uint256) {
        return money+earning;
    }
}