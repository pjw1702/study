(function($) {
  'use strict';
  $(function() {
    var todoListItem = $('.todo-list');
    var todoListInput = $('.todo-list-input');
    $('.todo-list-add-btn').on("click", function(event) { // add 버튼을 클릭했을 때의 이벤트에 대한 처리 부분으로, 이벤트 처리 로직을 백엔드에서 구현해야 함
      event.preventDefault();
      
      var item = $(this).prevAll('.todo-list-input').val(); // item: 텍스트 박스에서 input 받는 내용(텍스트)
      
      if (item) {
        // $.post("/todos", {name:item}, function(e){  // 서버가 보내준 응답이 모두 e에 JSON 오브젝트로 담겨있다(name:item과 같이 작성하면 서버에서 보내준 id 값을 읽어올 수 없다)
        //   addItem({name:item, complited:false}) // /todos 핸들러로 post 요청 시, 응답을 받았을 때 add를 하게 된다
        // })
        $.post("/todos", {name:item}, addItem)  // 서버가 JSON 오브젝트로 응답하므로 addItem 호출을 이용해 바로 사용할 수 있다
        // todoListItem.append() 메서드를 수행할 필요가 없다
        //todoListItem.append("<li><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
        todoListInput.val("");
  }
  
  });

  // /todos 페이지를 호출 시 item에 input된 텍스트 내용이 add되면서 웹 페이지에 출력
  // todos에 대한 핸들러를 백엔드에서 구현해 주어야 함
  var addItem = function(item){
    if (item.completed)
      todoListItem.append("<li class='completed'" + " id='" + item.id + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' checked='checked' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
    else
      todoListItem.append("<li" + " id='" + item.id + "'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
  };

  $.get('/todos', function(items) {
    items.forEach(e => {
      addItem(e)
    });
  });
  
  // 토글 이벤트(체크 박스)에 대한 처리
  todoListItem.on('change', '.checkbox', function() {
      var id = $(this).closest("li").attr('id');  
      var $self = $(this);
      var complete = true;
      if ($self.attr('checked')) {
        complete = true;
      }
      $.get("complete-todo/" + id + "?complete=" + complete, function(data) {
        // if ($self.attr('checked')) {
        //   $self.removeAttr('checked');
        // } else {
        //   $self.attr('checked', 'checked');
        // }
        if (complete) {
          $self.attr('checked', 'checked');
        } else {
          $self.removeAttr('checked');
        }
        
        $self.closest("li").toggleClass('completed');
      })
    });
    
    // remove 버튼을 클릭 했을 때의 이벤트 처리
    todoListItem.on('click', '.remove', function() {
      // url: todos/id method: DELETE
      var id = $(this).closest("li").attr('id');  // this의 가장 가까운 li의 id 값
      var $self = $(this);
      $.ajax({
        url: "todos/" + id,
        type: "DELETE",
        success: function(data) {
          if (data.success)
          $self.parent().remove();
        }
      })
      // $(this).parent().remove();
    });
  
  });
})(jQuery);