		<div id="secondary" class="widget-area main-sidebar col-md-3" role="complementary">
		  <aside class="widget widget_search sidebar-widget clearfix">
			<h3 class="widget-title">搜索</h3>
			<form class="search" action="/search" method="get">
			  <fieldset>
				<div class="text">
				  <input name="keyword" id="keyword" type="text" placeholder="Search ...">
				  <button class="fa fa-search">Search</button>
				</div>
			  	<img src="http://moefq.com/images/2015/11/22/cd9ff66328c3383f232402ad166cffdb.jpg" alt="cd9ff66328c3383f232402ad166cffdb.jpg" border="0">
			  </fieldset>
			</form>
		  </aside>
		  <aside class="widget widget_categories sidebar-widget clearfix">
			<h3 class="widget-title">分类目录</h3>
			<ul>
			  {{range $i, $tag := .Category}}
				<li class="cat-item"><a href="/category?name={{$tag.Name}}">{{ $tag.Name }}</a></li>
				{{end}}
			</ul>
		  </aside>
		  <aside class="widget widget_archive sidebar-widget clearfix">
			<h3 class="widget-title">文章归档</h3>
			<ul>
			  {{range .MonthBlog}}
				<li><a href="/{{setURLMonth .}}">{{.}}</a></li>
				{{end}}
			</ul>
		  </aside>	<aside class="widget widget_recent_entries sidebar-widget clearfix">
			<h3 class="widget-title">近期文章</h3>
			<ul>
			  {{range $i, $blog := .LastestBlogs}}
			  <li><a href="/blog/{{showObjectID $blog.ID}}">{{$blog.Title}}</a></li>
			  {{end}}
    		</ul>
		  </aside>
		  <aside class="widget widget_tag_cloud sidebar-widget clearfix">
			<h3 class="widget-title">标签</h3>
			<div class="tagcloud">
			  {{range $i,$tag := .Tags}}
				<a href="/tag?name={{$tag.Name}}" title="{{$tag.Name}}">{{$tag.Name}}</a>
				{{end}}
			</div>
		  </aside>
		  <aside class="widget widget_text sidebar-widget clearfix">
			<h3 class="widget-title">介绍</h3>
			<div class="textwidget">
			  <p>关注领域：安全学校，社会稳评，风险评估，风险管理，舆情分析，企业风险，等等。目前本站尚在初步建设中，更多功能将逐步完善与开放，敬请持续关注！</p>
			</div>
		  </aside>
		  <aside class="widget widget_categories sidebar-widget clearfix">
			<h3 class="widget-title">友情链接</h3>
			<ul>
			  {{range $key, $val := .FriendLinks}}
				<li class="cat-item"><a href="{{$val}}" target="_blank">{{ $key }}</a></li>
				{{end}}
			</ul>
		  </aside>
		</div>
