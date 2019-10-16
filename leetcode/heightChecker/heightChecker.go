func heightChecker(heights []int) int {
    tmp := make([]int, len(heights))
    copy(tmp, heights)
    sort.Ints(tmp)
    
    wrongPositions := 0
    for i := 0; i < len(tmp); i++ {
        if heights[i] != tmp[i] {
            wrongPositions++
        }
    }
    return wrongPositions
}