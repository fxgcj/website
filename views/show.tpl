<div id="primary" class="content-area single col-md-9">
  <div id="main" class="site-main" role="main">
	<article class="post hentry">
	  <header class="entry-header">
		<h1 class="post-title"><a href="/{{.Blog.Name}}.html" rel="bookmark">{{.Blog.Title}}</a></h1>
		<div class="entry-meta">
		  <time class="post-date"><i class="fa fa-clock-o"></i>{{ showDate .Blog.Created}}</time>
		  <span class="seperator">/</span>
	      <span><i class="fa fa-user"></i> {{.Blog.Author}}</span>
	    </div><!-- .entry-meta -->
	  </header><!-- .entry-header -->
	  <div class="entry-content">
		{{str2html .BContent}}
	  </div><!-- .entry-content -->
	  <footer class="entry-footer">
		<ul class="post-categories">
          {{range .Blog.Category}}
			<li><a href="/category?name={{.}}" rel="category">{{.}}</a></li>
			{{end}}
        </ul>
		
		<ul class="post-tags">
		  {{range .Blog.Tags}}
			<li><a href="/tag?name={{.}}" rel="tag">{{.}}</a></li>
			{{end}}
        </ul>
		
	  </footer><!-- .entry-footer -->
	</article><!-- #post-## -->

<!-- 多说评论框 start -->
	<div class="ds-thread" data-thread-key="{{showObjectID .Blog.ID}}" data-title="{{.Blog.Title}}" data-url="{{.HostUrl}}blog/{{showObjectID .Blog.ID}}"></div>
<script type="text/javascript">
var duoshuoQuery = {short_name:"fxgcj"};
	(function() {
		var ds = document.createElement('script');
		ds.type = 'text/javascript';ds.async = true;
		ds.src = (document.location.protocol == 'https:' ? 'https:' : 'http:') + '//static.duoshuo.com/embed.js';
		ds.charset = 'UTF-8';
		(document.getElementsByTagName('head')[0] 
		 || document.getElementsByTagName('body')[0]).appendChild(ds);
	})();
	</script>
<!-- 多说公共JS代码 end -->

  </div>
  <!-- #main -->
</div>