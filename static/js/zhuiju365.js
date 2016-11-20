$(function(){
    $('#login-form').validate({
        ignore:'',
        rules : {
            username:{ required: true},
            password:{required: true}
        },
        messages : {
            username : {required: '请填写用户名'},
            password : {required: '请填写密码'}
        },
        submitHandler:function(form) {
            var url = '/login';
            $(form).ajaxSubmit({
                url:url,
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {
                       setTimeout(function(){window.location.href="/"}, 2000);
                    } else {
                       setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
                    }
                }
            });
        }
    });
    function dialogInfo(msg) {
        $('#dialogInfo').remove();
        var html = '';
        html = '<div class="modal fade" id="dialogInfo" tabindex="-1" role="dialog" aria-labelledby="dialogInfoTitle">';
        html += '<div class="modal-dialog" role="document">';
        html += '<div class="modal-content">';
        html += '<div class="modal-header">';
        html += '<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>';
        html += '<h4 class="modal-title" id="dialogInfoTitle">zhuiju365提示</h4>';
        html += ' </div>';
        html += '<div class="modal-body">';
        html += '<p>'+msg+'</p>';
        html += '</div>';
        html += '</div>';
        html += '</div>';
        html += '</div>';
        $('body').append(html);
        $('#dialogInfo').modal('show')  
    }
})
