func runningSum(nums []int) []int {
    var res = []int{}
    for i, v := range nums {
        if i == 0 {
            res = append(res, v)
        } else {
            res = append(res, res[i-1] + v)
        }
    }
    
    return res
}