// SPDX-License-Identifier: GPL-30
pragma solidity >= 0.7.0 < 0.9.0;

// 매핑(Mapping): Key-Value 쌍으로 이루어져 있다
// 특정 Key를 입력하면, 해당 Key에 대응하는 Value를 반환
// mapping(Key의 타입=>Value의 타입) <접근 제어자> <매핑 후 반환받을 변수>;

// 매핑은 Key-Value 쌍으로 이루어져 있으므로 길이(Length)를 구할 수 없다
contract lec17 {

    mapping(uint256=>uint256) private ageList;
    mapping(string=>uint256) private priceList;
    mapping(uint256=>string) private nameList;

    function setAgeList(uint256 _index, uint256 _age) public {
        ageList[_index] = _age;
    }

    function getAge(uint256 _index) public view returns(uint256) {
        return ageList[_index];
    }

    function setNameList(uint256 _index, string memory _name) public {
        nameList[_index] = _name;
    }

    function getName(uint256 _index) public view returns(string memory) {
        return nameList[_index];
    }

    function setPriceList(string memory _itemName, uint256 _price) public {
        priceList[_itemName] = _price;
    }

    function getPriceList(string memory _index) public view returns(uint256) {
        return priceList[_index];
    }
}