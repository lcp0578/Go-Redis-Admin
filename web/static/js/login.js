$(function(){
	$(".passcode").click(function(){
		$(this).attr("src", API_URI + "reloadcaptcha?t=" + Date.parse(new Date()));
	});
	$('.btn-submit').on('click', function(){
		var username = $('#username').val();
		var password = $('#password').val();
		var captcha = $('#captcha').val();
        if(username == ""){
            layer.msg('请输入用户名');
            return false;
        }
        if(password == ""){
            layer.msg('请输入用户名');
            return false;
        }
        if(captcha == ""){
            layer.msg('请输入用户名');
            return false;
        }
        var jsonData = {
            username: username,
            password: password,
            captcha: captcha
        }
		$.post(API_URI + 'login', jsonData, function(data){
            var jobj = JSON.parse(data);
            if(jobj.Code == 1){
                location.href = "/";
            }else{
                $('.passcode').attr("src", API_URI + "newcaptcha?t=" + Date.parse(new Date()));
                layer.msg(jobj.Msg);
                return false;
            }
        });
        return false;
	});
})