package plugin

import (
    "gopkg.in/sensorbee/sensorbee.v0/bql"
    "gopkg.in/sensorbee/sensorbee.v0/bql/udf"
	"github.com/yukimori/test_udsf"
    // "github.com/sensorbee/examples/udsfs"
)

func init() {
	bql.MustRegisterGlobalSourceCreator("lorem", bql.SourceCreatorFunc(udsfs.CreateLoremSource))
	udf.MustRegisterGlobalUDSFCreator("word_splitter", udf.MustConvertToUDSFCreator(udsfs.CreateWordSplitter))
    udf.MustRegisterGlobalUDSFCreator("ticker", udf.MustConvertToUDSFCreator(udsfs.CreateTicker))
}
