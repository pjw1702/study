// SPDX-License-Identifier: GPL-30
pragma solidity >= 0.7.0 < 0.9.0;

contract lec7{
    /** 솔리디티에서의 영역에 따른 데이터 저장
    stroage: 대부분의 변수, 함수들이 저장되며, 영속적으로 저장이되어 가스 비용이 비싸다 (storage라는 블록에 저장되며, 해당 블록에 대한 사용자들의 다운로드 횟수가 빈번하므로 영속적으로 저장해야 하며, 그에 따른 가스 비용도 상승된다)
    memory: 함수 내에서의 변수, 함수의 파라미터, 리턴 값, 레퍼런스 타입이 주로 저장된다. 그러나 storage 처럼 영속적이지 않고, 함수내에서만 유효하기에 storage보다 가스 비용이 싸다
    Colldata: 주로 external function의 파라미터에서 사용된다
    stack: EVM(Etherium Virtual Machine)에서 stack data를 관리할 때 쓰는 영역이며, 1024MB로 제한적이다
    */

    // string은 레퍼런스 타입으로, 기본 데이터 타입(primitive type)이 아니다
    // string 타입 등의 레퍼런스 타입은 타입 선언 뒤에 memory 키워드를 명시해야 한다

     // 아래와 같이 uint256 타입과 같은 primitive 타입은 memory 키워드가 default 설정이 되어있으므로 명시해 줄 필요 없다
    // function get_uint(uint256 memory _uint) public pure returns(uint256 memory) {
    //     return _uint;
    // }

    // function -string
    function get_string(string memory _str) public pure returns(string memory) {
        return _str;
    }


}