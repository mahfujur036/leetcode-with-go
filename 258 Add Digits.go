/// Problem 258 https://leetcode.com/problems/add-digits
/// Seems very naive at first and very easy to mistake at first
/// Can be solved using direct formula of calculating 'Digital Root'. https://en.wikipedia.org/wiki/Digital_root

func addDigits(num int) int {
    if num == 0 {
        return 0
    }
    return 1 + (num - 1) % 9

}
