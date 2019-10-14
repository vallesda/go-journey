func flipAndInvertImage(A [][]int) [][]int {
    s := len(A)
    for i := 0; i < s; i++ {
        l := len(A[i])
        for j := 0; j <  (l + 1) / 2; j++ {
            temp := A[i][j] ^ 1
            A[i][j] = A[i][l - j - 1] ^ 1
            A[i][l - j - 1] = temp
        }
    }
    return A
}