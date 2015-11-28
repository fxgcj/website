(function($){
	var load_add_tag = function(input_id, tag_class){
		$("#add_"+input_id).click(function(event) {
			add_tag(input_id, tag_class);
		});
		$("#group_"+input_id+" .label_article").each(function() {
			$(this).click(function(event) {
				 $(this).remove();
			});
		});
	};
	var add_tag = function(input_id, tag_class){
		var input_var = $("#input_"+input_id).val();

		// 是否已存在
		$("#group_"+input_id+" .label_article").each(function(){
		    if (input_var == $(this).val() ) {
		    	input_var="" ;
		    }
		});
		if (input_var=="") {
			$("#input_"+input_id).val("");
			return;
		};
		var group_class ='btn btn-xs label_article '+ tag_class;
		var input_html = '<button type="button" class="'+group_class+'" value="'+input_var+'">' + input_var + '   &times;</button>';
		$("#input_"+input_id).val("");
		$("#group_"+input_id).append(input_html);

		// 绑定删除事件
		$("#group_"+input_id+" [value='"+input_var+"']").click(function(event) {
			console.log("remove "+input_var);
			$("#group_"+input_id+" [value='"+input_var+"']").remove();
		});
	};
	var get_article_input = function(){
		var title = $("#article_title").val();
		var tags = new Array();
		$("#group_article_tag .label_article").each(function(){
			    tags.push( $(this).val() );
			});
		var categories = new Array(); 
		$("#group_article_category .label_article").each(function(){
			    categories.push( $(this).val() );
			});
		var summary =  $("#article_summary").val();
		var content =  $("#article_content").val();
		var link = $("#article_link").val();
		var secret = $.md5($('#commit_secret').val());

		var article = {
				"title":title,
				"tags":tags,
				"category":categories,
				"summary":summary,
				"content":content,
				"link":link,
				"secret":secret
			};
		return article;
	};
	var create_article = function(){
		var	objData = get_article_input();
		$.ajax({
			type: 'POST',
			url: '/admin',
			data: objData,
			dataType: "text",
			success: function(result){
				alert(result);
			}
		});
	};
	var get_URL_param = function(name){
		var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
		var r = window.location.search.substr(1).match(reg);  //匹配目标参数
		if (r!=null) 
			return unescape(r[2]); 
		return null; //返回参数值
	};
	var update_article = function(){
		var	objData = get_article_input();
		$.ajax({
			type: 'PATCH',
			url: '/admin?id='+get_URL_param("id"),
			data: objData,
			dataType: "text",
			success: function(result){
				alert(result);
			}
		});
	};
	var login_submit = function(){
		var objData = {
			"password":$.md5($('#login_passowrd').val()),
			"remember":$("#login_remember").val()
		};
		console.log(objData);
		$.ajax({
			type: 'POST',
			url: '/admin/login',
			data: objData,
			dataType: "text",
			success: function(result){
				alert(result);
				window.location.reload();
			}
		});
	}
	$.extend({
		losd_edit: function() {
			load_add_tag("article_tag","btn-danger");
			load_add_tag("article_category","btn-info");
			$("#article_commit_create").click(function(event) {
				create_article();
			});
			$("#article_commit_update").click(function(event) {
				update_article();
			});
			$(".group_article_category_useful").click(function(event) {
				$("#input_article_category").val($(this).text());
				add_tag("article_category","btn-info");
			});
			$("#login_submit").click(function(event) {
				login_submit();
			});
			console.log("load...");
		}
	});
})(jQuery);

$(document).ready(function() {
	$.losd_edit();
});