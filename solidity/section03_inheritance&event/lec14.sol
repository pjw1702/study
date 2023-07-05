// SPDX-License-Identifier: GPL-30
pragma solidity >= 0.7.0 < 0.9.0;

// Index: 이벤트의 키워드
// 특정 이벤트의 값들을 가져올 때 사용
// 500개의 이벤트가 발생할 때, 그 중 '박정우'에 대한 이벤트만 필요할 경우 사용
// 필터 값이라고 생각할 수 있다
contract lec14 {

    event numberTracker(uint256 num, string str);
    event numberTracker2(uint256 indexed num, string str);

    uint256 num = 0;
    function PushEvent(string memory _str) public {
        emit numberTracker(num, _str);
        emit numberTracker2(num, _str);
        num++;
    }
}