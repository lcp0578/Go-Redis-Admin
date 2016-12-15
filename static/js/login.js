$(function(){
	$(".passcode").click(function(){
		$(this).attr("src", "/api/v1/reloadcaptcha?t=" + Date.parse(new Date()));
	});
})