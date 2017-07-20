package test_udsf

import (
    "fmt"
    "strings"

    "gopkg.in/sensorbee/sensorbee.v0/bql/udf"
    "gopkg.in/sensorbee/sensorbee.v0/core"
    "gopkg.in/sensorbee/sensorbee.v0/data"
)

type WordSplitter struct {
    field string
}

// ProcessはUDSFインタフェースのメソッド
// ctx 処理のコンテキスト情報
// t 入力情報 他のストリームから情報をうけとるストリーム型UDSF <- 他のストリームからの入力を有する新しいtupleが到着する度にProcessメソッドが呼び出される
// 出力は複数のtuple（入力tupleのデータをスペースで分割するため）
// w タプルが出力される宛先
// TODO: core.Wrierとは？
func (w *WordSplitter) Process(ctx *core.Context,
    t *core.Tuple, writer core.Writer) error {
    var kwd []string
	// v 値か？ 値がnullならtrueになると思われる
    if v, ok := t.Data[w.field]; !ok {
        return fmt.Errorf("the tuple doesn't have the required field: %v", w.field)
    } else if s, err := data.AsString(v); err != nil {  // 値が文字列にならない -> この分岐
        return fmt.Errorf("'%v' field must be string: %v", w.field, err)
    } else {
        kwd = strings.Split(s, " ")
    }

	// tをoutにcopyしている
	// 空白で分割して単語数分繰り返される
	// w.fieldをキーとして値が複数格納される？ -> YESっぽい 複数のtupleを生成しているg
    for _, k := range kwd {
        out := t.Copy()
        out.Data[w.field] = data.String(k)
		// writerを使ってctxにoutを書き込んでいる ctxはコンテキストか？
        if err := writer.Write(ctx, out); err != nil {
            return err
        }
    }
    return nil
}

// TerminateはUDSFインタフェースのメソッド
func (w *WordSplitter) Terminate(ctx *core.Context) error {
    return nil
}

// このメソッドはUDSFDelareerを持っている
// このメソッドをRegisterGlobalUDSFCreator関数またはMustRegisterGlobalUDSFCreator関数で
// UDSFCreatorを登録することでBQLで利用できる
func CreateWordSplitter(decl udf.UDSFDeclarer,
    inputStream, field string) (udf.UDSF, error) {
    if err := decl.Input(inputStream, nil); err != nil {
        return nil, err
    }
    return &WordSplitter{
        field: field,
    }, nil
}
