package markdown

import (
	"fmt"
	"testing"
)

func TestMarkdown(t *testing.T) {
	mark := []byte(`# Test

* hahah[click](http://www.ckeyer.com/)
* hahah[click](http://www.ckeyer.com/)

<h2>logo</h2>
<image src="http://www.ckeyer.com/static/img/logo.png"/>

### hello world

hello, this  is a test...

		`)
	html := Trans2html(mark)
	fmt.Println(string(html))
}
