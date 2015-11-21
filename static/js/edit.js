(function($){
	var load_add_tag = function(input_id, tag_class){
		$("#add_"+input_id).click(function(event) {
			add_tag(input_id, tag_class);
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
	var post_article = function(){
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
		// console.log("hero");
		// console.log("title: "+title);
		// console.log(tags);
		// console.log(categories);
		// console.log(summary);
		// console.log(content);
		// console.log(secret);
		$.ajax({
			type: 'POST',
			url: '/admin',
			data: {
				"title":title,
				"tags":tags,
				"category":categories,
				"summary":summary,
				"content":content,
				"link":link,
				"secret":secret
			},
			dataType: "text",
			success: function(result){
				console.log(result);
			}
		});
	};
	$.extend({
		losd_edit: function() {
			load_add_tag("article_tag","btn-danger");
			load_add_tag("article_category","btn-info");
			$("#article_commit").click(function(event) {
				post_article();
			});
			$(".group_article_category_useful").click(function(event) {
				$("#input_article_category").val($(this).text());
				add_tag("article_category","btn-info");
			});
			console.log("load...");
		}
	});
})(jQuery);

$(document).ready(function() {
	$.losd_edit();
});