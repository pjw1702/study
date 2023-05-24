(function($) {
  'use strict';
  $(function() {
    var todoListItem = $('.todo-list');
    var todoListInput = $('.todo-list-input');
    $('.todo-list-add-btn').on("click", function(event) {
      event.preventDefault();
      
      var item = $(this).prevAll('.todo-list-input').val(); // item: 텍스트 박스에서 input 받는 내용(텍스트)
      
      if (item) {
        todoListItem.append("<li><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
        todoListInput.val("");
  }
  
  });

  // /todos 페이지를 호출 시 item에 input된 텍스트 내용이 add되면서 웹 페이지에 출력
  // todos에 대한 핸들러를 백엔드에서 구현해 주어야 함
  var addItem = function(item){
    if (item.completed)
      todoListItem.append("<li class='completed'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' checked='checked' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
    else
      todoListItem.append("<li><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
  };

  $.get('/todos', function(items) {
    items.forEach(e => {
      addItem(e)
    });
  });
  
  todoListItem.on('change', '.checkbox', function() {
      if ($(this).attr('checked')) {
        $(this).removeAttr('checked');
      } else {
        $(this).attr('checked', 'checked');
      }
      
      $(this).closest("li").toggleClass('completed');
    });
    
    todoListItem.on('click', '.remove', function() {
      $(this).parent().remove();
    });
  
  });
})(jQuery);