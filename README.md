What's this
===========

A branch of [github.com/golang/glog](https://github.com/golang/glog)
with the addition of a Context feature:

    ctx := glog.NewCtx("module")
    ctx.Info("Starting my module")

or:

    ctx := glog.NewCtxf("Request %d", requestNo)
    ctx.Warn("Caught fire, abort transaction")

The Ctx object also has a V() method:

    ctx.V(2).Info("Blah!")
    
    if ctx.V(2) != nil {
        ctx.Info("Send in the bees")
    }



