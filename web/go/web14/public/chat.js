$(function(){
  if (!window.EventSource) {
    alert("No EventSource!")
    return
  }

  var $chatlog = $('#chat-log') // html 파일에서 'chat-log' Element를 찾아옴
  var $chatmsg = $('#chat-msg')

  var isBlank = function(string) {
    return string == null || string.trim() === "";  // "=="는 값만 검사, "==="는 타입까지 검사
  };
  var username;
  while (isBlank(username)) {
    username = prompt("What's your name"); // 문자열 입력받음
    if (!isBlank(username)) {
      $('#user-name').html('<b>' + username + '</b>');
    }
  }
  // html에서 'input-form' Element에 'submit' 이벤트가 발생했을 때의 처리: 콜백 함수 function(e) 호출
  $('#input-form').on('submit', function(e){
    // $: jQuery 라이브러리를 의미
    $.post('/messages', {
      msg: $chatmsg.val(),  // chatmsg에 해당하는 값(valuable)
      name: username
    });  // jQuer 라이브러리의 post() 호출
    $chatmsg.val("");  // 메세지를 보냈으므로 지워 줌
    $chatmsg.focus();  // 문자열을 재 입력할 수 있게 포커싱
    return false; // submit 버튼에 대한 이벤트를 처리하는 함수에서 true를 반환하면 다른 페이지로 넘어간다
  });

  var addMessage = function(data) {
    var text = "";
    if (!isBlank(data.name)) {
      text = '<strong>' + data.name + ':</strong>';
    }
    text += data.msg;
    $chatlog.prepend('<div><span>' + text + '</span></div>');
  };
  
  // addMessage({
  //   msg: 'hello',
  //   name: 'aaa'
  // })
  
  // addMessage({
  //   msg: 'hello2'
  // })

  // EventSource open
  var es = new EventSource('/stream');
  es.onopen = function(e) { // 이벤트 소스가 오픈될 때의 이벤트 등록
    $.post('users/', {
      name: username
    });
  }
  // 이벤트 소스를 통해 메세지가 올 때의 이벤트 등록
  es.onmessage = function(e) {
    var msg = JSON.parse(e.data);
    addMessage(msg)
  }

  // 채팅방 나갈 때의 이벤트 처리
  window.onbeforeunload = function() {
    $.ajax({
      url: "/users?username=" + username,
      type: "DELETE"
    });
    es.close()
  };
})