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
</div>