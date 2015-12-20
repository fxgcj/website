<div id="primary" class="content-area col-md-9">
	<div id="main" class="site-main" role="main">
		<article class="post hentry">
			<table class="table table-striped">
			<thead>
				<tr>
					<td>#</td>
					<td>Title</td>
					<td>Created_At</td>
					<td>Option</td>
				</tr>
			</thead>
			<tbody>
			{{ range $index, $blog := .Blogs }}
				<tr>
					<td>{{ $index }}</td>
					<td><a href="/admin/edit?id={{ showObjectID $blog.ID }}"></a>{{ $blog.Title }}</td>
					<td>{{ showDate $blog.Created }}</td>
					<td>
						<ul class="post-categories">
							<li><a href="/admin/edit?id={{ showObjectID $blog.ID }}">编辑</a></li>
							<li><a href="/admin/delete?id={{ showObjectID $blog.ID }}">删除</a></li>
						</ul>
					</td>
				</tr>
			{{end}}
			</tbody>
			</table>
		</article>
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