<div id="primary" class="content-area col-md-9">
  <div id="main" class="site-main" role="main">
	{{ range $index, $blog := .Blogs }}
	  <article class="post hentry">
		<header class="entry-header">
		  <h1 class="post-title"><a href="/blog/{{ showObjectID $blog.ID }}" rel="bookmark">{{ $blog.Title }}</a></h1>
		  <div class="entry-meta">
			<time class="post-date"><i class="fa fa-clock-o"></i>{{ showDate $blog.Created }}</time>
			<span class="seperator">/</span>
			<span><i class="fa fa-user"></i> {{ $blog.Author }}</span>
	      </div>
		</header>
		<div class="entry-content">
		  <p>{{ $blog.Summary }}</p>
		</div>
		<footer class="entry-footer">
			  <ul class="post-categories">
					{{ range $i, $categ := $blog.Category }}
					  <li><a href="/category?name={{$categ}}" rel="category">{{$categ}}</a></li>
					  {{end}}
		          </ul>
				  
				  <ul class="post-tags">
					{{range $i, $tag:= $blog.Tags }}
					  <li><a href="/tag?name={{$tag}}" rel="tag">{{$tag}}</a></li>
					  {{end}}
		          </ul>
				  
		  <div class="read-more">
			<a href="/blog/{{ showObjectID $blog.ID }}">阅读全文<i class="fa fa-angle-double-right "></i></a>
		  </div>
		</footer>
	  </article>
	  {{end}}
  </div>
	{{if .IsPaging}}
		<div class="center">
		  <ul class="pagination">
			<li><a class="prev page-numbers" href="{{ str2html .FirstPage }}"><i class="fa fa-angle-double-left"></i></a></li>
			{{range $index,$value := .Query}}
			<p></p>
				<li><a class="prev page-numbers" href='{{str2html $value}}'><i class="fa">{{Add $index 1}}</i></a></li>
			{{end}}
			<li><a class="next page-numbers" href="{{str2html .LastPage}}"><i class="fa fa-angle-double-right"></i></a></li>
		  </ul>
		</div>
	{{end}}
</div>


