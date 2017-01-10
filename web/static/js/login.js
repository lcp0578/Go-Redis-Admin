$(function(){
	$(".passcode").click(function(){
		$(this).attr("src", "/api/v1/reloadcaptcha?t=" + Date.parse(new Date()));
	});
	$('.btn-submit').on('click', function(){
		var username = $('#username').val();
		var password = $('#password').val();
		var captcha = $('#captcha').val();
		$.post('');
	});
})