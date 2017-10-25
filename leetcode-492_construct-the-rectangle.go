//给定一个矩形面积, 求长和宽. 长必须大于宽, 并且长宽尽量接近
func constructRectangle(area int) []int {
    w := int(math.Sqrt(float64(area)))
    for area % w != 0 {
        w--
    }
    l := area / w
    return []int{l, w}
}
