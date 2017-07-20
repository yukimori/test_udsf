// ソースタイプのUDSFの例

type Ticker struct {
    interval time.Duration
    stopped  int32
}

// terminateメソッドが呼ばれるまで {"tick": カウント}を発行し続ける
func (t *Ticker) Process(ctx *core.Context, tuple *core.Tuple, w core.Writer) error {
    var i int64
	// terminateメソッドでStoreInt32を実行している
    for ; atomic.LoadInt32(&t.stopped) == 0; i++ {
        newTuple := core.NewTuple(data.Map{"tick": data.Int(i)})
        if err := w.Write(ctx, newTuple); err != nil {
            return err
        }
        time.Sleep(t.interval)
    }
    return nil
}

func (t *Ticker) Terminate(ctx *core.Context) error {
    atomic.StoreInt32(&t.stopped, 1)
    return nil
}
