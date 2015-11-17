(function($){
	var load_add_tag = function(input_id, tag_class){
		$("#add_"+input_id).click(function(event) {
			var input_var = $("#input_"+input_id).val();
			var group_class ='btn btn-xs label_article '+ tag_class;
			var input_html = '<button type="button" class="'+group_class+'" value="'+input_var+'">' + input_var + '   &times;</button>';
			$("#group_"+input_id).append(input_html);
			$("#input_"+input_id).val("");
			load_del_tag(input_var);
		});
	};
	var load_del_tag = function(tag_value){
		$(".label_article").click(function(event) {
			$("[value='"+tag_value+"']").remove();
		});
	};
	$.extend({
		losd_edit: function() {
			load_add_tag("article_tag","btn-danger");
			load_add_tag("article_category","btn-info");
			console.log("fucke");
		}
	});
})(jQuery);

$(document).ready(function() {
	$.losd_edit();
});