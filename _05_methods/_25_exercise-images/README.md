* Exercise: Images

Remember the [[/tour/moretypes/18][picture generator]] you wrote earlier? Let's write another one, but this time it will return an implementation of `image.Image` instead of a slice of data.

Define your own `Image` type, implement [[/pkg/image/#Image][the necessary methods]], and call `pic.ShowImage`.

`Bounds` should return a `image.Rectangle`, like `image.Rect(0,`0,`w,`h)`.

`ColorModel` should return `color.RGBAModel`.

`At` should return a color; the value `v` in the last picture generator corresponds to `color.RGBA{v,`v,`255,`255}` in this one.

.play methods/exercise-images.go

* Congratulations!

You finished this lesson!

You can go back to the list of [[/tour/list][modules]] to find what to learn next, or continue with the [[javascript:click('.next-page')][next lesson]].
