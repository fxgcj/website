<div class="form-horizontal">
  <div class="form-group" align="center">
    <label class="col-sm-6 control-label"><h2>身份验证中心</h2></label>
  </div>
  <div class="form-group">
    <label for="inputPassword3" class="col-sm-2 control-label">切口</label>
    <div class="col-sm-10">
      <input type="password" id="login_passowrd" class="form-control" placeholder="切口： {{.a}}+{{.b}}=？">
    </div>
  </div>
  <div class="form-group">
    <div class="col-sm-offset-2 col-sm-10">
      <div class="checkbox">
        <label>
          <input type="checkbox" id="login_remember"> Remember me
        </label>
      </div>
    </div>
  </div>
  <div class="form-group">
    <div class="col-sm-offset-2 col-sm-10">
      <button type="submit" id="login_submit" class="btn btn-default">认证</button>
    </div>
  </div>
</div>