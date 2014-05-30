What's this
===========

A branch of [github.com/golang/glog](https://github.com/golang/glog)
with the addition of a context feature:

    ctx := glog.NewCtx("module")
    ctx.Info("Starting ...")
    =>
    I0530 12:38:38.935866 25167 example.go:11] [module] Starting ...

or:

    ctx := glog.NewCtxf("Request %d", requestNo)
    ctx.Warn("Disk on fire, aborting transaction")
    =>
    W0530 12:38:38.935866 25167 example.go:11] [Request 666] Disk on fire, aborting transaction ...

The `Ctx` type also has `V()` methods:

    ctx.V(2).Info("Blah!")
    
    if ctx.VB(2) {  // VB returns a bool
        ctx.Info("The value of pi is", calcPi())
    }

The first example works because `V(level)` method returns a valid
`*Ctx` if the verbosity level is `>= level` or `nil` if it is less and
the `Info` method can be called with a `nil` receiver (it does nothing).
The second example is perfered when evaluating the arguments to Info
is expensive.


