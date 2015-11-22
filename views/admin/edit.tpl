<div id="primary" class="content-area single col-md-9">
	<div id="main" class="site-main" role="main">
		<article class="post hentry">

			<div class="input-group">
				<span class="input-group-addon">标题</span>
				<input type="text" id="article_title" class="form-control" placeholder="标题" value="{{ .Blog.Title }}">
			</div>
			<p>
				<div class="input-group">
					<span class="input-group-btn">
						<button id="add_article_category" class="btn btn-default" type="button">添加</button>
					</span>
					<span class="input-group-btn">
					  <button class="btn btn-default dropdown-toggle" type="button" id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
					    常用项
					    <span class="caret"></span>
					  </button>
					  <ul class="dropdown-menu" aria-labelledby="dropdownMenu1">
					    <li><a class="group_article_category_useful">分享</a></li>
					    <li><a class="group_article_category_useful">原创</a></li>
					  </ul>
					</span>
					<input type="text" id="input_article_category" class="form-control" placeholder="添加所属类别">

				</div>
				<span id="group_article_category">
					{{range $i, $tag := .Blog.Category }}
						<button type="button" class="btn btn-xs label_article btn-info" value="{{$tag}}">{{$tag}}   ×</button>
					{{end}}
				</span>
			</p>

			<p>
				<div class="input-group">
					<span class="input-group-btn">
						<button id="add_article_tag" class="btn btn-default" type="button">添加</button>
					</span>
					<input type="text" id="input_article_tag" class="form-control" placeholder="添加关键词标签">
				</div>
				<span id="group_article_tag">
					{{range $i, $tag := .Blog.Tags }}
						<button type="button" class="btn btn-xs label_article btn-danger" value="{{$tag}}">{{$tag}}   ×</button>
					{{end}}
				</span>
			</p>

			<p>
				<h4>
				</h4>
				<textarea class="form-control" id="article_summary" rows="2" placeholder="摘要">{{.Blog.Summary}}</textarea>
			</p>

			<p>
				<h4>
				</h4>
				<textarea class="form-control" id="article_content" rows="30" placeholder="正文">{{.Blog.Source}}</textarea>
			</p>

			<p>
				<div class="input-group">
					<span class="input-group-addon">来源</span>
					<input type="text" id="article_link" class="form-control" placeholder="来源(原创请为空)" value="{{.Blog.Link}}">
				</div>

				<div class="input-group">
					<span class="input-group-btn">
						<button id="article_commit_update" class="btn btn-default" type="button">确认提交</button>
					</span>
					<input type="text" id="commit_secret" placeholder="切口： {{.a}}+{{.b}}=？">
				</div>
			</p>
		</article>
		<!-- #post-## -->
	</div>
	<!-- #main -->
</div>