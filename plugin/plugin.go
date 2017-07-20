package plugin

// yukimori/test_udsfをimportしないとtest_udsf.CreateLoremSourceなどが参照できない
import (
    "gopkg.in/sensorbee/sensorbee.v0/bql"
    "gopkg.in/sensorbee/sensorbee.v0/bql/udf"
	"github.com/yukimori/test_udsf"
    // "github.com/sensorbee/examples/udsfs"
)

func init() {
	bql.MustRegisterGlobalSourceCreator("lorem", bql.SourceCreatorFunc(test_udsf.CreateLoremSource))
	udf.MustRegisterGlobalUDSFCreator("word_splitter", udf.MustConvertToUDSFCreator(test_udsf.CreateWordSplitter))
    udf.MustRegisterGlobalUDSFCreator("ticker", udf.MustConvertToUDSFCreator(test_udsf.CreateTicker))
}
