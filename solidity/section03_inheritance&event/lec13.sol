// SPDX-License-Identifier: GPL-30
pragma solidity >= 0.7.0 < 0.9.0;

// 솔리디티에는 print 함수가 없으므로 event 함수를 통해 값을 출력할 수 있다
// event호출 시에는 블록내에 값이 저장된다
// 언제든지 블록에서 event 함수 호출로 인해 저장된 값을 저장할 수 있다
// event <이벤트 이름>(출력할 값)과 같이 호출
// event호출 후에는 emit 함수를 호출해야 값이 출력된다

contract lec13 {
    event info(string name, uint256 money);

    function sendMoney() public {
        emit info("ParkJeongWoo", 1000);
    }
}

// 배포 후, 아래 로그를 통해 출력을 확인
// "args": {
// 			"0": "ParkJeongWoo",
// 			"1": "1000",
// 			"name": "ParkJeongWoo",
// 			"money": "1000"
// 	}