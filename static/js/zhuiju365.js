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
})

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

function is_mobile() { 
    var regex_match = /(nokia|iphone|android|motorola|^mot-|softbank|foma|docomo|kddi|up.browser|up.link|htc|dopod|blazer|netfront|helio|hosin|huawei|novarra|CoolPad|webos|techfaith|palmsource|blackberry|alcatel|amoi|ktouch|nexian|samsung|^sam-|s[cg]h|^lge|ericsson|philips|sagem|wellcom|bunjalloo|maui|symbian|smartphone|midp|wap|phone|windows ce|iemobile|^spice|^bird|^zte-|longcos|pantech|gionee|^sie-|portalmmm|jigs browser|hiptop|^benq|haier|^lct|operas*mobi|opera*mini|320x320|240x320|176x220)/i; 
    var u = navigator.userAgent; 
    if (null == u) { 
        return true; 
    } 
    var result = regex_match.exec(u);
    if (null == result) { 
        return false 
    } else { 
        return true 
    } 
}