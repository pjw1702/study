// SPDX-License-Identifier: GPL-30
pragma solidity >= 0.7.0 < 0.9.0;

// 생성자: 인스턴스화 할 때, 변수를 초기화하는 함수(스마트 컨트렉트 배포 시 가장 먼저 실행되는 함수)
// 모두 인스턴스화 해서 접근하면 인스턴스마다 가스비가 많이 소비되므로 clone factory 패턴을 이용해 효율적인 방안을 모색해야 한다
// 각 블록마다 소비할 수 있는 가스가 제한되어 있다: 비용적 측면 이전에 보안적 측면에서, 스마트 컨트렉트 배포 에러에 대해 고려해야 한다

contract A{
    
    string public name;
    uint256 public age;
    
     constructor(string memory _name, uint256 _age){
         name = _name;
         age = _age;
     }
        
    function change(string memory _name, uint256 _age) public  {
         name = _name;
         age = _age;
    }
}

contract B{

  A public instance = new A("Alice", 52);
  
  function change(string memory _name, uint256 _age) public  {
        instance.change(_name,_age);
    }
  
  function get() public view returns(string memory, uint256) {
        return (instance.name(), instance.age());
    }

}